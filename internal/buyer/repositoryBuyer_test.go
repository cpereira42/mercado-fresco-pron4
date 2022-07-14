package buyer_test

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/stretchr/testify/assert"
)

var (
	buyer1 = buyer.Buyer{ID: 1, Card_number_ID: "111", First_name: "Adrians", Last_name: "Rosa"}
	buyer2 = buyer.Buyer{ID: 2, Card_number_ID: "222", First_name: "Adriana", Last_name: "Rosaa"}
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	buyerRepo := buyer.NewRepository(db)

	t.Run("Find all", func(t *testing.T) {
		buyersList := []buyer.Buyer{buyer1, buyer2}
		rows := sqlmock.NewRows([]string{
			"ID",
			"Card_number_ID",
			"First_name",
			"Last_name",
		}).AddRow(
			buyersList[0].ID,
			buyersList[0].Card_number_ID,
			buyersList[0].First_name,
			buyersList[0].Last_name,
		).AddRow(
			buyersList[1].ID,
			buyersList[1].Card_number_ID,
			buyersList[1].First_name,
			buyersList[1].Last_name,
		)
		mock.ExpectQuery(buyer.GET_ALL_BUYERS).WillReturnRows(rows)
		result, err := buyerRepo.GetAll()
		assert.NoError(t, err)

		assert.Equal(t, buyersList[0].Card_number_ID, result[0].Card_number_ID)
		assert.Equal(t, buyersList[1].Card_number_ID, result[1].Card_number_ID)
	})

	t.Run("Find all fail", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"ID",
			"Card_number_ID",
			"First_name",
			"Last_name",
		}).AddRow("", "", "", "")

		mock.ExpectQuery(buyer.GET_ALL_BUYERS).WillReturnRows(rows)
		_, err = buyerRepo.GetAll()

		assert.Error(t, err)
	})
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	buyerRepo := buyer.NewRepository(db)

	t.Run("Find by ID OK", func(t *testing.T) {
		buyersList := []buyer.Buyer{buyer1, buyer2}
		rows := sqlmock.NewRows([]string{
			"ID",
			"Card_number_ID",
			"First_name",
			"Last_name",
		}).AddRow(
			buyersList[0].ID,
			buyersList[0].Card_number_ID,
			buyersList[0].First_name,
			buyersList[0].Last_name,
		)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(buyer.GET_BUYER_BY_ID))
		stmt.ExpectQuery().WithArgs(1).WillReturnRows(rows)

		result, _ := buyerRepo.GetId(1)
		assert.NoError(t, err)
		assert.Equal(t, buyersList[0], result)
	})

	t.Run("Find by ID - query fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectQuery(buyer.GET_BUYER_BY_ID).WithArgs(1).WillReturnError(fmt.Errorf("fail to prepare query"))
		_, err = buyerRepo.GetId(1)

		assert.Equal(t, fmt.Errorf("fail to prepare query"), err)
	})

	t.Run("Find by ID fail", func(t *testing.T) {
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(buyer.GET_BUYER_BY_ID))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf("fail to get id"))

		_, err := buyerRepo.GetId(1)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("fail to get id"), err)
	})
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	t.Run("Create OK", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(buyer.CREATE_BUYER)).WithArgs(
			buyer1.Card_number_ID,
			buyer1.First_name,
			buyer1.Last_name,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		buyerRepo := buyer.NewRepository(db)
		_, err := buyerRepo.Create(buyer1.Card_number_ID, buyer1.First_name, buyer1.Last_name)
		assert.NoError(t, err)
	})

	t.Run("Create fail", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(buyer.CREATE_BUYER)).WithArgs(
			buyer1.Card_number_ID,
			buyer1.First_name,
			buyer1.Last_name,
		).WillReturnError(fmt.Errorf("fail to create"))
		buyerRepo := buyer.NewRepository(db)
		_, err = buyerRepo.Create(buyer1.Card_number_ID, buyer1.First_name, buyer1.Last_name)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("fail to create"), err)
	})
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	buyerRepo := buyer.NewRepository(db)

	t.Run("Update OK", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(buyer.UPDATE_BUYER)).WithArgs(
			buyer1.ID,
			buyer1.Card_number_ID,
			buyer1.First_name,
			buyer1.Last_name,
		).WillReturnResult(driver.RowsAffected(1))

		up, err := buyerRepo.Update(buyer1.ID, buyer1.Card_number_ID, buyer1.First_name, buyer1.Last_name)

		assert.NoError(t, err)
		assert.Equal(t, buyer1.Card_number_ID, up.Card_number_ID)
	})

	t.Run("Update fail", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(buyer.UPDATE_BUYER)).WillReturnError(fmt.Errorf("fail to find id"))

		_, err := buyerRepo.Update(buyer1.ID, buyer1.Card_number_ID, buyer1.First_name, buyer1.Last_name)

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("fail to find id"), err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		buyerRepo := buyer.NewRepository(db)
		mock.ExpectExec(regexp.QuoteMeta(buyer.DELETE_BUYER)).WithArgs(buyer1.ID).WillReturnResult(driver.RowsAffected(1))
		err = buyerRepo.Delete(buyer1.ID)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})

	t.Run("Delete fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		buyerRepo := buyer.NewRepository(db)
		mock.ExpectExec(regexp.QuoteMeta(buyer.DELETE_BUYER)).WillReturnResult(driver.ResultNoRows)

		err = buyerRepo.Delete(buyer1.ID)

		assert.Error(t, err)
		assert.NotNil(t, err)
	})

}
