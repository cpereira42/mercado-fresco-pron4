package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
