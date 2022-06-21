package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"

	"net/http"
	"net/http/httptest"

	//"os"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	//"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	mocks "github.com/cpereira42/mercado-fresco-pron4/internal/section/mock"

	//"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

  
/* 
 * inicia uma request
 * @param method
 * @param url
 * @param body
 */
func createRequestServer(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")
	return req, httptest.NewRecorder()
}

/* 
 * cria um server
 * @param mock
 * @param method
 * @param url
 * @param body
 */
func createServer(serv *mocks.SectionService, method, url, body string) *httptest.ResponseRecorder {
	sectionController := handler.NewSectionController(serv)
	router := gin.Default()
	gp := router.Group("/api/v1/sections")
	gp.GET("/", sectionController.ListarSectionAll())
	gp.GET("/:id", sectionController.ListarSectionOne())
	gp.POST("/", sectionController.CreateSection())
	gp.PATCH("/:id", sectionController.UpdateSection())
	gp.DELETE("/:id", sectionController.DeleteSection())
	req, rr := createRequestServer(method, url, body)
	router.ServeHTTP(rr, req)
	return rr
}
 

func getData(data []byte, v any) error {
	return json.Unmarshal(data, v)
}


func TestCreateSection(t *testing.T) {
	mockService := &mocks.SectionService{}
	t.Run("criar section, sucesso (ok 201)", func(t *testing.T) {
		
		newSectionCreate := section.SectionRequestCreate{
			SectionNumber: 1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity: 1,
			MinimumCapacity: 1,
			MaximumCapacity: 1,
			WarehouseId: 1,
			ProductTypeId: 1,
		}

		newSectionRes := section.Section{
			Id: 1,
			SectionNumber: 1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity: 1,
			MinimumCapacity: 1,
			MaximumCapacity: 1,
			WarehouseId: 1,
			ProductTypeId: 1,
		}
		
		
		mockService.On("CreateSection", 
			mock.AnythingOfType("section.SectionRequestCreate"),
		).Return(newSectionRes, nil).Once()

		newSectionCreateByte, _ := json.Marshal(newSectionCreate)

		rr := createServer(mockService, 
			http.MethodPost, "/api/v1/sections/", string(newSectionCreateByte))
		
		objRes := struct {
			Code int 
			Data section.Section
		}{}

		getData(rr.Body.Bytes(), &objRes)

		assert.Equal(t, objRes.Code, rr.Code)		
		assert.ObjectsAreEqual(newSectionRes, objRes.Data)		
	})
}

func TestListarSectionAll(t *testing.T) {
	mockService := &mocks.SectionService{} 		
	t.Run("lista todos section, sucesso (ok 200)", func(t *testing.T) {
		sectionList := []section.Section{
			{
				Id: 1,
				SectionNumber: 1,
				CurrentTemperature: 1,
				MinimumTemperature: 1,
				CurrentCapacity: 1,
				MinimumCapacity: 1,
				MaximumCapacity: 1,
				WarehouseId: 1,
				ProductTypeId: 1,
			},
		}
		mockService.On("ListarSectionAll").
			Return(sectionList, nil).
			Once()
		
		rr := createServer(mockService, http.MethodGet, "/api/v1/sections/", "")
		
		assert.Equal(t, 200, rr.Code) 

		objResponse := struct {
			Code int 
			Data []section.Section
		}{}
		 
		errListResponse := getData(rr.Body.Bytes(), &objResponse)
		assert.True(t, len(objResponse.Data) != 0)
		assert.Equal(t, sectionList, objResponse.Data)
		assert.Equal(t, 200, objResponse.Code)
		assert.Nil(t, errListResponse)
	})
	t.Run("lista todos section, error (bad request 400)", func(t *testing.T) {
		 
		sectionListNil := []section.Section{}
		errListNil := errors.New("não há sections registrados")
		mockService.On("ListarSectionAll").
			Return(sectionListNil, errListNil).
			Once()
		
		rr := createServer(mockService, http.MethodGet, "/api/v1/sections/", "")
		
		assert.Equal(t, 400, rr.Code) 

		objResponse := struct {
			Code int 
			Data []section.Section
		}{}
			
		errListResponse := json.Unmarshal(rr.Body.Bytes(), &objResponse.Data)

		assert.True(t, len(objResponse.Data) == 0)
		assert.Error(t, errListResponse)
	})
}

func TestListarSectionOne(t *testing.T) {
	// criar mock do service
	var mockService *mocks.SectionService = &mocks.SectionService{}

	// objeto esperado no retorno
	newSectionRes := section.Section{
		Id: 1,
		SectionNumber: 1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity: 1,
		MinimumCapacity: 1,
		MaximumCapacity: 1,
		WarehouseId: 1,
		ProductTypeId: 1,
	}

	t.Run("lista section, sucesso(ok 200)", func(t *testing.T) {
		// chama metodo mockado que será testado
		mockService.On("ListarSectionOne", mock.AnythingOfType("int")). 
			Return(newSectionRes, nil).
			Once()

		rr := createServer(
			mockService, 
			http.MethodGet, 
			"/api/v1/sections/1", 
			"",
		)
		
		assert.Equal(t, http.StatusOK, rr.Code)
	})
	/*t.Run("lista section, error(not found 404)", func(t *testing.T) {
		req, rr := getRequest(http.MethodGet, "/api/v1/sections/3", "")	
  
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
	t.Run("lista section, error(not found 404)", func(t *testing.T) {
		req, rr := getRequest(http.MethodGet, "/api/v1/sections/2s", "")	
 
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})*/
}


func TestUpdateSection(t *testing.T) {
	// cria um objeto de mock do service
	var mockService *mocks.SectionService = &mocks.SectionService{}

	// cria as datas de test
	updateSection := section.SectionRequestUpdate{
		SectionNumber: 1,
		CurrentTemperature: 23, 
		MinimumTemperature: 23,
		CurrentCapacity: 23,
		MinimumCapacity: 23,
		MaximumCapacity: 23,
		WarehouseId: 23,
		ProductTypeId: 23,
	} 
	updateSectionRes := section.Section{
		Id: 1,
		SectionNumber: 1,
		CurrentTemperature: 23, 
		MinimumTemperature: 23,
		CurrentCapacity: 23,
		MinimumCapacity: 23,
		MaximumCapacity: 23,
		WarehouseId: 23,
		ProductTypeId: 23,
	}
	sectionListRes := []section.Section{
		{
			Id: 1,
			SectionNumber: 1,
			CurrentTemperature: 2, 
			MinimumTemperature: 2,
			CurrentCapacity: 2,
			MinimumCapacity: 2,
			MaximumCapacity: 2,
			WarehouseId: 2,
			ProductTypeId: 2,
		},
	}

	// inicia o caso de test de sucesso
	t.Run("update section, sucesso (ok 200)", func(t *testing.T) {
		// chamada dos métodos de teste
		mockService.On("ListarSectionAll").Return(sectionListRes, nil).Once()
		mockService.On("UpdateSection", 
			mock.AnythingOfType("int"),
			mock.AnythingOfType("section.SectionRequestUpdate"), 
		).Return(updateSectionRes, nil).Once()

		// prepara os dados do body da request
		updateSectionByte, _ := json.Marshal(updateSection)
		// realizar a chamada de request
		rr := createServer(
			mockService, 
			http.MethodPatch, 
			"/api/v1/sections/1", 
			string(updateSectionByte),
		)
		// análisa o retorno da response
		assert.Equal(t, 200, rr.Code)
		// gera uma estrutura compativel com a que sera retornado no body do response
		objRes := struct {
			Code int 
			Data section.Section
		}{}
		// obtem os dados da response
		err := getData(rr.Body.Bytes(), &objRes) 
		assert.Nil(t, err)
		assert.Equal(t, objRes.Code, rr.Code)
		assert.ObjectsAreEqual(updateSectionRes, objRes.Data)
	})
	/*
	t.Run("update section, error (unprocessableEntity 422)", func(t *testing.T) {
		updateSectionError := struct {
			SectionNumber int
			CurrentTemperature int
			MinimumTemperature int
			CurrentCapacity int
			MinimumCapacity int
			MaximumCapacity int
			WarehouseId float64
			ProductTypeId float64
		}{}
		updateSectionError.SectionNumber = 123
		updateSectionError.CurrentTemperature = 123
		updateSectionError.MinimumTemperature = 123
		updateSectionError.CurrentCapacity = 123
		updateSectionError.MinimumCapacity = 123
		updateSectionError.MaximumCapacity = 123
		updateSectionError.WarehouseId = 10.2
		updateSectionError.ProductTypeId = 1.12
		updateSectionByte, _ = json.Marshal(updateSectionError)

		req, rr := getRequest(http.MethodPatch, "/api/v1/sections/1", string(updateSectionByte))
		
		defer req.Body.Close()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)

		// obtem o corpo do response(um section atualizado)
		var reqSection sectionTest

		err := getData(rr.Body.Bytes(), reqSection)
		assert.Error(t, err)
		assert.ObjectsAreEqual(updateSection, reqSection)
	})
	t.Run("update section, error (not found 404)", func(t *testing.T) {
		req, rr := getRequest(http.MethodPatch, "/api/v1/sections/2", string(updateSectionByte))
		
		defer req.Body.Close()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
	t.Run("update section, sucesso (not found 404)", func(t *testing.T) {
		req, rr := getRequest(http.MethodPatch, "/api/v1/sections/2s", string(updateSectionByte))
		
		defer req.Body.Close()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
*/
}


func TestSectionDelete(t *testing.T) {
	// criar um mock do service 
	var mockService *mocks.SectionService = &mocks.SectionService{}
	var sectionListRes = []section.Section{
		{
			Id: 1,
			SectionNumber: 1,
			CurrentTemperature: 2, 
			MinimumTemperature: 2,
			CurrentCapacity: 2,
			MinimumCapacity: 2,
			MaximumCapacity: 2,
			WarehouseId: 2,
			ProductTypeId: 2,
		},
	}
	
	t.Run("delete sucesso, (not content 204)", func(t *testing.T) {
		// realizar a chamada do metodo que será testado
		mockService.On("ListarSectionAll").Return(sectionListRes, nil).Once()
		mockService.On("DeleteSection",  mock.AnythingOfType("int")).
			Return(nil). 
			Once()
		
		// realiza a chamado no server
		rr := createServer(mockService, http.MethodDelete, "/api/v1/sections/1", "")

		assert.Equal(t, http.StatusNoContent, rr.Code)
	})
	t.Run("delete error, (not found 404)", func(t *testing.T) {
		errNotFound :=  errors.New("section not found")
		// realizar a chamada do metodo que será testado
		mockService.On("ListarSectionAll").Return(sectionListRes, nil).Once()
		mockService.On("DeleteSection",  mock.AnythingOfType("int")).
			Return(errNotFound). 
			Once()
		
		// realiza a chamado no server
		rr := createServer(mockService, http.MethodDelete, "/api/v1/sections/1", "")

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
	t.Run("delete error, (not found 404)", func(t *testing.T) {
		// realizar a chamada do metodo que será testado
		mockService.On("ListarSectionAll").Return(sectionListRes, nil).Once()
		mockService.On("DeleteSection",  mock.AnythingOfType("int")).
			Return(nil). 
			Once()
		
		// realiza a chamado no server
		rr := createServer(mockService, http.MethodDelete, "/api/v1/sections/1s", "")

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

}
