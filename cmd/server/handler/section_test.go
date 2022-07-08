package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	mocks "github.com/cpereira42/mercado-fresco-pron4/internal/section/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// * Inicia uma request
// @param method string
// @param url string
// @param body string
func CreateRequestServer(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")
	return req, httptest.NewRecorder()
}

/*
 * cria um server
 * @param mock mocks.SectionService
 * @param method string
 * @param url string
 * @param body string
 */
func CreateServerSection(serv *mocks.Service, method, url, body string) *httptest.ResponseRecorder {
	sectionController := NewSectionController(serv)
	router := gin.Default()
	gp := router.Group("/api/v1/sections")
	gp.GET("/", sectionController.ListarSectionAll())
	gp.GET("/:id", sectionController.ListarSectionOne())
	gp.POST("/", sectionController.CreateSection())
	gp.PATCH("/:id", sectionController.UpdateSection())
	gp.DELETE("/:id", sectionController.DeleteSection())
	req, rr := CreateRequestServer(method, url, body)
	router.ServeHTTP(rr, req)
	return rr
}

var ObjetoResponse struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func getData(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

var sectionList []section.Section = []section.Section{
	{
		Id:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	},
	{
		Id:                 2,
		SectionNumber:      3,
		CurrentTemperature: 79845,
		MinimumTemperature: 4,
		CurrentCapacity:    135,
		MinimumCapacity:    23,
		MaximumCapacity:    456,
		WarehouseId:        78,
		ProductTypeId:      456,
	},
}

var newSectionCreate = section.SectionRequestCreate{
	SectionNumber:      1,
	CurrentTemperature: 1,
	MinimumTemperature: 1,
	CurrentCapacity:    1,
	MinimumCapacity:    1,
	MaximumCapacity:    1,
	WarehouseId:        1,
	ProductTypeId:      1,
}

var newSectionRes = section.SectionRequestCreate{
	SectionNumber:      1,
	CurrentTemperature: 1,
	MinimumTemperature: 1,
	CurrentCapacity:    1,
	MinimumCapacity:    1,
	MaximumCapacity:    1,
	WarehouseId:        1,
	ProductTypeId:      1,
}

func TestCreateSection(t *testing.T) {
	t.Run("criar section, sucesso (ok 201)", func(t *testing.T) {
		mockService := &mocks.Service{}
		mockService.On("CreateSection",
			mock.AnythingOfType("section.SectionRequestCreate"),
		).Return(newSectionRes, nil).Once()
		newSectionCreateByte, _ := json.Marshal(newSectionCreate)
		rr := CreateServerSection(
			mockService,
			http.MethodPost,
			"/api/v1/sections/",
			string(newSectionCreateByte),
		)
		getData(rr.Body.Bytes(), &ObjetoResponse)

		dtByte, _ := json.Marshal(ObjetoResponse.Data)
		object := section.SectionRequestCreate{}
		getData(dtByte, &object)

		assert.Equal(t, ObjetoResponse.Code, rr.Code)
		assert.Equal(t, newSectionRes.SectionNumber, object.SectionNumber)
		assert.Equal(t, newSectionRes.CurrentCapacity, object.CurrentCapacity)
		assert.Equal(t, newSectionRes.CurrentTemperature, object.CurrentTemperature)
		assert.Equal(t, newSectionRes.MaximumCapacity, object.MaximumCapacity)
		assert.Equal(t, newSectionRes.MinimumCapacity, object.MinimumCapacity)
		assert.Equal(t, newSectionRes.MinimumTemperature, object.MinimumTemperature)
		assert.Equal(t, newSectionRes.WarehouseId, object.WarehouseId)
		assert.Equal(t, newSectionRes.ProductTypeId, object.ProductTypeId)
		assert.ObjectsAreEqual(newSectionRes, ObjetoResponse.Data)
	})
	t.Run("criar section, error (conflic 409)", func(t *testing.T) {
		mockService := &mocks.Service{}
		errNewSection := errors.New("section invalid, section_number field must be unique")

		mockService.On("CreateSection",
			mock.AnythingOfType("section.SectionRequestCreate")).
			Return(newSectionRes, errNewSection).
			Once()
		newSectionCreateByte, _ := json.Marshal(newSectionCreate)
		rr := CreateServerSection(
			mockService,
			http.MethodPost,
			"/api/v1/sections/",
			string(newSectionCreateByte),
		)
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, ObjetoResponse.Code, rr.Code)
		assert.ObjectsAreEqual(newSectionRes, ObjetoResponse.Data)
	})
	t.Run("criar section, error (unprocessableEntity 422)", func(t *testing.T) {
		mockService := &mocks.Service{}
		newSectionCreate := section.SectionRequestCreate{
			SectionNumber:      3,
			CurrentTemperature: 1,
			MaximumCapacity:    1,
		}
		newSectionRes := section.Section{}
		errNewSection := errors.New("This field is required")
		mockService.On("CreateSection",
			mock.AnythingOfType("section.SectionRequestCreate")).
			Return(newSectionRes, errNewSection).
			Once()
		newSectionCreateByte, _ := json.Marshal(newSectionCreate)
		rr := CreateServerSection(
			mockService,
			http.MethodPost,
			"/api/v1/sections/",
			string(newSectionCreateByte),
		)
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, 422, rr.Code)
		assert.ObjectsAreEqual(newSectionRes, ObjetoResponse.Data)
	})
}

func TestListarSectionAll(t *testing.T) {
	t.Run("lista todos section, sucesso (ok 200)", func(t *testing.T) {
		mockService := &mocks.Service{}
		mockService.On("ListarSectionAll").
			Return(sectionList, nil).
			Once()
		rr := CreateServerSection(
			mockService,
			http.MethodGet,
			"/api/v1/sections/",
			"")
		assert.Equal(t, 200, rr.Code)
		errListResponse := getData(rr.Body.Bytes(), &ObjetoResponse)

		var sList []section.Section

		objData, _ := json.Marshal(ObjetoResponse.Data)
		getData(objData, &sList)

		assert.True(t, len(sList) > 0)
		assert.Equal(t, sectionList, sList)
		assert.Equal(t, 200, ObjetoResponse.Code)
		assert.NoError(t, errListResponse)
	})
	t.Run("lista todos section, error (500)", func(t *testing.T) {
		mockService := &mocks.Service{}
		errListNil := errors.New("não há sections registrados")
		mockService.On("ListarSectionAll").
			Return([]section.Section{}, errListNil).
			Once()
		rr := CreateServerSection(mockService, http.MethodGet, "/api/v1/sections/", "")

		errListResponse := json.Unmarshal(rr.Body.Bytes(), &ObjetoResponse)
		assert.NoError(t, errListResponse)
		assert.Equal(t, rr.Code, ObjetoResponse.Code)
		assert.Equal(t, errListNil.Error(), ObjetoResponse.Error)
	})
}

func TestListarSectionOne(t *testing.T) {
	t.Run("lista section, sucesso(ok 200)", func(t *testing.T) {
		var mockService *mocks.Service = &mocks.Service{}
		newSectionRes := section.Section{
			Id:                 1,
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseId:        1,
			ProductTypeId:      1,
		}
		mockService.On("ListarSectionOne", mock.AnythingOfType("int64")).
			Return(newSectionRes, nil).
			Once()
		rr := CreateServerSection(
			mockService,
			http.MethodGet,
			"/api/v1/sections/1",
			"",
		)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
	t.Run("lista section, error(not found 404)", func(t *testing.T) {
		var mockService *mocks.Service = &mocks.Service{}
		var sectionNil section.Section = section.Section{}
		expectErr := errors.New("Section is not registered")
		mockService.On("ListarSectionOne", mock.Anything).
			Return(sectionNil, expectErr).
			Once()
		rr := CreateServerSection(mockService, http.MethodGet, "/api/v1/sections/3", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, expectErr.Error(), ObjetoResponse.Error)
	})
	t.Run("lista section, error(not found 404)", func(t *testing.T) {
		var mockService *mocks.Service = &mocks.Service{}
		var searchSection section.Section = section.Section{}
		var errSection error = errors.New("strconv.ParseInt: parsing \"2s\": invalid syntax")
		mockService.On("ListarSectionOne",
			mock.AnythingOfType("int64")).
			Return(searchSection, errSection).
			Once()
		rr := CreateServerSection(
			mockService,
			http.MethodGet,
			"/api/v1/sections/2s",
			"",
		)
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, errSection.Error(), ObjetoResponse.Error)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestUpdateSection(t *testing.T) {
	var sectionListResponse []section.Section = []section.Section{}
	var mockService *mocks.Service = &mocks.Service{}

	t.Run("update section, sucesso (ok 200)", func(t *testing.T) {
		updateSection := section.SectionRequestUpdate{
			SectionNumber:      1,
			CurrentTemperature: 23,
			MinimumTemperature: 23,
			CurrentCapacity:    23,
			MinimumCapacity:    23,
			MaximumCapacity:    23,
			WarehouseId:        23,
			ProductTypeId:      23,
		}
		updateSectionRes := section.Section{
			Id:                 1,
			SectionNumber:      1,
			CurrentTemperature: 23,
			MinimumTemperature: 23,
			CurrentCapacity:    23,
			MinimumCapacity:    23,
			MaximumCapacity:    23,
			WarehouseId:        23,
			ProductTypeId:      23,
		}
		mockService.On("ListarSectionAll").Return(sectionListResponse, nil).Once()
		mockService.On("UpdateSection",
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("section.SectionRequestUpdate"),
		).Return(updateSectionRes, nil).Once()
		updateSectionByte, _ := json.Marshal(updateSection)
		rr := CreateServerSection(
			mockService,
			http.MethodPatch,
			"/api/v1/sections/1",
			string(updateSectionByte),
		)
		assert.Equal(t, 200, rr.Code)
		err := getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Nil(t, err)
		assert.Equal(t, ObjetoResponse.Code, rr.Code)
		assert.ObjectsAreEqual(updateSectionRes, ObjetoResponse.Data)
	})
	t.Run("update section, error (500)", func(t *testing.T) {
		updateSection := section.SectionRequestUpdate{
			SectionNumber:      1,
			CurrentTemperature: 23,
			MinimumTemperature: 23,
			CurrentCapacity:    23,
			MinimumCapacity:    23,
			MaximumCapacity:    23,
			WarehouseId:        23,
			ProductTypeId:      23,
		}
		mockService.On("ListarSectionAll").
			Return(sectionListResponse, nil).
			Once()
		mockService.On("UpdateSection",
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("section.SectionRequestUpdate")).
			Return(section.Section{}, errors.New("o tipo do parâmentro está invalido")).
			Once()
		updateSectionByte, _ := json.Marshal(updateSection)
		rr := CreateServerSection(
			mockService,
			http.MethodPatch,
			"/api/v1/sections/1s",
			string(updateSectionByte),
		)
		assert.Equal(t, 500, rr.Code)
		_ = getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, ObjetoResponse.Code, rr.Code)
	})
	t.Run("update section, error (not found 404)", func(t *testing.T) {
		updateSection := section.SectionRequestUpdate{
			SectionNumber:      1,
			CurrentTemperature: 23,
			MinimumTemperature: 23,
			MaximumCapacity:    23,
			WarehouseId:        23,
			ProductTypeId:      23,
		}
		mockService.On("ListarSectionAll").
			Return([]section.Section{}, fmt.Errorf("não há sections registrados")).
			Once()
		mockService.On("UpdateSection",
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("section.SectionRequestUpdate")).
			Return(section.Section{}, errors.New("section is not found")).
			Once()
		updateSectionByte, _ := json.Marshal(updateSection)
		rr := CreateServerSection(
			mockService,
			http.MethodPatch,
			"/api/v1/sections/1",
			string(updateSectionByte),
		)
		assert.Equal(t, 404, rr.Code)
	})
	t.Run("update section, error (unprocessableEntity 422)", func(t *testing.T) {
		sectionListRes := []section.Section{
			{
				Id:                 1,
				SectionNumber:      1,
				CurrentTemperature: 1,
				MinimumTemperature: 1,
				CurrentCapacity:    1,
				MinimumCapacity:    1,
				MaximumCapacity:    1,
				WarehouseId:        1,
				ProductTypeId:      1,
			},
		}
		mockService.On("ListarSectionAll").
			Return(sectionListRes, nil).
			Once()
		mockService.On("UpdateSection",
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("section.SectionRequestUpdate")).
			Return(section.Section{}, errors.New("This field is required")).
			Once()
		updateSectionByte := `{
			"section_number":1,
			"current_temperature":23,
			"minimum_temperature":23,
			"current_capacity":"0",
			"minimum_capacity":"0",
			"maximum_capacity":23,
			"warehouse_id":23,
			"product_type_id":23
		}`
		rr := CreateServerSection(
			mockService,
			http.MethodPatch,
			"/api/v1/sections/1",
			string(updateSectionByte),
		)
		assert.Equal(t, 422, rr.Code)
	})
}

func TestSectionDelete(t *testing.T) {
	var mockService *mocks.Service = &mocks.Service{}

	t.Run("delete sucesso, (not content 204)", func(t *testing.T) {
		mockService.On("ListarSectionAll").Return(sectionList, nil).Once()
		mockService.On("DeleteSection", mock.AnythingOfType("int64")).
			Return(nil).
			Once()
		rr := CreateServerSection(mockService, http.MethodDelete, "/api/v1/sections/1", "")
		assert.Equal(t, http.StatusNoContent, rr.Code)
	})
	t.Run("delete error, (not found 404)", func(t *testing.T) {
		errNotFound := errors.New("section not found")
		mockService.On("ListarSectionAll").Return(sectionList, nil).Once()
		mockService.On("DeleteSection", mock.AnythingOfType("int64")).
			Return(errNotFound).
			Once()
		rr := CreateServerSection(mockService, http.MethodDelete, "/api/v1/sections/6", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
	t.Run("delete error, (not found 500)", func(t *testing.T) {
		mockService.On("ListarSectionAll").Return(sectionList, nil).Once()
		mockService.On("DeleteSection", mock.AnythingOfType("int64")).
			Return(nil).
			Once()
		rr := CreateServerSection(mockService, http.MethodDelete, "/api/v1/sections/6s", "")
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
