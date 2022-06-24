package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section/entites"
	mocks "github.com/cpereira42/mercado-fresco-pron4/internal/section/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getData(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func TestCreateSection(t *testing.T) {
	t.Run("criar section, sucesso (ok 201)", func(t *testing.T) {
		mockService := &mocks.SectionService{}
		newSectionCreate := entites.SectionRequestCreate{
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseId:        1,
			ProductTypeId:      1,
		}
		newSectionRes := entites.Section{
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
		var sectionList []entites.Section = []entites.Section{
			{
				Id:                 1,
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
		mockService.On("ListarSectionAll").Return(sectionList, nil).Once()
		mockService.On("CreateSection",
			mock.AnythingOfType("entites.SectionRequestCreate"),
		).Return(newSectionRes, nil).Once()
		newSectionCreateByte, _ := json.Marshal(newSectionCreate)
		rr := CreateServerSection(mockService,
			http.MethodPost, "/api/v1/sections/", string(newSectionCreateByte))
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, ObjetoResponse.Code, rr.Code)
		assert.ObjectsAreEqual(newSectionRes, ObjetoResponse.Data)
	})
	t.Run("criar section, error (conflic 409)", func(t *testing.T) {
		mockService := &mocks.SectionService{}
		newSectionCreate := entites.SectionRequestCreate{
			SectionNumber:      3,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseId:        1,
			ProductTypeId:      1,
		}
		newSectionRes := entites.Section{}
		errNewSection := errors.New("section invalid, section_number field must be unique")
		var sectionList []entites.Section = []entites.Section{
			{
				Id:                 1,
				SectionNumber:      3,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity:    135,
				MinimumCapacity:    23,
				MaximumCapacity:    456,
				WarehouseId:        78,
				ProductTypeId:      456,
			}, {
				Id:                 2,
				SectionNumber:      313,
				CurrentTemperature: 745,
				MinimumTemperature: 344,
				CurrentCapacity:    1345,
				MinimumCapacity:    243,
				MaximumCapacity:    43456,
				WarehouseId:        784,
				ProductTypeId:      43456,
			}, {
				Id:                 3,
				SectionNumber:      490,
				CurrentTemperature: 795,
				MinimumTemperature: 3,
				CurrentCapacity:    15,
				MinimumCapacity:    23,
				MaximumCapacity:    3,
				WarehouseId:        78,
				ProductTypeId:      456,
			}, {
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
			mock.AnythingOfType("entites.SectionRequestCreate")).
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
		mockService := &mocks.SectionService{}
		newSectionCreate := entites.SectionRequestCreate{
			SectionNumber:      3,
			CurrentTemperature: 1,
			MaximumCapacity:    1,
		}
		newSectionRes := entites.Section{}
		errNewSection := errors.New("This field is required")
		var sectionList []entites.Section = []entites.Section{
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
			mock.AnythingOfType("entites.SectionRequestCreate")).
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
	mockService := &mocks.SectionService{}
	t.Run("lista todos section, sucesso (ok 200)", func(t *testing.T) {
		var ObjetoResponse struct {
			Code  int               `json:"code"`
			Data  []entites.Section `json:"data"`
			Error error             `json:"error"`
		}
		sectionList := []entites.Section{
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
			Data  []entites.Section `json:"data"`
			Error error             `json:"error"`
		}
		sectionListNil := []entites.Section{}
		errListNil := errors.New("não há sections registrados")
		mockService.On("ListarSectionAll").
			Return(sectionListNil, errListNil).
			Once()
		rr := CreateServerSection(mockService, http.MethodGet, "/api/v1/sections/", "")
		assert.Equal(t, 400, rr.Code)

		errListResponse := json.Unmarshal(rr.Body.Bytes(), &ObjetoResponse.Data)
		assert.True(t, len(ObjetoResponse.Data) == 0)
		assert.Error(t, errListResponse)
	})
}

func TestListarSectionOne(t *testing.T) {
	t.Run("lista section, sucesso(ok 200)", func(t *testing.T) {
		var mockService *mocks.SectionService = &mocks.SectionService{}
		newSectionRes := entites.Section{
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
		var mockService *mocks.SectionService = &mocks.SectionService{}
		var sectionNil entites.Section = entites.Section{}
		mockService.On("ListarSectionOne", mock.Anything).
			Return(sectionNil, errors.New("Section is not registered")).
			Once()
		rr := CreateServerSection(mockService, http.MethodGet, "/api/v1/sections/3", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
	t.Run("lista section, error(not found 404)", func(t *testing.T) {
		var mockService *mocks.SectionService = &mocks.SectionService{}
		var searchSection entites.Section = entites.Section{}
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
	var sectionListResponse []entites.Section = []entites.Section{}
	var mockService *mocks.SectionService = &mocks.SectionService{}

	t.Run("update section, sucesso (ok 200)", func(t *testing.T) {
		updateSection := entites.SectionRequestUpdate{
			SectionNumber:      1,
			CurrentTemperature: 23,
			MinimumTemperature: 23,
			CurrentCapacity:    23,
			MinimumCapacity:    23,
			MaximumCapacity:    23,
			WarehouseId:        23,
			ProductTypeId:      23,
		}
		updateSectionRes := entites.Section{
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
			mock.AnythingOfType("entites.SectionRequestUpdate"),
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
		updateSection := entites.SectionRequestUpdate{
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
			mock.AnythingOfType("entites.SectionRequestUpdate")).
			Return(entites.Section{}, errors.New("o tipo do parâmentro está invalido")).
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
		updateSection := entites.SectionRequestUpdate{
			SectionNumber:      1,
			CurrentTemperature: 23,
			MinimumTemperature: 23,
			MaximumCapacity:    23,
			WarehouseId:        23,
			ProductTypeId:      23,
		}
		mockService.On("ListarSectionAll").
			Return([]entites.Section{}, fmt.Errorf("não há sections registrados")).
			Once()
		mockService.On("UpdateSection",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("entites.SectionRequestUpdate")).
			Return(entites.Section{}, errors.New("section is not found")).
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
		sectionListRes := []entites.Section{
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
			mock.AnythingOfType("entites.SectionRequestUpdate")).
			Return(entites.Section{}, errors.New("This field is required")).
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
	var mockService *mocks.SectionService = &mocks.SectionService{}
	sectionListRes := []entites.Section{
		{
			Id:                 1,
			SectionNumber:      1,
			CurrentTemperature: 2,
			MinimumTemperature: 2,
			CurrentCapacity:    2,
			MinimumCapacity:    2,
			MaximumCapacity:    2,
			WarehouseId:        2,
			ProductTypeId:      2,
		},
	}
	t.Run("delete sucesso, (not content 204)", func(t *testing.T) {
		mockService.On("ListarSectionAll").Return(sectionListRes, nil).Once()
		mockService.On("DeleteSection", mock.AnythingOfType("int")).
			Return(nil).
			Once()
		rr := CreateServerSection(mockService, http.MethodDelete, "/api/v1/sections/1", "")
		assert.Equal(t, http.StatusNoContent, rr.Code)
	})
	t.Run("delete error, (not found 404)", func(t *testing.T) {
		errNotFound := errors.New("section not found")
		mockService.On("ListarSectionAll").Return(sectionListRes, nil).Once()
		mockService.On("DeleteSection", mock.AnythingOfType("int")).
			Return(errNotFound).
			Once()
		rr := CreateServerSection(mockService, http.MethodDelete, "/api/v1/sections/1", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
	t.Run("delete error, (not found 404)", func(t *testing.T) {
		mockService.On("ListarSectionAll").Return(sectionListRes, nil).Once()
		mockService.On("DeleteSection", mock.AnythingOfType("int")).
			Return(nil).
			Once()
		rr := CreateServerSection(mockService, http.MethodDelete, "/api/v1/sections/1s", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
