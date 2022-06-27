// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	warehouse "github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature
func (_m *Repository) Create(id int, address string, telephone string, warehouse_code string, minimum_capacity int, minimum_temperature int) (warehouse.Warehouse, error) {
	ret := _m.Called(id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature)

	var r0 warehouse.Warehouse
	if rf, ok := ret.Get(0).(func(int, string, string, string, int, int) warehouse.Warehouse); ok {
		r0 = rf(id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature)
	} else {
		r0 = ret.Get(0).(warehouse.Warehouse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string, string, int, int) error); ok {
		r1 = rf(id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) error {
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
func (_m *Repository) GetAll() ([]warehouse.Warehouse, error) {
	ret := _m.Called()

	var r0 []warehouse.Warehouse
	if rf, ok := ret.Get(0).(func() []warehouse.Warehouse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]warehouse.Warehouse)
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

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id int) (warehouse.Warehouse, error) {
	ret := _m.Called(id)

	var r0 warehouse.Warehouse
	if rf, ok := ret.Get(0).(func(int) warehouse.Warehouse); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(warehouse.Warehouse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LastID provides a mock function with given fields:
func (_m *Repository) LastID() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature
func (_m *Repository) Update(id int, address string, telephone string, warehouse_code string, minimum_capacity int, minimum_temperature int) (warehouse.Warehouse, error) {
	ret := _m.Called(id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature)

	var r0 warehouse.Warehouse
	if rf, ok := ret.Get(0).(func(int, string, string, string, int, int) warehouse.Warehouse); ok {
		r0 = rf(id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature)
	} else {
		r0 = ret.Get(0).(warehouse.Warehouse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string, string, int, int) error); ok {
		r1 = rf(id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}