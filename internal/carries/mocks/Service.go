// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	carries "github.com/cpereira42/mercado-fresco-pron4/internal/carries"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: cid, companyName, address, telephone, localityID
func (_m *Service) Create(cid string, companyName string, address string, telephone string, localityID int) (carries.Carries, error) {
	ret := _m.Called(cid, companyName, address, telephone, localityID)

	var r0 carries.Carries
	if rf, ok := ret.Get(0).(func(string, string, string, string, int) carries.Carries); ok {
		r0 = rf(cid, companyName, address, telephone, localityID)
	} else {
		r0 = ret.Get(0).(carries.Carries)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, int) error); ok {
		r1 = rf(cid, companyName, address, telephone, localityID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllReport provides a mock function with given fields:
func (_m *Service) GetAllReport() ([]carries.Localities, error) {
	ret := _m.Called()

	var r0 []carries.Localities
	if rf, ok := ret.Get(0).(func() []carries.Localities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]carries.Localities)
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

// GetByIDReport provides a mock function with given fields: id
func (_m *Service) GetByIDReport(id int) (carries.Localities, error) {
	ret := _m.Called(id)

	var r0 carries.Localities
	if rf, ok := ret.Get(0).(func(int) carries.Localities); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(carries.Localities)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
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
