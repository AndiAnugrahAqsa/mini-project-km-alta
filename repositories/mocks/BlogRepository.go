// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "mini-project/models"

	mock "github.com/stretchr/testify/mock"
)

// BlogRepository is an autogenerated mock type for the BlogRepository type
type BlogRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: blogRequest
func (_m *BlogRepository) Create(blogRequest models.BlogRequest) models.Blog {
	ret := _m.Called(blogRequest)

	var r0 models.Blog
	if rf, ok := ret.Get(0).(func(models.BlogRequest) models.Blog); ok {
		r0 = rf(blogRequest)
	} else {
		r0 = ret.Get(0).(models.Blog)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *BlogRepository) Delete(id int) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *BlogRepository) GetAll() []models.Blog {
	ret := _m.Called()

	var r0 []models.Blog
	if rf, ok := ret.Get(0).(func() []models.Blog); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Blog)
		}
	}

	return r0
}

// GetByCategoryID provides a mock function with given fields: category_id
func (_m *BlogRepository) GetByCategoryID(category_id int) []models.Blog {
	ret := _m.Called(category_id)

	var r0 []models.Blog
	if rf, ok := ret.Get(0).(func(int) []models.Blog); ok {
		r0 = rf(category_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Blog)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *BlogRepository) GetByID(id int) models.Blog {
	ret := _m.Called(id)

	var r0 models.Blog
	if rf, ok := ret.Get(0).(func(int) models.Blog); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Blog)
	}

	return r0
}

// GetByUserID provides a mock function with given fields: user_id
func (_m *BlogRepository) GetByUserID(user_id int) []models.Blog {
	ret := _m.Called(user_id)

	var r0 []models.Blog
	if rf, ok := ret.Get(0).(func(int) []models.Blog); ok {
		r0 = rf(user_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Blog)
		}
	}

	return r0
}

// Update provides a mock function with given fields: id, blogRequest
func (_m *BlogRepository) Update(id int, blogRequest models.BlogRequest) models.Blog {
	ret := _m.Called(id, blogRequest)

	var r0 models.Blog
	if rf, ok := ret.Get(0).(func(int, models.BlogRequest) models.Blog); ok {
		r0 = rf(id, blogRequest)
	} else {
		r0 = ret.Get(0).(models.Blog)
	}

	return r0
}

type mockConstructorTestingTNewBlogRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlogRepository creates a new instance of BlogRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlogRepository(t mockConstructorTestingTNewBlogRepository) *BlogRepository {
	mock := &BlogRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}