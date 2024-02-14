// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	mock "github.com/stretchr/testify/mock"
)

// CardServiceServer is an autogenerated mock type for the CardServiceServer type
type CardServiceServer struct {
	mock.Mock
}

type CardServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *CardServiceServer) EXPECT() *CardServiceServer_Expecter {
	return &CardServiceServer_Expecter{mock: &_m.Mock}
}

// AddCard provides a mock function with given fields: _a0, _a1
func (_m *CardServiceServer) AddCard(_a0 context.Context, _a1 *pb.AddCardRequest) (*pb.AddCardResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddCard")
	}

	var r0 *pb.AddCardResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.AddCardRequest) (*pb.AddCardResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.AddCardRequest) *pb.AddCardResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AddCardResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.AddCardRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CardServiceServer_AddCard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddCard'
type CardServiceServer_AddCard_Call struct {
	*mock.Call
}

// AddCard is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *pb.AddCardRequest
func (_e *CardServiceServer_Expecter) AddCard(_a0 interface{}, _a1 interface{}) *CardServiceServer_AddCard_Call {
	return &CardServiceServer_AddCard_Call{Call: _e.mock.On("AddCard", _a0, _a1)}
}

func (_c *CardServiceServer_AddCard_Call) Run(run func(_a0 context.Context, _a1 *pb.AddCardRequest)) *CardServiceServer_AddCard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pb.AddCardRequest))
	})
	return _c
}

func (_c *CardServiceServer_AddCard_Call) Return(_a0 *pb.AddCardResponse, _a1 error) *CardServiceServer_AddCard_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CardServiceServer_AddCard_Call) RunAndReturn(run func(context.Context, *pb.AddCardRequest) (*pb.AddCardResponse, error)) *CardServiceServer_AddCard_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCard provides a mock function with given fields: _a0, _a1
func (_m *CardServiceServer) DeleteCard(_a0 context.Context, _a1 *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCard")
	}

	var r0 *pb.DeleteCardResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteCardRequest) *pb.DeleteCardResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeleteCardResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeleteCardRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CardServiceServer_DeleteCard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCard'
type CardServiceServer_DeleteCard_Call struct {
	*mock.Call
}

// DeleteCard is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *pb.DeleteCardRequest
func (_e *CardServiceServer_Expecter) DeleteCard(_a0 interface{}, _a1 interface{}) *CardServiceServer_DeleteCard_Call {
	return &CardServiceServer_DeleteCard_Call{Call: _e.mock.On("DeleteCard", _a0, _a1)}
}

func (_c *CardServiceServer_DeleteCard_Call) Run(run func(_a0 context.Context, _a1 *pb.DeleteCardRequest)) *CardServiceServer_DeleteCard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pb.DeleteCardRequest))
	})
	return _c
}

func (_c *CardServiceServer_DeleteCard_Call) Return(_a0 *pb.DeleteCardResponse, _a1 error) *CardServiceServer_DeleteCard_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CardServiceServer_DeleteCard_Call) RunAndReturn(run func(context.Context, *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error)) *CardServiceServer_DeleteCard_Call {
	_c.Call.Return(run)
	return _c
}

// GetCard provides a mock function with given fields: _a0, _a1
func (_m *CardServiceServer) GetCard(_a0 context.Context, _a1 *pb.GetCardRequest) (*pb.GetCardResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetCard")
	}

	var r0 *pb.GetCardResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetCardRequest) (*pb.GetCardResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetCardRequest) *pb.GetCardResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetCardResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetCardRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CardServiceServer_GetCard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCard'
type CardServiceServer_GetCard_Call struct {
	*mock.Call
}

// GetCard is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *pb.GetCardRequest
func (_e *CardServiceServer_Expecter) GetCard(_a0 interface{}, _a1 interface{}) *CardServiceServer_GetCard_Call {
	return &CardServiceServer_GetCard_Call{Call: _e.mock.On("GetCard", _a0, _a1)}
}

func (_c *CardServiceServer_GetCard_Call) Run(run func(_a0 context.Context, _a1 *pb.GetCardRequest)) *CardServiceServer_GetCard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pb.GetCardRequest))
	})
	return _c
}

func (_c *CardServiceServer_GetCard_Call) Return(_a0 *pb.GetCardResponse, _a1 error) *CardServiceServer_GetCard_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CardServiceServer_GetCard_Call) RunAndReturn(run func(context.Context, *pb.GetCardRequest) (*pb.GetCardResponse, error)) *CardServiceServer_GetCard_Call {
	_c.Call.Return(run)
	return _c
}

// ListCard provides a mock function with given fields: _a0, _a1
func (_m *CardServiceServer) ListCard(_a0 context.Context, _a1 *pb.ListCardRequest) (*pb.ListCardResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ListCard")
	}

	var r0 *pb.ListCardResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListCardRequest) (*pb.ListCardResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListCardRequest) *pb.ListCardResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ListCardResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.ListCardRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CardServiceServer_ListCard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCard'
type CardServiceServer_ListCard_Call struct {
	*mock.Call
}

// ListCard is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *pb.ListCardRequest
func (_e *CardServiceServer_Expecter) ListCard(_a0 interface{}, _a1 interface{}) *CardServiceServer_ListCard_Call {
	return &CardServiceServer_ListCard_Call{Call: _e.mock.On("ListCard", _a0, _a1)}
}

func (_c *CardServiceServer_ListCard_Call) Run(run func(_a0 context.Context, _a1 *pb.ListCardRequest)) *CardServiceServer_ListCard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pb.ListCardRequest))
	})
	return _c
}

func (_c *CardServiceServer_ListCard_Call) Return(_a0 *pb.ListCardResponse, _a1 error) *CardServiceServer_ListCard_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CardServiceServer_ListCard_Call) RunAndReturn(run func(context.Context, *pb.ListCardRequest) (*pb.ListCardResponse, error)) *CardServiceServer_ListCard_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedCardServiceServer provides a mock function with given fields:
func (_m *CardServiceServer) mustEmbedUnimplementedCardServiceServer() {
	_m.Called()
}

// CardServiceServer_mustEmbedUnimplementedCardServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedCardServiceServer'
type CardServiceServer_mustEmbedUnimplementedCardServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedCardServiceServer is a helper method to define mock.On call
func (_e *CardServiceServer_Expecter) mustEmbedUnimplementedCardServiceServer() *CardServiceServer_mustEmbedUnimplementedCardServiceServer_Call {
	return &CardServiceServer_mustEmbedUnimplementedCardServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedCardServiceServer")}
}

func (_c *CardServiceServer_mustEmbedUnimplementedCardServiceServer_Call) Run(run func()) *CardServiceServer_mustEmbedUnimplementedCardServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CardServiceServer_mustEmbedUnimplementedCardServiceServer_Call) Return() *CardServiceServer_mustEmbedUnimplementedCardServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *CardServiceServer_mustEmbedUnimplementedCardServiceServer_Call) RunAndReturn(run func()) *CardServiceServer_mustEmbedUnimplementedCardServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewCardServiceServer creates a new instance of CardServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCardServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *CardServiceServer {
	mock := &CardServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
