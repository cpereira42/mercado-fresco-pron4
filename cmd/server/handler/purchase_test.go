package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/purchaseorders"
	"github.com/cpereira42/mercado-fresco-pron4/internal/purchaseorders/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	purchase1 = purchaseorders.Purchase{
		ID:                1,
		Order_date:        "2022-07-11",
		Order_number:      "1",
		Tracking_code:     "123",
		Buyer_id:          1,
		Product_record_id: 1,
		Order_status_id:   1,
	}

	purchase2 = purchaseorders.Purchase{
		ID:                2,
		Order_date:        "2022-07-11",
		Order_number:      "2",
		Tracking_code:     "222",
		Buyer_id:          2,
		Product_record_id: 2,
		Order_status_id:   2,
	}
)

func createRequestTestPurchase(
	method string,
	url string,
	body string,
) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func createServerPurchase(serv *mocks.Service, method string, url string, body string) *httptest.ResponseRecorder {
	p := handler.NewPurchase(serv)
	r := gin.Default()
	pr := r.Group("/api/v1/purchase")
	pr.GET("/:id", p.GetById())
	pr.POST("/", p.Create())
	req, rr := createRequestTest(method, url, body)
	r.ServeHTTP(rr, req)
	return rr
}

func Test_GetByIdPurchase(t *testing.T) {

	t.Run("Test find_by_id_existent", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("GetById", 1).Return(purchase1, nil)
		rr := createServerPurchase(servMock, http.MethodGet, "/api/v1/purchase/1", "")

		assert.Equal(t, http.StatusOK, rr.Code)

		response := struct {
			Code int
			Data purchaseorders.Purchase
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, purchase1, response.Data)
	})

	t.Run("Test find_by_id_non_existent", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("GetById", 4).Return(purchaseorders.Purchase{}, fmt.Errorf("purchase_order_not_found"))
		rr := createServerPurchase(servMock, http.MethodGet, "/api/v1/purchase/4", "")

		response := struct {
			Code  int
			Error string
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "purchase_order_not_found", response.Error)
	})

	t.Run("Test invalid_id", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("GetById", "A").Return(purchaseorders.Purchase{}, fmt.Errorf("Invalid ID"))
		rr := createServerPurchase(servMock, http.MethodGet, "/api/v1/purchase/A", "")

		response := struct {
			Code  int
			Error string
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "Invalid ID", response.Error)
	})
}

func Test_CreatePurchase(t *testing.T) {

	t.Run("Test create_ok", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("Create",
			purchase2.Order_date,
			purchase2.Order_number,
			purchase2.Tracking_code,
			purchase2.Buyer_id,
			purchase2.Product_record_id,
			purchase2.Order_status_id).
			Return(purchase2, nil)

		rr := createServerPurchase(servMock, http.MethodPost, "/api/v1/purchase/", `{
			"order_date": "2022-07-11",
			"order_number": "2",			
			"tracking_code": "222",
			"buyer_id": 2,
			"product_record_id": 2,
			"order_status_id": 2
			}`)

		response := struct {
			Code int
			Data purchaseorders.Purchase
		}{}
		log.Println(string(rr.Body.Bytes()))
		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, purchase2, response.Data)
		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("Test create_fail", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("Create",
			purchase2.Order_date,
			purchase2.Order_number,
			purchase2.Tracking_code,
			purchase2.Buyer_id,
			purchase2.Product_record_id,
			purchase2.Order_status_id).
			Return(purchaseorders.Purchase{}, fmt.Errorf("create_fail"))

		rr := createServerPurchase(servMock, http.MethodPost, "/api/v1/purchase/", `{
			"order_date": "2022-07-11",
			"order_number": "2",			
			"tracking_code": "222",
			"buyer_id": 2,
			"product_record_id": 2,
			"order_status_id": 2
			}`)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)

		response := struct {
			Code  int
			Error string
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, "create_fail", response.Error)
	})

	t.Run("Test create_conflict", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("Create",
			purchase2.Order_date,
			purchase2.Order_number,
			purchase2.Tracking_code,
			purchase2.Buyer_id,
			purchase2.Product_record_id,
			purchase2.Order_status_id).
			Return(purchaseorders.Purchase{}, fmt.Errorf("create_conflict: this purchase order already exists"))

		rr := createServerPurchase(servMock, http.MethodPost, "/api/v1/purchase/", `{
			"order_date": "2022-07-11",
			"order_number": "2",			
			"tracking_code": "222",
			"buyer_id": 2,
			"product_record_id": 2,
			"order_status_id": 2
			}`)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)

		response := struct {
			Code  int
			Error string
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, "create_conflict: this purchase order already exists", response.Error)
	})
}
