package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	//"github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/mocks"
	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	//products "github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/repository"
)

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
	SellerId:                       9}

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
	SellerId:                       9}

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
	SellerId:                       9}
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
	SellerId:                       9}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func Test_RepositoryGetAll(t *testing.T) {
	produtos := []products.Product{prod1, prod2}

	t.Run("Get All", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodGet, "/api/v1/products/", "")
		serv := &mocks.Service{}
		serv.On("GetAll").Return(produtos, nil)

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.GET("/", p.GetAll())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)

		objRes := struct {
			Code int
			Data []products.Product
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Equal(t, objRes.Data, produtos)
		assert.Nil(t, err)
		assert.True(t, len(objRes.Data) > 0)
	})

	t.Run("Get All Fail", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodGet, "/api/v1/products/", "")
		serv := &mocks.Service{}
		serv.On("GetAll").Return([]products.Product{}, fmt.Errorf("error"))

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.GET("/", p.GetAll())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 401, rr.Code)

		objRes := struct {
			Code int
			Data []products.Product
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Equal(t, objRes.Code, 401)
		//assert.Equal(t, objRes.Data, produtos)
		assert.Nil(t, err)
		//assert.True(t, len(objRes.Data) > 0)
	})
}

func Test_RepositoryGetId(t *testing.T) {
	produtos := []products.Product{prod1, prod2}

	t.Run("Get Id", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodGet, "/api/v1/products/1", "")
		serv := &mocks.Service{}
		serv.On("GetId", 1).Return(produtos[0], nil)

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.GET("/:id", p.GetId())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)

		objRes := struct {
			Code int
			Data products.Product
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Equal(t, objRes.Data, produtos[0])
		assert.Nil(t, err)
	})

	t.Run("Get Id Wrong id ", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodGet, "/api/v1/products/A", "")
		serv := &mocks.Service{}
		serv.On("GetId", "A").Return(products.Product{}, fmt.Errorf("error"))

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.GET("/:id", p.GetId())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		objRes := struct {
			Code int
			Data products.Product
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
	})

	t.Run("Get Id Non Exist ", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodGet, "/api/v1/products/9", "")
		serv := &mocks.Service{}
		serv.On("GetId", 9).Return(products.Product{}, fmt.Errorf("error"))

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.GET("/:id", p.GetId())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		objRes := struct {
			Code int
			Data products.Product
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
	})

	/*t.Run("Get All Fail", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodGet, "/api/v1/products/", "")
		serv := &mocks.Service{}
		serv.On("GetAll").Return([]products.Product{}, fmt.Errorf("error"))

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.GET("/", p.GetAll())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 401, rr.Code)

		objRes := struct {
			Code int
			Data []products.Product
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Equal(t, objRes.Code, 401)
		//assert.Equal(t, objRes.Data, produtos)
		assert.Nil(t, err)
		//assert.True(t, len(objRes.Data) > 0)
	})*/

}

func Test_RepositoryDelete(t *testing.T) {

	t.Run("Delete Ok", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodDelete, "/api/v1/products/1", "")
		serv := &mocks.Service{}

		serv.On("Delete", 1).Return(nil)

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.DELETE("/:id", p.Delete())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 204, rr.Code)
	})

	t.Run("Delete Fail id = A", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodDelete, "/api/v1/products/A", "")
		serv := &mocks.Service{}

		serv.On("Delete").Return(nil)

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.DELETE("/:id", p.Delete())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		objRes := struct {
			Code  int
			Error string
		}{}
		log.Println(string(rr.Body.Bytes()))

		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.NotNil(t, objRes.Error)
		assert.Equal(t, "Invalid ID", objRes.Error)
	})

	t.Run("Delete Fail id non exist", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodDelete, "/api/v1/products/9", "")
		serv := &mocks.Service{}

		serv.On("Delete", 9).Return(fmt.Errorf("Product 9 not found"))

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.DELETE("/:id", p.Delete())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		objRes := struct {
			Code  int
			Error string
		}{}
		log.Println(string(rr.Body.Bytes()))

		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.NotNil(t, objRes.Error)
		assert.Equal(t, "Product 9 not found", objRes.Error)
	})

	/*t.Run("Get All Fail", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodGet, "/api/v1/products/", "")
		serv := &mocks.Service{}
		serv.On("GetAll").Return([]products.Product{}, fmt.Errorf("error"))

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.GET("/", p.GetAll())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 401, rr.Code)

		objRes := struct {
			Code int
			Data []products.Product
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Equal(t, objRes.Code, 401)
		//assert.Equal(t, objRes.Data, produtos)
		assert.Nil(t, err)
		//assert.True(t, len(objRes.Data) > 0)
	})*/
}

func Test_RepositoryCreate(t *testing.T) {

	t.Run("Creat Ok", func(t *testing.T) {
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
			SellerId:                       9}

		req, rr := createRequestTest(http.MethodPost, "/api/v1/products/", `{
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

		serv := &mocks.Service{}
		serv.On("CheckCode", 0, "prod1").Return(false)
		serv.On("Create", prodNew).Return(prod1, nil)

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.POST("/", p.Create())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 201, rr.Code)

		objRes := struct {
			Code int
			Data products.Product
		}{}
		log.Println(string(rr.Body.Bytes()))
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		log.Println(objRes, err)
		assert.Equal(t, prod1, objRes.Data)
	})

	t.Run("Creat Fail - less field", func(t *testing.T) {
		var prodNew = products.RequestProductsCreate{
			Description:                    "prod1",
			ExpirationRate:                 1,
			FreezingRate:                   2,
			Height:                         3.3,
			Length:                         4.3,
			ProductCode:                    "prod1",
			RecommendedFreezingTemperature: 6.6,
			Width:                          7.7,
			ProductTypeId:                  8,
			SellerId:                       9}

		req, rr := createRequestTest(http.MethodPost, "/api/v1/products/", `{
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

		serv := &mocks.Service{}
		serv.On("CheckCode", 0, "prod1").Return(false)
		serv.On("Create", prodNew).Return(prod1, fmt.Errorf("This field is required"))

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.POST("/", p.Create())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 422, rr.Code)

		objRes := struct {
			Code int
			Data []struct {
				Field   string
				Message string
			}
		}{}
		log.Println(string(rr.Body.Bytes()))
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		log.Println(objRes, err)
		assert.Equal(t, "This field is required", objRes.Data[0].Message)
	})

	t.Run("Creat Fail - Duplicated", func(t *testing.T) {
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
			SellerId:                       9}

		req, rr := createRequestTest(http.MethodPost, "/api/v1/products/", `{
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

		serv := &mocks.Service{}
		serv.On("CheckCode", 0, "prod1").Return(true)
		serv.On("Create", prodNew).Return(products.Product{}, fmt.Errorf("Product prod1 already registred"))

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.POST("/", p.Create())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 409, rr.Code)

		objRes := struct {
			Code  int
			Error string
		}{}
		log.Println(string(rr.Body.Bytes()))
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		log.Println(objRes, err)
		assert.Equal(t, "Product prod1 already registred", objRes.Error)
	})

	t.Run("Creat Fail", func(t *testing.T) {
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
			SellerId:                       9}

		req, rr := createRequestTest(http.MethodPost, "/api/v1/products/", `{
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

		serv := &mocks.Service{}
		serv.On("CheckCode", 0, "prod10").Return(false)
		serv.On("Create", prodNew).Return(products.Product{}, fmt.Errorf("Error to save"))

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.POST("/", p.Create())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 422, rr.Code)

		objRes := struct {
			Code int
			Data products.Product
		}{}
		log.Println(string(rr.Body.Bytes()))
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		log.Println(objRes, err)
		//assert.Equal(t, prod1, objRes.Data)
	})
}

func Test_RepositoryUpdate(t *testing.T) {

	t.Run("Update Ok", func(t *testing.T) {
		var prodNew = products.RequestProductsUpdate{
			Description: "prod10",
		}

		req, rr := createRequestTest(http.MethodPatch, "/api/v1/products/1", `{
				"Description":                    "prod10"
				}`)

		serv := &mocks.Service{}
		serv.On("getId", 1).Return(prod1, nil)
		serv.On("CheckCode", 1, "prod1").Return(false)
		prod1.Description = "prod10"
		serv.On("Update", 1, prodNew).Return(prod1, nil)

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.PATCH("/:id", p.Update())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)

		objRes := struct {
			Code int
			Data products.Product
		}{}
		log.Println(string(rr.Body.Bytes()))
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		log.Println(objRes, err)
		assert.Equal(t, prod1, objRes.Data)
	})

	t.Run("Update ID A", func(t *testing.T) {
		var prodNew = products.RequestProductsUpdate{
			Description: "prod10",
		}

		req, rr := createRequestTest(http.MethodPatch, "/api/v1/products/A", `{
				"Description":                    "prod10"
				}`)

		serv := &mocks.Service{}

		serv.On("Update", 1, prodNew).Return(prod1, nil)

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.PATCH("/:id", p.Update())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		objRes := struct {
			Code  int
			Error string
		}{}
		log.Println(string(rr.Body.Bytes()))
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		log.Println(objRes, err)
		assert.Equal(t, "Invali ID", objRes.Error)
	})

	t.Run("Update ID non exist", func(t *testing.T) {
		var prodNew = products.RequestProductsUpdate{
			Description: "prod10",
		}

		req, rr := createRequestTest(http.MethodPatch, "/api/v1/products/99", `{
				"Description":                    "prod10"
				}`)

		serv := &mocks.Service{}

		serv.On("Update", 99, prodNew).Return(products.Product{}, fmt.Errorf("Product not found"))

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.PATCH("/:id", p.Update())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		objRes := struct {
			Code  int
			Error string
		}{}
		log.Println(string(rr.Body.Bytes()))
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		log.Println(objRes, err)
		assert.Equal(t, "Product not found", objRes.Error)
	})

	t.Run("Update Fail", func(t *testing.T) {
		var prodNew = products.RequestProductsUpdate{
			Description: "prod10",
		}

		req, rr := createRequestTest(http.MethodPatch, "/api/v1/products/99", `{
				"Description":                    "prod10"
				}`)

		serv := &mocks.Service{}

		serv.On("Update", 99, prodNew).Return(products.Product{}, fmt.Errorf("Error to save"))

		p := handler.NewProduct(serv)
		r := gin.Default()
		pr := r.Group("/api/v1/products")
		pr.PATCH("/:id", p.Update())

		r.ServeHTTP(rr, req)
		assert.Equal(t, 422, rr.Code)

		objRes := struct {
			Code  int
			Error string
		}{}
		log.Println(string(rr.Body.Bytes()))
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		log.Println(objRes, err)
		assert.Equal(t, "Error to save", objRes.Error)
	})

}
