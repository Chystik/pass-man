package adapters

import (
	"bytes"
	"context"
	"io"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/user/entities"
	"github.com/Chystik/pass-man/internal/vault/file/adapters/converter"
	"github.com/Chystik/pass-man/internal/vault/file/usecases"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	maxFileSize = 5 << 30 // 5 GiB
	bufferSize  = 4 * 1024
)

type fileHandlers struct {
	usecases usecases.FileUsecases
	pb.UnimplementedFileServiceServer
}

func NewFileHandlers(fu usecases.FileUsecases) *fileHandlers {
	return &fileHandlers{
		usecases: fu,
	}
}

func (fh *fileHandlers) Upload(fStream pb.FileService_UploadServer) error {
	ctx := fStream.Context()

	userID, err := entities.GetLoginFromContext(ctx)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	// First receive file info
	req, err := fStream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive file info")
	}

	fm := req.GetFile().GetMeta()
	fn := req.GetFile().GetFullName()

	file := converter.ToDomainFile(&pb.File{Meta: fm, FullName: fn})

	file.Data = &bytes.Buffer{}
	fileSize := 0

	for {
		// Now we can receive chunks
		req, err = fStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		fileSize += size
		if fileSize > maxFileSize {
			return status.Errorf(codes.InvalidArgument, "file is too large: %d > %d", fileSize, maxFileSize)
		}

		_, err = file.Data.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}
	}

	_, err = fh.usecases.Upload(ctx, userID, file)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot save file to the store: %v", err)
	}

	res := pb.UploadFileResponse{
		Size: uint32(fileSize),
	}

	err = fStream.SendAndClose(&res)
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot send response: %v", err)
	}

	return nil
}
func (fh *fileHandlers) Download(f *pb.DownloadFileRequest, fStream pb.FileService_DownloadServer) error {
	ctx := fStream.Context()

	userID, err := entities.GetLoginFromContext(ctx)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	file := converter.ToDomainFile(f.File)
	file.Data = &bytes.Buffer{}

	_, err = fh.usecases.Download(ctx, userID, file)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	res := &pb.DownloadFileResponse{
		Data: &pb.DownloadFileResponse_File{
			File: converter.FromDomainFile(file),
		},
	}

	err = fStream.Send(res)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	buffer := make([]byte, bufferSize)

	for {
		n, err := file.Data.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Error(codes.Unknown, err.Error())
		}

		res := &pb.DownloadFileResponse{
			Data: &pb.DownloadFileResponse_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = fStream.Send(res)
		if err != nil {
			return status.Error(codes.Unknown, err.Error())
		}
	}

	return nil
}

func (fh *fileHandlers) ListFiles(ctx context.Context, f *pb.ListFileRequest) (*pb.ListFileResponse, error) {
	var response pb.ListFileResponse

	userID, err := entities.GetLoginFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	filesList, err := fh.usecases.ListFiles(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list files error: %s", err.Error())
	}

	response.FileList = converter.FromDomainFiles(filesList)

	return &response, nil
}

func (fh *fileHandlers) Delete(ctx context.Context, f *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	var response pb.DeleteFileResponse

	userID, err := entities.GetLoginFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	err = fh.usecases.Delete(ctx, userID, converter.ToDomainFile(f.File))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete file error: %s", err.Error())
	}

	return &response, nil
}
