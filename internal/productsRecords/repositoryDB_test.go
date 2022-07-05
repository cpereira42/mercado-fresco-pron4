package productsRecords_test

import (
	"database/sql"
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
