// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isUploadFileRequest_Data is an autogenerated mock type for the isUploadFileRequest_Data type
type isUploadFileRequest_Data struct {
	mock.Mock
}

type isUploadFileRequest_Data_Expecter struct {
	mock *mock.Mock
}

func (_m *isUploadFileRequest_Data) EXPECT() *isUploadFileRequest_Data_Expecter {
	return &isUploadFileRequest_Data_Expecter{mock: &_m.Mock}
}

// isUploadFileRequest_Data provides a mock function with given fields:
func (_m *isUploadFileRequest_Data) isUploadFileRequest_Data() {
	_m.Called()
}

// isUploadFileRequest_Data_isUploadFileRequest_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isUploadFileRequest_Data'
type isUploadFileRequest_Data_isUploadFileRequest_Data_Call struct {
	*mock.Call
}

// isUploadFileRequest_Data is a helper method to define mock.On call
func (_e *isUploadFileRequest_Data_Expecter) isUploadFileRequest_Data() *isUploadFileRequest_Data_isUploadFileRequest_Data_Call {
	return &isUploadFileRequest_Data_isUploadFileRequest_Data_Call{Call: _e.mock.On("isUploadFileRequest_Data")}
}

func (_c *isUploadFileRequest_Data_isUploadFileRequest_Data_Call) Run(run func()) *isUploadFileRequest_Data_isUploadFileRequest_Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *isUploadFileRequest_Data_isUploadFileRequest_Data_Call) Return() *isUploadFileRequest_Data_isUploadFileRequest_Data_Call {
	_c.Call.Return()
	return _c
}

func (_c *isUploadFileRequest_Data_isUploadFileRequest_Data_Call) RunAndReturn(run func()) *isUploadFileRequest_Data_isUploadFileRequest_Data_Call {
	_c.Call.Return(run)
	return _c
}

// newIsUploadFileRequest_Data creates a new instance of isUploadFileRequest_Data. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newIsUploadFileRequest_Data(t interface {
	mock.TestingT
	Cleanup(func())
}) *isUploadFileRequest_Data {
	mock := &isUploadFileRequest_Data{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
