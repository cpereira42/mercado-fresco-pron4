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

var employeeResponseJson = struct {
	Code  int
	Data  employee.Employee
	Error string
}{}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func createServer(serv *mocks.Service, method string, url string, body string) *httptest.ResponseRecorder {
	e := handler.NewEmployee(serv)
	r := gin.Default()
	pr := r.Group("/api/v1/employees")
	pr.GET("/", e.GetAll())
	pr.GET("/:id", e.GetByID())
	pr.DELETE("/:id", e.Delete())
	pr.POST("/", e.Create())
	pr.PATCH("/:id", e.Update())
	req, rr := createRequestTest(method, url, body)
	r.ServeHTTP(rr, req)
	return rr
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
			`{
		"card_number_id": "123456",
		"first_name": "Marta",
		"last_name" : "Gomes",
		"warehouse_id": 3
			}`)
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

func TestHandlerUpdate(t *testing.T) {
	t.Run("If request Update is OK, it should return status code 200 and an employee", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodPatch, "/api/v1/employees/1",
			`{
			"card_number_id": "123",
			"first_name": "Gustavo",
			"last_name" : "Junior",
			"warehouse_id": 1
			}`)
		serviceMock := &mocks.Service{}
		serviceMock.On("Update",
			tmock.AnythingOfType("int"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(employee1Update, nil)

		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees")
		employeeGroup.PATCH("/:id", e.Update())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)

		jsonResponse := struct {
			Code int
			Data employee.Employee
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)
		assert.Equal(t, jsonResponse.Data, employee1Update)

	})
	t.Run("If user is not passing a number in the parameter of the url to Update, it should return an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("invalid ID")
		req, rr := createRequestTest(http.MethodPatch, "/api/v1/employees/a",
			`{
			"card_number_id": "123",
			"first_name": "Gustavo",
			"last_name" : "Junior",
			"warehouse_id": 1
			}`)
		serviceMock := &mocks.Service{}
		serviceMock.On("Update",
			tmock.AnythingOfType("int"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(employee.Employee{}, errorMsg)

		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees")
		employeeGroup.PATCH("/:id", e.Update())

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
	t.Run("If invalid JSON in the request Update, it should return status code 422 and an error", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodPatch, "/api/v1/employees/1",
			`{
		"card_number_id": "123456"
		"last_name" : "Gomes",
		"warehouse_id": 3
  		}`)
		serviceMock := &mocks.Service{}
		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees")
		employeeGroup.PATCH("/:id", e.Update())
		r.ServeHTTP(rr, req)
		assert.Equal(t, 422, rr.Code)

		jsonResponse := struct {
			Code  int
			Error string
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)
		assert.Equal(t, "invalid character '\"' after object key:value pair", jsonResponse.Error)
	})
	t.Run("If there is an error to Update, it should return status code 404 and an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("error to update employee")
		req, rr := createRequestTest(http.MethodPatch, "/api/v1/employees/1",
			`{
		"card_number_id": "123456",
		"first_name": "Marta",
		"last_name" : "Gomes",
		"warehouse_id": 3
			}`)
		serviceMock := &mocks.Service{}
		serviceMock.On("Update",
			tmock.AnythingOfType("int"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int"),
		).Return(employee.Employee{}, errorMsg)

		e := handler.NewEmployee(serviceMock)
		r := gin.Default()
		employeeGroup := r.Group("/api/v1/employees")
		employeeGroup.PATCH("/:id", e.Update())

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

func TestHandlerDelete(t *testing.T) {
	t.Run("If request Delete is OK, it should return status code 204", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodDelete, "/api/v1/employees/1", "")
		service := &mocks.Service{}
		e := handler.NewEmployee(service)
		r := gin.Default()

		service.On("Delete", 1).Return(nil)
		employeeGroup := r.Group("/api/v1/employees")
		employeeGroup.DELETE("/:id", e.Delete())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 204, rr.Code)
	})
	t.Run(
		"If user is not passing a number in the parameter of the url to GetByID, it should return an error", func(t *testing.T) {
			errorMsg := fmt.Errorf("invalid ID")
			req, rr := createRequestTest(http.MethodDelete, "/api/v1/employees/a", "")
			service := &mocks.Service{}
			e := handler.NewEmployee(service)
			r := gin.Default()
			employeeGroup := r.Group("/api/v1/employees")
			employeeGroup.DELETE("/:id", e.Delete())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 404, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, errorMsg.Error(), objResp.Error)
		})
	t.Run(
		"If the request Delete has an invalid ID, it should return status code  404 and an error", func(t *testing.T) {
			errorMsg := fmt.Errorf("employee with id 10 not found")
			serv := &mocks.Service{}
			serv.On("Delete", 10).Return(errorMsg)

			rr := createServer(serv, http.MethodDelete, "/api/v1/employees/10", "")
			assert.Equal(t, 404, rr.Code)

			err := json.Unmarshal(rr.Body.Bytes(), &employeeResponseJson)
			assert.Nil(t, err)
			assert.NotNil(t, employeeResponseJson.Error)
			assert.Equal(t, errorMsg.Error(), employeeResponseJson.Error)

		})
}
