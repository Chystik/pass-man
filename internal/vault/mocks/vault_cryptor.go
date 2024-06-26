// Code generated by mockery v2.23.4. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// VaultCryptor is an autogenerated mock type for the VaultCryptor type
type VaultCryptor struct {
	mock.Mock
}

type VaultCryptor_Expecter struct {
	mock *mock.Mock
}

func (_m *VaultCryptor) EXPECT() *VaultCryptor_Expecter {
	return &VaultCryptor_Expecter{mock: &_m.Mock}
}

// Decrypt provides a mock function with given fields: in, out, userID
func (_m *VaultCryptor) Decrypt(in io.Reader, out io.Writer, userID string) (int, error) {
	ret := _m.Called(in, out, userID)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(io.Reader, io.Writer, string) (int, error)); ok {
		return rf(in, out, userID)
	}
	if rf, ok := ret.Get(0).(func(io.Reader, io.Writer, string) int); ok {
		r0 = rf(in, out, userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(io.Reader, io.Writer, string) error); ok {
		r1 = rf(in, out, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VaultCryptor_Decrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decrypt'
type VaultCryptor_Decrypt_Call struct {
	*mock.Call
}

// Decrypt is a helper method to define mock.On call
//   - in io.Reader
//   - out io.Writer
//   - userID string
func (_e *VaultCryptor_Expecter) Decrypt(in interface{}, out interface{}, userID interface{}) *VaultCryptor_Decrypt_Call {
	return &VaultCryptor_Decrypt_Call{Call: _e.mock.On("Decrypt", in, out, userID)}
}

func (_c *VaultCryptor_Decrypt_Call) Run(run func(in io.Reader, out io.Writer, userID string)) *VaultCryptor_Decrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(io.Reader), args[1].(io.Writer), args[2].(string))
	})
	return _c
}

func (_c *VaultCryptor_Decrypt_Call) Return(_a0 int, _a1 error) *VaultCryptor_Decrypt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VaultCryptor_Decrypt_Call) RunAndReturn(run func(io.Reader, io.Writer, string) (int, error)) *VaultCryptor_Decrypt_Call {
	_c.Call.Return(run)
	return _c
}

// Encrypt provides a mock function with given fields: in, out, userID
func (_m *VaultCryptor) Encrypt(in io.Reader, out io.Writer, userID string) (int, error) {
	ret := _m.Called(in, out, userID)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(io.Reader, io.Writer, string) (int, error)); ok {
		return rf(in, out, userID)
	}
	if rf, ok := ret.Get(0).(func(io.Reader, io.Writer, string) int); ok {
		r0 = rf(in, out, userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(io.Reader, io.Writer, string) error); ok {
		r1 = rf(in, out, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VaultCryptor_Encrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encrypt'
type VaultCryptor_Encrypt_Call struct {
	*mock.Call
}

// Encrypt is a helper method to define mock.On call
//   - in io.Reader
//   - out io.Writer
//   - userID string
func (_e *VaultCryptor_Expecter) Encrypt(in interface{}, out interface{}, userID interface{}) *VaultCryptor_Encrypt_Call {
	return &VaultCryptor_Encrypt_Call{Call: _e.mock.On("Encrypt", in, out, userID)}
}

func (_c *VaultCryptor_Encrypt_Call) Run(run func(in io.Reader, out io.Writer, userID string)) *VaultCryptor_Encrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(io.Reader), args[1].(io.Writer), args[2].(string))
	})
	return _c
}

func (_c *VaultCryptor_Encrypt_Call) Return(_a0 int, _a1 error) *VaultCryptor_Encrypt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VaultCryptor_Encrypt_Call) RunAndReturn(run func(io.Reader, io.Writer, string) (int, error)) *VaultCryptor_Encrypt_Call {
	_c.Call.Return(run)
	return _c
}

// NewVaultCryptor creates a new instance of VaultCryptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVaultCryptor(t interface {
	mock.TestingT
	Cleanup(func())
}) *VaultCryptor {
	mock := &VaultCryptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
