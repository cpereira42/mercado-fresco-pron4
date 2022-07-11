package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee/mocks"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
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
var employeesResponseJson = struct {
	Code int
	Data []employee.Employee
}{}

var employeeResponseJson = struct {
	Code  int
	Data  employee.Employee
	Error string
}{}

var employeeFieldsResponseJson = struct {
	Code int
	Data []struct {
		Field   string
		Message string
	}
}{}

func createServerEmployee(serverMock *mocks.Service, method string, url string, body string) *httptest.ResponseRecorder {
	r := gin.Default()
	handler.NewEmployee(r, serverMock)
	req, rr := util.CreateRequestTest(method, url, body)
	r.ServeHTTP(rr, req)
	return rr
}

func TestHandlerGetAll(t *testing.T) {
	employees := []employee.Employee{employee1, employee2, employee3}

	t.Run("If request GetAll is OK, it should return status code 200 and a list of employees", func(t *testing.T) {
		serviceMock := &mocks.Service{}
		serviceMock.On("GetAll").Return(employees, nil)
		rr := createServerEmployee(serviceMock, http.MethodGet, "/api/v1/employees/", "")

		err := json.Unmarshal(rr.Body.Bytes(), &employeesResponseJson)

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, employeesResponseJson.Data, employees)
		assert.Nil(t, err)
		assert.True(t, len(employeesResponseJson.Data) > 0)
	})
	t.Run("If request GetAll has an error, it should return status code 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("error to get all employees")
		serviceMock := &mocks.Service{}
		serviceMock.On("GetAll").Return([]employee.Employee{}, errorMsg)
		rr := createServerEmployee(serviceMock, http.MethodGet, "/api/v1/employees/", "")
		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 404, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), employeeResponseJson.Error)
	})
}
func TestHandlerGetByID(t *testing.T) {

	t.Run("If request GetByID is OK, it should return status code 200 and an employee", func(t *testing.T) {
		serviceMock := &mocks.Service{}
		serviceMock.On("GetByID", 1).Return(employee1, nil)
		rr := createServerEmployee(serviceMock, http.MethodGet, "/api/v1/employees/1", "")

		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 200, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, employee1, employeeResponseJson.Data)

	})
	t.Run("If user is not passing a number in the parameter of the url to GetByID, it should return status code 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("invalid ID")
		serviceMock := &mocks.Service{}
		serviceMock.On("GetByID").Return(employee.Employee{}, errorMsg)
		rr := createServerEmployee(serviceMock, http.MethodGet, "/api/v1/employees/a", "")

		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 404, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), employeeResponseJson.Error)
	})
	t.Run("If request GetByID does not find a employee, it should return 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("employee with id 10 not found")
		serviceMock := &mocks.Service{}
		serviceMock.On("GetByID", 10).Return(employee.Employee{}, errorMsg)
		rr := createServerEmployee(serviceMock, http.MethodGet, "/api/v1/employees/10", "")

		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 404, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), employeeResponseJson.Error)
	})
}
func TestHandlerCreate(t *testing.T) {

	t.Run("If request Create is OK, it should return status code 201 and an employee", func(t *testing.T) {
		serviceMock := &mocks.Service{}
		serviceMock.On("Create",
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(newEmployee, nil)
		rr := createServerEmployee(serviceMock, http.MethodPost, "/api/v1/employees/",
			`
		{
		"card_number_id": "123456",
		"first_name": "Marta",
		"last_name" : "Gomes",
		"warehouse_id": 3
  		}
			`)

		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 201, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, newEmployee, employeeResponseJson.Data)

	})
	t.Run("If user is not passing the correct body to the Create , it should return status code 422 and an error", func(t *testing.T) {
		serviceMock := &mocks.Service{}
		rr := createServerEmployee(serviceMock, http.MethodPost, "/api/v1/employees/",
			`{
		"card_number_id": "123456",
		"last_name" : "Gomes",
		"warehouse_id": 3
  		}`)

		assert.Equal(t, 422, rr.Code)

		err := json.Unmarshal(rr.Body.Bytes(), &employeeFieldsResponseJson)
		assert.Nil(t, err)
		assert.Equal(t, "This field is required", employeeFieldsResponseJson.Data[0].Message)
		assert.Equal(t, "first_name", employeeFieldsResponseJson.Data[0].Field)
	})
	t.Run("If there is an error to Create, it should return status code 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("error to create employee")
		serviceMock := &mocks.Service{}
		serviceMock.On("Create",
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(employee.Employee{}, errorMsg)
		rr := createServerEmployee(serviceMock, http.MethodPost, "/api/v1/employees/",
			`{
		"card_number_id": "123456",
		"first_name": "Marta",
		"last_name" : "Gomes",
		"warehouse_id": 3
			}`)

		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 404, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), employeeResponseJson.Error)
	})
	t.Run("If the CardNumberID already exists  is the Create request , it should return status code 409 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("employee with this card number id 123 exists")
		serviceMock := &mocks.Service{}
		serviceMock.On("Create",
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(employee.Employee{}, errorMsg)
		rr := createServerEmployee(serviceMock, http.MethodPost, "/api/v1/employees/",
			`{
		"card_number_id": "123",
		"first_name": "Gustavo",
		"last_name" : "Junior",
		"warehouse_id": 1
			}`)

		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 409, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), employeeResponseJson.Error)
	})

}

func TestHandlerUpdate(t *testing.T) {
	t.Run("If request Update is OK, it should return status code 200 and an employee", func(t *testing.T) {
		serviceMock := &mocks.Service{}
		serviceMock.On("Update",
			tmock.AnythingOfType("int"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(employee1Update, nil)
		rr := createServerEmployee(serviceMock, http.MethodPatch, "/api/v1/employees/1",
			`{
			"card_number_id": "123",
			"first_name": "Gustavo",
			"last_name" : "Junior",
			"warehouse_id": 1
			}`)

		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 200, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, employee1Update, employeeResponseJson.Data)

	})
	t.Run("If user is not passing a number in the parameter of the url to Update, it should return an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("invalid ID")
		serviceMock := &mocks.Service{}
		serviceMock.On("Update",
			tmock.AnythingOfType("int"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(employee.Employee{}, errorMsg)
		rr := createServerEmployee(serviceMock, http.MethodPatch, "/api/v1/employees/a",
			`{
			"card_number_id": "123",
			"first_name": "Gustavo",
			"last_name" : "Junior",
			"warehouse_id": 1
			}`)

		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 404, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), employeeResponseJson.Error)

	})
	t.Run("If invalid JSON in the request Update, it should return status code 422 and an error", func(t *testing.T) {
		serviceMock := &mocks.Service{}
		rr := createServerEmployee(serviceMock, http.MethodPatch, "/api/v1/employees/1",
			`{
		"card_number_id": "123456"
		"last_name" : "Gomes",
		"warehouse_id": 3
  		}`)

		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 422, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, "invalid character '\"' after object key:value pair", employeeResponseJson.Error)
	})
	t.Run("If there is an error to Update, it should return status code 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("error to update employee")
		serviceMock := &mocks.Service{}
		serviceMock.On("Update",
			tmock.AnythingOfType("int"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(employee.Employee{}, errorMsg)
		rr := createServerEmployee(serviceMock, http.MethodPatch, "/api/v1/employees/1",
			`{
		"card_number_id": "123456",
		"first_name": "Marta",
		"last_name" : "Gomes",
		"warehouse_id": 3
			}`)

		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 404, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), employeeResponseJson.Error)
	})
	t.Run("If the CardNumberID already exists  is the Update request , it should return status code 409 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("employee with this card number id 123 exists")
		serviceMock := &mocks.Service{}
		serviceMock.On("Update", 2,
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(employee.Employee{}, errorMsg)
		rr := createServerEmployee(serviceMock, http.MethodPatch, "/api/v1/employees/2",
			`{
		"card_number_id": "123",
		"first_name": "Gustavo",
		"last_name" : "Junior",
		"warehouse_id": 1
			}`)

		err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

		assert.Equal(t, 409, rr.Code)
		assert.Nil(t, err)
		assert.Equal(t, errorMsg.Error(), employeeResponseJson.Error)
	})
}

func TestHandlerDelete(t *testing.T) {
	t.Run("If request Delete is OK, it should return status code 204", func(t *testing.T) {
		serviceMock := &mocks.Service{}
		serviceMock.On("Delete", 1).Return(nil)
		rr := createServerEmployee(serviceMock, http.MethodDelete, "/api/v1/employees/1", "")

		assert.Equal(t, 204, rr.Code)
	})
	t.Run(
		"If user is not passing a number in the parameter of the url to GetByID, it should return status code 404 and  an error", func(t *testing.T) {
			errorMsg := fmt.Errorf("invalid ID")
			serviceMock := &mocks.Service{}
			rr := createServerEmployee(serviceMock, http.MethodDelete, "/api/v1/employees/a", "")

			err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

			assert.Nil(t, err)
			assert.Equal(t, 404, rr.Code)
			assert.Equal(t, errorMsg.Error(), employeeResponseJson.Error)
		})
	t.Run(
		"If the request Delete has an invalid ID, it should return status code  404 and an error", func(t *testing.T) {
			errorMsg := fmt.Errorf("employee with id 10 not found")
			serviceMock := &mocks.Service{}
			serviceMock.On("Delete", 10).Return(errorMsg)

			rr := createServerEmployee(serviceMock, http.MethodDelete, "/api/v1/employees/10", "")

			err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)

			assert.Equal(t, 404, rr.Code)
			assert.Nil(t, err)
			assert.NotNil(t, employeeResponseJson.Error)
			assert.Equal(t, errorMsg.Error(), employeeResponseJson.Error)

		})
}
