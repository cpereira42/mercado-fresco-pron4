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
			carry1.LocalityID).WillReturnError(fmt.Errorf("failed to create carry"))
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
