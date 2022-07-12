package carries_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpereira42/mercado-fresco-pron4/internal/carries"
	"github.com/stretchr/testify/assert"
)

var carry1 = carries.Carries{"cid1", "companyname1", "rua 1", "11112222", 1}
var carry2 = carries.Carries{"cid2", "companyname2", "rua 2", "11112222", 1}

var locality1 = carries.Localities{1, "São Paulo", 1}
var locality2 = carries.Localities{2, "Nova York", 2}

var carriesReports = []carries.Localities{locality1, locality2}

func TestCreate(t *testing.T) {

	t.Run("Should create a new carries", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		repository := carries.NewRepository(db)
		mock.ExpectExec(regexp.QuoteMeta(carries.CreateCarry)).WithArgs(
			carry1.Cid,
			carry1.CompanyName,
			carry1.Address,
			carry1.Telephone,
			carry1.LocalityID).WillReturnResult(sqlmock.NewResult(1, 1))
		_, err = repository.Create(carry1.Cid, carry1.CompanyName, carry1.Address, carry1.Telephone, carry1.LocalityID)
		assert.NoError(t, err)

	})

	t.Run("Create error saving", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		repository := carries.NewRepository(db)
		mock.ExpectExec(regexp.QuoteMeta(carries.CreateCarry)).WithArgs(
			carry1.Cid,
			carry1.CompanyName,
			carry1.Address,
			carry1.Telephone,
			carry1.LocalityID).WillReturnError(fmt.Errorf(carries.FailedToCreateCarry))
		_, err = repository.Create(carry1.Cid, carry1.CompanyName, carry1.Address, carry1.Telephone, carry1.LocalityID)
		assert.Error(t, err)

	})

	t.Run("Create - No rows affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		repository := carries.NewRepository(db)
		mock.ExpectExec(regexp.QuoteMeta(carries.CreateCarry)).WithArgs(
			carry1.Cid,
			carry1.CompanyName,
			carry1.Address,
			carry1.Telephone,
			carry1.LocalityID).WillReturnResult(sqlmock.NewResult(0, 0))
		_, err = repository.Create(carry1.Cid, carry1.CompanyName, carry1.Address, carry1.Telephone, carry1.LocalityID)
		assert.NotNil(t, err)

	})

}
func TestGetByIDReport(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repository := carries.NewRepository(db)

	rows := sqlmock.NewRows([]string{"LocalityID", "LocalityName", "Count"}).
		AddRow(locality1.LocalityID, locality1.LocalityName, locality1.Count)

	t.Run("Generate Report of Localites by ID", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(carries.GetByIDReport)).WillReturnRows(rows)

		result, err := repository.GetByIDReport(1)
		assert.NoError(t, err)
		assert.ObjectsAreEqual(locality1, result)

	})
	t.Run("Generate Report and ID not Found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(carries.GetByIDReport)).WithArgs(1).WillReturnRows(rows)

		_, err := repository.GetByIDReport(1)

		assert.Equal(t, fmt.Errorf(carries.FailedIdNotFound), err)

	})

}
func TestGetAllReport(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repository := carries.NewRepository(db)

	rows := sqlmock.NewRows([]string{"LocalityID", "LocalityName", "Count"}).
		AddRow(carriesReports[0].LocalityID, carriesReports[0].LocalityName, carriesReports[0].Count).
		AddRow(carriesReports[1].LocalityID, carriesReports[1].LocalityName, carriesReports[1].Count)

	t.Run("Generate Report of All Localities - Should be OK", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(carries.GetAllReport)).WillReturnRows(rows)
		assert.NoError(t, err)

		result, err := repository.GetAllReport()
		assert.NoError(t, err)
		assert.ObjectsAreEqual(carriesReports[0], result[0])
		assert.ObjectsAreEqual(carriesReports[1], result[1])

	})

	t.Run("Generate Report of All Localities - Should be Fail", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"LocalityId",
			"LocalityName",
			"SellersCount",
		}).AddRow("", "", "")
		mock.ExpectQuery(regexp.QuoteMeta(carries.GetAllReport)).WillReturnRows(rows)

		_, err := repository.GetAllReport()
		assert.Error(t, err)

	})

	t.Run("Generate Report of All Localities - Should be Fail to Read", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(carries.GetAllReport)).WillReturnError(err)
		_, err := repository.GetAllReport()
		assert.Error(t, err)

	})

}
