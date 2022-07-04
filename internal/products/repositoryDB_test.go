package products_test

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/stretchr/testify/assert"
)

var prodNewDB = products.Product{
	Description:                    "prod1",
	ExpirationRate:                 1,
	FreezingRate:                   2,
	Height:                         3.3,
	Length:                         4.3,
	NetWeight:                      5.5,
	ProductCode:                    "prod12s",
	RecommendedFreezingTemperature: 6.6,
	Width:                          7.7,
	ProductTypeId:                  8,
	SellerId:                       9}

var prod1DB = products.Product{
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

var prod2DB = products.Product{
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

var prodUpdate = products.Product{
	ProductCode: "prod12s",
	Description: "prod1",
	Width:       7.8,
	Length:      4.5,
	Height:      3.4}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	produtos := []products.Product{prod1DB, prod2DB}

	productsRepo := products.NewRepositoryProductsDB(db)

	query := "SELECT \\* FROM products"

	t.Run("GetAll Fail Scan", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"Id",
			"ProductCode",
			"Description",
			"Width",
			"Length",
			"Height",
			"NetWeight",
			"ExpirationRate",
			"RecommendedFreezingTemperature",
			"FreezingRate",
			"ProductTypeId",
			"SellerId",
		}).AddRow(
			produtos[0].Id,
			produtos[0].ProductCode,
			produtos[0].Description,
			produtos[0].Width,
			produtos[0].Length,
			produtos[0].Height,
			produtos[0].NetWeight,
			produtos[0].ExpirationRate,
			produtos[0].RecommendedFreezingTemperature,
			produtos[0].FreezingRate,
			produtos[0].ProductTypeId,
			produtos[0].SellerId,
		).AddRow(
			produtos[1].Id,
			produtos[1].ProductCode,
			produtos[1].Description,
			produtos[1].Width,
			produtos[1].Length,
			produtos[1].Height,
			produtos[1].NetWeight,
			produtos[1].ExpirationRate,
			produtos[1].RecommendedFreezingTemperature,
			produtos[1].FreezingRate,
			produtos[1].ProductTypeId,
			produtos[1].SellerId,
		)

		mock.ExpectQuery(query).WillReturnRows(rows)

		result, err := productsRepo.GetAll()
		assert.NoError(t, err)

		assert.Equal(t, result[0].Description, produtos[0].Description)
		assert.Equal(t, result[1].Description, produtos[1].Description)
	})

	t.Run("GetAll Fail Scan", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"Id",
			"ProductCode",
			"Description",
			"Width",
			"Length",
			"Height",
			"NetWeight",
			"ExpirationRate",
			"RecommendedFreezingTemperature",
			"FreezingRate",
			"ProductTypeId",
			"SellerId",
		}).AddRow("", "", "", "", "", "", "", "", "", "", "", "")

		mock.ExpectQuery(query).WillReturnRows(rows)

		_, err = productsRepo.GetAll()
		assert.Error(t, err)
	})

	t.Run("GetAll Fail Select", func(t *testing.T) {
		_, err = productsRepo.GetAll()
		assert.Error(t, err)
		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)
	})
}

func TestGetId(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	produtos := []products.Product{prod1DB, prod2DB}
	productsRepo := products.NewRepositoryProductsDB(db)
	rows := sqlmock.NewRows([]string{
		"Id",
		"ProductCode",
		"Description",
		"Width",
		"Length",
		"Height",
		"NetWeight",
		"ExpirationRate",
		"RecommendedFreezingTemperature",
		"FreezingRate",
		"ProductTypeId",
		"SellerId",
	}).AddRow(
		prod1DB.Id,
		prod1DB.ProductCode,
		prod1DB.Description,
		prod1DB.Width,
		prod1DB.Length,
		prod1DB.Height,
		prod1DB.NetWeight,
		prod1DB.ExpirationRate,
		prod1DB.RecommendedFreezingTemperature,
		prod1DB.FreezingRate,
		prod1DB.ProductTypeId,
		prod1DB.SellerId,
	)

	t.Run("Get ID - OK", func(t *testing.T) {

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.QUERY_GETID))
		stmt.ExpectQuery().WithArgs(1).WillReturnRows(rows)
		result, _ := productsRepo.GetId(1)
		assert.NoError(t, err)

		assert.Equal(t, produtos[0].Description, result.Description)
	})

	t.Run("Get ID - Fail prepar query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		productsRepo := products.NewRepositoryProductsDB(db)
		mock.ExpectQuery(products.QUERY_GETID).WithArgs(1).WillReturnError(fmt.Errorf("Fail to prepar query"))

		_, err = productsRepo.GetId(1)
		assert.Equal(t, fmt.Errorf("Fail to prepar query"), err)

	})

	t.Run("Get ID - Fail, product not found", func(t *testing.T) {
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.QUERY_GETID))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf("error"))

		_, err := productsRepo.GetId(1)
		assert.NotNil(t, err)

	})

}

func TestDelete(t *testing.T) {

	t.Run("Delete Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		productsRepo := products.NewRepositoryProductsDB(db)
		mock.ExpectPrepare(regexp.QuoteMeta(products.QUERY_DELETE)).WillReturnError(fmt.Errorf("error"))

		err = productsRepo.Delete(1)
		defer db.Close()
		assert.NotNil(t, err)

	})

	t.Run("Delete falha boa pergunta", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		productsRepo := products.NewRepositoryProductsDB(db)
		mock.ExpectPrepare(regexp.QuoteMeta(products.QUERY_DELETE)).ExpectExec().WithArgs(1)

		err = productsRepo.Delete(1)
		defer db.Close()
		assert.NotNil(t, err)
	})

	t.Run("Delete Ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.QUERY_DELETE))
		stmt.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
		productsRepo := products.NewRepositoryProductsDB(db)
		err = productsRepo.Delete(1)
		assert.NoError(t, err)
	})

	t.Run("Delete not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.QUERY_DELETE))
		stmt.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 0))
		productsRepo := products.NewRepositoryProductsDB(db)
		err = productsRepo.Delete(1)
		assert.Equal(t, err, fmt.Errorf("product not found"))
	})
}

func TestCreate(t *testing.T) {

	t.Run("Create Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		productsRepo := products.NewRepositoryProductsDB(db)
		mock.ExpectPrepare(regexp.QuoteMeta(products.QUERY_INSERT)).WillReturnError(fmt.Errorf("error"))
		_, err = productsRepo.Create(prod1DB)
		defer db.Close()
		assert.NotNil(t, err)
	})

	t.Run("create a boa pergunta", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		productsRepo := products.NewRepositoryProductsDB(db)
		mock.ExpectPrepare(regexp.QuoteMeta(products.QUERY_INSERT)).ExpectExec().WithArgs(1)

		_, err = productsRepo.Create(prod1DB)
		defer db.Close()
		assert.NotNil(t, err)
	})

	t.Run("Create Ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.QUERY_INSERT))
		stmt.ExpectExec().WithArgs(
			prodNewDB.ProductCode,
			prodNewDB.Description,
			prodNewDB.Width,
			prodNewDB.Length,
			prodNewDB.Height,
			prodNewDB.NetWeight,
			prodNewDB.ExpirationRate,
			prodNewDB.RecommendedFreezingTemperature,
			prodNewDB.FreezingRate,
			prodNewDB.ProductTypeId,
			prodNewDB.SellerId).WillReturnResult(sqlmock.NewResult(0, 1))
		productsRepo := products.NewRepositoryProductsDB(db)
		_, err = productsRepo.Create(prodNewDB)
		assert.NoError(t, err)

	})

	t.Run("Rows 0", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.QUERY_INSERT))
		stmt.ExpectExec().WithArgs(
			prodNewDB.ProductCode,
			prodNewDB.Description,
			prodNewDB.Width,
			prodNewDB.Length,
			prodNewDB.Height,
			prodNewDB.NetWeight,
			prodNewDB.ExpirationRate,
			prodNewDB.RecommendedFreezingTemperature,
			prodNewDB.FreezingRate,
			prodNewDB.ProductTypeId,
			prodNewDB.SellerId).WillReturnResult(sqlmock.NewResult(0, 0))
		productsRepo := products.NewRepositoryProductsDB(db)
		_, err = productsRepo.Create(prodNewDB)

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("Fail to save"), err)

	})
}

func TestUpdate(t *testing.T) {

	t.Run("Create Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		productsRepo := products.NewRepositoryProductsDB(db)
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(products.QUERY_UPDATE))
		stmt.ExpectExec().WithArgs(
			prodUpdate.ProductCode,
			prodUpdate.Description,
			prodUpdate.Width,
			prodUpdate.Length,
			prodUpdate.Height,
		).WillReturnResult(sqlmock.NewResult(0, 1))
		_, err = productsRepo.Update(1, prodUpdate)
		defer db.Close()
		assert.NotNil(t, err)
	})

}
