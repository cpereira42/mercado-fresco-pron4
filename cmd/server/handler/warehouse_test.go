package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var (
	warehouse1                              = warehouse.Warehouse{ID: 1, Address: "Rua 1", Telephone: "11111-1111", Warehouse_code: "W1", Minimum_capacity: 10, Minimum_temperature: 20}
	warehouse2                              = warehouse.Warehouse{ID: 2, Address: "Rua 2", Telephone: "22222-2222", Warehouse_code: "W2", Minimum_capacity: 20, Minimum_temperature: 30}
	warehouse3                              = warehouse.Warehouse{ID: 3, Address: "Rua 3", Telephone: "33333-3333", Warehouse_code: "W3", Minimum_capacity: 30, Minimum_temperature: 40}
	warehouse1Updated warehouse.Warehouse   = warehouse.Warehouse{ID: 1, Address: "Rua 4", Telephone: "11111-1111", Warehouse_code: "W1", Minimum_capacity: 10, Minimum_temperature: 20}
	warehouse4        warehouse.Warehouse   = warehouse.Warehouse{ID: 4, Address: "Rua 4", Telephone: "11111-1111", Warehouse_code: "W4", Minimum_capacity: 10, Minimum_temperature: 20}
	warehouseList     []warehouse.Warehouse = []warehouse.Warehouse{warehouse1, warehouse2, warehouse3}
)

func createRequestTests(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestControllerGetAll(t *testing.T) {
	t.Run(
		"should return all warehouses", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodGet, "/api/v1/warehouse/", "")
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			serviceMock.On("GetAll").Return(warehouseList, nil)
			wr := r.Group("/api/v1/warehouse")
			wr.GET("/", w.GetAll)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 200, rr.Code)
			objResp := struct {
				Code int
				Data []warehouse.Warehouse
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, warehouseList, objResp.Data)
		})
	t.Run(
		"should return error 500", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodGet, "/api/v1/warehouse/", "")
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			msgError := "Could not read file"
			serviceMock.On("GetAll").Return([]warehouse.Warehouse{}, fmt.Errorf(msgError))
			wr := r.Group("/api/v1/warehouse")
			wr.GET("/", w.GetAll)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 500, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, msgError, objResp.Error)
		})
}
func TestControllerUpdate(t *testing.T) {
	t.Run(
		"should updated warehouse", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodPatch, "/api/v1/warehouse/1",
				`{
				"address": "Rua 4",
				"telephone": "11111111",
				"warehouse_code": "W1",
				"minimum_capacity": 10,
				"minimum_temperature": 20
			}`)
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			serviceMock.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse1Updated, nil)
			wr := r.Group("/api/v1/warehouse")
			wr.PATCH("/:id", w.Update)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 200, rr.Code)
			objResp := struct {
				Code int
				Data warehouse.Warehouse
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.ObjectsAreEqual(warehouse1Updated, objResp.Data)
		})
	t.Run(
		"Test Update - should return 404", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodPatch, "/api/v1/warehouse/a",
				`{
				"address": "Rua 4",
				"telephone": "11111111",
				"warehouse_code": "W1",
				"minimum_capacity": 10,
				"minimum_temperature": 20
			}`)
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			wr := r.Group("/api/v1/warehouse")
			wr.PATCH("/:id", w.Update)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 404, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "Warehouse Not Found", objResp.Error)
		})
	t.Run(
		"Test Update - Invalid JSON", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodPatch, "/api/v1/warehouse/1",
				`{
				"telephone": "11111111"
				"warehouse_code": "W1",
				"minimum_capacity": 10,
				"minimum_temperature": 20
			}`)
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			wr := r.Group("/api/v1/warehouse")
			wr.PATCH("/:id", w.Update)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 422, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			log.Println(rr.Body.Bytes())
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "invalid character '\"' after object key:value pair", objResp.Error)
		})
	t.Run(
		"Test Update - id not found", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodPatch, "/api/v1/warehouse/1",
				`{
				"address": "Rua 4",
				"telephone": "11111111",
				"warehouse_code": "W1",
				"minimum_capacity": 10,
				"minimum_temperature": 20
			}`)
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			msgError := fmt.Errorf("Warehouse Not Found")
			serviceMock.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse.Warehouse{}, msgError)
			wr := r.Group("/api/v1/warehouse")
			wr.PATCH("/:id", w.Update)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 404, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "Warehouse Not Found", objResp.Error)
		})

}

func TestControllerDelete(t *testing.T) {
	t.Run(
		"Test Delete - should return 204", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodDelete, "/api/v1/warehouse/1", "")
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			serviceMock.On("Delete", 1).Return(nil)
			wr := r.Group("/api/v1/warehouse")
			wr.DELETE("/:id", w.Delete)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 204, rr.Code)
		})
	t.Run(
		"Test Delete - should return 404", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodDelete, "/api/v1/warehouse/a", "")
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			wr := r.Group("/api/v1/warehouse")
			wr.DELETE("/:id", w.Delete)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 404, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "Invalid ID", objResp.Error)
		})
	t.Run(
		"Test Delete - ID not found", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodDelete, "/api/v1/warehouse/5", "")
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			msgError := fmt.Errorf("Warehouse Not Found")
			serviceMock.On("Delete", 5).Return(msgError)
			wr := r.Group("/api/v1/warehouse")
			wr.DELETE("/:id", w.Delete)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 404, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "Warehouse Not Found", objResp.Error)
		})

}

func TestControllerGet(t *testing.T) {
	t.Run(
		"Test GetByID - should return 200", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodGet, "/api/v1/warehouse/1", "")
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			serviceMock.On("GetByID", 1).Return(warehouse1, nil)
			wr := r.Group("/api/v1/warehouse")
			wr.GET("/:id", w.GetByID)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 200, rr.Code)
			objResp := struct {
				Code int
				Data warehouse.Warehouse
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, warehouse1, objResp.Data)
		})
	t.Run(
		"Test GetByID - Invalid ID - should return 404", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodGet, "/api/v1/warehouse/a", "")
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			wr := r.Group("/api/v1/warehouse")
			wr.GET("/:id", w.GetByID)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 404, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "Invalid ID", objResp.Error)
		})

	t.Run(
		"Test GetByID - ID not found - should return 404", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodGet, "/api/v1/warehouse/5", "")
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			msgError := fmt.Errorf("Warehouse Not Found")
			serviceMock.On("GetByID", 5).Return(warehouse.Warehouse{}, msgError)
			wr := r.Group("/api/v1/warehouse")
			wr.GET("/:id", w.GetByID)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 404, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "Warehouse Not Found", objResp.Error)
		})

}

func TestControllerCreate(t *testing.T) {
	t.Run(
		"Test Create - should return 201", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodPost, "/api/v1/warehouse/",
				`{
				"address": "Rua 1",
				"telephone": "11111111",
				"warehouse_code": "W1",
				"minimum_capacity": 10,
				"minimum_temperature": 20
			}`)
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			serviceMock.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse1, nil)
			wr := r.Group("/api/v1/warehouse")
			wr.POST("/", w.Create)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 201, rr.Code)
			objResp := struct {
				Code int
				Data warehouse.Warehouse
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, warehouse1, objResp.Data)
		})

	t.Run(
		"Test Create - Request Body error - without Telephone", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodPost, "/api/v1/warehouse/",
				`{
				"address": "Rua 1",
				"warehouse_code": "W1",
				"minimum_capacity": 10,
				"minimum_temperature": 20
			}`)
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			wr := r.Group("/api/v1/warehouse")
			wr.POST("/", w.Create)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 422, rr.Code)
			objResp := struct {
				Code int
				Data []struct {
					Field   string
					Message string
				}
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "telephone", objResp.Data[0].Field)
		})

	t.Run(
		"Test Create - should return 409", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodPost, "/api/v1/warehouse/",
				`{
				"address": "Rua 1",
				"telephone": "11111111",
				"warehouse_code": "W1",
				"minimum_capacity": 10,
				"minimum_temperature": 20
			}`)
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			msgError := fmt.Errorf("Warehouse already exists")
			serviceMock.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse.Warehouse{}, msgError)
			wr := r.Group("/api/v1/warehouse")
			wr.POST("/", w.Create)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 409, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "Warehouse already exists", objResp.Error)
		})
	t.Run(
		"Test Create - Error to save", func(t *testing.T) {
			req, rr := createRequestTests(http.MethodPost, "/api/v1/warehouse/",
				`{
				"address": "Rua 1",
				"telephone": "11111111",
				"warehouse_code": "W1",
				"minimum_capacity": 10,
				"minimum_temperature": 20
			}`)
			serviceMock := new(mocks.Service)
			w := handler.NewWarehouse(serviceMock)
			r := gin.Default()
			msgError := fmt.Errorf("Error to save")
			serviceMock.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse.Warehouse{}, msgError)
			wr := r.Group("/api/v1/warehouse")
			wr.POST("/", w.Create)
			r.ServeHTTP(rr, req)
			assert.Equal(t, 422, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "Error to save", objResp.Error)
		})

}
