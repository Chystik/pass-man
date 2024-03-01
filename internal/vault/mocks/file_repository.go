// Code generated by mockery v2.23.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/Chystik/pass-man/internal/vault/file/entities"
	mock "github.com/stretchr/testify/mock"
)

// FileRepository is an autogenerated mock type for the FileRepository type
type FileRepository struct {
	mock.Mock
}

type FileRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *FileRepository) EXPECT() *FileRepository_Expecter {
	return &FileRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, userID, file
func (_m *FileRepository) Create(ctx context.Context, userID string, file *entities.File) (int, error) {
	ret := _m.Called(ctx, userID, file)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *entities.File) (int, error)); ok {
		return rf(ctx, userID, file)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *entities.File) int); ok {
		r0 = rf(ctx, userID, file)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *entities.File) error); ok {
		r1 = rf(ctx, userID, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type FileRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - file *entities.File
func (_e *FileRepository_Expecter) Create(ctx interface{}, userID interface{}, file interface{}) *FileRepository_Create_Call {
	return &FileRepository_Create_Call{Call: _e.mock.On("Create", ctx, userID, file)}
}

func (_c *FileRepository_Create_Call) Run(run func(ctx context.Context, userID string, file *entities.File)) *FileRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*entities.File))
	})
	return _c
}

func (_c *FileRepository_Create_Call) Return(_a0 int, _a1 error) *FileRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileRepository_Create_Call) RunAndReturn(run func(context.Context, string, *entities.File) (int, error)) *FileRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, userID, file
func (_m *FileRepository) Delete(ctx context.Context, userID string, file *entities.File) error {
	ret := _m.Called(ctx, userID, file)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *entities.File) error); ok {
		r0 = rf(ctx, userID, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type FileRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - file *entities.File
func (_e *FileRepository_Expecter) Delete(ctx interface{}, userID interface{}, file interface{}) *FileRepository_Delete_Call {
	return &FileRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, userID, file)}
}

func (_c *FileRepository_Delete_Call) Run(run func(ctx context.Context, userID string, file *entities.File)) *FileRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*entities.File))
	})
	return _c
}

func (_c *FileRepository_Delete_Call) Return(_a0 error) *FileRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileRepository_Delete_Call) RunAndReturn(run func(context.Context, string, *entities.File) error) *FileRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: ctx, userID
func (_m *FileRepository) GetList(ctx context.Context, userID string) ([]*entities.File, error) {
	ret := _m.Called(ctx, userID)

	var r0 []*entities.File
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*entities.File, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entities.File); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.File)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileRepository_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type FileRepository_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *FileRepository_Expecter) GetList(ctx interface{}, userID interface{}) *FileRepository_GetList_Call {
	return &FileRepository_GetList_Call{Call: _e.mock.On("GetList", ctx, userID)}
}

func (_c *FileRepository_GetList_Call) Run(run func(ctx context.Context, userID string)) *FileRepository_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FileRepository_GetList_Call) Return(_a0 []*entities.File, _a1 error) *FileRepository_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileRepository_GetList_Call) RunAndReturn(run func(context.Context, string) ([]*entities.File, error)) *FileRepository_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// GetOne provides a mock function with given fields: ctx, userID, file
func (_m *FileRepository) GetOne(ctx context.Context, userID string, file *entities.File) (int, error) {
	ret := _m.Called(ctx, userID, file)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *entities.File) (int, error)); ok {
		return rf(ctx, userID, file)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *entities.File) int); ok {
		r0 = rf(ctx, userID, file)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *entities.File) error); ok {
		r1 = rf(ctx, userID, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileRepository_GetOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOne'
type FileRepository_GetOne_Call struct {
	*mock.Call
}

// GetOne is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - file *entities.File
func (_e *FileRepository_Expecter) GetOne(ctx interface{}, userID interface{}, file interface{}) *FileRepository_GetOne_Call {
	return &FileRepository_GetOne_Call{Call: _e.mock.On("GetOne", ctx, userID, file)}
}

func (_c *FileRepository_GetOne_Call) Run(run func(ctx context.Context, userID string, file *entities.File)) *FileRepository_GetOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*entities.File))
	})
	return _c
}

func (_c *FileRepository_GetOne_Call) Return(_a0 int, _a1 error) *FileRepository_GetOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileRepository_GetOne_Call) RunAndReturn(run func(context.Context, string, *entities.File) (int, error)) *FileRepository_GetOne_Call {
	_c.Call.Return(run)
	return _c
}

// NewFileRepository creates a new instance of FileRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFileRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *FileRepository {
	mock := &FileRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
