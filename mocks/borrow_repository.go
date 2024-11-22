// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	context "context"
	dtos "library-api/dtos"

	mock "github.com/stretchr/testify/mock"

	models "library-api/models"
)

// BorrowRepository is an autogenerated mock type for the BorrowRepository type
type BorrowRepository struct {
	mock.Mock
}

// IsAlreadyReturnBook provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *BorrowRepository) IsAlreadyReturnBook(_a0 context.Context, _a1 int64, _a2 int64, _a3 int64) bool {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for IsAlreadyReturnBook")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int64) bool); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsBorrowIdValid provides a mock function with given fields: _a0, _a1, _a2
func (_m *BorrowRepository) IsBorrowIdValid(_a0 context.Context, _a1 int64, _a2 int64) bool {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for IsBorrowIdValid")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsUserBorrowNow provides a mock function with given fields: _a0, _a1
func (_m *BorrowRepository) IsUserBorrowNow(_a0 context.Context, _a1 int64) bool {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for IsUserBorrowNow")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int64) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PostNewBorrow provides a mock function with given fields: _a0, _a1, _a2
func (_m *BorrowRepository) PostNewBorrow(_a0 context.Context, _a1 dtos.RequestBorrowBook, _a2 int64) (*models.Borrow, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for PostNewBorrow")
	}

	var r0 *models.Borrow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dtos.RequestBorrowBook, int64) (*models.Borrow, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dtos.RequestBorrowBook, int64) *models.Borrow); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Borrow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dtos.RequestBorrowBook, int64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostReturnBook provides a mock function with given fields: _a0, _a1
func (_m *BorrowRepository) PostReturnBook(_a0 context.Context, _a1 dtos.RequestReturnBook) (*models.Borrow, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for PostReturnBook")
	}

	var r0 *models.Borrow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dtos.RequestReturnBook) (*models.Borrow, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dtos.RequestReturnBook) *models.Borrow); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Borrow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dtos.RequestReturnBook) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBorrowRepository creates a new instance of BorrowRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBorrowRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BorrowRepository {
	mock := &BorrowRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
