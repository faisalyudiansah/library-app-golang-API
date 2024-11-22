// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	context "context"
	dtos "library-api/dtos"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// PostLoginUserService provides a mock function with given fields: _a0, _a1
func (_m *UserService) PostLoginUserService(_a0 context.Context, _a1 dtos.RequestLoginUser) (*dtos.ResponseDataUser, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for PostLoginUserService")
	}

	var r0 *dtos.ResponseDataUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dtos.RequestLoginUser) (*dtos.ResponseDataUser, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dtos.RequestLoginUser) *dtos.ResponseDataUser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.ResponseDataUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dtos.RequestLoginUser) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostRegisterUserService provides a mock function with given fields: _a0, _a1
func (_m *UserService) PostRegisterUserService(_a0 context.Context, _a1 dtos.RequestRegisterUser) (*dtos.ResponseDataUser, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for PostRegisterUserService")
	}

	var r0 *dtos.ResponseDataUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dtos.RequestRegisterUser) (*dtos.ResponseDataUser, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dtos.RequestRegisterUser) *dtos.ResponseDataUser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.ResponseDataUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dtos.RequestRegisterUser) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
