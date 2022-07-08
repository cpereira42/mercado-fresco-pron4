package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// * Inicia uma request
// @param method string
// @param url string
// @param body string
func CreateRequestServerPB(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
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
func CreateServerPB(serv *mocks.ServicePB, method, url, body string) *httptest.ResponseRecorder {
	sectionController := NewProductBatChesController(serv)
	router := gin.Default()
	router.GET("/api/v1/sections/reportProducts", sectionController.ReadPB())
	router.POST("/api/v1/productBatches", sectionController.CreatePB())
	req, rr := CreateRequestServerPB(method, url, body)
	router.ServeHTTP(rr, req)
	return rr
}

var productBatchesList []productbatch.ProductBatchesResponse = []productbatch.ProductBatchesResponse{
	{
		SectionId:     1,
		SectionNumber: 1,
		ProductsCount: 1,
	},
}

var productBatchesRes = productbatch.ProductBatchesResponse{
	SectionId:     1,
	SectionNumber: 1,
	ProductsCount: 1,
}

var productBatches productbatch.ProductBatches = productbatch.ProductBatches{
	Id:                 1,
	BatchNumber:        "111",
	CurrentQuantity:    1,
	CurrentTemperature: 1,
	DueDate:            "2022-04-04",
	InitialQuantity:    1,
	ManufacturingDate:  "2020-04-04 14:30:10",
	ManufacturingHour:  "2020-05-01 14:20:15",
	MinimumTemperature: 1,
	ProductId:          1,
	SectionId:          1,
}

func TestServiceCreatePB(t *testing.T) {
	t.Run("create PB, sucesso (201)", func(t *testing.T) {
		mockPB := &mocks.ServicePB{}
		mockPB.On("CreatePB", mock.Anything).
			Return(productBatches, nil).
			Once()
		byteProductBatches, err := json.Marshal(productBatches)
		assert.NoError(t, err)
		rr := CreateServerPB(
			mockPB,
			http.MethodPost,
			"/api/v1/productBatches",
			string(byteProductBatches),
		)
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, ObjetoResponse.Code, rr.Code)
		assert.Empty(t, ObjetoResponse.Error)
		assert.ObjectsAreEqual(productBatches, ObjetoResponse.Data)
	})
	t.Run("create PB, error (409)", func(t *testing.T) {
		mockPB := &mocks.ServicePB{}
		errNewPB := errors.New("batch_number_UNIQUE is unique, and 111 already registered")

		mockPB.On("CreatePB", mock.Anything).
			Return(productBatches, errNewPB).
			Once()

		byteProductBatches, err := json.Marshal(productBatches)
		assert.NoError(t, err)
		rr := CreateServerPB(
			mockPB,
			http.MethodPost,
			"/api/v1/productBatches",
			string(byteProductBatches),
		)
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.NotEmpty(t, ObjetoResponse.Error)
	})
	t.Run("create PB, error (422)", func(t *testing.T) {
		mockPB := &mocks.ServicePB{}
		errNewPB := errors.New("This field is required")
		pbNil := productbatch.ProductBatches{}
		mockPBReq := productbatch.ProductBatches{
			BatchNumber: "312243",
		}
		mockPB.On("CreatePB", mock.Anything).
			Return(pbNil, errNewPB).
			Once()

		byteProductBatches, err := json.Marshal(mockPBReq)
		assert.NoError(t, err)
		rr := CreateServerPB(
			mockPB,
			http.MethodPost,
			"/api/v1/productBatches",
			string(byteProductBatches),
		)
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, 422, rr.Code)
		assert.NotEmpty(t, ObjetoResponse.Error)
		assert.ObjectsAreEqual(pbNil, ObjetoResponse.Data)
	})
}

func TestServiceReadPBGetAll(t *testing.T) {
	t.Run("metodo GetAll, (200)", func(t *testing.T) {
		mockServicePB := &mocks.ServicePB{}
		mockServicePB.On("GetAll").
			Return(productBatchesList, nil).
			Once()
		rr := CreateServerPB(
			mockServicePB,
			http.MethodGet,
			"/api/v1/sections/reportProducts",
			"")
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, 200, rr.Code)
		dtByte, _ := json.Marshal(ObjetoResponse.Data)
		pdrList := []productbatch.ProductBatchesResponse{}
		getData(dtByte, &pdrList)
		assert.True(t, len(pdrList) > 0)
	})
	t.Run("metodo GetAll, (500)", func(t *testing.T) {
		mockServicePB := &mocks.ServicePB{}
		pdrList := []productbatch.ProductBatchesResponse{}
		messageErr := errors.New("failed to serialize product_batches_response fields")
		mockServicePB.On("GetAll").
			Return(pdrList, messageErr).
			Once()
		rr := CreateServerPB(
			mockServicePB,
			http.MethodGet,
			"/api/v1/sections/reportProducts",
			"")
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, 500, rr.Code)
		assert.NotEmpty(t, ObjetoResponse.Error)
		assert.Equal(t, messageErr.Error(), ObjetoResponse.Error)
	})
}

func TestServiceReadPBGetId(t *testing.T) {
	t.Run("Método GetId (200)", func(t *testing.T) {
		// retorno do GetId = productbatch.ProductBatchesResponse, error

		mockServicePB := &mocks.ServicePB{}
		objPbResponse := productbatch.ProductBatchesResponse{
			SectionId:     32,
			SectionNumber: 111,
			ProductsCount: 1,
		}

		mockServicePB.On("GetId", mock.AnythingOfType("int64")).
			Return(objPbResponse, nil).
			Once()
		rr := CreateServerPB(
			mockServicePB,
			http.MethodGet,
			"/api/v1/sections/reportProducts?id=1",
			"")
		responseObj := struct {
			Code  int                                 `json:"code"`
			Data  productbatch.ProductBatchesResponse `json:"data"`
			Error string                              `json:"error"`
		}{}

		getData(rr.Body.Bytes(), &responseObj)
		assert.Equal(t, 200, rr.Code)
		assert.ObjectsAreEqual(objPbResponse, responseObj.Data)
		assert.True(t, responseObj.Error == "")
	})
	t.Run("Método GetId (404)", func(t *testing.T) {
		// retorno do GetId = productbatch.ProductBatchesResponse, error
		productBatResponse := productbatch.ProductBatchesResponse{}
		messageErr := errors.New("section_id not found")
		mockServicePB := &mocks.ServicePB{}
		mockServicePB.On("GetId", mock.AnythingOfType("int64")).
			Return(productBatResponse, messageErr).
			Once()
		rr := CreateServerPB(
			mockServicePB,
			http.MethodGet,
			"/api/v1/sections/reportProducts?id=5",
			"")
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, 404, rr.Code)
		assert.ObjectsAreEqual(productBatchesRes, ObjetoResponse.Data)
		assert.NotEmpty(t, ObjetoResponse.Error)
	})
	t.Run("Método GetId (paraId incorrento)", func(t *testing.T) {
		// retorno do GetId = productbatch.ProductBatchesResponse, error
		productBatResponse := productbatch.ProductBatchesResponse{}
		messageErr := errors.New("strconv.ParseInt: parsing \"5s\": invalid syntax")
		mockServicePB := &mocks.ServicePB{}
		mockServicePB.On("GetId", mock.AnythingOfType("int64")).
			Return(productBatResponse, messageErr).
			Once()
		rr := CreateServerPB(
			mockServicePB,
			http.MethodGet,
			"/api/v1/sections/reportProducts?id=5s",
			"")
		getData(rr.Body.Bytes(), &ObjetoResponse)
		assert.Equal(t, 500, rr.Code)
		assert.ObjectsAreEqual(productBatchesRes, ObjetoResponse.Data)
		assert.Equal(t, messageErr.Error(), ObjetoResponse.Error)
	})
}
