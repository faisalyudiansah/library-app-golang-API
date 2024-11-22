// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "library-api/models"

	mock "github.com/stretchr/testify/mock"
)

// AuthorRepository is an autogenerated mock type for the AuthorRepository type
type AuthorRepository struct {
	mock.Mock
}

// GetAuthorById provides a mock function with given fields: _a0, _a1
func (_m *AuthorRepository) GetAuthorById(_a0 context.Context, _a1 int64) (*models.Author, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorById")
	}

	var r0 *models.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*models.Author, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Author); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Author)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthorRepository creates a new instance of AuthorRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthorRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthorRepository {
	mock := &AuthorRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
