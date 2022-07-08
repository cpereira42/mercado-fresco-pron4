package handler_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/locality"
	"github.com/cpereira42/mercado-fresco-pron4/internal/locality/mocks"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var locality1 = locality.Locality{1, "Itabaiana", "Sergipe", "Brasil"}
var report1 = locality.LocalityReport{1, "Itabaiana", 2}
var report2 = locality.LocalityReport{2, "Aracaju", 4}
var reportList = []locality.LocalityReport{report1, report2}

func TestControllerCreateLocality(t *testing.T) {
	t.Run(
		"Test Create - OK", func(t *testing.T) {
			req, rr := util.CreateRequestTest(http.MethodPost, "/api/v1/localities/",
				`{
					"id": 1000,
			"locality_name": "Itabaiana",
			"province_name": "Sergipe",
			"country_name": "Brasil"
		}`)
			serviceMock := new(mocks.Service)
			s := handler.NewLocality(serviceMock)
			r := gin.Default()
			serviceMock.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(locality1, nil)
			localities := r.Group("/api/v1/localities")
			localities.POST("/", s.Create())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 201, rr.Code)
			objResp := struct {
				Code int
				Data locality.Locality
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, locality1, objResp.Data)
		})
	t.Run(
		"Test Create - Fail - Invalid Arguments", func(t *testing.T) {
			req, rr := util.CreateRequestTest(http.MethodPost, "/api/v1/localities/",
				`{
				"locality_name": "Itabaiana",
				"province_name": "Sergipe",
				"country_name": "Brasil"
			}`)
			serviceMock := new(mocks.Service)
			s := handler.NewLocality(serviceMock)
			r := gin.Default()
			localities := r.Group("/api/v1/localities")
			localities.POST("/", s.Create())
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
			assert.Equal(t, "id", objResp.Data[0].Field)
		})
	t.Run(
		"Test Create - ID already registered", func(t *testing.T) {
			req, rr := util.CreateRequestTest(http.MethodPost, "/api/v1/localities/",
				`{
				"id": 1000,
				"locality_name": "Itabaiana",
				"province_name": "Sergipe",
				"country_name": "Brasil"
			}`)
			serviceMock := new(mocks.Service)
			s := handler.NewLocality(serviceMock)
			r := gin.Default()
			msgError := fmt.Errorf("Locality already registered")
			serviceMock.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(locality.Locality{}, msgError)
			localities := r.Group("/api/v1/localities")
			localities.POST("/", s.Create())
			r.ServeHTTP(rr, req)
			log.Println(rr.Body.String())
			assert.Equal(t, 409, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, msgError.Error(), objResp.Error)
		})
}

func TestControllerGenerateReportAll(t *testing.T) {
	t.Run(
		"Test GenerateReportAll - OK", func(t *testing.T) {
			req, rr := util.CreateRequestTest(http.MethodGet, "/api/v1/localities/", "")
			serviceMock := new(mocks.Service)
			s := handler.NewLocality(serviceMock)
			r := gin.Default()
			serviceMock.On("GenerateReportAll").Return(reportList, nil)
			localities := r.Group("/api/v1/localities")
			localities.GET("/", s.GenerateReportAll())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 200, rr.Code)
			objResp := struct {
				Code int
				Data []locality.LocalityReport
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, reportList, objResp.Data)
		})
	t.Run(
		"Test GenerateReportAll - DB Error", func(t *testing.T) {
			req, rr := util.CreateRequestTest(http.MethodGet, "/api/v1/localities/", "")
			serviceMock := new(mocks.Service)
			s := handler.NewLocality(serviceMock)
			r := gin.Default()
			msgError := fmt.Errorf("DB Connection Failed")
			serviceMock.On("GenerateReportAll").Return([]locality.LocalityReport{}, msgError)
			localities := r.Group("/api/v1/localities")
			localities.GET("/", s.GenerateReportAll())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 400, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, msgError.Error(), objResp.Error)
		})
}

func TestControllerGenerateReportById(t *testing.T) {
	t.Run(
		"Test GenerateReportById - OK", func(t *testing.T) {
			req, rr := util.CreateRequestTest(http.MethodGet, "/api/v1/localities/1", "")
			serviceMock := new(mocks.Service)
			s := handler.NewLocality(serviceMock)
			r := gin.Default()
			serviceMock.On("GenerateReportById", 1).Return(report1, nil)
			localities := r.Group("/api/v1/localities")
			localities.GET("/:id", s.GenerateReportById())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 200, rr.Code)
			objResp := struct {
				Code int
				Data locality.LocalityReport
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, report1, objResp.Data)
		})
	t.Run(
		"Test GenerateReportById - Invalid ID", func(t *testing.T) {
			req, rr := util.CreateRequestTest(http.MethodGet, "/api/v1/localities/a", "")
			serviceMock := new(mocks.Service)
			s := handler.NewLocality(serviceMock)
			r := gin.Default()
			msgError := fmt.Errorf("Invalid ID")
			serviceMock.On("GenerateReportById", "a").Return(locality.LocalityReport{}, msgError)
			localities := r.Group("/api/v1/localities")
			localities.GET("/:id", s.GenerateReportById())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 400, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, msgError.Error(), objResp.Error)
		})
	t.Run(
		"Test GenerateReportById - ID not found", func(t *testing.T) {
			req, rr := util.CreateRequestTest(http.MethodGet, "/api/v1/localities/4", "")
			serviceMock := new(mocks.Service)
			s := handler.NewLocality(serviceMock)
			r := gin.Default()
			msgError := fmt.Errorf("Locality 4 not found")
			serviceMock.On("GenerateReportById", 4).Return(locality.LocalityReport{}, msgError)
			localities := r.Group("/api/v1/localities")
			localities.GET("/:id", s.GenerateReportById())
			r.ServeHTTP(rr, req)
			assert.Equal(t, 404, rr.Code)
			objResp := struct {
				Code  int
				Error string
			}{}
			err := json.Unmarshal(rr.Body.Bytes(), &objResp)
			assert.Nil(t, err)
			assert.Equal(t, msgError.Error(), objResp.Error)
		})
}
