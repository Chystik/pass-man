// Code generated by mockery v2.23.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// VaultKeyStore is an autogenerated mock type for the VaultKeyStore type
type VaultKeyStore struct {
	mock.Mock
}

type VaultKeyStore_Expecter struct {
	mock *mock.Mock
}

func (_m *VaultKeyStore) EXPECT() *VaultKeyStore_Expecter {
	return &VaultKeyStore_Expecter{mock: &_m.Mock}
}

// GetKey provides a mock function with given fields: login
func (_m *VaultKeyStore) GetKey(login string) ([]byte, error) {
	ret := _m.Called(login)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(login)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VaultKeyStore_GetKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetKey'
type VaultKeyStore_GetKey_Call struct {
	*mock.Call
}

// GetKey is a helper method to define mock.On call
//   - login string
func (_e *VaultKeyStore_Expecter) GetKey(login interface{}) *VaultKeyStore_GetKey_Call {
	return &VaultKeyStore_GetKey_Call{Call: _e.mock.On("GetKey", login)}
}

func (_c *VaultKeyStore_GetKey_Call) Run(run func(login string)) *VaultKeyStore_GetKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *VaultKeyStore_GetKey_Call) Return(_a0 []byte, _a1 error) *VaultKeyStore_GetKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VaultKeyStore_GetKey_Call) RunAndReturn(run func(string) ([]byte, error)) *VaultKeyStore_GetKey_Call {
	_c.Call.Return(run)
	return _c
}

// Lock provides a mock function with given fields: login
func (_m *VaultKeyStore) Lock(login string) error {
	ret := _m.Called(login)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(login)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VaultKeyStore_Lock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Lock'
type VaultKeyStore_Lock_Call struct {
	*mock.Call
}

// Lock is a helper method to define mock.On call
//   - login string
func (_e *VaultKeyStore_Expecter) Lock(login interface{}) *VaultKeyStore_Lock_Call {
	return &VaultKeyStore_Lock_Call{Call: _e.mock.On("Lock", login)}
}

func (_c *VaultKeyStore_Lock_Call) Run(run func(login string)) *VaultKeyStore_Lock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *VaultKeyStore_Lock_Call) Return(_a0 error) *VaultKeyStore_Lock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VaultKeyStore_Lock_Call) RunAndReturn(run func(string) error) *VaultKeyStore_Lock_Call {
	_c.Call.Return(run)
	return _c
}

// Unlock provides a mock function with given fields: login, key
func (_m *VaultKeyStore) Unlock(login string, key []byte) error {
	ret := _m.Called(login, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(login, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VaultKeyStore_Unlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unlock'
type VaultKeyStore_Unlock_Call struct {
	*mock.Call
}

// Unlock is a helper method to define mock.On call
//   - login string
//   - key []byte
func (_e *VaultKeyStore_Expecter) Unlock(login interface{}, key interface{}) *VaultKeyStore_Unlock_Call {
	return &VaultKeyStore_Unlock_Call{Call: _e.mock.On("Unlock", login, key)}
}

func (_c *VaultKeyStore_Unlock_Call) Run(run func(login string, key []byte)) *VaultKeyStore_Unlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]byte))
	})
	return _c
}

func (_c *VaultKeyStore_Unlock_Call) Return(_a0 error) *VaultKeyStore_Unlock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VaultKeyStore_Unlock_Call) RunAndReturn(run func(string, []byte) error) *VaultKeyStore_Unlock_Call {
	_c.Call.Return(run)
	return _c
}

// NewVaultKeyStore creates a new instance of VaultKeyStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVaultKeyStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *VaultKeyStore {
	mock := &VaultKeyStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}