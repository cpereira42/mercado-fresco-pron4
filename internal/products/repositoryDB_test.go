package products_test

import (
	"database/sql"
	"regexp"
	"testing"

	//"github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/mocks"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/stretchr/testify/assert"
	//products "github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/repository"
)

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

	query := "SELECT * FROM products Where id = ?"

	t.Run("Get ID - OK", func(t *testing.T) {
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

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectQuery().WithArgs(1).WillReturnRows(rows)

		/*mock.ExpectQuery(query).
		WithArgs(1).
		WillReturnRows(rows)*/

		result, _ := productsRepo.GetId(1)
		assert.NoError(t, err)

		assert.Equal(t, produtos[0].Description, result.Description)

	})

}
