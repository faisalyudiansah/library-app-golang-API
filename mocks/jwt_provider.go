// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	helpers "library-api/helpers"

	mock "github.com/stretchr/testify/mock"
)

// JWTProvider is an autogenerated mock type for the JWTProvider type
type JWTProvider struct {
	mock.Mock
}

// CreateToken provides a mock function with given fields: userID
func (_m *JWTProvider) CreateToken(userID int64) (string, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for CreateToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (string, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyToken provides a mock function with given fields: token
func (_m *JWTProvider) VerifyToken(token string) (helpers.JWTClaims, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyToken")
	}

	var r0 helpers.JWTClaims
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (helpers.JWTClaims, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) helpers.JWTClaims); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(helpers.JWTClaims)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJWTProvider creates a new instance of JWTProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWTProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWTProvider {
	mock := &JWTProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
