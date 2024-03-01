// Code generated by mockery v2.23.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/Chystik/pass-man/internal/vault/note/entities"
	mock "github.com/stretchr/testify/mock"
)

// NoteRepository is an autogenerated mock type for the NoteRepository type
type NoteRepository struct {
	mock.Mock
}

type NoteRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *NoteRepository) EXPECT() *NoteRepository_Expecter {
	return &NoteRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, userID, note
func (_m *NoteRepository) Create(ctx context.Context, userID string, note entities.Note) error {
	ret := _m.Called(ctx, userID, note)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entities.Note) error); ok {
		r0 = rf(ctx, userID, note)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NoteRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type NoteRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - note entities.Note
func (_e *NoteRepository_Expecter) Create(ctx interface{}, userID interface{}, note interface{}) *NoteRepository_Create_Call {
	return &NoteRepository_Create_Call{Call: _e.mock.On("Create", ctx, userID, note)}
}

func (_c *NoteRepository_Create_Call) Run(run func(ctx context.Context, userID string, note entities.Note)) *NoteRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entities.Note))
	})
	return _c
}

func (_c *NoteRepository_Create_Call) Return(_a0 error) *NoteRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NoteRepository_Create_Call) RunAndReturn(run func(context.Context, string, entities.Note) error) *NoteRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, userID, meta
func (_m *NoteRepository) Delete(ctx context.Context, userID string, meta string) error {
	ret := _m.Called(ctx, userID, meta)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userID, meta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NoteRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type NoteRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - meta string
func (_e *NoteRepository_Expecter) Delete(ctx interface{}, userID interface{}, meta interface{}) *NoteRepository_Delete_Call {
	return &NoteRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, userID, meta)}
}

func (_c *NoteRepository_Delete_Call) Run(run func(ctx context.Context, userID string, meta string)) *NoteRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *NoteRepository_Delete_Call) Return(_a0 error) *NoteRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NoteRepository_Delete_Call) RunAndReturn(run func(context.Context, string, string) error) *NoteRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: ctx, userID
func (_m *NoteRepository) GetList(ctx context.Context, userID string) ([]entities.Note, error) {
	ret := _m.Called(ctx, userID)

	var r0 []entities.Note
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entities.Note, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entities.Note); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Note)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NoteRepository_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type NoteRepository_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *NoteRepository_Expecter) GetList(ctx interface{}, userID interface{}) *NoteRepository_GetList_Call {
	return &NoteRepository_GetList_Call{Call: _e.mock.On("GetList", ctx, userID)}
}

func (_c *NoteRepository_GetList_Call) Run(run func(ctx context.Context, userID string)) *NoteRepository_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *NoteRepository_GetList_Call) Return(_a0 []entities.Note, _a1 error) *NoteRepository_GetList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteRepository_GetList_Call) RunAndReturn(run func(context.Context, string) ([]entities.Note, error)) *NoteRepository_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// GetOne provides a mock function with given fields: ctx, userID, meta
func (_m *NoteRepository) GetOne(ctx context.Context, userID string, meta string) (entities.Note, error) {
	ret := _m.Called(ctx, userID, meta)

	var r0 entities.Note
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (entities.Note, error)); ok {
		return rf(ctx, userID, meta)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) entities.Note); ok {
		r0 = rf(ctx, userID, meta)
	} else {
		r0 = ret.Get(0).(entities.Note)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NoteRepository_GetOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOne'
type NoteRepository_GetOne_Call struct {
	*mock.Call
}

// GetOne is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - meta string
func (_e *NoteRepository_Expecter) GetOne(ctx interface{}, userID interface{}, meta interface{}) *NoteRepository_GetOne_Call {
	return &NoteRepository_GetOne_Call{Call: _e.mock.On("GetOne", ctx, userID, meta)}
}

func (_c *NoteRepository_GetOne_Call) Run(run func(ctx context.Context, userID string, meta string)) *NoteRepository_GetOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *NoteRepository_GetOne_Call) Return(_a0 entities.Note, _a1 error) *NoteRepository_GetOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteRepository_GetOne_Call) RunAndReturn(run func(context.Context, string, string) (entities.Note, error)) *NoteRepository_GetOne_Call {
	_c.Call.Return(run)
	return _c
}

// NewNoteRepository creates a new instance of NoteRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNoteRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *NoteRepository {
	mock := &NoteRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
