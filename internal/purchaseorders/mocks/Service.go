// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	purchaseorders "github.com/cpereira42/mercado-fresco-pron4/internal/purchaseorders"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: order_number, order_date, tracking_code, buyer_id, product_record_id, order_status_id
func (_m *Service) Create(order_number string, order_date string, tracking_code string, buyer_id int, product_record_id int, order_status_id int) (purchaseorders.Purchase, error) {
	ret := _m.Called(order_number, order_date, tracking_code, buyer_id, product_record_id, order_status_id)

	var r0 purchaseorders.Purchase
	if rf, ok := ret.Get(0).(func(string, string, string, int, int, int) purchaseorders.Purchase); ok {
		r0 = rf(order_number, order_date, tracking_code, buyer_id, product_record_id, order_status_id)
	} else {
		r0 = ret.Get(0).(purchaseorders.Purchase)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, int, int, int) error); ok {
		r1 = rf(order_number, order_date, tracking_code, buyer_id, product_record_id, order_status_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *Service) GetById(id int) (purchaseorders.Purchase, error) {
	ret := _m.Called(id)

	var r0 purchaseorders.Purchase
	if rf, ok := ret.Get(0).(func(int) purchaseorders.Purchase); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(purchaseorders.Purchase)
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
