// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeUserServiceServer is an autogenerated mock type for the UnsafeUserServiceServer type
type UnsafeUserServiceServer struct {
	mock.Mock
}

type UnsafeUserServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *UnsafeUserServiceServer) EXPECT() *UnsafeUserServiceServer_Expecter {
	return &UnsafeUserServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedUserServiceServer provides a mock function with given fields:
func (_m *UnsafeUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	_m.Called()
}

// UnsafeUserServiceServer_mustEmbedUnimplementedUserServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedUserServiceServer'
type UnsafeUserServiceServer_mustEmbedUnimplementedUserServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedUserServiceServer is a helper method to define mock.On call
func (_e *UnsafeUserServiceServer_Expecter) mustEmbedUnimplementedUserServiceServer() *UnsafeUserServiceServer_mustEmbedUnimplementedUserServiceServer_Call {
	return &UnsafeUserServiceServer_mustEmbedUnimplementedUserServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedUserServiceServer")}
}

func (_c *UnsafeUserServiceServer_mustEmbedUnimplementedUserServiceServer_Call) Run(run func()) *UnsafeUserServiceServer_mustEmbedUnimplementedUserServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UnsafeUserServiceServer_mustEmbedUnimplementedUserServiceServer_Call) Return() *UnsafeUserServiceServer_mustEmbedUnimplementedUserServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *UnsafeUserServiceServer_mustEmbedUnimplementedUserServiceServer_Call) RunAndReturn(run func()) *UnsafeUserServiceServer_mustEmbedUnimplementedUserServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewUnsafeUserServiceServer creates a new instance of UnsafeUserServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeUserServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeUserServiceServer {
	mock := &UnsafeUserServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
