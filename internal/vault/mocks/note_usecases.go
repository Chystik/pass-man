// Code generated by mockery v2.23.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/Chystik/pass-man/internal/vault/note/entities"
	mock "github.com/stretchr/testify/mock"
)

// NoteUsecases is an autogenerated mock type for the NoteUsecases type
type NoteUsecases struct {
	mock.Mock
}

type NoteUsecases_Expecter struct {
	mock *mock.Mock
}

func (_m *NoteUsecases) EXPECT() *NoteUsecases_Expecter {
	return &NoteUsecases_Expecter{mock: &_m.Mock}
}

// AddNote provides a mock function with given fields: ctx, userID, note
func (_m *NoteUsecases) AddNote(ctx context.Context, userID string, note entities.Note) error {
	ret := _m.Called(ctx, userID, note)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entities.Note) error); ok {
		r0 = rf(ctx, userID, note)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NoteUsecases_AddNote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddNote'
type NoteUsecases_AddNote_Call struct {
	*mock.Call
}

// AddNote is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - note entities.Note
func (_e *NoteUsecases_Expecter) AddNote(ctx interface{}, userID interface{}, note interface{}) *NoteUsecases_AddNote_Call {
	return &NoteUsecases_AddNote_Call{Call: _e.mock.On("AddNote", ctx, userID, note)}
}

func (_c *NoteUsecases_AddNote_Call) Run(run func(ctx context.Context, userID string, note entities.Note)) *NoteUsecases_AddNote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entities.Note))
	})
	return _c
}

func (_c *NoteUsecases_AddNote_Call) Return(_a0 error) *NoteUsecases_AddNote_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NoteUsecases_AddNote_Call) RunAndReturn(run func(context.Context, string, entities.Note) error) *NoteUsecases_AddNote_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteNote provides a mock function with given fields: ctx, userID, meta
func (_m *NoteUsecases) DeleteNote(ctx context.Context, userID string, meta string) error {
	ret := _m.Called(ctx, userID, meta)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userID, meta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NoteUsecases_DeleteNote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteNote'
type NoteUsecases_DeleteNote_Call struct {
	*mock.Call
}

// DeleteNote is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - meta string
func (_e *NoteUsecases_Expecter) DeleteNote(ctx interface{}, userID interface{}, meta interface{}) *NoteUsecases_DeleteNote_Call {
	return &NoteUsecases_DeleteNote_Call{Call: _e.mock.On("DeleteNote", ctx, userID, meta)}
}

func (_c *NoteUsecases_DeleteNote_Call) Run(run func(ctx context.Context, userID string, meta string)) *NoteUsecases_DeleteNote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *NoteUsecases_DeleteNote_Call) Return(_a0 error) *NoteUsecases_DeleteNote_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NoteUsecases_DeleteNote_Call) RunAndReturn(run func(context.Context, string, string) error) *NoteUsecases_DeleteNote_Call {
	_c.Call.Return(run)
	return _c
}

// GetNote provides a mock function with given fields: ctx, userID, meta
func (_m *NoteUsecases) GetNote(ctx context.Context, userID string, meta string) (entities.Note, error) {
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

// NoteUsecases_GetNote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNote'
type NoteUsecases_GetNote_Call struct {
	*mock.Call
}

// GetNote is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - meta string
func (_e *NoteUsecases_Expecter) GetNote(ctx interface{}, userID interface{}, meta interface{}) *NoteUsecases_GetNote_Call {
	return &NoteUsecases_GetNote_Call{Call: _e.mock.On("GetNote", ctx, userID, meta)}
}

func (_c *NoteUsecases_GetNote_Call) Run(run func(ctx context.Context, userID string, meta string)) *NoteUsecases_GetNote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *NoteUsecases_GetNote_Call) Return(_a0 entities.Note, _a1 error) *NoteUsecases_GetNote_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteUsecases_GetNote_Call) RunAndReturn(run func(context.Context, string, string) (entities.Note, error)) *NoteUsecases_GetNote_Call {
	_c.Call.Return(run)
	return _c
}

// ListNote provides a mock function with given fields: ctx, userID
func (_m *NoteUsecases) ListNote(ctx context.Context, userID string) ([]entities.Note, error) {
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

// NoteUsecases_ListNote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListNote'
type NoteUsecases_ListNote_Call struct {
	*mock.Call
}

// ListNote is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *NoteUsecases_Expecter) ListNote(ctx interface{}, userID interface{}) *NoteUsecases_ListNote_Call {
	return &NoteUsecases_ListNote_Call{Call: _e.mock.On("ListNote", ctx, userID)}
}

func (_c *NoteUsecases_ListNote_Call) Run(run func(ctx context.Context, userID string)) *NoteUsecases_ListNote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *NoteUsecases_ListNote_Call) Return(_a0 []entities.Note, _a1 error) *NoteUsecases_ListNote_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteUsecases_ListNote_Call) RunAndReturn(run func(context.Context, string) ([]entities.Note, error)) *NoteUsecases_ListNote_Call {
	_c.Call.Return(run)
	return _c
}

// NewNoteUsecases creates a new instance of NoteUsecases. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNoteUsecases(t interface {
	mock.TestingT
	Cleanup(func())
}) *NoteUsecases {
	mock := &NoteUsecases{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
