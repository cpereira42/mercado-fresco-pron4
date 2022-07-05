// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	locality "github.com/cpereira42/mercado-fresco-pron4/internal/locality"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryLocality is an autogenerated mock type for the RepositoryLocality type
type RepositoryLocality struct {
	mock.Mock
}

// Create provides a mock function with given fields: localityName, provinceName, countryName
func (_m *RepositoryLocality) Create(localityName string, provinceName string, countryName string) (locality.Locality, error) {
	ret := _m.Called(localityName, provinceName, countryName)

	var r0 locality.Locality
	if rf, ok := ret.Get(0).(func(string, string, string) locality.Locality); ok {
		r0 = rf(localityName, provinceName, countryName)
	} else {
		r0 = ret.Get(0).(locality.Locality)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(localityName, provinceName, countryName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateReportAll provides a mock function with given fields:
func (_m *RepositoryLocality) GenerateReportAll() ([]locality.LocalityReport, error) {
	ret := _m.Called()

	var r0 []locality.LocalityReport
	if rf, ok := ret.Get(0).(func() []locality.LocalityReport); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]locality.LocalityReport)
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

// GenerateReportById provides a mock function with given fields: id
func (_m *RepositoryLocality) GenerateReportById(id int) (locality.LocalityReport, error) {
	ret := _m.Called(id)

	var r0 locality.LocalityReport
	if rf, ok := ret.Get(0).(func(int) locality.LocalityReport); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(locality.LocalityReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *RepositoryLocality) GetAll() ([]locality.Locality, error) {
	ret := _m.Called()

	var r0 []locality.Locality
	if rf, ok := ret.Get(0).(func() []locality.Locality); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]locality.Locality)
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

type mockConstructorTestingTNewRepositoryLocality interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryLocality creates a new instance of RepositoryLocality. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryLocality(t mockConstructorTestingTNewRepositoryLocality) *RepositoryLocality {
	mock := &RepositoryLocality{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}