// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	productsRecords "github.com/cpereira42/mercado-fresco-pron4/internal/productsRecords"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: p
func (_m *Service) Create(p productsRecords.RequestProductRecordsCreate) (productsRecords.ProductRecords, error) {
	ret := _m.Called(p)

	var r0 productsRecords.ProductRecords
	if rf, ok := ret.Get(0).(func(productsRecords.RequestProductRecordsCreate) productsRecords.ProductRecords); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(productsRecords.ProductRecords)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(productsRecords.RequestProductRecordsCreate) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllRecords provides a mock function with given fields:
func (_m *Service) GetAllRecords() ([]productsRecords.ReturnProductRecords, error) {
	ret := _m.Called()

	var r0 []productsRecords.ReturnProductRecords
	if rf, ok := ret.Get(0).(func() []productsRecords.ReturnProductRecords); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]productsRecords.ReturnProductRecords)
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

// GetIdRecords provides a mock function with given fields: id
func (_m *Service) GetIdRecords(id int) (productsRecords.ReturnProductRecords, error) {
	ret := _m.Called(id)

	var r0 productsRecords.ReturnProductRecords
	if rf, ok := ret.Get(0).(func(int) productsRecords.ReturnProductRecords); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(productsRecords.ReturnProductRecords)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
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