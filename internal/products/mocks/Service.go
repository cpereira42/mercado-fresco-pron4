// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	products "github.com/cpereira42/mercado-fresco-pron4/internal/products"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CheckCode provides a mock function with given fields: id, code
func (_m *Service) CheckCode(id int, code string) bool {
	ret := _m.Called(id, code)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, string) bool); ok {
		r0 = rf(id, code)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Create provides a mock function with given fields: p
func (_m *Service) Create(p products.RequestProductsCreate) (products.Product, error) {
	ret := _m.Called(p)

	var r0 products.Product
	if rf, ok := ret.Get(0).(func(products.RequestProductsCreate) products.Product); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(products.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(products.RequestProductsCreate) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Service) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Service) GetAll() ([]products.Product, error) {
	ret := _m.Called()

	var r0 []products.Product
	if rf, ok := ret.Get(0).(func() []products.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetId provides a mock function with given fields: id
func (_m *Service) GetId(id int) (products.Product, error) {
	ret := _m.Called(id)

	var r0 products.Product
	if rf, ok := ret.Get(0).(func(int) products.Product); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(products.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, p
func (_m *Service) Update(id int, p products.RequestProductsUpdate) (products.Product, error) {
	ret := _m.Called(id, p)

	var r0 products.Product
	if rf, ok := ret.Get(0).(func(int, products.RequestProductsUpdate) products.Product); ok {
		r0 = rf(id, p)
	} else {
		r0 = ret.Get(0).(products.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, products.RequestProductsUpdate) error); ok {
		r1 = rf(id, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewServiceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t NewServiceT) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}