// Code generated by mockery v2.23.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/Chystik/pass-man/internal/vault/password/entities"
	mock "github.com/stretchr/testify/mock"
)

// PasswordRepository is an autogenerated mock type for the PasswordRepository type
type PasswordRepository struct {
	mock.Mock
}

type PasswordRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *PasswordRepository) EXPECT() *PasswordRepository_Expecter {
	return &PasswordRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, userID, password
func (_m *PasswordRepository) Create(ctx context.Context, userID string, password entities.Password) error {
	ret := _m.Called(ctx, userID, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entities.Password) error); ok {
		r0 = rf(ctx, userID, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PasswordRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type PasswordRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - password entities.Password
func (_e *PasswordRepository_Expecter) Create(ctx interface{}, userID interface{}, password interface{}) *PasswordRepository_Create_Call {
	return &PasswordRepository_Create_Call{Call: _e.mock.On("Create", ctx, userID, password)}
}

func (_c *PasswordRepository_Create_Call) Run(run func(ctx context.Context, userID string, password entities.Password)) *PasswordRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entities.Password))
	})
	return _c
}

func (_c *PasswordRepository_Create_Call) Return(_a0 error) *PasswordRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PasswordRepository_Create_Call) RunAndReturn(run func(context.Context, string, entities.Password) error) *PasswordRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, userID, meta
func (_m *PasswordRepository) Delete(ctx context.Context, userID string, meta string) error {
	ret := _m.Called(ctx, userID, meta)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userID, meta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PasswordRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type PasswordRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - meta string
func (_e *PasswordRepository_Expecter) Delete(ctx interface{}, userID interface{}, meta interface{}) *PasswordRepository_Delete_Call {
	return &PasswordRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, userID, meta)}
}

func (_c *PasswordRepository_Delete_Call) Run(run func(ctx context.Context, userID string, meta string)) *PasswordRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *PasswordRepository_Delete_Call) Return(_a0 error) *PasswordRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PasswordRepository_Delete_Call) RunAndReturn(run func(context.Context, string, string) error) *PasswordRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: ctx, userID
func (_m *PasswordRepository) GetList(ctx context.Context, userID string) ([]entities.Password, error) {
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

// PasswordRepository_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type PasswordRepository_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *PasswordRepository_Expecter) GetList(ctx interface{}, userID interface{}) *PasswordRepository_GetList_Call {
	return &PasswordRepository_GetList_Call{Call: _e.mock.On("GetList", ctx, userID)}
}

func (_c *PasswordRepository_GetList_Call) Run(run func(ctx context.Context, userID string)) *PasswordRepository_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *PasswordRepository_GetList_Call) Return(_a0 []entities.Password, _a1 error) *PasswordRepository_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PasswordRepository_GetList_Call) RunAndReturn(run func(context.Context, string) ([]entities.Password, error)) *PasswordRepository_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// GetOne provides a mock function with given fields: ctx, userID, meta
func (_m *PasswordRepository) GetOne(ctx context.Context, userID string, meta string) (entities.Password, error) {
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

// PasswordRepository_GetOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOne'
type PasswordRepository_GetOne_Call struct {
	*mock.Call
}

// GetOne is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - meta string
func (_e *PasswordRepository_Expecter) GetOne(ctx interface{}, userID interface{}, meta interface{}) *PasswordRepository_GetOne_Call {
	return &PasswordRepository_GetOne_Call{Call: _e.mock.On("GetOne", ctx, userID, meta)}
}

func (_c *PasswordRepository_GetOne_Call) Run(run func(ctx context.Context, userID string, meta string)) *PasswordRepository_GetOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *PasswordRepository_GetOne_Call) Return(_a0 entities.Password, _a1 error) *PasswordRepository_GetOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PasswordRepository_GetOne_Call) RunAndReturn(run func(context.Context, string, string) (entities.Password, error)) *PasswordRepository_GetOne_Call {
	_c.Call.Return(run)
	return _c
}

// NewPasswordRepository creates a new instance of PasswordRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPasswordRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PasswordRepository {
	mock := &PasswordRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
