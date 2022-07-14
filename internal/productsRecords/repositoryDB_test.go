package productsRecords_test

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productsRecords"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	productsRepo := productsRecords.NewRepositoryProductsRecordsDB(db)

	t.Run("GetAll Fail Select", func(t *testing.T) {

		mock.ExpectQuery(regexp.QuoteMeta(productsRecords.QUERY_GETALL)).WithArgs().WillReturnError(sql.ErrNoRows)

		_, err = productsRepo.GetAllRecords()

		assert.Error(t, err)
		assert.Equal(t, "data not found", err.Error())
	})

	t.Run("GetAll Ok", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"product_id",
			"description",
			"records_count",
		}).AddRow(2, "noteboo2k", 1).AddRow(6, "celular", 8)
		mock.ExpectQuery(regexp.QuoteMeta(productsRecords.QUERY_GETALL)).WithArgs().WillReturnRows(rows)

		result, err := productsRepo.GetAllRecords()
		assert.NoError(t, err)
		assert.Equal(t, "noteboo2k", result[0].Description)
		assert.Equal(t, "celular", result[1].Description)
	})

	t.Run("GetAll Fail Scan", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"product_id",
			"records_count",
			"description",
		}).AddRow(2, 1, "noteboo2k").AddRow(6, 8, "celular")
		mock.ExpectQuery(regexp.QuoteMeta(productsRecords.QUERY_GETALL)).WithArgs().WillReturnRows(rows)
		_, err = productsRepo.GetAllRecords()
		assert.Error(t, err)
		assert.Equal(t, "Fail to scan", err.Error())

	})

}

func TestGetId(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	productsRepo := productsRecords.NewRepositoryProductsRecordsDB(db)

	t.Run("GetId Ok", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"product_id",
			"description",
			"records_count",
		}).AddRow(6, "celular", 8)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productsRecords.QUERY_GETID))
		stmt.ExpectQuery().WithArgs(1).WillReturnRows(rows)
		result, err := productsRepo.GetIdRecords(1)

		assert.NoError(t, err)
		assert.Equal(t, 6, result.ProductId)
		assert.Equal(t, "celular", result.Description)
		assert.Equal(t, 8, result.RecordsCount)
	})

	t.Run("GetId Fail prepare", func(t *testing.T) {

		mock.ExpectPrepare(regexp.QuoteMeta(productsRecords.QUERY_GETID)).WillReturnError(fmt.Errorf("Fail to prepare"))
		_, err := productsRepo.GetIdRecords(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Fail to prepare", err.Error())
	})

	t.Run("GetId No Records ", func(t *testing.T) {

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productsRecords.QUERY_GETID))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(sql.ErrNoRows)
		_, err := productsRepo.GetIdRecords(1)

		assert.NotNil(t, err)
		assert.Equal(t, "data not found", err.Error())
	})
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	productsRepo := productsRecords.NewRepositoryProductsRecordsDB(db)

	var prodNewDB = productsRecords.ProductRecords{
		LastUpdateDate: "2022-07-06 10:02:00",
		PurchasePrice:  105,
		SalePrice:      155,
		ProductId:      6}

	t.Run("Create Fail prepare", func(t *testing.T) {

		mock.ExpectPrepare(regexp.QuoteMeta(productsRecords.QUERY_INSERT)).WillReturnError(fmt.Errorf("Fail to prepare"))
		_, err := productsRepo.Create(prodNewDB)

		assert.NotNil(t, err)
		assert.Equal(t, "Fail to prepare", err.Error())
	})

	t.Run("Create Ok", func(t *testing.T) {

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productsRecords.QUERY_INSERT))
		stmt.ExpectExec().WithArgs(
			prodNewDB.LastUpdateDate,
			prodNewDB.PurchasePrice,
			prodNewDB.SalePrice,
			prodNewDB.ProductId).WillReturnResult(sqlmock.NewResult(0, 1))

		result, err := productsRepo.Create(prodNewDB)

		assert.NoError(t, err)
		assert.Equal(t, prodNewDB, result)
	})

	t.Run("Create fail", func(t *testing.T) {

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productsRecords.QUERY_INSERT))
		stmt.ExpectExec().WithArgs(
			prodNewDB.LastUpdateDate,
			prodNewDB.PurchasePrice,
			prodNewDB.SalePrice,
			prodNewDB.ProductId).WillReturnResult(sqlmock.NewResult(0, 0))

		result, err := productsRepo.Create(prodNewDB)

		assert.NotNil(t, err)
		assert.Equal(t, productsRecords.ProductRecords{}, result)
		assert.Equal(t, "Fail to save", err.Error())
	})

	t.Run("Create fail exec", func(t *testing.T) {

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(productsRecords.QUERY_INSERT))
		stmt.ExpectExec().WithArgs(
			prodNewDB.LastUpdateDate,
			"prodNewDB.PurchasePrice",
			prodNewDB.SalePrice,
			prodNewDB.ProductId).WillReturnResult(sqlmock.NewResult(0, 0))

		result, err := productsRepo.Create(prodNewDB)

		assert.NotNil(t, err)
		assert.Equal(t, productsRecords.ProductRecords{}, result)
		assert.Equal(t, "Query Fail", err.Error())
	})
}
