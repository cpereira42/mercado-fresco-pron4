package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/stretchr/testify/assert"

	inboundOrders "github.com/cpereira42/mercado-fresco-pron4/internal/inbound_orders"
	"github.com/cpereira42/mercado-fresco-pron4/internal/inbound_orders/mocks"
	"github.com/gin-gonic/gin"
)

var (
	inboundHandler1              = inboundOrders.ReportInboundOrders{ID: 1, CardNumberID: "1", FirstName: "mercado", LastName: "livre", WarehouseID: 1, InboundOrdersCount: 5}
	inboundHandler2              = inboundOrders.ReportInboundOrders{ID: 2, CardNumberID: "121", FirstName: "edu", LastName: "araujo", WarehouseID: 2, InboundOrdersCount: 4}
	inboundOrderHandlerCreate    = inboundOrders.InboundOrdersCreate{OrderNumber: "Order2", EmployeeID: 3, ProductBatchID: 1, WarehouseID: 1}
	inboundOrdersHandlerResponse = inboundOrders.InboundOrdersResponse{OrderDate: "2022-07-08 10:48:22", OrderNumber: "Order2", EmployeeID: 3, ProductBatchID: 1, WarehouseID: 1}
)

var employeesReportResponseJson = struct {
	Code int
	Data []inboundOrders.ReportInboundOrders
}{}

var employeeReportResponseJson = struct {
	Code int
	Data inboundOrders.ReportInboundOrders
}{}

var inboundOrderResponseJson = struct {
	Code int
	Data inboundOrders.InboundOrdersResponse
}{}

var inboundOrderResponseJsonError = struct {
	Code  int
	Data  inboundOrders.InboundOrdersResponse
	Error string
}{}

var employeeReportResponseJsonError = struct {
	Code  int
	Data  inboundOrders.ReportInboundOrders
	Error string
}{}

var inboundOrdersFieldsResponseJson = struct {
	Code int
	Data []struct {
		Field   string
		Message string
	}
}{}

func createServerInboundOrders(serverMock *mocks.Service, method string, url string, body string) *httptest.ResponseRecorder {
	r := gin.Default()
	handler.NewInboundOrders(r, serverMock)
	req, rr := util.CreateRequestTest(method, url, body)
	r.ServeHTTP(rr, req)
	return rr
}

func TestHandlerReportInboundOrders(t *testing.T) {
	employees := []inboundOrders.ReportInboundOrders{inboundHandler1, inboundHandler2}

	t.Run("If request GetAll is OK, it should return status code 200 and a list of employees report", func(t *testing.T) {
		serviceMock := &mocks.Service{}
		serviceMock.On("GetAll").Return(employees, nil)
		rr := createServerInboundOrders(serviceMock, http.MethodGet, "/api/v1/employees/reportInboundOrders", "")

		err := json.Unmarshal(rr.Body.Bytes(), &employeesReportResponseJson)

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, employeesReportResponseJson.Data, employees)
		assert.Nil(t, err)
		assert.True(t, len(employeesReportResponseJson.Data) > 0)
	})

	t.Run("If request GetAll has an error, it should return status code 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("error to get all employees report")
		serviceMock := &mocks.Service{}
		serviceMock.On("GetAll").Return([]inboundOrders.ReportInboundOrders{}, errorMsg)
		rr := createServerInboundOrders(serviceMock, http.MethodGet, "/api/v1/employees/reportInboundOrders", "")
		err := json.Unmarshal(rr.Body.Bytes(), &employeeReportResponseJsonError)

		assert.Equal(t, 404, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), employeeReportResponseJsonError.Error)
	})

	t.Run("If request GetByID is OK, it should return status code 200 and an employee report", func(t *testing.T) {
		serviceMock := &mocks.Service{}
		serviceMock.On("GetByID", 1).Return(inboundHandler1, nil)
		rr := createServerInboundOrders(serviceMock, http.MethodGet, "/api/v1/employees/reportInboundOrders?id=1", "")

		err := json.Unmarshal(rr.Body.Bytes(), &employeeReportResponseJson)

		assert.Equal(t, 200, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, inboundHandler1, employeeReportResponseJson.Data)

	})
	t.Run("If user is not passing a number in the parameter of the url to GetByID, it should return status code 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("invalid ID")
		serviceMock := &mocks.Service{}
		serviceMock.On("GetByID").Return(inboundOrders.ReportInboundOrders{}, errorMsg)
		rr := createServerInboundOrders(serviceMock, http.MethodGet, "/api/v1/employees/reportInboundOrders?id=a", "")

		err := json.Unmarshal(rr.Body.Bytes(), &employeeReportResponseJsonError)

		assert.Equal(t, 404, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), employeeReportResponseJsonError.Error)
	})
	t.Run("If request GetByID does not find a employee, it should return 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("employee not found")
		serviceMock := &mocks.Service{}
		serviceMock.On("GetByID", 100).Return(inboundOrders.ReportInboundOrders{}, errorMsg)
		rr := createServerInboundOrders(serviceMock, http.MethodGet, "/api/v1/employees/reportInboundOrders?id=100", "")

		err := json.Unmarshal(rr.Body.Bytes(), &employeeReportResponseJsonError)

		assert.Equal(t, 404, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), employeeReportResponseJsonError.Error)
	})
}

func TestHandlerCreateInboundOrder(t *testing.T) {

	t.Run("If request Create is OK, it should return status code 201 and an inbound order", func(t *testing.T) {
		serviceMock := &mocks.Service{}
		serviceMock.On("Create", inboundOrderHandlerCreate).Return(inboundOrdersHandlerResponse, nil)
		rr := createServerInboundOrders(serviceMock, http.MethodPost, "/api/v1/inboundOrders",
			`
			{
				"order_number":"Order2",
				"employee_id": 3,
				"product_batch_id":1,
				"warehouse_id":1
			 }
			`)

		err := json.Unmarshal(rr.Body.Bytes(), &inboundOrderResponseJson)

		assert.Equal(t, 201, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, inboundOrdersHandlerResponse, inboundOrderResponseJson.Data)

	})
	t.Run("If user is not passing the correct body to the Create , it should return status code 422 and an error", func(t *testing.T) {
		serviceMock := &mocks.Service{}
		rr := createServerInboundOrders(serviceMock, http.MethodPost, "/api/v1/inboundOrders",
			`{
			"order_number":"Order2",
			"employee_id": 3,
			"warehouse_id":1
		 }`,
		)

		assert.Equal(t, 422, rr.Code)

		err := json.Unmarshal(rr.Body.Bytes(), &inboundOrdersFieldsResponseJson)
		assert.Nil(t, err)
		assert.Equal(t, "This field is required", inboundOrdersFieldsResponseJson.Data[0].Message)
		assert.Equal(t, "product_batch_id", inboundOrdersFieldsResponseJson.Data[0].Field)
	})
	t.Run("If there is an error to Create, it should return status code 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("error to create inbound order")
		serviceMock := &mocks.Service{}
		serviceMock.On("Create", inboundOrderHandlerCreate).Return(inboundOrders.InboundOrdersResponse{}, errorMsg)
		rr := createServerInboundOrders(serviceMock, http.MethodPost, "/api/v1/inboundOrders",
			`{
				"order_number":"Order2",
				"employee_id": 3,
				"product_batch_id":1,
				"warehouse_id":1
			 }`,
		)

		err := json.Unmarshal(rr.Body.Bytes(), &inboundOrderResponseJsonError)

		assert.Equal(t, 404, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), inboundOrderResponseJsonError.Error)
	})
}
