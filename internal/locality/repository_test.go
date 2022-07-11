package locality_test

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/locality"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var locality1 = locality.Locality{1, "Itabaiana", "Sergipe", "Brasil"}
var locality2 = locality.Locality{2, "Nova York", "Nova York", "EUA"}

var localityReport1 = locality.LocalityReport{1, "Itabaiana", 200}
var localityReport2 = locality.LocalityReport{2, "Nova York", 3000}

func TestCreate(t *testing.T) {
	query := `INSERT INTO localities 
	(id,
	locality_name,
	province_name,
	country_name) 
	VALUES(?,?,?,?)`
	t.Run("Create Error saving", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		localityRepo := locality.NewRepositoryLocality(db)
		mock.ExpectPrepare(query).WillReturnError(fmt.Errorf("error"))
		_, err = localityRepo.Create(locality1.Id, locality1.LocalityName, locality1.ProvinceName, locality1.CountryName)
		defer db.Close()
		assert.NotNil(t, err)
	})
	t.Run("Create Ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectExec().WithArgs(locality1.Id, locality1.LocalityName, locality1.ProvinceName, locality1.CountryName).WillReturnResult(sqlmock.NewResult(0, 1))
		localityRepo := locality.NewRepositoryLocality(db)
		_, err = localityRepo.Create(locality1.Id, locality1.LocalityName, locality1.ProvinceName, locality1.CountryName)
		defer db.Close()
		assert.NoError(t, err)
	})
	t.Run("Create - if no rows affected in the Create, it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectExec().WithArgs(locality1.Id, locality1.LocalityName, locality1.ProvinceName, locality1.CountryName).WillReturnError(fmt.Errorf("Fail to Save"))
		localityRepo := locality.NewRepositoryLocality(db)
		_, err = localityRepo.Create(locality1.Id, locality1.LocalityName, locality1.ProvinceName, locality1.CountryName)
		assert.NotNil(t, err)
		assert.NotNil(t, err)
	})
}

func TestGenerateReportAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	reports := []locality.LocalityReport{localityReport1, localityReport2}
	localityRepo := locality.NewRepositoryLocality(db)

	query := `SELECT localities.id, localities.locality_name, COUNT(sellers.locality_id) 
	FROM mercadofresco.localities
	INNER JOIN sellers ON localities.id = sellers.locality_id
	GROUP BY locality_id;`
	t.Run("GenerateReportAll - OK", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"LocalityId",
			"LocalityName",
			"SellersCount",
		}).AddRow(
			reports[0].LocalityId,
			reports[0].LocalityName,
			reports[0].SellersCount,
		).AddRow(
			reports[1].LocalityId,
			reports[1].LocalityName,
			reports[1].SellersCount,
		)

		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		result, err := localityRepo.GenerateReportAll()
		assert.NoError(t, err)

		assert.Equal(t, result[0].LocalityId, reports[0].LocalityId)
		assert.Equal(t, result[1].LocalityId, reports[1].LocalityId)
	})
	t.Run("GenerateReportAll - Fail Scan", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"LocalityId",
			"LocalityName",
			"SellersCount",
		}).AddRow("", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		_, err = localityRepo.GenerateReportAll()
		assert.Error(t, err)
	})
	t.Run("GetAll - Fail Select/Read", func(t *testing.T) {
		_, err = localityRepo.GenerateReportAll()
		assert.Error(t, err)
		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)
	})
}

func TestGenerateReportById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	localityReports := []locality.LocalityReport{localityReport1, localityReport2}
	localityRepo := locality.NewRepositoryLocality(db)
	rows := sqlmock.NewRows([]string{
		"LocalityId",
		"LocalityName",
		"SellersCount",
	}).AddRow(
		localityReports[0].LocalityId,
		localityReports[0].LocalityName,
		localityReports[0].SellersCount,
	)

	query := `SELECT localities.id, localities.locality_name, COUNT(sellers.locality_id) 
	FROM mercadofresco.localities
	INNER JOIN sellers ON localities.id = sellers.locality_id
	WHERE localities.id = ?
	GROUP BY locality_id;`
	t.Run("Generate report by ID - OK", func(t *testing.T) {

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectQuery().WithArgs(1).WillReturnRows(rows)
		result, _ := localityRepo.GenerateReportById(1)
		assert.NoError(t, err)

		assert.Equal(t, localityReports[0].LocalityId, result.LocalityId)
	})
	t.Run("Generate report by ID - Fail prepar query", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query)).WillReturnError(fmt.Errorf("Fail to prepar query"))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf("Fail to prepar query"))
		localityRepo := locality.NewRepositoryLocality(db)
		_, err = localityRepo.GenerateReportById(1)
		assert.Equal(t, fmt.Errorf("Fail to prepar query"), err)

	})
	t.Run("Get ID - Locality Not found", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf("Locality 1 not found"))
		localityRepo := locality.NewRepositoryLocality(db)
		_, err = localityRepo.GenerateReportById(1)
		assert.Equal(t, fmt.Errorf("Locality 1 not found"), err)

	})
}
