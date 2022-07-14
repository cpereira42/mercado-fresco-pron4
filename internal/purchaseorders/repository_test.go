package purchaseorders_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpereira42/mercado-fresco-pron4/internal/purchaseorders"
	"github.com/stretchr/testify/assert"
)

var (
	purchaseR1 = purchaseorders.Purchase{
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

func TestRepositoryGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	purchaseRepo := purchaseorders.NewRepository(db)

	t.Run("Find by ID OK", func(t *testing.T) {
		purchaseList := []purchaseorders.Purchase{purchaseR1, purchase2}
		rows := sqlmock.NewRows([]string{
			"ID",
			"Order_date",
			"Order_number",
			"Tracking_code",
			"Buyer_id",
			"Product_record_id",
			"Order_status_id",
		}).AddRow(
			purchaseList[0].ID,
			purchaseList[0].Order_date,
			purchaseList[0].Order_number,
			purchaseList[0].Tracking_code,
			purchaseList[0].Buyer_id,
			purchaseList[0].Product_record_id,
			purchaseList[0].Order_status_id,
		)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(purchaseorders.GET_PURCHASE_BY_ID))
		stmt.ExpectQuery().WithArgs(1).WillReturnRows(rows)

		result, _ := purchaseRepo.GetById(1)
		assert.NoError(t, err)
		assert.Equal(t, purchaseList[0], result)
	})

	t.Run("Find by ID - query fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectQuery(purchaseorders.GET_PURCHASE_BY_ID).WithArgs(1).WillReturnError(fmt.Errorf("fail to prepare query"))
		_, err = purchaseRepo.GetById(1)

		assert.Equal(t, fmt.Errorf("fail to prepare query"), err)
	})

	t.Run("Find by ID fail", func(t *testing.T) {
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(purchaseorders.GET_PURCHASE_BY_ID))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf("fail to get id"))

		_, err := purchaseRepo.GetById(1)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("fail to get id"), err)
	})
}

func TestRepositoryCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	t.Run("Create OK", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(purchaseorders.CREATE_PURCHASE)).WithArgs(
			purchaseR1.Order_date,
			purchaseR1.Order_number,
			purchaseR1.Tracking_code,
			purchaseR1.Buyer_id,
			purchaseR1.Product_record_id,
			purchaseR1.Order_status_id,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		purchaseRepo := purchaseorders.NewRepository(db)
		_, err := purchaseRepo.Create(
			purchaseR1.Order_date,
			purchaseR1.Order_number,
			purchaseR1.Tracking_code,
			purchaseR1.Buyer_id,
			purchaseR1.Product_record_id,
			purchaseR1.Order_status_id)
		assert.NoError(t, err)
	})

	t.Run("Create fail", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(purchaseorders.CREATE_PURCHASE)).WithArgs(
			purchaseR1.Order_date,
			purchaseR1.Order_number,
			purchaseR1.Tracking_code,
			purchaseR1.Buyer_id,
			purchaseR1.Product_record_id,
			purchaseR1.Order_status_id,
		).WillReturnError(fmt.Errorf("fail to create"))
		purchaseRepo := purchaseorders.NewRepository(db)
		_, err = purchaseRepo.Create(
			purchaseR1.Order_date,
			purchaseR1.Order_number,
			purchaseR1.Tracking_code,
			purchaseR1.Buyer_id,
			purchaseR1.Product_record_id,
			purchaseR1.Order_status_id)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("fail to create"), err)
	})

	t.Run("Create fail - 0 rows affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(purchaseorders.CREATE_PURCHASE)).WithArgs(
			purchaseR1.Order_date,
			purchaseR1.Order_number,
			purchaseR1.Tracking_code,
			purchaseR1.Buyer_id,
			purchaseR1.Product_record_id,
			purchaseR1.Order_status_id,
		).WillReturnResult(sqlmock.NewResult(0, 0))
		purchaseRepo := purchaseorders.NewRepository(db)
		_, err = purchaseRepo.Create(
			purchaseR1.Order_date,
			purchaseR1.Order_number,
			purchaseR1.Tracking_code,
			purchaseR1.Buyer_id,
			purchaseR1.Product_record_id,
			purchaseR1.Order_status_id)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("fail to create"), err)
	})
}
