// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	context "context"
	dtos "library-api/dtos"

	mock "github.com/stretchr/testify/mock"

	models "library-api/models"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// GetAllRepository provides a mock function with given fields: _a0, _a1
func (_m *BookRepository) GetAllRepository(_a0 context.Context, _a1 string) ([]models.AuthorBook, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAllRepository")
	}

	var r0 []models.AuthorBook
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]models.AuthorBook, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.AuthorBook); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.AuthorBook)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBookByID provides a mock function with given fields: _a0, _a1
func (_m *BookRepository) GetBookByID(_a0 context.Context, _a1 int64) (*models.Book, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetBookByID")
	}

	var r0 *models.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*models.Book, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Book); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsBookHasTheSameTitle provides a mock function with given fields: _a0, _a1
func (_m *BookRepository) IsBookHasTheSameTitle(_a0 context.Context, _a1 string) bool {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for IsBookHasTheSameTitle")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsBookOutOfStock provides a mock function with given fields: _a0
func (_m *BookRepository) IsBookOutOfStock(_a0 int64) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for IsBookOutOfStock")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PostBookRepository provides a mock function with given fields: _a0, _a1, _a2
func (_m *BookRepository) PostBookRepository(_a0 context.Context, _a1 dtos.RequestPostBook, _a2 models.Author) (*models.AuthorBook, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for PostBookRepository")
	}

	var r0 *models.AuthorBook
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dtos.RequestPostBook, models.Author) (*models.AuthorBook, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dtos.RequestPostBook, models.Author) *models.AuthorBook); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AuthorBook)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dtos.RequestPostBook, models.Author) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutQuantityBook provides a mock function with given fields: _a0, _a1, _a2
func (_m *BookRepository) PutQuantityBook(_a0 context.Context, _a1 int, _a2 int64) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for PutQuantityBook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int64) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBookRepository creates a new instance of BookRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBookRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BookRepository {
	mock := &BookRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
