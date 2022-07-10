package seller_test

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"regexp"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var seller1 = seller.Seller{1, "cid1", "MELI", "Andromeda 860", "444-555", 1}
var seller2 = seller.Seller{2, "cid2", "MELI", "Andromeda 970", "344-556", 1}

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
		}).AddRow("", "", "", "", "", "")

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
	)

	query := "SELECT \\* FROM sellers WHERE id = \\?"
	t.Run("Get ID - OK", func(t *testing.T) {

		stmt := mock.ExpectPrepare(query)
		stmt.ExpectQuery().WithArgs(1).WillReturnRows(rows)
		result, _ := sellersRepo.GetId(1)
		assert.NoError(t, err)

		assert.Equal(t, sellers[0].Cid, result.Cid)
	})
	t.Run("Get ID - Fail prepar query", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		sellersRepo := seller.NewRepositorySeller(db)
		stmt := mock.ExpectPrepare(query).WillReturnError(fmt.Errorf("Fail to prepar query"))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf("Fail to prepar query"))

		_, err = sellersRepo.GetId(1)
		assert.Equal(t, fmt.Errorf("Fail to prepar query"), err)

	})
	t.Run("Get ID - Seller Not found", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		sellersRepo := seller.NewRepositorySeller(db)
		stmt := mock.ExpectPrepare(query)
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf("Seller 1 not found"))

		_, err = sellersRepo.GetId(1)
		assert.Equal(t, fmt.Errorf("Seller 1 not found"), err)

	})
}

func TestCreate(t *testing.T) {
	t.Run("Create Error saving", func(t *testing.T) {
		query := `INSERT INTO sellers 
		(cid,
		company_name,
		address,
		telephone,
		locality_id) 
		VALUES(?,?,?,?,?)`
		db, mock, err := sqlmock.New()
		sellersRepo := seller.NewRepositorySeller(db)
		mock.ExpectPrepare(query).WillReturnError(fmt.Errorf("error"))
		_, err = sellersRepo.Create(seller1.Cid, seller1.CompanyName, seller1.Address, seller1.Telephone, seller1.LocalityId)
		defer db.Close()
		assert.NotNil(t, err)
	})
	t.Run("Create Ok", func(t *testing.T) {
		query := `INSERT INTO sellers 
			(cid,
		company_name,
		address,
		telephone,
		locality_id) 
		VALUES(?,?,?,?,?)`
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectExec().WithArgs(seller1.Cid, seller1.CompanyName, seller1.Address, seller1.Telephone, seller1.LocalityId).WillReturnResult(sqlmock.NewResult(0, 1))
		sellersRepo := seller.NewRepositorySeller(db)
		_, err = sellersRepo.Create(seller1.Cid, seller1.CompanyName, seller1.Address, seller1.Telephone, seller1.LocalityId)
		defer db.Close()
		assert.NoError(t, err)
	})
	t.Run("Create - if no rows affected in the Create, it should return an error", func(t *testing.T) {
		query := `INSERT INTO sellers 
			(cid,
		company_name,
		address,
		telephone,
		locality_id) 
		VALUES(?,?,?,?,?)`
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectExec().WithArgs(seller1.Cid, seller1.CompanyName, seller1.Address, seller1.Telephone, seller1.LocalityId).WillReturnResult((sqlmock.NewErrorResult(fmt.Errorf("Fail to save"))))
		sellersRepo := seller.NewRepositorySeller(db)
		_, err = sellersRepo.Create(seller1.Cid, seller1.CompanyName, seller1.Address, seller1.Telephone, seller1.LocalityId)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("Fail to save"), err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update Error saving", func(t *testing.T) {
		query := `UPDATE sellers SET 
		cid=?,
		company_name=?,
		address=?,
		telephone=?,
		locality_id=? WHERE id=?`
		db, mock, err := sqlmock.New()
		sellersRepo := seller.NewRepositorySeller(db)
		mock.ExpectPrepare(query).WillReturnError(fmt.Errorf("error"))
		_, err = sellersRepo.Update(seller1.Id, seller1.Cid, seller1.CompanyName, seller1.Address, seller1.Telephone, seller1.LocalityId)
		defer db.Close()
		assert.NotNil(t, err)
	})
	t.Run("Udpdate Ok", func(t *testing.T) {
		query := `UPDATE sellers SET 
		cid=?,
		company_name=?,
		address=?,
		telephone=?,
		locality_id=? WHERE id=?`
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectExec().WithArgs(seller1.Cid, seller1.CompanyName, seller1.Address, seller1.Telephone, seller1.LocalityId, seller1.Id).WillReturnResult(sqlmock.NewResult(0, 1))
		sellersRepo := seller.NewRepositorySeller(db)
		_, err = sellersRepo.Update(seller1.Id, seller1.Cid, seller1.CompanyName, seller1.Address, seller1.Telephone, seller1.LocalityId)
		defer db.Close()
		assert.NoError(t, err)
	})
	t.Run("Update - if no rows affected in the Update, it should return an error", func(t *testing.T) {
		query := `UPDATE sellers SET 
		cid=?,
		company_name=?,
		address=?,
		telephone=?,
		locality_id=? WHERE id=?`
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectExec().WithArgs(seller1.Cid, seller1.CompanyName, seller1.Address, seller1.Telephone, seller1.LocalityId, seller1.Id).WillReturnResult((sqlmock.NewErrorResult(fmt.Errorf("Fail to save"))))
		sellersRepo := seller.NewRepositorySeller(db)
		_, err = sellersRepo.Update(seller1.Id, seller1.Cid, seller1.CompanyName, seller1.Address, seller1.Telephone, seller1.LocalityId)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("Fail to save"), err)
	})
}

func TestDelete(t *testing.T) {
	query := `DELETE FROM sellers WHERE id=?`

	t.Run("if Delete is OK, it should not return error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		sellersRepo := seller.NewRepositorySeller(db)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectExec().WithArgs(seller1.Id).WillReturnResult(driver.RowsAffected(1))

		err = sellersRepo.Delete(seller1.Id)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
	t.Run("if there is an error in the exec of Delete, it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		sellersRepo := seller.NewRepositorySeller(db)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectExec().WithArgs(seller1.Id).WillReturnError(fmt.Errorf("error"))

		err = sellersRepo.Delete(seller1.Id)

		assert.Error(t, err)

		assert.NotNil(t, err)
	})
	t.Run("if there is an error in the Prepare of Delete, it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		sellersRepo := seller.NewRepositorySeller(db)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query)).WillReturnError(fmt.Errorf("error"))
		stmt.ExpectExec().WithArgs(seller1.Id).WillReturnError(fmt.Errorf("error"))

		err = sellersRepo.Delete(seller1.Id)

		assert.Error(t, err)

		assert.NotNil(t, err)
	})
	t.Run("if there is no rows affected in Delete, it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sellersRepo := seller.NewRepositorySeller(db)
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectExec().WithArgs(seller1.Id).WillReturnResult(driver.ResultNoRows)

		err = sellersRepo.Delete(seller1.Id)

		assert.Error(t, err)

		assert.NotNil(t, err)
	})
}
