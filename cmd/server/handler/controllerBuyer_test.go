package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	buyer1                   = buyer.Buyer{ID: 1, Card_number_ID: "111", First_name: "Adrians", Last_name: "Rosa"}
	buyer2                   = buyer.Buyer{ID: 2, Card_number_ID: "222", First_name: "Adriana", Last_name: "Rosaa"}
	buyer3                   = buyer.Buyer{ID: 3, Card_number_ID: "333", First_name: "Adriane", Last_name: "Rosaaa"}
	buyer4                   = buyer.Buyer{ID: 4, Card_number_ID: "444", First_name: "Ana", Last_name: "Banana"}
	buyersList []buyer.Buyer = []buyer.Buyer{buyer1, buyer2, buyer3}
)

func createRequestTestBuyer(
	method string,
	url string,
	body string,
) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func createServerBuyer(serv *mocks.Service, method string, url string, body string) *httptest.ResponseRecorder {
	p := handler.NewBuyer(serv)
	r := gin.Default()
	pr := r.Group("/api/v1/buyers")
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetID())
	pr.DELETE("/:id", p.Delete())
	pr.POST("/", p.Create())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.Update())
	req, rr := createRequestTest(method, url, body)
	r.ServeHTTP(rr, req)
	return rr
}

func Test_GetAllBuyer(t *testing.T) {

	t.Run("Test find_all", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("GetAll").Return(buyersList, nil)
		rr := createServerBuyer(servMock, http.MethodGet, "/api/v1/buyers/", "")

		assert.Equal(t, http.StatusOK, rr.Code)

		response := struct {
			Code int
			Data []buyer.Buyer
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, buyersList, response.Data)
	})

	t.Run("Test find_all_NotFound", func(t *testing.T) {
		servMock := &mocks.Service{}
		servMock.On("GetAll").Return([]buyer.Buyer{}, fmt.Errorf("not_found"))
		rr := createServerBuyer(servMock, http.MethodGet, "/api/v1/buyers/", "")

		assert.Equal(t, http.StatusNotFound, rr.Code)

		response := struct {
			Code  int
			Error string
			Data  buyer.Buyer
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, "not_found", response.Error)
	})
}

func Test_GetByIdBuyer(t *testing.T) {

	t.Run("Test find_by_id_existent", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("GetId", 1).Return(buyer1, nil)
		rr := createServerBuyer(servMock, http.MethodGet, "/api/v1/buyers/1", "")

		assert.Equal(t, http.StatusOK, rr.Code)

		response := struct {
			Code int
			Data buyer.Buyer
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, buyer1, response.Data)
	})

	t.Run("Test find_by_id_non_existent", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("GetId", 4).Return(buyer.Buyer{}, fmt.Errorf("buyer_not_found"))
		rr := createServerBuyer(servMock, http.MethodGet, "/api/v1/buyers/4", "")

		response := struct {
			Code  int
			Error string
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "buyer_not_found", response.Error)
	})

	t.Run("Test invalid_id", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("GetId", "A").Return(buyer.Buyer{}, fmt.Errorf("Invalid ID"))
		rr := createServerBuyer(servMock, http.MethodGet, "/api/v1/buyers/A", "")

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

func Test_CreateBuyer(t *testing.T) {

	t.Run("Test create_ok", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("Create",
			buyer4.Card_number_ID,
			buyer4.First_name,
			buyer4.Last_name).
			Return(buyer4, nil)

		rr := createServerBuyer(servMock, http.MethodPost, "/api/v1/buyers/", `{
			"card_number_id": "444",
			"first_name": "Ana",
			"last_name": "Banana"			
			}`)

		response := struct {
			Code int
			Data buyer.Buyer
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, buyer4, response.Data)
		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("Test create_fail", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("Create",
			buyer4.Card_number_ID,
			buyer4.First_name,
			buyer4.Last_name).
			Return(buyer.Buyer{}, fmt.Errorf("create_fail"))

		rr := createServerBuyer(servMock, http.MethodPost, "/api/v1/buyers/", `{
			"card_number_id": "444",
			"first_name": "Ana",
			"last_name": "Banana"			
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
			buyer4.Card_number_ID,
			buyer4.First_name,
			buyer4.Last_name).
			Return(buyer.Buyer{}, fmt.Errorf("create_conflict: this buyer already exists"))

		rr := createServerBuyer(servMock, http.MethodPost, "/api/v1/buyers/", `{
			"card_number_id": "444",
			"first_name": "Ana",
			"last_name": "Banana"			
			}`)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)

		response := struct {
			Code  int
			Error string
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, "create_conflict: this buyer already exists", response.Error)
	})
}

func Test_UpdateBuyer(t *testing.T) {
	t.Run("Test update_ok", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("Update",
			buyer4.ID,
			buyer4.Card_number_ID,
			buyer4.First_name,
			buyer4.Last_name).
			Return(buyer4, nil)

		rr := createServerBuyer(servMock, http.MethodPatch, "/api/v1/buyers/4", `{
				"card_number_id": "444",
				"first_name": "Ana",
				"last_name": "Banana"
				}`)

		response := struct {
			Code int
			Data buyer.Buyer
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, buyer4, response.Data)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Test update_non_existent", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("Update",
			buyer4.ID,
			buyer4.Card_number_ID,
			buyer4.First_name,
			buyer4.Last_name).
			Return(buyer.Buyer{}, fmt.Errorf("Invalid ID"))

		rr := createServerBuyer(servMock, http.MethodPatch, "/api/v1/buyers/A", `{
				"card_number_id": "444",
				"first_name": "Ana",
				"last_name": "Banana"
				}`)

		response := struct {
			Code  int
			Error string
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, "Invalid ID", response.Error)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("Test update_not_found", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("Update",
			buyer4.ID,
			buyer4.Card_number_ID,
			buyer4.First_name,
			buyer4.Last_name).
			Return(buyer.Buyer{}, fmt.Errorf("not_found"))

		rr := createServerBuyer(servMock, http.MethodPatch, "/api/v1/buyers/4", `{
			"card_number_id": "444",
			"first_name": "Ana",
			"last_name": "Banana"
			}`)

		response := struct {
			Code  int
			Error string
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, "not_found", response.Error)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func Test_DeleteBuyer(t *testing.T) {

	t.Run("Test delete_ok", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("Delete", 1).Return(nil)
		rr := createServerBuyer(servMock, http.MethodDelete, "/api/v1/buyers/1", "")

		assert.Equal(t, http.StatusNoContent, rr.Code)
	})

	t.Run("Test delete_non_existent", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("Delete", "A").Return(fmt.Errorf("Invalid ID"), nil)
		rr := createServerBuyer(servMock, http.MethodDelete, "/api/v1/buyers/A", "")

		response := struct {
			Code  int
			Error string
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, "Invalid ID", response.Error)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("Test delete_not_found", func(t *testing.T) {
		servMock := new(mocks.Service)
		servMock.On("Delete", 9).Return(fmt.Errorf("not_found"), nil)
		rr := createServerBuyer(servMock, http.MethodDelete, "/api/v1/buyers/9", "")

		response := struct {
			Code  int
			Error string
		}{}

		err := json.Unmarshal(rr.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, "not_found", response.Error)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
