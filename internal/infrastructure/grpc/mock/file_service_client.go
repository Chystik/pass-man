// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
)

// FileServiceClient is an autogenerated mock type for the FileServiceClient type
type FileServiceClient struct {
	mock.Mock
}

type FileServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *FileServiceClient) EXPECT() *FileServiceClient_Expecter {
	return &FileServiceClient_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, in, opts
func (_m *FileServiceClient) Delete(ctx context.Context, in *pb.DeleteFileRequest, opts ...grpc.CallOption) (*pb.DeleteFileResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *pb.DeleteFileResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteFileRequest, ...grpc.CallOption) (*pb.DeleteFileResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteFileRequest, ...grpc.CallOption) *pb.DeleteFileResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeleteFileResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeleteFileRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileServiceClient_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type FileServiceClient_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - in *pb.DeleteFileRequest
//   - opts ...grpc.CallOption
func (_e *FileServiceClient_Expecter) Delete(ctx interface{}, in interface{}, opts ...interface{}) *FileServiceClient_Delete_Call {
	return &FileServiceClient_Delete_Call{Call: _e.mock.On("Delete",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *FileServiceClient_Delete_Call) Run(run func(ctx context.Context, in *pb.DeleteFileRequest, opts ...grpc.CallOption)) *FileServiceClient_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*pb.DeleteFileRequest), variadicArgs...)
	})
	return _c
}

func (_c *FileServiceClient_Delete_Call) Return(_a0 *pb.DeleteFileResponse, _a1 error) *FileServiceClient_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileServiceClient_Delete_Call) RunAndReturn(run func(context.Context, *pb.DeleteFileRequest, ...grpc.CallOption) (*pb.DeleteFileResponse, error)) *FileServiceClient_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Download provides a mock function with given fields: ctx, in, opts
func (_m *FileServiceClient) Download(ctx context.Context, in *pb.DownloadFileRequest, opts ...grpc.CallOption) (pb.FileService_DownloadClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Download")
	}

	var r0 pb.FileService_DownloadClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DownloadFileRequest, ...grpc.CallOption) (pb.FileService_DownloadClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DownloadFileRequest, ...grpc.CallOption) pb.FileService_DownloadClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pb.FileService_DownloadClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.DownloadFileRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileServiceClient_Download_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Download'
type FileServiceClient_Download_Call struct {
	*mock.Call
}

// Download is a helper method to define mock.On call
//   - ctx context.Context
//   - in *pb.DownloadFileRequest
//   - opts ...grpc.CallOption
func (_e *FileServiceClient_Expecter) Download(ctx interface{}, in interface{}, opts ...interface{}) *FileServiceClient_Download_Call {
	return &FileServiceClient_Download_Call{Call: _e.mock.On("Download",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *FileServiceClient_Download_Call) Run(run func(ctx context.Context, in *pb.DownloadFileRequest, opts ...grpc.CallOption)) *FileServiceClient_Download_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*pb.DownloadFileRequest), variadicArgs...)
	})
	return _c
}

func (_c *FileServiceClient_Download_Call) Return(_a0 pb.FileService_DownloadClient, _a1 error) *FileServiceClient_Download_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileServiceClient_Download_Call) RunAndReturn(run func(context.Context, *pb.DownloadFileRequest, ...grpc.CallOption) (pb.FileService_DownloadClient, error)) *FileServiceClient_Download_Call {
	_c.Call.Return(run)
	return _c
}

// ListFiles provides a mock function with given fields: ctx, in, opts
func (_m *FileServiceClient) ListFiles(ctx context.Context, in *pb.ListFileRequest, opts ...grpc.CallOption) (*pb.ListFileResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListFiles")
	}

	var r0 *pb.ListFileResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListFileRequest, ...grpc.CallOption) (*pb.ListFileResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListFileRequest, ...grpc.CallOption) *pb.ListFileResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ListFileResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.ListFileRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileServiceClient_ListFiles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListFiles'
type FileServiceClient_ListFiles_Call struct {
	*mock.Call
}

// ListFiles is a helper method to define mock.On call
//   - ctx context.Context
//   - in *pb.ListFileRequest
//   - opts ...grpc.CallOption
func (_e *FileServiceClient_Expecter) ListFiles(ctx interface{}, in interface{}, opts ...interface{}) *FileServiceClient_ListFiles_Call {
	return &FileServiceClient_ListFiles_Call{Call: _e.mock.On("ListFiles",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *FileServiceClient_ListFiles_Call) Run(run func(ctx context.Context, in *pb.ListFileRequest, opts ...grpc.CallOption)) *FileServiceClient_ListFiles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*pb.ListFileRequest), variadicArgs...)
	})
	return _c
}

func (_c *FileServiceClient_ListFiles_Call) Return(_a0 *pb.ListFileResponse, _a1 error) *FileServiceClient_ListFiles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileServiceClient_ListFiles_Call) RunAndReturn(run func(context.Context, *pb.ListFileRequest, ...grpc.CallOption) (*pb.ListFileResponse, error)) *FileServiceClient_ListFiles_Call {
	_c.Call.Return(run)
	return _c
}

// Upload provides a mock function with given fields: ctx, opts
func (_m *FileServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (pb.FileService_UploadClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Upload")
	}

	var r0 pb.FileService_UploadClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...grpc.CallOption) (pb.FileService_UploadClient, error)); ok {
		return rf(ctx, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...grpc.CallOption) pb.FileService_UploadClient); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pb.FileService_UploadClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileServiceClient_Upload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upload'
type FileServiceClient_Upload_Call struct {
	*mock.Call
}

// Upload is a helper method to define mock.On call
//   - ctx context.Context
//   - opts ...grpc.CallOption
func (_e *FileServiceClient_Expecter) Upload(ctx interface{}, opts ...interface{}) *FileServiceClient_Upload_Call {
	return &FileServiceClient_Upload_Call{Call: _e.mock.On("Upload",
		append([]interface{}{ctx}, opts...)...)}
}

func (_c *FileServiceClient_Upload_Call) Run(run func(ctx context.Context, opts ...grpc.CallOption)) *FileServiceClient_Upload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *FileServiceClient_Upload_Call) Return(_a0 pb.FileService_UploadClient, _a1 error) *FileServiceClient_Upload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileServiceClient_Upload_Call) RunAndReturn(run func(context.Context, ...grpc.CallOption) (pb.FileService_UploadClient, error)) *FileServiceClient_Upload_Call {
	_c.Call.Return(run)
	return _c
}

// NewFileServiceClient creates a new instance of FileServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFileServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *FileServiceClient {
	mock := &FileServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
