package adapters

import (
	"bufio"
	"context"
	"errors"
	"io"
	"os"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/vault/file/adapters/converter"
	"github.com/Chystik/pass-man/internal/vault/file/entities"

	"google.golang.org/grpc"
)

type FileAPIClient interface {
	UploadFile(ctx context.Context, file *entities.File, filePath string) error
	DownloadFile(ctx context.Context, file *entities.File, filePath string) error
	ListFiles(ctx context.Context) ([]*entities.File, error)
	DeleteFile(ctx context.Context, file *entities.File) error
}

const (
	bufferSize = 4 * 1024
)

type fileAPIClient struct {
	conn *grpc.ClientConn
	file pb.FileServiceClient
	FileAPIClient
}

func NewFileAPIClient(conn *grpc.ClientConn, file pb.FileServiceClient) *fileAPIClient {
	return &fileAPIClient{
		conn: conn,
		file: file,
	}
}

func (fa *fileAPIClient) UploadFile(ctx context.Context, file *entities.File, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	file.Name = f.Name()

	stream, err := fa.file.Upload(ctx)
	if err != nil {
		return err
	}

	req := &pb.UploadFileRequest{
		Data: &pb.UploadFileRequest_File{
			File: converter.FromDomainFile(file),
		},
	}

	err = stream.Send(req)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(f)
	buffer := make([]byte, bufferSize)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		req := &pb.UploadFileRequest{
			Data: &pb.UploadFileRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			return err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	if res.Error != nil {
		return errors.New(res.Error.String())
	}

	return nil
}

func (fa *fileAPIClient) DownloadFile(ctx context.Context, file *entities.File, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	req := &pb.DownloadFileRequest{
		File: converter.FromDomainFile(file),
	}

	stream, err := fa.file.Download(ctx, req)
	if err != nil {
		return err
	}

	res, err := stream.Recv()
	if err != nil {
		return err
	}

	fm := res.GetFile().GetMeta()
	_ = res.GetFile().GetFullName()

	file.Meta = fm
	fileSize := 0

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		chunk := res.GetChunkData()
		size := len(chunk)

		fileSize += size

		_, err = f.Write(chunk)
		if err != nil {
			return nil
		}
	}

	return stream.CloseSend()
}

func (fa *fileAPIClient) ListFiles(ctx context.Context) ([]*entities.File, error) {
	f := []*entities.File{}
	req := &pb.ListFileRequest{}

	res, err := fa.file.ListFiles(ctx, req)
	if err != nil {
		return f, err
	}

	if res.Error != nil {
		return f, errors.New(res.Error.String())
	}

	return converter.ToDomainFiles(res.FileList), nil
}

func (fa *fileAPIClient) DeleteFile(ctx context.Context, file *entities.File) error {
	req := &pb.DeleteFileRequest{
		File: converter.FromDomainFile(file),
	}

	res, err := fa.file.Delete(ctx, req)
	if err != nil {
		return err
	}

	if res.Error != nil {
		return errors.New(res.Error.String())
	}

	return nil
}
