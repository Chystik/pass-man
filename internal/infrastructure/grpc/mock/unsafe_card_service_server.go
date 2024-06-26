// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeCardServiceServer is an autogenerated mock type for the UnsafeCardServiceServer type
type UnsafeCardServiceServer struct {
	mock.Mock
}

type UnsafeCardServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *UnsafeCardServiceServer) EXPECT() *UnsafeCardServiceServer_Expecter {
	return &UnsafeCardServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedCardServiceServer provides a mock function with given fields:
func (_m *UnsafeCardServiceServer) mustEmbedUnimplementedCardServiceServer() {
	_m.Called()
}

// UnsafeCardServiceServer_mustEmbedUnimplementedCardServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedCardServiceServer'
type UnsafeCardServiceServer_mustEmbedUnimplementedCardServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedCardServiceServer is a helper method to define mock.On call
func (_e *UnsafeCardServiceServer_Expecter) mustEmbedUnimplementedCardServiceServer() *UnsafeCardServiceServer_mustEmbedUnimplementedCardServiceServer_Call {
	return &UnsafeCardServiceServer_mustEmbedUnimplementedCardServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedCardServiceServer")}
}

func (_c *UnsafeCardServiceServer_mustEmbedUnimplementedCardServiceServer_Call) Run(run func()) *UnsafeCardServiceServer_mustEmbedUnimplementedCardServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UnsafeCardServiceServer_mustEmbedUnimplementedCardServiceServer_Call) Return() *UnsafeCardServiceServer_mustEmbedUnimplementedCardServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *UnsafeCardServiceServer_mustEmbedUnimplementedCardServiceServer_Call) RunAndReturn(run func()) *UnsafeCardServiceServer_mustEmbedUnimplementedCardServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewUnsafeCardServiceServer creates a new instance of UnsafeCardServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeCardServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeCardServiceServer {
	mock := &UnsafeCardServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
