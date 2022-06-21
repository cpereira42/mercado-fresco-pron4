package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var (
	employee1 = employee.Employee{ID: 1, CardNumberID: "123", FirstName: "Eduardo", LastName: "Araujo", WarehouseID: 1}

	employee2 = employee.Employee{ID: 2, CardNumberID: "1234", FirstName: "Maria", LastName: "Silva", WarehouseID: 2}

	employee3 = employee.Employee{ID: 3, CardNumberID: "12345", FirstName: "Jefferson", LastName: "Filho", WarehouseID: 1}

	newEmployee = employee.Employee{CardNumberID: "123456", FirstName: "Marta", LastName: "Gomes", WarehouseID: 3}

	employee1Update = employee.Employee{ID: 1, CardNumberID: "123", FirstName: "Gustavo", LastName: "Junior", WarehouseID: 1}

	employeeUpdateSameCardNumberID = employee.Employee{ID: 2, CardNumberID: "123", FirstName: "Gustavo", LastName: "Junior", WarehouseID: 1}
)

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestHandlerGetAll(t *testing.T) {
	employees := []employee.Employee{employee1, employee2, employee3}

	t.Run("If request GetAll is OK, it should return status code 200 and a list of employees", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodGet, "/api/v1/employees/", "")
		serviceMock := &mocks.Service{}
		serviceMock.On("GetAll").Return(employees, nil)

		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees/")
		employeeGroup.GET("/", e.GetAll())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)

		jsonResponse := struct {
			Code int
			Data []employee.Employee
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
		assert.Equal(t, jsonResponse.Data, employees)
		assert.Nil(t, err)
		assert.True(t, len(jsonResponse.Data) > 0)
	})
	t.Run("If request GetAll has an error, it should return status code 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("error to get all employees")
		req, rr := createRequestTest(http.MethodGet, "/api/v1/employees/", "")
		serviceMock := &mocks.Service{}
		serviceMock.On("GetAll").Return([]employee.Employee{}, errorMsg)

		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees/")
		employeeGroup.GET("/", e.GetAll())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		jsonResponse := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)
		assert.Equal(t, jsonResponse.Error, errorMsg.Error())
	})
}
func TestHandlerGetByID(t *testing.T) {

	t.Run("If request GetByID is OK, it should return status code 200 and an employee", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodGet, "/api/v1/employees/1", "")
		serviceMock := &mocks.Service{}
		serviceMock.On("GetByID", 1).Return(employee1, nil)

		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees/")
		employeeGroup.GET("/:id", e.GetByID())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)

		jsonResponse := struct {
			Code int
			Data employee.Employee
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)
		assert.Equal(t, jsonResponse.Data, employee1)

	})
	t.Run("If user is not passing a number in the parameter of the url to GetByID, it should return an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("invalid ID")
		req, rr := createRequestTest(http.MethodGet, "/api/v1/employees/a", "")
		serviceMock := &mocks.Service{}
		serviceMock.On("GetByID").Return(employee.Employee{}, errorMsg)

		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees/")
		employeeGroup.GET("/:id", e.GetByID())
		r.ServeHTTP(rr, req)

		jsonResponse := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)
		assert.Equal(t, jsonResponse.Error, errorMsg.Error())
	})
	t.Run("If request GetByID does not find a employee, it should return 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("employee with id 10 not found")
		req, rr := createRequestTest(http.MethodGet, "/api/v1/employees/10", "")
		serviceMock := &mocks.Service{}
		serviceMock.On("GetByID", 10).Return(employee.Employee{}, errorMsg)

		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees/")
		employeeGroup.GET("/:id", e.GetByID())
		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		jsonResponse := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)
		assert.Equal(t, jsonResponse.Error, errorMsg.Error())
	})
}
func TestHandlerCreate(t *testing.T) {

	t.Run("If request Create is OK, it should return status code 201 and an employee", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodPost, "/api/v1/employees/",
			`
		{
		"card_number_id": "123456",
		"first_name": "Marta",
		"last_name" : "Gomes",
		"warehouse_id": 3
  		}
			`)
		serviceMock := &mocks.Service{}
		serviceMock.On("Create",
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(newEmployee, nil)

		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees/")
		employeeGroup.POST("/", e.Create())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 201, rr.Code)

		jsonResponse := struct {
			Code int
			Data employee.Employee
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)
		assert.Equal(t, jsonResponse.Data, newEmployee)

	})
	t.Run("If user is not passing the correct body to the Create , it should return status code 422 and an error", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodPost, "/api/v1/employees/",
			`{
		"card_number_id": "123456",
		"last_name" : "Gomes",
		"warehouse_id": 3
  		}`)
		serviceMock := &mocks.Service{}
		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees")
		employeeGroup.POST("/", e.Create())
		r.ServeHTTP(rr, req)
		assert.Equal(t, 422, rr.Code)

		jsonResponse := struct {
			Code  int
			Error []struct {
				Field   string
				Message string
			}
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)
		assert.Equal(t, "This field is required", jsonResponse.Error[0].Message)
		assert.Equal(t, "FirstName", jsonResponse.Error[0].Field)
	})
	t.Run("If there is an error to Create, it should return status code 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("error to create employee")
		req, rr := createRequestTest(http.MethodPost, "/api/v1/employees/",
			`
	{
	"card_number_id": "123456",
	"first_name": "Marta",
	"last_name" : "Gomes",
	"warehouse_id": 3
	  }
		`)
		serviceMock := &mocks.Service{}
		serviceMock.On("Create",
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(employee.Employee{}, errorMsg)

		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees")
		employeeGroup.POST("/", e.Create())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		jsonResponse := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)
		assert.Equal(t, jsonResponse.Error, errorMsg.Error())
	})

}
