package adapters

import (
	"bufio"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	mocks "github.com/Chystik/pass-man/internal/infrastructure/grpc/mock"
	"github.com/Chystik/pass-man/internal/vault/file/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUploadFile_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	stream := &mocks.FileService_UploadClient{}
	tmpFile, err := os.CreateTemp("", "")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write([]byte("temporary file test data"))
	require.NoError(t, err)

	mk.file.EXPECT().Upload(ctx, mock.Anything).Return(stream, nil)
	stream.EXPECT().Send(mock.Anything).Return(nil).Once()
	stream.EXPECT().Send(mock.Anything).Return(nil).Once()
	stream.EXPECT().CloseAndRecv().Return(&pb.UploadFileResponse{}, nil)
	err = cl.UploadFile(ctx, &entities.File{}, tmpFile.Name())

	assert.NoError(t, err)
}

func TestUploadFile_WhenClientUploadReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	stream := &mocks.FileService_UploadClient{}
	tmpFile, err := os.CreateTemp("", "")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write([]byte("temporary file test data"))
	require.NoError(t, err)

	mk.file.EXPECT().Upload(ctx, mock.Anything).Return(stream, errors.New("err"))
	err = cl.UploadFile(ctx, &entities.File{}, tmpFile.Name())

	assert.Error(t, err)
}

func TestUploadFile_WhenFirstSendReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	stream := &mocks.FileService_UploadClient{}
	tmpFile, err := os.CreateTemp("", "")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write([]byte("temporary file test data"))
	require.NoError(t, err)

	mk.file.EXPECT().Upload(ctx, mock.Anything).Return(stream, nil)
	stream.EXPECT().Send(mock.Anything).Return(errors.New("err"))
	err = cl.UploadFile(ctx, &entities.File{}, tmpFile.Name())

	assert.Error(t, err)
}

func TestUploadFile_WhenSecondSendReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	stream := &mocks.FileService_UploadClient{}
	tmpFile, err := os.CreateTemp("", "")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write([]byte("temporary file test data"))
	require.NoError(t, err)

	mk.file.EXPECT().Upload(ctx, mock.Anything).Return(stream, nil)
	stream.EXPECT().Send(mock.Anything).Return(nil).Once()
	stream.EXPECT().Send(mock.Anything).Return(errors.New("err")).Once()
	err = cl.UploadFile(ctx, &entities.File{}, tmpFile.Name())

	assert.Error(t, err)
}

func TestUploadFile_WhenCloseAndRecvReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	stream := &mocks.FileService_UploadClient{}
	tmpFile, err := os.CreateTemp("", "")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write([]byte("temporary file test data"))
	require.NoError(t, err)

	mk.file.EXPECT().Upload(ctx, mock.Anything).Return(stream, nil)
	stream.EXPECT().Send(mock.Anything).Return(nil)
	stream.EXPECT().CloseAndRecv().Return(&pb.UploadFileResponse{}, errors.New("err"))
	err = cl.UploadFile(ctx, &entities.File{}, tmpFile.Name())

	assert.Error(t, err)
}

func TestUploadFile_WhenStreamReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	stream := &mocks.FileService_UploadClient{}
	tmpFile, err := os.CreateTemp("", "")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write([]byte("temporary file test data"))
	require.NoError(t, err)

	mk.file.EXPECT().Upload(ctx, mock.Anything).Return(stream, nil)
	stream.EXPECT().Send(mock.Anything).Return(nil)
	stream.EXPECT().CloseAndRecv().Return(&pb.UploadFileResponse{
		Error: &pb.Error{},
	}, nil)
	err = cl.UploadFile(ctx, &entities.File{}, tmpFile.Name())

	assert.Error(t, err)
}

func TestUploadFile_WhenOpenFileReturnsErr(t *testing.T) {
	t.Parallel()

	_, cl := newFileAPIClientMock()
	ctx := context.Background()
	err := cl.UploadFile(ctx, &entities.File{}, "wrong path")

	assert.Error(t, err)
}

func TestDownloadFile_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	stream := &mocks.FileService_DownloadClient{}
	chunk := []byte("test file content")
	tmpFile, err := os.CreateTemp("", "")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())

	resp := &pb.DownloadFileResponse{
		Data: &pb.DownloadFileResponse_File{
			File: &pb.File{
				Meta: "test",
			},
		},
	}
	respChunks := &pb.DownloadFileResponse{
		Data: &pb.DownloadFileResponse_ChunkData{
			ChunkData: chunk,
		},
	}

	mk.file.EXPECT().Download(ctx, mock.Anything).Return(stream, nil)
	stream.EXPECT().Recv().Return(resp, nil).Once()
	stream.EXPECT().Recv().Return(respChunks, nil).Once()
	stream.EXPECT().Recv().Return(nil, io.EOF).Once()
	stream.EXPECT().CloseSend().Return(nil)
	err = cl.DownloadFile(ctx, &entities.File{Meta: "test"}, tmpFile.Name())

	assert.NoError(t, err)

	downloaded := make([]byte, len(chunk))
	reader := bufio.NewReader(tmpFile)
	_, err = reader.Read(downloaded)
	require.NoError(t, err)

	assert.Equal(t, chunk, downloaded)
}

func TestDownloadFile_WhenClientDownloadReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	tmpFile, err := os.CreateTemp("", "")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())

	mk.file.EXPECT().Download(ctx, mock.Anything).Return(nil, errors.New("err"))
	err = cl.DownloadFile(ctx, &entities.File{Meta: "test"}, tmpFile.Name())

	assert.Error(t, err)
}

func TestDownloadFile_WhenStreamRecvReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	stream := &mocks.FileService_DownloadClient{}
	tmpFile, err := os.CreateTemp("", "")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())

	mk.file.EXPECT().Download(ctx, mock.Anything).Return(stream, nil)
	stream.EXPECT().Recv().Return(nil, errors.New("err")).Once()
	err = cl.DownloadFile(ctx, &entities.File{Meta: "test"}, tmpFile.Name())

	assert.Error(t, err)
}

func TestDownloadFile_WhenCreateFileReturnsErr(t *testing.T) {
	t.Parallel()

	_, cl := newFileAPIClientMock()
	ctx := context.Background()

	err := cl.DownloadFile(ctx, &entities.File{Meta: "test"}, "")

	assert.Error(t, err)
}

func TestListFiles_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	resp := &pb.ListFileResponse{}

	mk.file.EXPECT().ListFiles(ctx, mock.Anything).Return(resp, nil)
	_, err := cl.ListFiles(ctx)

	assert.NoError(t, err)
}

func TestListFiles_WhenFileClientReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()

	mk.file.EXPECT().ListFiles(ctx, mock.Anything).Return(nil, errors.New("err"))
	_, err := cl.ListFiles(ctx)

	assert.Error(t, err)
}

func TestListFiles_WhenFileClientRespReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	resp := &pb.ListFileResponse{
		Error: &pb.Error{},
	}

	mk.file.EXPECT().ListFiles(ctx, mock.Anything).Return(resp, nil)
	_, err := cl.ListFiles(ctx)

	assert.Error(t, err)
}

func TestDeleteFile_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	resp := &pb.DeleteFileResponse{}

	mk.file.EXPECT().Delete(ctx, mock.Anything).Return(resp, nil)
	err := cl.DeleteFile(ctx, &entities.File{})

	assert.NoError(t, err)
}

func TestDeleteFile_WhenFileClientReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	errExp := errors.New("err")

	mk.file.EXPECT().Delete(ctx, mock.Anything).Return(nil, errExp)
	err := cl.DeleteFile(ctx, &entities.File{})

	assert.Error(t, err)
	assert.EqualError(t, errExp, err.Error())
}

func TestDeleteFile_WhenFileClientRespReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	errExp := &pb.Error{
		Message: "err",
	}
	resp := &pb.DeleteFileResponse{
		Error: errExp,
	}

	mk.file.EXPECT().Delete(ctx, mock.Anything).Return(resp, nil)
	err := cl.DeleteFile(ctx, &entities.File{})

	assert.Error(t, err)
	assert.Equal(t, errExp.String(), err.Error())
}

type fileAPIClientMock struct {
	file *mocks.FileServiceClient
}

func newFileAPIClientMock() (*fileAPIClientMock, *fileAPIClient) {
	m := &fileAPIClientMock{
		file: &mocks.FileServiceClient{},
	}
	return m, NewFileAPIClient(m.file)
}
