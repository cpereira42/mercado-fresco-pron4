// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	seller "github.com/cpereira42/mercado-fresco-pron4/internal/seller"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CheckCid provides a mock function with given fields: cid
func (_m *Service) CheckCid(cid int) (bool, error) {
	ret := _m.Called(cid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(cid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(cid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: cid, company, adress, telephone
func (_m *Service) Create(cid int, company string, adress string, telephone string) (seller.Seller, error) {
	ret := _m.Called(cid, company, adress, telephone)

	var r0 seller.Seller
	if rf, ok := ret.Get(0).(func(int, string, string, string) seller.Seller); ok {
		r0 = rf(cid, company, adress, telephone)
	} else {
		r0 = ret.Get(0).(seller.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string, string) error); ok {
		r1 = rf(cid, company, adress, telephone)
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
func (_m *Service) GetAll() ([]seller.Seller, error) {
	ret := _m.Called()

	var r0 []seller.Seller
	if rf, ok := ret.Get(0).(func() []seller.Seller); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]seller.Seller)
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
func (_m *Service) GetId(id int) (seller.Seller, error) {
	ret := _m.Called(id)

	var r0 seller.Seller
	if rf, ok := ret.Get(0).(func(int) seller.Seller); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(seller.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, cid, company, adress, telephone
func (_m *Service) Update(id int, cid int, company string, adress string, telephone string) (seller.Seller, error) {
	ret := _m.Called(id, cid, company, adress, telephone)

	var r0 seller.Seller
	if rf, ok := ret.Get(0).(func(int, int, string, string, string) seller.Seller); ok {
		r0 = rf(id, cid, company, adress, telephone)
	} else {
		r0 = ret.Get(0).(seller.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string, string, string) error); ok {
		r1 = rf(id, cid, company, adress, telephone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
