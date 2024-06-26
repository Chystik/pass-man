// Code generated by mockery v2.23.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/Chystik/pass-man/internal/vault/file/entities"
	mock "github.com/stretchr/testify/mock"
)

// FileUsecases is an autogenerated mock type for the FileUsecases type
type FileUsecases struct {
	mock.Mock
}

type FileUsecases_Expecter struct {
	mock *mock.Mock
}

func (_m *FileUsecases) EXPECT() *FileUsecases_Expecter {
	return &FileUsecases_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, userID, file
func (_m *FileUsecases) Delete(ctx context.Context, userID string, file *entities.File) error {
	ret := _m.Called(ctx, userID, file)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *entities.File) error); ok {
		r0 = rf(ctx, userID, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileUsecases_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type FileUsecases_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - file *entities.File
func (_e *FileUsecases_Expecter) Delete(ctx interface{}, userID interface{}, file interface{}) *FileUsecases_Delete_Call {
	return &FileUsecases_Delete_Call{Call: _e.mock.On("Delete", ctx, userID, file)}
}

func (_c *FileUsecases_Delete_Call) Run(run func(ctx context.Context, userID string, file *entities.File)) *FileUsecases_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*entities.File))
	})
	return _c
}

func (_c *FileUsecases_Delete_Call) Return(_a0 error) *FileUsecases_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FileUsecases_Delete_Call) RunAndReturn(run func(context.Context, string, *entities.File) error) *FileUsecases_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Download provides a mock function with given fields: ctx, userID, file
func (_m *FileUsecases) Download(ctx context.Context, userID string, file *entities.File) (int, error) {
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

// FileUsecases_Download_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Download'
type FileUsecases_Download_Call struct {
	*mock.Call
}

// Download is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - file *entities.File
func (_e *FileUsecases_Expecter) Download(ctx interface{}, userID interface{}, file interface{}) *FileUsecases_Download_Call {
	return &FileUsecases_Download_Call{Call: _e.mock.On("Download", ctx, userID, file)}
}

func (_c *FileUsecases_Download_Call) Run(run func(ctx context.Context, userID string, file *entities.File)) *FileUsecases_Download_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*entities.File))
	})
	return _c
}

func (_c *FileUsecases_Download_Call) Return(_a0 int, _a1 error) *FileUsecases_Download_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileUsecases_Download_Call) RunAndReturn(run func(context.Context, string, *entities.File) (int, error)) *FileUsecases_Download_Call {
	_c.Call.Return(run)
	return _c
}

// ListFiles provides a mock function with given fields: ctx, userID
func (_m *FileUsecases) ListFiles(ctx context.Context, userID string) ([]*entities.File, error) {
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

// FileUsecases_ListFiles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListFiles'
type FileUsecases_ListFiles_Call struct {
	*mock.Call
}

// ListFiles is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *FileUsecases_Expecter) ListFiles(ctx interface{}, userID interface{}) *FileUsecases_ListFiles_Call {
	return &FileUsecases_ListFiles_Call{Call: _e.mock.On("ListFiles", ctx, userID)}
}

func (_c *FileUsecases_ListFiles_Call) Run(run func(ctx context.Context, userID string)) *FileUsecases_ListFiles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FileUsecases_ListFiles_Call) Return(_a0 []*entities.File, _a1 error) *FileUsecases_ListFiles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileUsecases_ListFiles_Call) RunAndReturn(run func(context.Context, string) ([]*entities.File, error)) *FileUsecases_ListFiles_Call {
	_c.Call.Return(run)
	return _c
}

// Upload provides a mock function with given fields: ctx, userID, file
func (_m *FileUsecases) Upload(ctx context.Context, userID string, file *entities.File) (int, error) {
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

// FileUsecases_Upload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upload'
type FileUsecases_Upload_Call struct {
	*mock.Call
}

// Upload is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - file *entities.File
func (_e *FileUsecases_Expecter) Upload(ctx interface{}, userID interface{}, file interface{}) *FileUsecases_Upload_Call {
	return &FileUsecases_Upload_Call{Call: _e.mock.On("Upload", ctx, userID, file)}
}

func (_c *FileUsecases_Upload_Call) Run(run func(ctx context.Context, userID string, file *entities.File)) *FileUsecases_Upload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*entities.File))
	})
	return _c
}

func (_c *FileUsecases_Upload_Call) Return(_a0 int, _a1 error) *FileUsecases_Upload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileUsecases_Upload_Call) RunAndReturn(run func(context.Context, string, *entities.File) (int, error)) *FileUsecases_Upload_Call {
	_c.Call.Return(run)
	return _c
}

// NewFileUsecases creates a new instance of FileUsecases. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFileUsecases(t interface {
	mock.TestingT
	Cleanup(func())
}) *FileUsecases {
	mock := &FileUsecases{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
