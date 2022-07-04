package seller_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var seller1 = seller.Seller{1, "cid1", "MELI", "Andromeda 860", "444-555", 1}
var seller2 = seller.Seller{2, "cid2", "MELI", "Andromeda 970", "344-556", 1}

// var seller3 = seller.Seller{3, "cid3", "MELI", "Bahia 555", "999-333", 2}
// var seller4 = seller.Seller{4, "cid4", "MELI", "Bahia 555", "999-333", 3}
// var seller5 = seller.Seller{4, "cid5", "MELI", "Bahia 555", "999-333", 2}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sellers := []seller.Seller{seller1, seller2}
	sellersRepo := seller.NewRepositorySeller(db)

	query := "SELECT \\* FROM sellers"
	t.Run("GetAll - OK", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"Id",
			"Cid",
			"CompanyName",
			"Address",
			"Telephone",
			"LocalityId",
		}).AddRow(
			sellers[0].Id,
			sellers[0].Cid,
			sellers[0].CompanyName,
			sellers[0].Address,
			sellers[0].Telephone,
			sellers[0].LocalityId,
		).AddRow(
			sellers[1].Id,
			sellers[1].Cid,
			sellers[1].CompanyName,
			sellers[1].Address,
			sellers[1].Telephone,
			sellers[1].LocalityId,
		)

		mock.ExpectQuery(query).WillReturnRows(rows)

		result, err := sellersRepo.GetAll()
		assert.NoError(t, err)

		assert.Equal(t, result[0].Cid, sellers[0].Cid)
		assert.Equal(t, result[1].Cid, sellers[1].Cid)
	})
	t.Run("GetAll - Fail Scan", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"Id",
			"Cid",
			"CompanyName",
			"Address",
			"Telephone",
			"LocalityId",
		}).AddRow("", "", "", "", "", 0)

		mock.ExpectQuery(query).WillReturnRows(rows)

		_, err = sellersRepo.GetAll()
		assert.Error(t, err)
	})
	t.Run("GetAll - Fail Select/Read", func(t *testing.T) {
		_, err = sellersRepo.GetAll()
		assert.Error(t, err)
		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)
	})
}

func TestGetId(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sellers := []seller.Seller{seller1, seller2}
	sellersRepo := seller.NewRepositorySeller(db)

	query := "SELECT \\* FROM sellers WHERE id = ?"
	t.Run("Get ID - OK", func(t *testing.T) {
		rows := []string{
			"Id",
			"Cid",
			"CompanyName",
			"Address",
			"Telephone",
			"LocalityId",
		}

		mock.ExpectQuery(query).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows(rows).AddRow(seller1))

		result, err := sellersRepo.GetId(1)

		log.Println(result)
		assert.NoError(t, err)

		assert.Equal(t, sellers[0].Cid, result.Cid)
	})
}
