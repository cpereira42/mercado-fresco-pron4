package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/carries"
	"github.com/cpereira42/mercado-fresco-pron4/internal/carries/mocks"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var carryOne = carries.Carries{"cid1", "companyname1", "rua 1, ", "11112222", 1}

//var carryOne = carries.RequestCarriesCreate{Cid: "cid1", CompanyName: "companyname1", Address: "rua 1,", Telephone: "11112222", LocalityID: 1}

var localityOne = carries.Localities{1, "São Paulo", 1}
var localityTwo = carries.Localities{2, "Nova York", 2}

var carriesGetReports = []carries.Localities{localityOne, localityTwo}

func createServerCarries(s *mocks.Service, method, url, body string) *httptest.ResponseRecorder {
	r := gin.Default()
	handler.NewCarry(r, s)
	req, rr := util.CreateRequestTest(method, url, body)
	r.ServeHTTP(rr, req)
	return rr

}

func TestCarriesCreate(t *testing.T) {
	t.Run("should return status code 201 and carry created", func(t *testing.T) {
		serviceMock := new(mocks.Service)
		serviceMock.On("Create",
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int")).Return(carryOne, nil)

		rr := createServerCarries(serviceMock, http.MethodPost, "/api/v1/carries/",
			`{
			"cid": "cid1", 
			"company_name": "companyname1", 
			"address": "rua 1", 
			"telephone": "11112222", 
			"locality_id": 1
		}`)

		objResp := struct {
			Code int
			Data carries.Carries
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &objResp)

		assert.Equal(t, 201, objResp.Code)
		assert.Nil(t, err)
		assert.Equal(t, carryOne, objResp.Data)

	})
	t.Run("Test Create - ID already Exists - should fail and return 409 and new Error", func(t *testing.T) {
		serviceMock := new(mocks.Service)
		errorMsg := fmt.Errorf("Carry already Exists")
		serviceMock.On("Create",
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int")).Return(carries.Carries{}, errorMsg)

		rr := createServerCarries(serviceMock, http.MethodPost, "/api/v1/carries/",
			`{
			"cid": "cid1", 
			"company_name": "companyname1", 
			"address": "rua 1", 
			"telephone": "11112222", 
			"locality_id": 1
		}`)

		objResp := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objResp)

		assert.Equal(t, 409, objResp.Code)
		assert.Nil(t, err)

	})

}

func TestGetReport(t *testing.T) {
	t.Run("GenerateReport - should return status code 200 and carries", func(t *testing.T) {
		serviceMock := new(mocks.Service)
		serviceMock.On("GetAllReport").Return(carriesGetReports, nil)

		rr := createServerCarries(serviceMock, http.MethodGet, "/api/v1/localities/", "")

		objResp := struct {
			Code int
			Data []carries.Localities
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objResp)

		assert.Equal(t, 200, objResp.Code)
		assert.Nil(t, err)
		assert.Equal(t, carriesGetReports, objResp.Data)

	})
	t.Run("GenerateReport By ID - should return 200 and carrry", func(t *testing.T) {
		serviceMock := new(mocks.Service)
		serviceMock.On("GetByIDReport", tmock.AnythingOfType("int")).Return(localityOne, nil)

		rr := createServerCarries(serviceMock, http.MethodGet, "/api/v1/localities/?id=1", "")

		objResp := struct {
			Code int
			Data carries.Localities
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objResp)

		assert.Equal(t, 200, objResp.Code)
		assert.Nil(t, err)
		assert.Equal(t, localityOne, objResp.Data)

	})
	t.Run("Generate ReportByID - Invalid ID - should return 404 and error", func(t *testing.T) {
		serviceMock := new(mocks.Service)
		errorMsg := fmt.Errorf("Invalid ID")
		serviceMock.On("GetByIDReport", tmock.AnythingOfType("int")).Return(carries.Localities{}, errorMsg)

		rr := createServerCarries(serviceMock, http.MethodGet, "/api/v1/localities/?id=1", "")

		objResp := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objResp)

		assert.Equal(t, 404, objResp.Code)
		assert.Nil(t, err)
		assert.Equal(t, "Invalid ID", objResp.Error)

	})

	t.Run("Generate ReportAll - Fail to read database - should retyurn 404 and error", func(t *testing.T) {
		serviceMock := new(mocks.Service)
		errorMsg := fmt.Errorf("Fail to read database")
		serviceMock.On("GetAllReport").Return([]carries.Localities{}, errorMsg)

		rr := createServerCarries(serviceMock, http.MethodGet, "/api/v1/localities/", "")

		objResp := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objResp)

		assert.Equal(t, 404, objResp.Code)
		assert.Nil(t, err)
		assert.Equal(t, "Fail to read database", objResp.Error)

	})

}
