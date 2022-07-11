package inbound_orders_test

import (
	"fmt"
	"testing"

	inboundOrders "github.com/cpereira42/mercado-fresco-pron4/internal/inbound_orders"
	"github.com/cpereira42/mercado-fresco-pron4/internal/inbound_orders/mocks"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var (
	inboundServDB1            = inboundOrders.ReportInboundOrders{1, "1", "mercado", "livre", 1, 5}
	inboundServDB2            = inboundOrders.ReportInboundOrders{2, "121", "edu", "araujo", 2, 4}
	inboundOrderServCreate    = inboundOrders.InboundOrdersCreate{"Order2", 3, 1, 1}
	inboundOrdersServResponse = inboundOrders.InboundOrdersResponse{"2022-07-08 10:48:22", "Order2", 3, 1, 1}
)

func TestServiceGetAll(t *testing.T) {
	inboundOrdersReport := []inboundOrders.ReportInboundOrders{inboundServDB1, inboundServDB2}
	t.Run("If GetAll is ok, it should return a list of employees reports", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetAll").Return(inboundOrdersReport, nil)

		service := inboundOrders.NewService(repo)
		employeesReport, err := service.GetAll()

		assert.Nil(t, err)
		assert.True(t, len(employeesReport) == 2)
		assert.Equal(t, inboundOrdersReport, employeesReport)
	})

	t.Run("If GetAll has an error, it should return an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("Failed to get all employees reports")
		repo := &mocks.Repository{}
		repo.On("GetAll").Return([]inboundOrders.ReportInboundOrders{}, errorMsg).Once()

		service := inboundOrders.NewService(repo)
		employeesReport, err := service.GetAll()

		assert.True(t, len(employeesReport) == 0)
		assert.Equal(t, fmt.Errorf("Failed to get all employees reports"), err)
		repo.AssertExpectations(t)

	})
}

func TestServiceGetByID(t *testing.T) {
	t.Run("If GetByID is success, it should return a employee report",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("GetByID", tmock.AnythingOfType("int")).Return(inboundServDB1, nil).Once()

			service := inboundOrders.NewService(repo)
			inboundOrdersReport, err := service.GetByID(1)

			assert.NoError(t, err)
			assert.Equal(t, inboundServDB1, inboundOrdersReport)

		})

	t.Run("If the id in the GetByID does not exists, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("employee not found")
			repo := &mocks.Repository{}
			repo.On("GetByID", tmock.AnythingOfType("int")).Return(inboundOrders.ReportInboundOrders{}, errorMsg).Once()

			service := inboundOrders.NewService(repo)
			_, err := service.GetByID(10)

			assert.Error(t, err)
			assert.EqualError(t, errorMsg, err.Error())

		})
}

func TestServiceCreate(t *testing.T) {
	t.Run("If Create is success, it should return a employee",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("Create", tmock.AnythingOfType("string"), inboundOrderServCreate).Return(inboundOrdersResponse, nil)

			service := inboundOrders.NewService(repo)
			employeeCreated, err := service.Create(inboundOrderServCreate)

			assert.NoError(t, err)
			assert.Equal(t, inboundOrdersResponse, employeeCreated)

		})

	t.Run("If Create has an error, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("error to create")
			repo := &mocks.Repository{}
			repo.On("Create", tmock.AnythingOfType("string"), inboundOrderServCreate).Return(inboundOrders.InboundOrdersResponse{}, errorMsg)

			service := inboundOrders.NewService(repo)
			_, err := service.Create(inboundOrderServCreate)

			assert.Error(t, err)
			assert.EqualError(t, errorMsg, err.Error())
		})
}
