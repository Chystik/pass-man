// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
)

// FileService_UploadClient is an autogenerated mock type for the FileService_UploadClient type
type FileService_UploadClient struct {
	mock.Mock
}

type FileService_UploadClient_Expecter struct {
	mock *mock.Mock
}

func (_m *FileService_UploadClient) EXPECT() *FileService_UploadClient_Expecter {
	return &FileService_UploadClient_Expecter{mock: &_m.Mock}
}

// CloseAndRecv provides a mock function with given fields:
func (_m *FileService_UploadClient) CloseAndRecv() (*pb.UploadFileResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CloseAndRecv")
	}

	var r0 *pb.UploadFileResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*pb.UploadFileResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *pb.UploadFileResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.UploadFileResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileService_UploadClient_CloseAndRecv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseAndRecv'
type FileService_UploadClient_CloseAndRecv_Call struct {
	*mock.Call
}

// CloseAndRecv is a helper method to define mock.On call
func (_e *FileService_UploadClient_Expecter) CloseAndRecv() *FileService_UploadClient_CloseAndRecv_Call {
	return &FileService_UploadClient_CloseAndRecv_Call{Call: _e.mock.On("CloseAndRecv")}
}

func (_c *FileService_UploadClient_CloseAndRecv_Call) Run(run func()) *FileService_UploadClient_CloseAndRecv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FileService_UploadClient_CloseAndRecv_Call) Return(_a0 *pb.UploadFileResponse, _a1 error) *FileService_UploadClient_CloseAndRecv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileService_UploadClient_CloseAndRecv_Call) RunAndReturn(run func() (*pb.UploadFileResponse, error)) *FileService_UploadClient_CloseAndRecv_Call {
	_c.Call.Return(run)
	return _c
}

// CloseSend provides a mock function with given fields:
func (_m *FileService_UploadClient) CloseSend() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CloseSend")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileService_UploadClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type FileService_UploadClient_CloseSend_Call struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *FileService_UploadClient_Expecter) CloseSend() *FileService_UploadClient_CloseSend_Call {
	return &FileService_UploadClient_CloseSend_Call{Call: _e.mock.On("CloseSend")}
}

func (_c *FileService_UploadClient_CloseSend_Call) Run(run func()) *FileService_UploadClient_CloseSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FileService_UploadClient_CloseSend_Call) Return(_a0 error) *FileService_UploadClient_CloseSend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileService_UploadClient_CloseSend_Call) RunAndReturn(run func() error) *FileService_UploadClient_CloseSend_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with given fields:
func (_m *FileService_UploadClient) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// FileService_UploadClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type FileService_UploadClient_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *FileService_UploadClient_Expecter) Context() *FileService_UploadClient_Context_Call {
	return &FileService_UploadClient_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *FileService_UploadClient_Context_Call) Run(run func()) *FileService_UploadClient_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FileService_UploadClient_Context_Call) Return(_a0 context.Context) *FileService_UploadClient_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileService_UploadClient_Context_Call) RunAndReturn(run func() context.Context) *FileService_UploadClient_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with given fields:
func (_m *FileService_UploadClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Header")
	}

	var r0 metadata.MD
	var r1 error
	if rf, ok := ret.Get(0).(func() (metadata.MD, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileService_UploadClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type FileService_UploadClient_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *FileService_UploadClient_Expecter) Header() *FileService_UploadClient_Header_Call {
	return &FileService_UploadClient_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *FileService_UploadClient_Header_Call) Run(run func()) *FileService_UploadClient_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FileService_UploadClient_Header_Call) Return(_a0 metadata.MD, _a1 error) *FileService_UploadClient_Header_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileService_UploadClient_Header_Call) RunAndReturn(run func() (metadata.MD, error)) *FileService_UploadClient_Header_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *FileService_UploadClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for RecvMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileService_UploadClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type FileService_UploadClient_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *FileService_UploadClient_Expecter) RecvMsg(m interface{}) *FileService_UploadClient_RecvMsg_Call {
	return &FileService_UploadClient_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *FileService_UploadClient_RecvMsg_Call) Run(run func(m interface{})) *FileService_UploadClient_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *FileService_UploadClient_RecvMsg_Call) Return(_a0 error) *FileService_UploadClient_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileService_UploadClient_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *FileService_UploadClient_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *FileService_UploadClient) Send(_a0 *pb.UploadFileRequest) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*pb.UploadFileRequest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileService_UploadClient_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type FileService_UploadClient_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - _a0 *pb.UploadFileRequest
func (_e *FileService_UploadClient_Expecter) Send(_a0 interface{}) *FileService_UploadClient_Send_Call {
	return &FileService_UploadClient_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *FileService_UploadClient_Send_Call) Run(run func(_a0 *pb.UploadFileRequest)) *FileService_UploadClient_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*pb.UploadFileRequest))
	})
	return _c
}

func (_c *FileService_UploadClient_Send_Call) Return(_a0 error) *FileService_UploadClient_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileService_UploadClient_Send_Call) RunAndReturn(run func(*pb.UploadFileRequest) error) *FileService_UploadClient_Send_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *FileService_UploadClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for SendMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileService_UploadClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type FileService_UploadClient_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *FileService_UploadClient_Expecter) SendMsg(m interface{}) *FileService_UploadClient_SendMsg_Call {
	return &FileService_UploadClient_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *FileService_UploadClient_SendMsg_Call) Run(run func(m interface{})) *FileService_UploadClient_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *FileService_UploadClient_SendMsg_Call) Return(_a0 error) *FileService_UploadClient_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileService_UploadClient_SendMsg_Call) RunAndReturn(run func(interface{}) error) *FileService_UploadClient_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with given fields:
func (_m *FileService_UploadClient) Trailer() metadata.MD {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Trailer")
	}

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// FileService_UploadClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type FileService_UploadClient_Trailer_Call struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *FileService_UploadClient_Expecter) Trailer() *FileService_UploadClient_Trailer_Call {
	return &FileService_UploadClient_Trailer_Call{Call: _e.mock.On("Trailer")}
}

func (_c *FileService_UploadClient_Trailer_Call) Run(run func()) *FileService_UploadClient_Trailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FileService_UploadClient_Trailer_Call) Return(_a0 metadata.MD) *FileService_UploadClient_Trailer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileService_UploadClient_Trailer_Call) RunAndReturn(run func() metadata.MD) *FileService_UploadClient_Trailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewFileService_UploadClient creates a new instance of FileService_UploadClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFileService_UploadClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *FileService_UploadClient {
	mock := &FileService_UploadClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
