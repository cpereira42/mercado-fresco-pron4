package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {

	dbProd := store.New(store.FileType, "../../internal/repositories/products.json")
	repoProd := products.NewRepositoryProducts(dbProd)
	serviceProd := products.NewService(repoProd)
	p := handler.NewProduct(serviceProd)
	r := gin.Default()
	pr := r.Group("/api/v1/products")
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetId())
	pr.DELETE("/:id", p.Delete())
	pr.POST("/", p.Create())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.Update())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")

	return req, httptest.NewRecorder()
}

func Test_Get(t *testing.T) {

	r := createServer()

	t.Run("test All", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodGet, "/api/v1/products/", "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)

		objRes := struct {
			Code int
			Data []products.Product
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.True(t, len(objRes.Data) > 0)
	})

	t.Run("Test ID valid", func(t *testing.T) {
		prod := products.Product{1, "prod1", "celular", 1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7, 8, 0}
		req, rr := createRequestTest(http.MethodGet, "/api/v1/products/1", "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)

		objRes := struct {
			Code int
			Data products.Product
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.Equal(t, prod, objRes.Data, "devem ser iguais")
	})

	t.Run("Test ID Invalid ", func(t *testing.T) {
		expected := "Product 99 not found"
		req, rr := createRequestTest(http.MethodGet, "/api/v1/products/99", "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		objRes := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.Equal(t, expected, objRes.Error, "devem ser iguais")
	})
}

func Test_Post(t *testing.T) {

	r := createServer()

	t.Run("Test Post Ok", func(t *testing.T) {
		prod := products.Product{
			Id:                             2,
			ProductCode:                    "prod0888",
			Description:                    "celular",
			Width:                          1.1,
			Length:                         2.2,
			Height:                         3.3,
			NetWeight:                      4.4,
			ExpirationRate:                 5.5,
			RecommendedFreezingTemperature: 6.6,
			FreezingRate:                   7,
			ProductTypeId:                  8}

		req, rr := createRequestTest(http.MethodPost, "/api/v1/products/", `{"product_code": "prod0888","description": "celular","width": 1.1,"length": 2.2,"height": 3.3,"net_weight": 4.4,"expiration_rate": 5.5,"recommended_freezing_temperature": 6.6,"freezing_rate": 7,"product_type_id": 8}`)
		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)

		objRes := struct {
			Code int
			Data products.Product
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.Equal(t, prod, objRes.Data, "devem ser iguais")
	})

	t.Run("Test Post Repet ProductCode", func(t *testing.T) {
		expected := "code Product prod0888 already registred"

		//r := createServer()
		req, rr := createRequestTest(http.MethodPost, "/api/v1/products/", `{"product_code": "prod0888","description": "celular","width": 1.1,"length": 2.2,"height": 3.3,"net_weight": 4.4,"expiration_rate": 5.5,"recommended_freezing_temperature": 6.6,"freezing_rate": 7,"product_type_id": 8}`)
		r.ServeHTTP(rr, req)
		assert.Equal(t, 422, rr.Code)

		objRes := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.Equal(t, expected, objRes.Error, "devem ser iguais")
	})

	t.Run("Test Post Product Less Field_Fail", func(t *testing.T) {
		expected := "This field is required"

		//r := createServer()
		req, rr := createRequestTest(http.MethodPost, "/api/v1/products/", `{"product_code": "prod0888","width": 1.1,"length": 2.2,"height": 3.3,"net_weight": 4.4,"expiration_rate": 5.5,"recommended_freezing_temperature": 6.6,"freezing_rate": 7,"product_type_id": 8}`)
		r.ServeHTTP(rr, req)
		assert.Equal(t, 422, rr.Code)

		objRes := struct {
			Code  int
			Error []struct {
				Field   string
				Message string
			}
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.Equal(t, expected, objRes.Error[0].Message, "devem ser iguais")
	})
}

func Test_Patch(t *testing.T) {

	r := createServer()
	t.Run("Test_PatchProduct_OK", func(t *testing.T) {
		prod := products.Product{
			Id:                             2,
			ProductCode:                    "prod0888",
			Description:                    "TV",
			Width:                          1.1,
			Length:                         2.2,
			Height:                         3.3,
			NetWeight:                      4.4,
			ExpirationRate:                 5.5,
			RecommendedFreezingTemperature: 6.6,
			FreezingRate:                   7,
			ProductTypeId:                  8}

		req, rr := createRequestTest(http.MethodPatch, "/api/v1/products/2", `{"description": "TV"}`)
		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)

		objRes := struct {
			Code int
			Data products.Product
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.Equal(t, prod, objRes.Data, "devem ser iguais")
	})

	t.Run("Test_PatchProductProductCodeRepeted_Fail", func(t *testing.T) {
		expected := "code Product prod1 already registred"
		req, rr := createRequestTest(http.MethodPatch, "/api/v1/products/2", `{"product_code": "prod1","description": "TV"}`)
		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		objRes := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.Equal(t, expected, objRes.Error, "devem ser iguais")
	})

	t.Run("Test_PatchProductInvalidId_Fail", func(t *testing.T) {
		expected := "Product 99 not found"
		//r := createServer()
		req, rr := createRequestTest(http.MethodPatch, "/api/v1/products/99", `{"product_code": "prod08","description": "TV"}`)
		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		objRes := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.Equal(t, expected, objRes.Error, "devem ser iguais")
	})

}

func Test_Delete(t *testing.T) {

	r := createServer()
	t.Run("Delete Ok", func(t *testing.T) {
		req, rr := createRequestTest(http.MethodDelete, "/api/v1/products/2", "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, 204, rr.Code)
	})

	t.Run("Delete Fail", func(t *testing.T) {
		expected := "Product 2 not found"
		req, rr := createRequestTest(http.MethodDelete, "/api/v1/products/2", "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, 404, rr.Code)

		objRes := struct {
			Code  int
			Error string
		}{}
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)
		assert.Equal(t, expected, objRes.Error, "devem ser iguais")
	})

}
