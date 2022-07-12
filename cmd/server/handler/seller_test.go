package handler_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"
	"github.com/cpereira42/mercado-fresco-pron4/internal/seller/mocks"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var objSellers = struct {
	Code int
	Data []seller.Seller
}{}

var objSeller = struct {
	Code int
	Data seller.Seller
}{}

var objError = struct {
	Code  int
	Error string
}{}

var objSellerWithError = struct {
	Code  int
	Data  seller.Seller
	Error string
}{}

var objMultipleErrors = struct {
	Code int
	Data []struct {
		Field   string
		Message string
	}
}{}

var seller1 seller.Seller = seller.Seller{Id: 1, Cid: "200", CompanyName: "MELI", Address: "Rua B", Telephone: "9999-8888", LocalityId: 1}

var sListSuccess []seller.Seller = []seller.Seller{
	seller1,
}

func createServerSeller(serv *mocks.Service, method string, url string, body string) *httptest.ResponseRecorder {

	r := gin.Default()
	handler.NewSeller(r, serv)
	req, rr := util.CreateRequestTest(method, url, body)
	r.ServeHTTP(rr, req)
	return rr
}

func TestControllerGetAllSeller(t *testing.T) {
	t.Run(
		"Test GetAll - OK", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("GetAll").Return(sListSuccess, nil)
			rr := createServerSeller(serviceMock, http.MethodGet, "/api/v1/sellers/", "")
			err := json.Unmarshal(rr.Body.Bytes(), &objSellers)
			assert.Equal(t, 200, objSellers.Code)
			assert.Nil(t, err)
			assert.Equal(t, sListSuccess, objSellers.Data)
		})
	t.Run(
		"Test GetAll - Error - Could not read file", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			msgError := "Could not read file"
			serviceMock.On("GetAll").Return([]seller.Seller{}, fmt.Errorf(msgError))
			rr := createServerSeller(serviceMock, http.MethodGet, "/api/v1/sellers/", "")

			err := json.Unmarshal(rr.Body.Bytes(), &objError)
			assert.Equal(t, 400, objError.Code)
			assert.Nil(t, err)
			assert.Equal(t, msgError, objError.Error)
		})
}

func TestControllerDeleteSeller(t *testing.T) {
	t.Run(
		"Test Delete - OK", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("Delete", 1).Return(nil)
			rr := createServerSeller(serviceMock, http.MethodDelete, "/api/v1/sellers/1", "")
			assert.Equal(t, 204, rr.Code)
		})
	t.Run(
		"Test Delete - Error - Invalid ID", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			rr := createServerSeller(serviceMock, http.MethodDelete, "/api/v1/sellers/a", "")
			err := json.Unmarshal(rr.Body.Bytes(), &objError)
			assert.Nil(t, err)
			assert.Equal(t, 400, objError.Code)
			assert.Equal(t, "invalid ID", objError.Error)
		})
	t.Run(
		"Test Delete - Error - ID not found", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("Delete", 2).Return(fmt.Errorf("Seller 2 not found"))
			rr := createServerSeller(serviceMock, http.MethodDelete, "/api/v1/sellers/2", "")
			err := json.Unmarshal(rr.Body.Bytes(), &objError)
			assert.Nil(t, err)
			assert.Equal(t, 404, rr.Code)
			assert.Equal(t, "Seller 2 not found", objError.Error)
		})
}

func TestControllerGetIdSeller(t *testing.T) {
	t.Run(
		"Test GetId - OK", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("GetId", 1).Return(seller1, nil)
			rr := createServerSeller(serviceMock, http.MethodGet, "/api/v1/sellers/1", "")
			err := json.Unmarshal(rr.Body.Bytes(), &objSeller)
			assert.Nil(t, err)
			assert.Equal(t, 200, objSeller.Code)
			assert.Equal(t, seller1, objSeller.Data)
		})
	t.Run(
		"Test GetId - Error - Invalid ID", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			rr := createServerSeller(serviceMock, http.MethodGet, "/api/v1/sellers/a", "")
			err := json.Unmarshal(rr.Body.Bytes(), &objError)
			assert.Nil(t, err)
			assert.Equal(t, 400, objError.Code)
			assert.Equal(t, "Invalid ID", objError.Error)
		})
	t.Run(
		"Test GetId - Error - ID not found", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("GetId", 2).Return(seller.Seller{}, fmt.Errorf("Seller 2 not found"))
			rr := createServerSeller(serviceMock, http.MethodGet, "/api/v1/sellers/2", "")
			err := json.Unmarshal(rr.Body.Bytes(), &objSellerWithError)
			assert.Nil(t, err)
			assert.Equal(t, 404, objSellerWithError.Code)
			assert.Equal(t, "Seller 2 not found", objSellerWithError.Error)
			assert.Equal(t, seller.Seller{}, objSellerWithError.Data)
		})
}

func TestControllerCreateSeller(t *testing.T) {
	t.Run(
		"Test Create - OK", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int")).
				Return(seller1, nil)
			rr := createServerSeller(serviceMock, http.MethodPost, "/api/v1/sellers/",
				`{
				"cid": "200", 
				"company_name": "MELI", 
				"address": "Rua B", 
				"telephone": "9999-8888",
				"locality_id": 1
				}`)

			err := json.Unmarshal(rr.Body.Bytes(), &objSeller)
			assert.Nil(t, err)
			assert.Equal(t, 201, objSeller.Code)
			assert.Equal(t, seller1, objSeller.Data)
		})
	t.Run(
		"Test Create - Requisition Body error - without Telephone", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			rr := createServerSeller(serviceMock, http.MethodPost, "/api/v1/sellers/",
				`{
					"cid": "200", 
					"company_name": "MELI", 
					"address": "Rua B"
					}`)

			err := json.Unmarshal(rr.Body.Bytes(), &objMultipleErrors)
			assert.Nil(t, err)
			assert.Equal(t, 422, objMultipleErrors.Code)
			assert.Equal(t, "This field is required", objMultipleErrors.Data[0].Message)
			assert.Equal(t, "telephone", objMultipleErrors.Data[0].Field)
		})
	t.Run(
		"Test Create - CID already registered", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int")).
				Return(seller.Seller{}, fmt.Errorf("Cid already registered"))

			rr := createServerSeller(serviceMock, http.MethodPost, "/api/v1/sellers/",
				`{
						"cid": "200", 
						"company_name": "MELI", 
						"address": "Rua B",
						"telephone": "9999-8888",
						"locality_id": 1
						}`)
			log.Println(string(rr.Body.Bytes()))
			err := json.Unmarshal(rr.Body.Bytes(), &objError)
			assert.Nil(t, err)
			assert.Equal(t, 409, objError.Code)
			assert.Equal(t, "Cid already registered", objError.Error)
		})
}

func TestControllerUpdateSeller(t *testing.T) {
	t.Run(
		"Test Update - OK", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int")).
				Return(seller1, nil)
			rr := createServerSeller(serviceMock, http.MethodPatch, "/api/v1/sellers/1",
				`{
				"cid": "200", 
				"company_name": "MELI", 
				"address": "Rua B", 
				"telephone": "9999-8888",
				"locality_id": 1
				}`)
			err := json.Unmarshal(rr.Body.Bytes(), &objSeller)
			assert.Equal(t, 200, objSeller.Code)
			assert.Nil(t, err)
			assert.Equal(t, seller1, objSeller.Data)
		})
	t.Run(
		"Test Update - Invalid ID", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			rr := createServerSeller(serviceMock, http.MethodPatch, "/api/v1/sellers/a",
				`{
					"cid": "200", 
					"company_name": "MELI", 
					"address": "Rua B", 
					"telephone": "9999-8888",
					"locality_id": 1
					}`)

			err := json.Unmarshal(rr.Body.Bytes(), &objError)
			assert.Equal(t, 400, objError.Code)
			assert.Nil(t, err)
			assert.Equal(t, "Invalid ID", objError.Error)
		})
	t.Run(
		"Test Update - Invalid JSON body", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			rr := createServerSeller(serviceMock, http.MethodPatch, "/api/v1/sellers/1",
				`{
				"cid": 200, 
				"company_name": "MELI" 
				"address": "Rua B" 
				"telephone": "9999-8888"
				}`)
			err := json.Unmarshal(rr.Body.Bytes(), &objError)
			assert.Equal(t, 400, objError.Code)
			assert.Nil(t, err)
			assert.Equal(t, "Invalid body arguments", objError.Error)
		})
	t.Run(
		"Test Update - ID not found", func(t *testing.T) {
			serviceMock := new(mocks.Service)
			serviceMock.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int")).
				Return(seller.Seller{}, fmt.Errorf("Seller 2 not found"))
			rr := createServerSeller(serviceMock, http.MethodPatch, "/api/v1/sellers/2",
				`{
				"cid": "200", 
				"company_name": "MELI", 
				"address": "Rua B", 
				"telephone": "9999-8888",
				"locality_id": 1
				}`)
			err := json.Unmarshal(rr.Body.Bytes(), &objError)
			assert.Nil(t, err)
			assert.Equal(t, 404, objError.Code)
			assert.Equal(t, "Seller 2 not found", objError.Error)
		})
}