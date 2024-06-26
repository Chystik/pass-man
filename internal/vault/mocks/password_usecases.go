// Code generated by mockery v2.23.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/Chystik/pass-man/internal/vault/password/entities"
	mock "github.com/stretchr/testify/mock"
)

// PasswordUsecases is an autogenerated mock type for the PasswordUsecases type
type PasswordUsecases struct {
	mock.Mock
}

type PasswordUsecases_Expecter struct {
	mock *mock.Mock
}

func (_m *PasswordUsecases) EXPECT() *PasswordUsecases_Expecter {
	return &PasswordUsecases_Expecter{mock: &_m.Mock}
}

// AddPassword provides a mock function with given fields: ctx, userID, password
func (_m *PasswordUsecases) AddPassword(ctx context.Context, userID string, password entities.Password) error {
	ret := _m.Called(ctx, userID, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entities.Password) error); ok {
		r0 = rf(ctx, userID, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PasswordUsecases_AddPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddPassword'
type PasswordUsecases_AddPassword_Call struct {
	*mock.Call
}

// AddPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - password entities.Password
func (_e *PasswordUsecases_Expecter) AddPassword(ctx interface{}, userID interface{}, password interface{}) *PasswordUsecases_AddPassword_Call {
	return &PasswordUsecases_AddPassword_Call{Call: _e.mock.On("AddPassword", ctx, userID, password)}
}

func (_c *PasswordUsecases_AddPassword_Call) Run(run func(ctx context.Context, userID string, password entities.Password)) *PasswordUsecases_AddPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entities.Password))
	})
	return _c
}

func (_c *PasswordUsecases_AddPassword_Call) Return(_a0 error) *PasswordUsecases_AddPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PasswordUsecases_AddPassword_Call) RunAndReturn(run func(context.Context, string, entities.Password) error) *PasswordUsecases_AddPassword_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePassword provides a mock function with given fields: ctx, userID, meta
func (_m *PasswordUsecases) DeletePassword(ctx context.Context, userID string, meta string) error {
	ret := _m.Called(ctx, userID, meta)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userID, meta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PasswordUsecases_DeletePassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePassword'
type PasswordUsecases_DeletePassword_Call struct {
	*mock.Call
}

// DeletePassword is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - meta string
func (_e *PasswordUsecases_Expecter) DeletePassword(ctx interface{}, userID interface{}, meta interface{}) *PasswordUsecases_DeletePassword_Call {
	return &PasswordUsecases_DeletePassword_Call{Call: _e.mock.On("DeletePassword", ctx, userID, meta)}
}

func (_c *PasswordUsecases_DeletePassword_Call) Run(run func(ctx context.Context, userID string, meta string)) *PasswordUsecases_DeletePassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *PasswordUsecases_DeletePassword_Call) Return(_a0 error) *PasswordUsecases_DeletePassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PasswordUsecases_DeletePassword_Call) RunAndReturn(run func(context.Context, string, string) error) *PasswordUsecases_DeletePassword_Call {
	_c.Call.Return(run)
	return _c
}

// GetPassword provides a mock function with given fields: ctx, userID, meta
func (_m *PasswordUsecases) GetPassword(ctx context.Context, userID string, meta string) (entities.Password, error) {
	ret := _m.Called(ctx, userID, meta)

	var r0 entities.Password
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (entities.Password, error)); ok {
		return rf(ctx, userID, meta)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) entities.Password); ok {
		r0 = rf(ctx, userID, meta)
	} else {
		r0 = ret.Get(0).(entities.Password)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PasswordUsecases_GetPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPassword'
type PasswordUsecases_GetPassword_Call struct {
	*mock.Call
}

// GetPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - meta string
func (_e *PasswordUsecases_Expecter) GetPassword(ctx interface{}, userID interface{}, meta interface{}) *PasswordUsecases_GetPassword_Call {
	return &PasswordUsecases_GetPassword_Call{Call: _e.mock.On("GetPassword", ctx, userID, meta)}
}

func (_c *PasswordUsecases_GetPassword_Call) Run(run func(ctx context.Context, userID string, meta string)) *PasswordUsecases_GetPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *PasswordUsecases_GetPassword_Call) Return(_a0 entities.Password, _a1 error) *PasswordUsecases_GetPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PasswordUsecases_GetPassword_Call) RunAndReturn(run func(context.Context, string, string) (entities.Password, error)) *PasswordUsecases_GetPassword_Call {
	_c.Call.Return(run)
	return _c
}

// ListPassword provides a mock function with given fields: ctx, userID
func (_m *PasswordUsecases) ListPassword(ctx context.Context, userID string) ([]entities.Password, error) {
	ret := _m.Called(ctx, userID)

	var r0 []entities.Password
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entities.Password, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entities.Password); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Password)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PasswordUsecases_ListPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPassword'
type PasswordUsecases_ListPassword_Call struct {
	*mock.Call
}

// ListPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *PasswordUsecases_Expecter) ListPassword(ctx interface{}, userID interface{}) *PasswordUsecases_ListPassword_Call {
	return &PasswordUsecases_ListPassword_Call{Call: _e.mock.On("ListPassword", ctx, userID)}
}

func (_c *PasswordUsecases_ListPassword_Call) Run(run func(ctx context.Context, userID string)) *PasswordUsecases_ListPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *PasswordUsecases_ListPassword_Call) Return(_a0 []entities.Password, _a1 error) *PasswordUsecases_ListPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PasswordUsecases_ListPassword_Call) RunAndReturn(run func(context.Context, string) ([]entities.Password, error)) *PasswordUsecases_ListPassword_Call {
	_c.Call.Return(run)
	return _c
}

// NewPasswordUsecases creates a new instance of PasswordUsecases. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPasswordUsecases(t interface {
	mock.TestingT
	Cleanup(func())
}) *PasswordUsecases {
	mock := &PasswordUsecases{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
