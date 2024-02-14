// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeNoteServiceServer is an autogenerated mock type for the UnsafeNoteServiceServer type
type UnsafeNoteServiceServer struct {
	mock.Mock
}

type UnsafeNoteServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *UnsafeNoteServiceServer) EXPECT() *UnsafeNoteServiceServer_Expecter {
	return &UnsafeNoteServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedNoteServiceServer provides a mock function with given fields:
func (_m *UnsafeNoteServiceServer) mustEmbedUnimplementedNoteServiceServer() {
	_m.Called()
}

// UnsafeNoteServiceServer_mustEmbedUnimplementedNoteServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedNoteServiceServer'
type UnsafeNoteServiceServer_mustEmbedUnimplementedNoteServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedNoteServiceServer is a helper method to define mock.On call
func (_e *UnsafeNoteServiceServer_Expecter) mustEmbedUnimplementedNoteServiceServer() *UnsafeNoteServiceServer_mustEmbedUnimplementedNoteServiceServer_Call {
	return &UnsafeNoteServiceServer_mustEmbedUnimplementedNoteServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedNoteServiceServer")}
}

func (_c *UnsafeNoteServiceServer_mustEmbedUnimplementedNoteServiceServer_Call) Run(run func()) *UnsafeNoteServiceServer_mustEmbedUnimplementedNoteServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UnsafeNoteServiceServer_mustEmbedUnimplementedNoteServiceServer_Call) Return() *UnsafeNoteServiceServer_mustEmbedUnimplementedNoteServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *UnsafeNoteServiceServer_mustEmbedUnimplementedNoteServiceServer_Call) RunAndReturn(run func()) *UnsafeNoteServiceServer_mustEmbedUnimplementedNoteServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewUnsafeNoteServiceServer creates a new instance of UnsafeNoteServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeNoteServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeNoteServiceServer {
	mock := &UnsafeNoteServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}