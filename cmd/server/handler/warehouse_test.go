package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse/mocks"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var (
	warehouse1                              = warehouse.Warehouse{ID: 1, Address: "Rua 1", Telephone: "11111-1111", Warehouse_code: "W1", Minimum_capacity: 10, Minimum_temperature: 20, Locality_id: 1}
	warehouse2                              = warehouse.Warehouse{ID: 2, Address: "Rua 2", Telephone: "22222-2222", Warehouse_code: "W2", Minimum_capacity: 20, Minimum_temperature: 30, Locality_id: 2}
	warehouse3                              = warehouse.Warehouse{ID: 3, Address: "Rua 3", Telephone: "33333-3333", Warehouse_code: "W3", Minimum_capacity: 30, Minimum_temperature: 40, Locality_id: 3}
	warehouse1Updated warehouse.Warehouse   = warehouse.Warehouse{ID: 1, Address: "Rua 4", Telephone: "11111-1111", Warehouse_code: "W1", Minimum_capacity: 10, Minimum_temperature: 20, Locality_id: 1}
	warehouse4        warehouse.Warehouse   = warehouse.Warehouse{ID: 4, Address: "Rua 4", Telephone: "11111-1111", Warehouse_code: "W4", Minimum_capacity: 10, Minimum_temperature: 20, Locality_id: 1}
	warehouseList     []warehouse.Warehouse = []warehouse.Warehouse{warehouse1, warehouse2, warehouse3}
)

func createRequestTests(serv *mocks.Service, method string, url string, body string) *httptest.ResponseRecorder {

	r := gin.Default()
	handler.NewWarehouse(r, serv)
	req, rr := util.CreateRequestTest(method, url, body)
	r.ServeHTTP(rr, req)
	return rr
}

func TestControllerGetAll(t *testing.T) {
	warehouseList := []warehouse.Warehouse{warehouse1, warehouse2, warehouse3}
	t.Run(
		"should return all warehouses", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("GetAll").Return(warehouseList, nil)
			rr := createRequestTests(serviceMock, http.MethodGet, "/api/v1/warehouse/", "")
			objResp := struct {
				Code int
				Data []warehouse.Warehouse
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Equal(t, 200, rr.Code)
			assert.Nil(t, err)
			assert.Equal(t, warehouseList, objResp.Data)
		})
	t.Run(
		"should return error 500", func(t *testing.T) {
			msgError := "Could not read file"
			serviceMock := new(mocks.Service)
			serviceMock.On("GetAll").Return([]warehouse.Warehouse{}, errors.New(msgError))
			rr := createRequestTests(serviceMock, http.MethodGet, "/api/v1/warehouse/", "")
			objResp := struct {
				Code  int
				Error string
			}{}
			fmt.Println(rr.Body.Bytes())
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, msgError, objResp.Error)
		})
}

func TestControllerUpdate(t *testing.T) {
	t.Run(
		"Update OK - should updated warehouse", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse1Updated, nil)
			rr := createRequestTests(serviceMock, http.MethodPatch, "/api/v1/warehouse/1",
				`{
				"address": "Rua 4",
				"telephone": "11111111",
				"warehouse_code": "W1",
				"minimum_capacity": 10,
				"minimum_temperature": 20,
				"locality_id": 1
			}`)
			objResp := struct {
				Code int
				Data warehouse.Warehouse
			}{}

			err := json.Unmarshal(rr.Body.Bytes(), &objResp)

			assert.Equal(t, 200, rr.Code)
			assert.Nil(t, err)
			assert.ObjectsAreEqual(warehouse1Updated, objResp.Data)
		})
	t.Run(
		"Test Update - ID Invalid - should return 404", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			rr := createRequestTests(serviceMock, http.MethodPatch, "/api/v1/warehouse/a",
				`{
				"address": "Rua 4",
				"telephone": "11111111",
				"warehouse_code": "W1",
				"minimum_capacity": 10,
				"minimum_temperature": 20,
				"locality_id": 1
			}`)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Equal(t, 404, rr.Code)
			assert.Nil(t, err)
			assert.Equal(t, "Warehouse Not Found", objResp.Error)
		})
	t.Run(
		"Test Update - Invalid JSON", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse.Warehouse{}, errors.New("invalid character '\"' after object key:value pair"))
			rr := createRequestTests(serviceMock, http.MethodPatch, "/api/v1/warehouse/1",
				`{
					"address": "Rua 4",
					"telephone": "11111111",
					"warehouse_code": "W1",
					"minimum_capacity": 10,
					"minimum_temperature": 20
				}`)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)

			assert.Equal(t, 404, rr.Code)
			assert.Nil(t, err)
			assert.Equal(t, "invalid character '\"' after object key:value pair", objResp.Error)
		})
}

func TestControllerDelete(t *testing.T) {
	t.Run(
		"Test Delete - should return 204", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("Delete", 1).Return(nil)

			rr := createRequestTests(serviceMock, http.MethodDelete, "/api/v1/warehouse/1", "")
			assert.Equal(t, 204, rr.Code)
		})
	t.Run(
		"Test Delete - should return 404", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			rr := createRequestTests(serviceMock, http.MethodDelete, "/api/v1/warehouse/a", "")
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Equal(t, 404, rr.Code)
			assert.Nil(t, err)
			assert.Equal(t, "Invalid ID", objResp.Error)
		})
	t.Run(
		"Test Delete - ID not found", func(t *testing.T) {
			msgError := fmt.Errorf("Warehouse Not Found")
			serviceMock := new(mocks.Service)
			serviceMock.On("Delete", 5).Return(msgError)
			rr := createRequestTests(serviceMock, http.MethodDelete, "/api/v1/warehouse/5", "")
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Equal(t, 404, rr.Code)
			assert.Nil(t, err)
			assert.Equal(t, "Warehouse Not Found", objResp.Error)
		})

}

func TestControllerGetByID(t *testing.T) {
	t.Run(
		"Test GetByID - should return 200", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("GetByID", 1).Return(warehouse1, nil)
			rr := createRequestTests(serviceMock, http.MethodGet, "/api/v1/warehouse/1", "")
			objResp := struct {
				Code int
				Data warehouse.Warehouse
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Equal(t, 200, rr.Code)
			assert.Nil(t, err)
			assert.Equal(t, warehouse1, objResp.Data)
		})
	t.Run(
		"Test GetByID - ID not found - should return 404", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			msgError := fmt.Errorf("Warehouse Not Found")
			serviceMock.On("GetByID", 5).Return(warehouse.Warehouse{}, msgError)
			rr := createRequestTests(serviceMock, http.MethodGet, "/api/v1/warehouse/5", "")
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Equal(t, 404, rr.Code)
			assert.Nil(t, err)
			assert.Equal(t, "Warehouse Not Found", objResp.Error)
		})
}

func TestControllerCreate(t *testing.T) {
	t.Run(
		"Test Create - should return 201", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse1, nil)
			rr := createRequestTests(serviceMock, http.MethodPost, "/api/v1/warehouse/",
				`{
					"address": "Rua 1",
					"telephone": "11111111",
					"warehouse_code": "W1",
					"minimum_capacity": 10,
					"minimum_temperature": 20,
					"locality_id": 1
				}`)
			objResp := struct {
				Code int
				Data warehouse.Warehouse
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Equal(t, 201, rr.Code)
			assert.Nil(t, err)
			assert.Equal(t, warehouse1, objResp.Data)
		})

	t.Run(
		"Test Create - should return 409", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			errorMsg := fmt.Errorf("Warehouse already exists")
			serviceMock.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse.Warehouse{}, errorMsg)
			rr := createRequestTests(serviceMock, http.MethodPost, "/api/v1/warehouse/",
				`{
					"address": "Rua 1",
					"telephone": "11111111",
					"warehouse_code": "W1",
					"minimum_capacity": 10,
					"minimum_temperature": 20,
					"locality_id": 1
				}`)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Equal(t, 409, rr.Code)
			assert.Nil(t, err)
			assert.Equal(t, "Warehouse already exists", objResp.Error)
		})
	t.Run(
		"Test Create - Error to save", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			errorMsg := fmt.Errorf("Error to save")
			serviceMock.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse.Warehouse{}, errorMsg)
			rr := createRequestTests(serviceMock, http.MethodPost, "/api/v1/warehouse/",
				`{
					"address": "Rua 1",
					"telephone": "11111111",
					"warehouse_code": "W1",
					"minimum_capacity": 10,
					"minimum_temperature": 20,
					"locality_id": 1
				}`)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Equal(t, 422, rr.Code)
			assert.Nil(t, err)
			assert.Equal(t, "Error to save", objResp.Error)
		})
}
