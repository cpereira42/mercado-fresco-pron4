package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productsRecords"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productsRecords/mocks"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var objProductsRecords = struct {
	Code  int
	Data  []productsRecords.ReturnProductRecords
	Error string
}{}

var objProductsRecordsRet = struct {
	Code  int
	Data  productsRecords.ProductRecords
	Error string
}{}

var objProductsRecord = struct {
	Code  int
	Data  productsRecords.ReturnProductRecords
	Error string
}{}

var prodRecordsReturn = productsRecords.ProductRecords{
	Id:             12,
	LastUpdateDate: "2022-07-05 17:04:08",
	PurchasePrice:  150,
	SalePrice:      150,
	ProductId:      1}

var prod1Records = productsRecords.ReturnProductRecords{

	ProductId:    2,
	RecordsCount: 1,
	Description:  "notebook"}

var prod2Records = productsRecords.ReturnProductRecords{

	ProductId:    6,
	RecordsCount: 8,
	Description:  "celular"}

var prodNewRecords = productsRecords.RequestProductRecordsCreate{
	PurchasePrice: 150,
	SalePrice:     150,
	ProductId:     1,
}

func createServerRecordsProducts(serv *mocks.Service, method string, url string, body string) *httptest.ResponseRecorder {

	r := gin.Default()
	handler.NewProductRecords(r, serv)
	req, rr := util.CreateRequestTest(method, url, body)
	r.ServeHTTP(rr, req)
	return rr
}

func Test_RepositoryIDRecordsProducts(t *testing.T) {
	t.Run("Get All Fail", func(t *testing.T) {
		produtos := []productsRecords.ReturnProductRecords(nil)
		serv := &mocks.Service{}
		serv.On("GetAllRecords").Return(produtos, fmt.Errorf("Invalid ID"))
		rr := createServerRecordsProducts(serv, http.MethodGet, "/api/v1/products/reportRecords/", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProductsRecords)
		assert.Equal(t, produtos, objProductsRecords.Data)
		assert.Nil(t, err)
		assert.Equal(t, "Invalid ID", objProductsRecords.Error)
		serv.AssertExpectations(t)
	})

	t.Run("Get All OK", func(t *testing.T) {
		produtos := []productsRecords.ReturnProductRecords{prod1Records, prod2Records}
		serv := &mocks.Service{}
		serv.On("GetAllRecords").Return(produtos, nil)
		rr := createServerRecordsProducts(serv, http.MethodGet, "/api/v1/products/reportRecords/", "")
		assert.Equal(t, http.StatusOK, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProductsRecords)
		assert.Equal(t, produtos, objProductsRecords.Data)
		assert.Nil(t, err)
		assert.True(t, len(objProductsRecords.Data) > 0)
		serv.AssertExpectations(t)
	})

	t.Run("Get ID Not Found", func(t *testing.T) {
		serv := &mocks.Service{}
		produto := productsRecords.ReturnProductRecords{}
		serv.On("GetIdRecords", 5).Return(produto, fmt.Errorf("Not Found"))
		rr := createServerRecordsProducts(serv, http.MethodGet, "/api/v1/products/reportRecords/?id=5", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProductsRecord)
		assert.Equal(t, produto, objProductsRecord.Data)
		assert.Nil(t, err)
		assert.Equal(t, "Not Found", objProductsRecord.Error)
		serv.AssertExpectations(t)

	})

	t.Run("Get ID A", func(t *testing.T) {
		serv := &mocks.Service{}
		produto := productsRecords.ReturnProductRecords{}
		rr := createServerRecordsProducts(serv, http.MethodGet, "/api/v1/products/reportRecords/?id=A", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProductsRecord)

		assert.Equal(t, produto, objProductsRecord.Data)
		assert.Nil(t, err)
		serv.AssertExpectations(t)

	})

	t.Run("Get ID Ok", func(t *testing.T) {
		serv := &mocks.Service{}
		serv.On("GetIdRecords", 6).Return(prod2Records, nil)
		rr := createServerRecordsProducts(serv, http.MethodGet, "/api/v1/products/reportRecords/?id=6", "")
		assert.Equal(t, http.StatusOK, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProductsRecord)

		assert.Equal(t, prod2Records, objProductsRecord.Data)
		assert.Nil(t, err)
		serv.AssertExpectations(t)

	})

}

func Test_RepositoryCreateRecordsProducts(t *testing.T) {
	t.Run("Create Fail less fields", func(t *testing.T) {
		serv := &mocks.Service{}
		rr := createServerRecordsProducts(serv, http.MethodPost, "/api/v1/productsRecords/", `{
			purchase_price: 150,
			product_id: 150,
		}`)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		serv.AssertExpectations(t)
	})

	t.Run("Create product not found", func(t *testing.T) {

		serv := &mocks.Service{}
		serv.On("Create", prodNewRecords).Return(productsRecords.ProductRecords{}, fmt.Errorf("product_id is not registered on products"))

		rr := createServerRecordsProducts(serv, http.MethodPost, "/api/v1/productsRecords/", `{
			"purchase_price": 150,
			"sale_price":150,
			"product_id": 1
		}`)
		assert.Equal(t, http.StatusConflict, rr.Code)
		serv.AssertExpectations(t)
	})

	t.Run("CreateFail other", func(t *testing.T) {

		serv := &mocks.Service{}
		serv.On("Create", prodNewRecords).Return(productsRecords.ProductRecords{}, fmt.Errorf("Error"))

		rr := createServerRecordsProducts(serv, http.MethodPost, "/api/v1/productsRecords/", `{
			"purchase_price": 150,
			"sale_price":150,
			"product_id": 1
		}`)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		serv.AssertExpectations(t)
	})

	t.Run("CreateOk", func(t *testing.T) {

		serv := &mocks.Service{}
		serv.On("Create", prodNewRecords).Return(prodRecordsReturn, nil)

		rr := createServerRecordsProducts(serv, http.MethodPost, "/api/v1/productsRecords/", `{
			"purchase_price": 150,
			"sale_price":150,
			"product_id": 1
		}`)
		assert.Equal(t, http.StatusCreated, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProductsRecordsRet)
		assert.NoError(t, err)
		assert.Equal(t, prodRecordsReturn, objProductsRecordsRet.Data)
		serv.AssertExpectations(t)
	})

}
