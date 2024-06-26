// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
)

// PasswordServiceClient is an autogenerated mock type for the PasswordServiceClient type
type PasswordServiceClient struct {
	mock.Mock
}

type PasswordServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *PasswordServiceClient) EXPECT() *PasswordServiceClient_Expecter {
	return &PasswordServiceClient_Expecter{mock: &_m.Mock}
}

// AddPassword provides a mock function with given fields: ctx, in, opts
func (_m *PasswordServiceClient) AddPassword(ctx context.Context, in *pb.AddPasswordRequest, opts ...grpc.CallOption) (*pb.AddPasswordResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddPassword")
	}

	var r0 *pb.AddPasswordResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.AddPasswordRequest, ...grpc.CallOption) (*pb.AddPasswordResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.AddPasswordRequest, ...grpc.CallOption) *pb.AddPasswordResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AddPasswordResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.AddPasswordRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PasswordServiceClient_AddPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddPassword'
type PasswordServiceClient_AddPassword_Call struct {
	*mock.Call
}

// AddPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - in *pb.AddPasswordRequest
//   - opts ...grpc.CallOption
func (_e *PasswordServiceClient_Expecter) AddPassword(ctx interface{}, in interface{}, opts ...interface{}) *PasswordServiceClient_AddPassword_Call {
	return &PasswordServiceClient_AddPassword_Call{Call: _e.mock.On("AddPassword",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *PasswordServiceClient_AddPassword_Call) Run(run func(ctx context.Context, in *pb.AddPasswordRequest, opts ...grpc.CallOption)) *PasswordServiceClient_AddPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*pb.AddPasswordRequest), variadicArgs...)
	})
	return _c
}

func (_c *PasswordServiceClient_AddPassword_Call) Return(_a0 *pb.AddPasswordResponse, _a1 error) *PasswordServiceClient_AddPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PasswordServiceClient_AddPassword_Call) RunAndReturn(run func(context.Context, *pb.AddPasswordRequest, ...grpc.CallOption) (*pb.AddPasswordResponse, error)) *PasswordServiceClient_AddPassword_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePassword provides a mock function with given fields: ctx, in, opts
func (_m *PasswordServiceClient) DeletePassword(ctx context.Context, in *pb.DeletePasswordRequest, opts ...grpc.CallOption) (*pb.DeletePasswordResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePassword")
	}

	var r0 *pb.DeletePasswordResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeletePasswordRequest, ...grpc.CallOption) (*pb.DeletePasswordResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeletePasswordRequest, ...grpc.CallOption) *pb.DeletePasswordResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeletePasswordResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeletePasswordRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PasswordServiceClient_DeletePassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePassword'
type PasswordServiceClient_DeletePassword_Call struct {
	*mock.Call
}

// DeletePassword is a helper method to define mock.On call
//   - ctx context.Context
//   - in *pb.DeletePasswordRequest
//   - opts ...grpc.CallOption
func (_e *PasswordServiceClient_Expecter) DeletePassword(ctx interface{}, in interface{}, opts ...interface{}) *PasswordServiceClient_DeletePassword_Call {
	return &PasswordServiceClient_DeletePassword_Call{Call: _e.mock.On("DeletePassword",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *PasswordServiceClient_DeletePassword_Call) Run(run func(ctx context.Context, in *pb.DeletePasswordRequest, opts ...grpc.CallOption)) *PasswordServiceClient_DeletePassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*pb.DeletePasswordRequest), variadicArgs...)
	})
	return _c
}

func (_c *PasswordServiceClient_DeletePassword_Call) Return(_a0 *pb.DeletePasswordResponse, _a1 error) *PasswordServiceClient_DeletePassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PasswordServiceClient_DeletePassword_Call) RunAndReturn(run func(context.Context, *pb.DeletePasswordRequest, ...grpc.CallOption) (*pb.DeletePasswordResponse, error)) *PasswordServiceClient_DeletePassword_Call {
	_c.Call.Return(run)
	return _c
}

// GetPassword provides a mock function with given fields: ctx, in, opts
func (_m *PasswordServiceClient) GetPassword(ctx context.Context, in *pb.GetPasswordRequest, opts ...grpc.CallOption) (*pb.GetPasswordResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPassword")
	}

	var r0 *pb.GetPasswordResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetPasswordRequest, ...grpc.CallOption) (*pb.GetPasswordResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetPasswordRequest, ...grpc.CallOption) *pb.GetPasswordResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetPasswordResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetPasswordRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PasswordServiceClient_GetPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPassword'
type PasswordServiceClient_GetPassword_Call struct {
	*mock.Call
}

// GetPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - in *pb.GetPasswordRequest
//   - opts ...grpc.CallOption
func (_e *PasswordServiceClient_Expecter) GetPassword(ctx interface{}, in interface{}, opts ...interface{}) *PasswordServiceClient_GetPassword_Call {
	return &PasswordServiceClient_GetPassword_Call{Call: _e.mock.On("GetPassword",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *PasswordServiceClient_GetPassword_Call) Run(run func(ctx context.Context, in *pb.GetPasswordRequest, opts ...grpc.CallOption)) *PasswordServiceClient_GetPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*pb.GetPasswordRequest), variadicArgs...)
	})
	return _c
}

func (_c *PasswordServiceClient_GetPassword_Call) Return(_a0 *pb.GetPasswordResponse, _a1 error) *PasswordServiceClient_GetPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PasswordServiceClient_GetPassword_Call) RunAndReturn(run func(context.Context, *pb.GetPasswordRequest, ...grpc.CallOption) (*pb.GetPasswordResponse, error)) *PasswordServiceClient_GetPassword_Call {
	_c.Call.Return(run)
	return _c
}

// ListPassword provides a mock function with given fields: ctx, in, opts
func (_m *PasswordServiceClient) ListPassword(ctx context.Context, in *pb.ListPasswordRequest, opts ...grpc.CallOption) (*pb.ListPasswordResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPassword")
	}

	var r0 *pb.ListPasswordResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListPasswordRequest, ...grpc.CallOption) (*pb.ListPasswordResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListPasswordRequest, ...grpc.CallOption) *pb.ListPasswordResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ListPasswordResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.ListPasswordRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PasswordServiceClient_ListPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPassword'
type PasswordServiceClient_ListPassword_Call struct {
	*mock.Call
}

// ListPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - in *pb.ListPasswordRequest
//   - opts ...grpc.CallOption
func (_e *PasswordServiceClient_Expecter) ListPassword(ctx interface{}, in interface{}, opts ...interface{}) *PasswordServiceClient_ListPassword_Call {
	return &PasswordServiceClient_ListPassword_Call{Call: _e.mock.On("ListPassword",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *PasswordServiceClient_ListPassword_Call) Run(run func(ctx context.Context, in *pb.ListPasswordRequest, opts ...grpc.CallOption)) *PasswordServiceClient_ListPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*pb.ListPasswordRequest), variadicArgs...)
	})
	return _c
}

func (_c *PasswordServiceClient_ListPassword_Call) Return(_a0 *pb.ListPasswordResponse, _a1 error) *PasswordServiceClient_ListPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PasswordServiceClient_ListPassword_Call) RunAndReturn(run func(context.Context, *pb.ListPasswordRequest, ...grpc.CallOption) (*pb.ListPasswordResponse, error)) *PasswordServiceClient_ListPassword_Call {
	_c.Call.Return(run)
	return _c
}

// NewPasswordServiceClient creates a new instance of PasswordServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPasswordServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *PasswordServiceClient {
	mock := &PasswordServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
