package adapters

import (
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

	mk.File.EXPECT().Upload(ctx, mock.Anything).Return(stream, nil)
	stream.EXPECT().Send(mock.Anything).Return(nil)
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

	mk.File.EXPECT().Upload(ctx, mock.Anything).Return(stream, errors.New("err"))
	err = cl.UploadFile(ctx, &entities.File{}, tmpFile.Name())

	assert.Error(t, err)
}

func TestUploadFile_WhenSendReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	stream := &mocks.FileService_UploadClient{}
	tmpFile, err := os.CreateTemp("", "")
	require.NoError(t, err)

	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write([]byte("temporary file test data"))
	require.NoError(t, err)

	mk.File.EXPECT().Upload(ctx, mock.Anything).Return(stream, nil)
	stream.EXPECT().Send(mock.Anything).Return(errors.New("err"))
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

	mk.File.EXPECT().Upload(ctx, mock.Anything).Return(stream, nil)
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

	mk.File.EXPECT().Upload(ctx, mock.Anything).Return(stream, nil)
	stream.EXPECT().Send(mock.Anything).Return(nil)
	stream.EXPECT().CloseAndRecv().Return(&pb.UploadFileResponse{
		Error: &pb.Error{},
	}, nil)
	err = cl.UploadFile(ctx, &entities.File{}, tmpFile.Name())

	assert.Error(t, err)
}

func TestDownloadFile_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mk, cl := newFileAPIClientMock()
	ctx := context.Background()
	stream := &mocks.FileService_DownloadClient{}
	//treamResp := &mocks.FileServiceClient_Download_Call{}
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
			ChunkData: []byte("test file data"),
		},
	}

	mk.File.EXPECT().Download(ctx, mock.Anything).Return(stream, nil)
	stream.EXPECT().Recv().Return(resp, nil).Once()
	stream.EXPECT().Recv().Return(respChunks, nil).Once()
	stream.EXPECT().Recv().Return(nil, io.EOF).Once()
	stream.EXPECT().CloseSend().Return(nil)

	err = cl.DownloadFile(ctx, &entities.File{Meta: "test"}, tmpFile.Name())

	assert.NoError(t, err)
}

type fileAPIClientMock struct {
	File *mocks.FileServiceClient
}

func newFileAPIClientMock() (*fileAPIClientMock, *fileAPIClient) {
	m := &fileAPIClientMock{
		File: &mocks.FileServiceClient{},
	}
	return m, NewFileAPIClient(m.File)
}
