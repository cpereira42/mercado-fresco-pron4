package purchaseorders_test

import (
	"fmt"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/purchaseorders"
	"github.com/cpereira42/mercado-fresco-pron4/internal/purchaseorders/mocks"

	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
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
)

func TestServiceCreate(t *testing.T) {
	t.Run("Test create_ok",
		func(t *testing.T) {
			repo := new(mocks.Repository)
			repo.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
			).Return(purchase1, nil)

			service := purchaseorders.NewService(repo)
			newPurchase, err := service.Create(
				purchase1.Order_date,
				purchase1.Order_number,
				purchase1.Tracking_code,
				purchase1.Buyer_id,
				purchase1.Product_record_id,
				purchase1.Order_status_id)

			assert.NoError(t, err)
			assert.Equal(t, purchase1, newPurchase)

		})

	t.Run("Test create_conflict",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("a purchase order with this id already exists")
			repo := new(mocks.Repository)
			repo.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
			).Return(purchaseorders.Purchase{}, errorMsg)

			service := purchaseorders.NewService(repo)
			_, err := service.Create(
				purchase1.Order_date,
				purchase1.Order_number,
				purchase1.Tracking_code,
				purchase1.Buyer_id,
				purchase1.Product_record_id,
				purchase1.Order_status_id)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())

		})

	t.Run("Test missing_field",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("missing field")
			repo := new(mocks.Repository)
			repo.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
			).Return(purchaseorders.Purchase{}, errorMsg).Maybe()

			service := purchaseorders.NewService(repo)
			_, err := service.Create(
				"2022-07-05",
				"111",
				"",
				8,
				9,
				0,
			)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())

		})
}

func TestServiceGetById(t *testing.T) {
	t.Run("Test getbyid_ok",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("GetById", tmock.AnythingOfType("int")).Return(purchase1, nil).Once()

			service := purchaseorders.NewService(repo)
			idPurchase, err := service.GetById(1)

			assert.NoError(t, err)
			assert.Equal(t, purchase1, idPurchase)
		})
	t.Run("Test getbyid_fail",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("the purchase order under this id dont exist")
			repo := &mocks.Repository{}
			repo.On("GetById", tmock.AnythingOfType("int")).Return(purchaseorders.Purchase{}, errorMsg).Once()

			service := purchaseorders.NewService(repo)
			_, err := service.GetById(45)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())
		})
}
