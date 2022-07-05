// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	productbatch "github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryProductBatches is an autogenerated mock type for the RepositoryProductBatches type
type RepositoryProductBatches struct {
	mock.Mock
}

// CreatePB provides a mock function with given fields: productBatches
func (_m *RepositoryProductBatches) CreatePB(productBatches productbatch.ProductBatches) (productbatch.ProductBatches, error) {
	ret := _m.Called(productBatches)

	var r0 productbatch.ProductBatches
	if rf, ok := ret.Get(0).(func(productbatch.ProductBatches) productbatch.ProductBatches); ok {
		r0 = rf(productBatches)
	} else {
		r0 = ret.Get(0).(productbatch.ProductBatches)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(productbatch.ProductBatches) error); ok {
		r1 = rf(productBatches)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByBatcheNumber provides a mock function with given fields: batch_number
func (_m *RepositoryProductBatches) GetByBatcheNumber(batch_number string) (bool, error) {
	ret := _m.Called(batch_number)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(batch_number)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(batch_number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadPBSectionId provides a mock function with given fields: id
func (_m *RepositoryProductBatches) ReadPBSectionId(id int64) (productbatch.ProductBatchesResponse, error) {
	ret := _m.Called(id)

	var r0 productbatch.ProductBatchesResponse
	if rf, ok := ret.Get(0).(func(int64) productbatch.ProductBatchesResponse); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(productbatch.ProductBatchesResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadPBSectionTodo provides a mock function with given fields:
func (_m *RepositoryProductBatches) ReadPBSectionTodo() ([]productbatch.ProductBatchesResponse, error) {
	ret := _m.Called()

	var r0 []productbatch.ProductBatchesResponse
	if rf, ok := ret.Get(0).(func() []productbatch.ProductBatchesResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]productbatch.ProductBatchesResponse)
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

// SearchProductById provides a mock function with given fields: id
func (_m *RepositoryProductBatches) SearchProductById(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchSectionId provides a mock function with given fields: id
func (_m *RepositoryProductBatches) SearchSectionId(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepositoryProductBatches interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryProductBatches creates a new instance of RepositoryProductBatches. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryProductBatches(t mockConstructorTestingTNewRepositoryProductBatches) *RepositoryProductBatches {
	mock := &RepositoryProductBatches{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
