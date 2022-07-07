package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products/mocks"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var objProducts = struct {
	Code int
	Data []products.Product
}{}

var objProduct = struct {
	Code  int
	Data  products.Product
	Error string
}{}

var prod1 = products.Product{
	Id:                             1,
	Description:                    "prod1",
	ExpirationRate:                 1,
	FreezingRate:                   2,
	Height:                         3.3,
	Length:                         4.3,
	NetWeight:                      5.5,
	ProductCode:                    "prod1",
	RecommendedFreezingTemperature: 6.6,
	Width:                          7.7,
	ProductTypeId:                  8,
	SellerId:                       9,
}

var prod2 = products.Product{
	Id:                             2,
	Description:                    "prod2",
	ExpirationRate:                 1,
	FreezingRate:                   2,
	Height:                         3.3,
	Length:                         4.3,
	NetWeight:                      5.5,
	ProductCode:                    "prod2",
	RecommendedFreezingTemperature: 6.6,
	Width:                          7.7,
	ProductTypeId:                  8,
	SellerId:                       9,
}

var prod3 = products.Product{
	Description:                    "prod3",
	ExpirationRate:                 1,
	FreezingRate:                   2,
	Height:                         3.3,
	Length:                         4.3,
	NetWeight:                      5.5,
	ProductCode:                    "prod3",
	RecommendedFreezingTemperature: 6.6,
	Width:                          7.7,
	ProductTypeId:                  8,
	SellerId:                       9,
}

var prodNew = products.RequestProductsCreate{
	Description:                    "prod1",
	ExpirationRate:                 1,
	FreezingRate:                   2,
	Height:                         3.3,
	Length:                         4.3,
	NetWeight:                      5.5,
	ProductCode:                    "prod1",
	RecommendedFreezingTemperature: 6.6,
	Width:                          7.7,
	ProductTypeId:                  8,
	SellerId:                       9,
}

func createServer(serv *mocks.Service, method string, url string, body string) *httptest.ResponseRecorder {
	p := handler.NewProduct(serv)
	r := gin.Default()
	pr := r.Group("/api/v1/products")
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetId())
	pr.DELETE("/:id", p.Delete())
	pr.POST("/", p.Create())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.Update())
	req, rr := util.CreateRequestTest(method, url, body)
	r.ServeHTTP(rr, req)
	return rr
}

func Test_RepositoryGetAll(t *testing.T) {
	produtos := []products.Product{prod1, prod2}

	t.Run("Get All", func(t *testing.T) {
		serv := &mocks.Service{}
		serv.On("GetAll").Return(produtos, nil)
		rr := createServer(serv, http.MethodGet, "/api/v1/products/", "")
		assert.Equal(t, http.StatusOK, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProducts)
		assert.Equal(t, objProducts.Data, produtos)
		assert.Nil(t, err)
		assert.True(t, len(objProducts.Data) > 0)
		serv.AssertExpectations(t)
	})

	t.Run("Get All Fail", func(t *testing.T) {
		serv := &mocks.Service{}
		serv.On("GetAll").Return([]products.Product{}, fmt.Errorf("error"))

		rr := createServer(serv, http.MethodGet, "/api/v1/products/", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProducts)
		assert.Equal(t, objProducts.Code, http.StatusNotFound)
		assert.Nil(t, err)
		serv.AssertExpectations(t)
	})
}

func Test_RepositoryGetId(t *testing.T) {
	produtos := []products.Product{prod1, prod2}

	t.Run("Get Id", func(t *testing.T) {
		serv := &mocks.Service{}
		serv.On("GetId", 1).Return(produtos[0], nil)

		rr := createServer(serv, http.MethodGet, "/api/v1/products/1", "")
		assert.Equal(t, http.StatusOK, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProduct)
		assert.Equal(t, objProduct.Data, produtos[0])
		assert.Nil(t, err)
		serv.AssertExpectations(t)
	})

	t.Run("Get Id Wrong id ", func(t *testing.T) {
		serv := &mocks.Service{}
		rr := createServer(serv, http.MethodGet, "/api/v1/products/A", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)

		err := json.Unmarshal(rr.Body.Bytes(), &objProduct)
		assert.Nil(t, err)
		serv.AssertExpectations(t)
	})

	t.Run("Get Id Non Exist ", func(t *testing.T) {
		serv := &mocks.Service{}
		serv.On("GetId", 9).Return(products.Product{}, fmt.Errorf("error"))
		rr := createServer(serv, http.MethodGet, "/api/v1/products/9", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProduct)
		assert.Nil(t, err)
		serv.AssertExpectations(t)
	})
}

func Test_RepositoryDelete(t *testing.T) {
	t.Run("Delete Ok", func(t *testing.T) {
		serv := &mocks.Service{}

		serv.On("Delete", 1).Return(nil)
		rr := createServer(serv, http.MethodDelete, "/api/v1/products/1", "")
		assert.Equal(t, http.StatusNoContent, rr.Code)
		serv.AssertExpectations(t)
	})

	t.Run("Delete Fail id = A", func(t *testing.T) {
		serv := &mocks.Service{}

		rr := createServer(serv, http.MethodDelete, "/api/v1/products/A", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)

		err := json.Unmarshal(rr.Body.Bytes(), &objProduct)
		assert.Nil(t, err)
		assert.NotNil(t, objProduct.Error)
		assert.Equal(t, "Invalid ID", objProduct.Error)
		serv.AssertExpectations(t)
	})

	t.Run("Delete Fail id non exist", func(t *testing.T) {
		serv := &mocks.Service{}

		serv.On("Delete", 9).Return(fmt.Errorf("Product 9 not found"))

		rr := createServer(serv, http.MethodDelete, "/api/v1/products/9", "")
		assert.Equal(t, http.StatusNotFound, rr.Code)

		err := json.Unmarshal(rr.Body.Bytes(), &objProduct)
		assert.Nil(t, err)
		assert.NotNil(t, objProduct.Error)
		assert.Equal(t, "Product 9 not found", objProduct.Error)
		serv.AssertExpectations(t)
	})
}

func Test_RepositoryCreate(t *testing.T) {
	t.Run("Creat Ok", func(t *testing.T) {
		serv := &mocks.Service{}
		serv.On("Create", prodNew).Return(prod1, nil)

		rr := createServer(serv, http.MethodPost, "/api/v1/products/", `{
			"Description":                    "prod1",
			"Expiration_rate":                 1,
			"Freezing_rate":                   2,
			"Height":                         3.3,
			"Length":                         4.3,
			"Net_Weight":                      5.5,
			"Product_code":                    "prod1",
			"Recommended_freezing_temperature": 6.6,
			"Width":                          7.7,
			"Product_type_id":                  8,
			"Seller_id":                       9}`)
		assert.Equal(t, http.StatusCreated, rr.Code)

		err := json.Unmarshal(rr.Body.Bytes(), &objProduct)
		assert.Equal(t, prod1, objProduct.Data)
		assert.Nil(t, err)
		serv.AssertExpectations(t)
	})

	t.Run("Creat Fail - less field", func(t *testing.T) {
		serv := &mocks.Service{}
		rr := createServer(serv, http.MethodPost, "/api/v1/products/", `{
					"Description":                    "prod1",
					"Expiration_rate":                 1,
					"Freezing_rate":                   2,
					"Height":                         3.3,
					"Length":                         4.3,
					"Product_code":                    "prod1",
					"Recommended_freezing_temperature": 6.6,
					"Width":                          7.7,
					"Product_type_id":                  8,
					"Seller_id":                       9}`)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)

		objRes := struct {
			Code int
			Data []struct {
				Field   string
				Message string
			}
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.Equal(t, "This field is required", objRes.Data[0].Message)
	})

	t.Run("Creat Fail - Duplicated", func(t *testing.T) {
		serv := &mocks.Service{}
		serv.On("Create", prodNew).Return(products.Product{}, fmt.Errorf("product_code is unique, and prod1 already registered"))
		rr := createServer(serv, http.MethodPost, "/api/v1/products/", `{
					"Description":                    "prod1",
					"Expiration_rate":                 1,
					"Freezing_rate":                   2,
					"Height":                         3.3,
					"Length":                         4.3,
					"Net_Weight":                      5.5,
					"Product_code":                    "prod1",
					"Recommended_freezing_temperature": 6.6,
					"Width":                          7.7,
					"Product_type_id":                  8,
					"Seller_id":                       9}`)

		assert.Equal(t, http.StatusConflict, rr.Code)

		err := json.Unmarshal(rr.Body.Bytes(), &objProduct)
		assert.Nil(t, err)
		assert.Equal(t, "product_code is unique, and prod1 already registered", objProduct.Error)
		serv.AssertExpectations(t)
	})

	t.Run("Creat Fail", func(t *testing.T) {
		serv := &mocks.Service{}

		serv.On("Create", prodNew).Return(products.Product{}, fmt.Errorf("Error to save"))

		rr := createServer(serv, http.MethodPost, "/api/v1/products/", `{
				"Description":                    "prod1",
				"Expiration_rate":                 1,
				"Freezing_rate":                   2,
				"Height":                         3.3,
				"Length":                         4.3,
				"Net_Weight":                      5.5,
				"Product_code":                    "prod1",
				"Recommended_freezing_temperature": 6.6,
				"Width":                          7.7,
				"Product_type_id":                  8,
				"Seller_id":                       9}`)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		serv.AssertExpectations(t)
	})
}

func Test_RepositoryUpdate(t *testing.T) {
	t.Run("Update Ok", func(t *testing.T) {
		prodNew := products.RequestProductsUpdate{
			Description: "prod10",
		}

		serv := &mocks.Service{}
		prod1.Description = "prod10"
		serv.On("Update", 1, prodNew).Return(prod1, nil)

		rr := createServer(serv, http.MethodPatch, "/api/v1/products/1", `{
			"Description":                    "prod10"
			}`)

		assert.Equal(t, http.StatusOK, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProduct)
		assert.Nil(t, err)
		assert.Equal(t, prod1, objProduct.Data)
		serv.AssertExpectations(t)
	})

	t.Run("Update ID A", func(t *testing.T) {
		serv := &mocks.Service{}
		rr := createServer(serv, http.MethodPatch, "/api/v1/products/A", `{
			"Description":                    "prod10"
			}`)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProduct)
		assert.Nil(t, err)
		assert.Equal(t, "Invali ID", objProduct.Error)
		serv.AssertExpectations(t)
	})

	t.Run("Update ID non exist", func(t *testing.T) {
		prodNew := products.RequestProductsUpdate{
			Description: "prod10",
		}

		serv := &mocks.Service{}

		serv.On("Update", 99, prodNew).Return(products.Product{}, fmt.Errorf("data not found"))
		rr := createServer(serv, http.MethodPatch, "/api/v1/products/99", `{
			"Description":                    "prod10"
			}`)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &objProduct)
		assert.Nil(t, err)
		assert.Equal(t, "data not found", objProduct.Error)
		serv.AssertExpectations(t)
	})

	t.Run("Update Fail", func(t *testing.T) {
		prodNew := products.RequestProductsUpdate{
			Description: "prod10",
		}

		serv := &mocks.Service{}

		serv.On("Update", 99, prodNew).Return(products.Product{}, fmt.Errorf("Error to save"))
		rr := createServer(serv, http.MethodPatch, "/api/v1/products/99", `{
			"Description":                    "prod10"
			}`)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)

		err := json.Unmarshal(rr.Body.Bytes(), &objProduct)
		assert.Nil(t, err)
		assert.Equal(t, "Error to save", objProduct.Error)
		serv.AssertExpectations(t)
	})

	t.Run("Update Fail - invalid args", func(t *testing.T) {
		serv := &mocks.Service{}
		rr := createServer(serv, http.MethodPatch, "/api/v1/products/99", `{
			"Description":                    10
			}`)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)

		objRes := struct {
			Code int
			Data struct {
				Field   string
				Message string
			}
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.Equal(t, "description", objRes.Data.Field)
		serv.AssertExpectations(t)
	})
}
