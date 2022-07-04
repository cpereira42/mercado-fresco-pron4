package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
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
	repoPB := productbatch.NewRepositoryProductBatches(nil)
	servicePB := productbatch.NewServiceProductBatches(repoPB)
	sectionController := NewSectionController(serv, servicePB)
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
	Code  int   `json:"code"`
	Data  any   `json:"data"`
	Error error `json:"error"`
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
	{
		Id:                 3,
		SectionNumber:      313,
		CurrentTemperature: 745,
		MinimumTemperature: 344,
		CurrentCapacity:    1345,
		MinimumCapacity:    243,
		MaximumCapacity:    43456,
		WarehouseId:        784,
		ProductTypeId:      43456,
	},
	{
		Id:                 4,
		SectionNumber:      490,
		CurrentTemperature: 795,
		MinimumTemperature: 3,
		CurrentCapacity:    15,
		MinimumCapacity:    23,
		MaximumCapacity:    3,
		WarehouseId:        78,
		ProductTypeId:      456,
	},
	{
		Id:                 5,
		SectionNumber:      495,
		CurrentTemperature: 795,
		MinimumTemperature: 3,
		CurrentCapacity:    15,
		MinimumCapacity:    23,
		MaximumCapacity:    456,
		WarehouseId:        78,
		ProductTypeId:      456,
	},
}

func TestCreateSection(t *testing.T) {
	t.Run("criar section, sucesso (ok 201)", func(t *testing.T) {
		mockService := &mocks.Service{}
		newSectionCreate := section.SectionRequestCreate{
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseId:        1,
			ProductTypeId:      1,
		}
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

		mockService.On("ListarSectionAll").Return(sectionList, nil).Once()
		mockService.On("CreateSection",
			mock.AnythingOfType("section.SectionRequestCreate"),
		).Return(newSectionRes, nil).Once()
		newSectionCreateByte, _ := json.Marshal(newSectionCreate)
		rr := CreateServerSection(mockService,
			http.MethodPost, "/api/v1/sections/", string(newSectionCreateByte))
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, ObjetoResponse.Code, rr.Code)
		assert.ObjectsAreEqual(newSectionRes, ObjetoResponse.Data)
	})
	t.Run("criar section, error (conflic 409)", func(t *testing.T) {
		mockService := &mocks.Service{}
		newSectionCreate := section.SectionRequestCreate{
			SectionNumber:      3,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseId:        1,
			ProductTypeId:      1,
		}
		newSectionRes := section.Section{}
		errNewSection := errors.New("section invalid, section_number field must be unique")

		mockService.On("ListarSectionAll").Return(sectionList, nil).Once()
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
		var sectionList []section.Section = []section.Section{
			{
				Id:                 4,
				SectionNumber:      495,
				CurrentTemperature: 795,
				MinimumTemperature: 3,
				CurrentCapacity:    15,
				MinimumCapacity:    23,
				MaximumCapacity:    456,
				WarehouseId:        78,
				ProductTypeId:      456,
			},
		}
		mockService.On("ListarSectionAll").Return(sectionList, nil).Once()
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
	mockService := &mocks.Service{}
	t.Run("lista todos section, sucesso (ok 200)", func(t *testing.T) {
		var ObjetoResponse struct {
			Code  int               `json:"code"`
			Data  []section.Section `json:"data"`
			Error error             `json:"error"`
		}

		mockService.On("ListarSectionAll").
			Return(sectionList, nil).
			Once()
		rr := CreateServerSection(mockService, http.MethodGet, "/api/v1/sections/", "")
		assert.Equal(t, 200, rr.Code)
		errListResponse := getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.True(t, len(ObjetoResponse.Data) != 0)
		assert.Equal(t, sectionList, ObjetoResponse.Data)
		assert.Equal(t, 200, ObjetoResponse.Code)
		assert.Nil(t, errListResponse)
	})

	t.Run("lista todos section, error (bad request 400)", func(t *testing.T) {
		var ObjetoResponse struct {
			Code  int               `json:"code"`
			Data  []section.Section `json:"data"`
			Error error             `json:"error"`
		}
		sectionListNil := []section.Section{}
		errListNil := errors.New("não há sections registrados")
		mockService.On("ListarSectionAll").
			Return(sectionListNil, errListNil).
			Once()
		rr := CreateServerSection(mockService, http.MethodGet, "/api/v1/sections/", "")
		assert.Equal(t, 500, rr.Code)

		errListResponse := json.Unmarshal(rr.Body.Bytes(), &ObjetoResponse.Data)
		assert.True(t, len(ObjetoResponse.Data) == 0)
		assert.Error(t, errListResponse)
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
		mockService.On("ListarSectionOne", mock.AnythingOfType("int")).
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
		mockService.On("ListarSectionOne", mock.Anything).
			Return(sectionNil, errors.New("Section is not registered")).
			Once()
		rr := CreateServerSection(mockService, http.MethodGet, "/api/v1/sections/3", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
	t.Run("lista section, error(not found 404)", func(t *testing.T) {
		var mockService *mocks.Service = &mocks.Service{}
		var searchSection section.Section = section.Section{}
		var errSection error = errors.New("Sectin is not registered")
		mockService.On("ListarSectionOne",
			mock.AnythingOfType("int")).
			Return(searchSection, errSection).
			Once()
		rr := CreateServerSection(
			mockService,
			http.MethodGet,
			"/api/v1/sections/2s",
			"",
		)
		assert.Equal(t, http.StatusNotFound, rr.Code)
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
			mock.AnythingOfType("int"),
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
	t.Run("update section, error (not found 404)", func(t *testing.T) {
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
			mock.AnythingOfType("int"),
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
		assert.Equal(t, 404, rr.Code)
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
			mock.AnythingOfType("int"),
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
			mock.AnythingOfType("int"),
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
