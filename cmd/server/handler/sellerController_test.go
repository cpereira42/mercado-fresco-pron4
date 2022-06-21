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
	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"
	"github.com/cpereira42/mercado-fresco-pron4/internal/seller/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var seller1 seller.Seller = seller.Seller{Id: 1, Cid: 200, CompanyName: "MELI", Adress: "Rua B", Telephone: "9999-8888"}

var sListSuccess []seller.Seller = []seller.Seller{
	seller1,
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestControllerGetAll(t *testing.T) {
	t.Run(
		"Test GetAll - OK", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodGet, "/api/v1/sellers/", "")
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			serviceMock.On("GetAll").Return(sListSuccess, nil)
			sellers := r.Group("/api/v1/sellers")
			sellers.GET("/", s.GetAll())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 200, rr.Code)
			objResp := struct {
				Code int
				Data []seller.Seller
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, sListSuccess, objResp.Data)
		})
	t.Run(
		"Test GetAll - Error - Could not read file", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodGet, "/api/v1/sellers/", "")
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			msgError := "Could not read file"
			serviceMock.On("GetAll").Return([]seller.Seller{}, fmt.Errorf(msgError))
			sellers := r.Group("/api/v1/sellers")
			sellers.GET("/", s.GetAll())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 400, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, msgError, objResp.Error)
		})
	t.Run(
		"Test GetAll - Error - Sellers length == 0", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodGet, "/api/v1/sellers/", "")
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			msgError := "Sellers is empty"
			serviceMock.On("GetAll").Return([]seller.Seller{}, nil)
			sellers := r.Group("/api/v1/sellers")
			sellers.GET("/", s.GetAll())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 400, rr.Code)
			log.Println(string(rr.Body.Bytes()))
			objResp := struct {
				Code  int
				Data  []seller.Seller
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, msgError, objResp.Error)
		})
}

func TestControllerDelete(t *testing.T) {
	t.Run(
		"Test Delete - OK", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodDelete, "/api/v1/sellers/1", "")
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			serviceMock.On("Delete", 1).Return(nil)
			sellers := r.Group("/api/v1/sellers")
			sellers.DELETE("/:id", s.Delete())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 204, rr.Code)
		})
	t.Run(
		"Test Delete - Error - Invalid ID", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodDelete, "/api/v1/sellers/a", "")
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			sellers := r.Group("/api/v1/sellers")
			sellers.DELETE("/:id", s.Delete())
			r.ServeHTTP(rr, req)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, 400, rr.Code)
			assert.Equal(t, "invalid ID", objResp.Error)
		})
	t.Run(
		"Test Delete - Error - ID not found", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodDelete, "/api/v1/sellers/2", "")
			serviceMock := new(mocks.Service)
			serviceMock.On("Delete", 2).Return(fmt.Errorf("Seller 2 not found"))
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			sellers := r.Group("/api/v1/sellers")
			sellers.DELETE("/:id", s.Delete())
			r.ServeHTTP(rr, req)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, 404, rr.Code)
			assert.Equal(t, "Seller 2 not found", objResp.Error)
		})
}

func TestControllerGetId(t *testing.T) {
	t.Run(
		"Test GetId - OK", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodGet, "/api/v1/sellers/1", "")
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			serviceMock.On("GetId", 1).Return(seller1, nil)
			sellers := r.Group("/api/v1/sellers")
			sellers.GET("/:id", s.GetId())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 200, rr.Code)
			objResp := struct {
				Code int
				Data seller.Seller
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, seller1, objResp.Data)
		})
	t.Run(
		"Test GetId - Error - Invalid ID", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodGet, "/api/v1/sellers/a", "")
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			sellers := r.Group("/api/v1/sellers")
			sellers.GET("/:id", s.GetId())
			r.ServeHTTP(rr, req)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, 400, rr.Code)
			assert.Equal(t, "Invalid ID", objResp.Error)
		})
	t.Run(
		"Test GetId - Error - ID not found", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodGet, "/api/v1/sellers/2", "")
			serviceMock := new(mocks.Service)
			serviceMock.On("GetId", 2).Return(seller.Seller{}, fmt.Errorf("Seller 2 not found"))
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			sellers := r.Group("/api/v1/sellers")
			sellers.GET("/:id", s.GetId())
			r.ServeHTTP(rr, req)
			objResp := struct {
				Code  int
				Data  seller.Seller
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, 404, rr.Code)
			assert.Equal(t, "Seller 2 not found", objResp.Error)
			assert.Equal(t, seller.Seller{}, objResp.Data)
		})
}

func TestControllerCreate(t *testing.T) {
	t.Run(
		"Test Create - OK", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodPost, "/api/v1/sellers/",
				`{
				"cid": 200, 
				"company_name": "MELI", 
				"address": "Rua B", 
				"telephone": "9999-8888"
				}`)
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			serviceMock.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(seller1, nil)
			sellers := r.Group("/api/v1/sellers")
			sellers.POST("/", s.Create())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 201, rr.Code)
			objResp := struct {
				Code int
				Data seller.Seller
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, seller1, objResp.Data)
		})
	t.Run(
		"Test Create - Requisition Body error - without Telephone", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodPost, "/api/v1/sellers/",
				`{
					"cid": 200, 
					"company_name": "MELI", 
					"address": "Rua B"
					}`)
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			sellers := r.Group("/api/v1/sellers")
			sellers.POST("/", s.Create())
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
			assert.Equal(t, "This field is required", objResp.Data[0].Message)
			assert.Equal(t, "telephone", objResp.Data[0].Field)
		})
	t.Run(
		"Test Create - CID already registered", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodPost, "/api/v1/sellers/",
				`{
						"cid": 200, 
						"company_name": "MELI", 
						"address": "Rua B",
						"telephone": "9999-8888"
						}`)
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			serviceMock.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(seller.Seller{}, fmt.Errorf("Cid already registered"))
			sellers := r.Group("/api/v1/sellers")
			sellers.POST("/", s.Create())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 409, rr.Code)
			log.Println(string(rr.Body.Bytes()))
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "Cid already registered", objResp.Error)
		})
}

func TestControllerUpdate(t *testing.T) {
	t.Run(
		"Test Update - OK", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodPatch, "/api/v1/sellers/1",
				`{
				"cid": 200, 
				"company_name": "MELI", 
				"address": "Rua B", 
				"telephone": "9999-8888"
				}`)
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			serviceMock.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(seller1, nil)
			sellers := r.Group("/api/v1/sellers")
			sellers.PATCH("/:id", s.Update())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 200, rr.Code)
			objResp := struct {
				Code int
				Data seller.Seller
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, seller1, objResp.Data)
		})
	t.Run(
		"Test Update - Invalid ID", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodPatch, "/api/v1/sellers/a",
				`{
					"cid": 200, 
					"company_name": "MELI", 
					"address": "Rua B", 
					"telephone": "9999-8888"
					}`)
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			sellers := r.Group("/api/v1/sellers")
			sellers.PATCH("/:id", s.Update())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 400, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "Invalid ID", objResp.Error)
		})
	t.Run(
		"Test Update - Invalid JSON body", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodPatch, "/api/v1/sellers/1",
				`{
				"cid": 200, 
				"company_name": "MELI" 
				"address": "Rua B" 
				"telephone": "9999-8888"
				}`)
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			sellers := r.Group("/api/v1/sellers")
			sellers.PATCH("/:id", s.Update())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 400, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, "Invalid body arguments", objResp.Error)
		})
	t.Run(
		"Test Update - ID not found", func(t *testing.T) {
			req, rr := createRequestTest(http.MethodPatch, "/api/v1/sellers/2",
				`{
					"cid": 200, 
					"company_name": "MELI", 
					"address": "Rua B", 
					"telephone": "9999-8888"
					}`)
			serviceMock := new(mocks.Service)
			s := handler.NewSeller(serviceMock)
			r := gin.Default()
			serviceMock.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(seller.Seller{}, fmt.Errorf("Seller 2 not found"))
			sellers := r.Group("/api/v1/sellers")
			sellers.PATCH("/:id", s.Update())
			r.ServeHTTP(rr, req)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, 404, rr.Code)
			assert.Equal(t, "Seller 2 not found", objResp.Error)
		})
}
