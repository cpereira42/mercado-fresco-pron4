package productbatch_test

import (
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
	"github.com/stretchr/testify/assert"
)

func NewConnectionMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("erro ao fazer conexão do mock, %s", err.Error())
	}
	return db, mock
}

var productBatchesList []productbatch.ProductBatchesResponse = []productbatch.ProductBatchesResponse{
	{
		SectionId:     1,
		SectionNumber: 1,
		ProductsCount: 1,
	},
}

var productBatchesRes = productbatch.ProductBatchesResponse{
	SectionId:     1,
	SectionNumber: 1,
	ProductsCount: 1,
}

var productBatches productbatch.ProductBatches = productbatch.ProductBatches{
	Id:                 1,
	BatchNumber:        "111",
	CurrentQuantity:    1,
	CurrentTemperature: 1,
	DueDate:            "2022-04-04",
	InitialQuantity:    1,
	ManufacturingDate:  "2020-04-04 14:30:10",
	ManufacturingHour:  "2020-05-01 14:20:15",
	MinimumTemperature: 1,
	ProductId:          1,
	SectionId:          1,
}

func TestRepositoryProductBatChesGetAll(t *testing.T) {
	rows := sqlmock.NewRows([]string{
		"SectionId",
		"SectionNumber",
		"ProductsCount",
	}).
		AddRow(
			productBatchesList[0].SectionId,
			productBatchesList[0].SectionNumber,
			productBatchesList[0].ProductsCount,
		)
	t.Run("sucesso", func(t *testing.T) {
		db, mock := NewConnectionMock()
		defer db.Close()
		mock.ExpectQuery(regexp.QuoteMeta(productbatch.SqlrelatorioTodo)).WillReturnRows(rows)
		repoSection := productbatch.NewRepositoryProductBatches(db)
		pbList, err := repoSection.GetAll()
		assert.NoError(t, err)
		assert.True(t, len(pbList) == 1)
	})
	t.Run("falha ao listar", func(t *testing.T) {
		db, mock := NewConnectionMock()
		defer db.Close()
		rows := sqlmock.NewRows([]string{
			"SectionId",
			"SectionNumber",
			"ProductsCount",
		}).
			AddRow(
				"", "", "",
			)
		mock.ExpectQuery(regexp.QuoteMeta(productbatch.SqlrelatorioTodo)).WillReturnRows(rows)
		repoSection := productbatch.NewRepositoryProductBatches(db)
		pbList, err := repoSection.GetAll()
		assert.Error(t, err)
		assert.True(t, len(pbList) == 0)
	})
	t.Run("falha ao executar query", func(t *testing.T) {
		db, mock := NewConnectionMock()
		defer db.Close()
		queryError := ` s.id AS 'section_id', s.section_number AS 'section_number', count(*) AS 'products_count' FROM mercadofresco.products_batches as pbs INNER JOIN mercadofresco.sections AS s ON s.id = pbs.section_id INNER JOIN mercadofresco.products AS pcts ON pbs.product_id = pcts.id GROUP BY s.id`
		mock.ExpectQuery(regexp.QuoteMeta(queryError))
		repoSection := productbatch.NewRepositoryProductBatches(db)
		productBatchesResList, err := repoSection.GetAll()
		assert.Error(t, err)
		expctErr := errors.New("query sql invalid")
		assert.Equal(t, expctErr, err)
		assert.Equal(t, []productbatch.ProductBatchesResponse{}, productBatchesResList)
	})
}

func TestRespositoryProductBatchesGetId(t *testing.T) {
	t.Run("GetId, sucesso", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()
		paramId := 1

		rows := sqlmock.NewRows([]string{
			"SectionId", "SectionNumber", "ProductsCount",
		}).AddRow(productBatchesRes.SectionId, productBatchesRes.SectionNumber, productBatchesRes.ProductsCount)

		mock.ExpectQuery(regexp.QuoteMeta(productbatch.SqlrelatorioSectioId)).WithArgs(paramId).WillReturnRows(rows)

		repoPB := productbatch.NewRepositoryProductBatches(db)
		pb, err := repoPB.GetId(int64(paramId))
		assert.NoError(t, err)
		assert.ObjectsAreEqual(productBatchesRes, pb)
	})
	t.Run("GetId, error section_id not found", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()
		paramId := 90

		messageErr := errors.New("section_id not found")

		rows := sqlmock.NewRows([]string{
			"SectionId", "SectionNumber", "ProductsCount",
		}).AddRow("", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(productbatch.SqlrelatorioSectioId)).WithArgs(paramId).WillReturnRows(rows)

		repoPB := productbatch.NewRepositoryProductBatches(db)
		pb, err := repoPB.GetId(int64(paramId))
		assert.Error(t, err)
		assert.Equal(t, messageErr, err)

		productBatchesres := productbatch.ProductBatchesResponse{}
		assert.ObjectsAreEqual(productBatchesres, pb)
	})
}

func TestRepositoryProductBatchesCreatePB(t *testing.T) {
	t.Run("create product_batches", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreatePB)).
			WithArgs(
				&productBatches.BatchNumber,
				&productBatches.CurrentQuantity,
				&productBatches.CurrentTemperature,
				&productBatches.DueDate,
				&productBatches.InitialQuantity,
				&productBatches.ManufacturingDate,
				&productBatches.ManufacturingHour,
				&productBatches.MinimumTemperature,
				&productBatches.ProductId,
				&productBatches.SectionId,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repoPB := productbatch.NewRepositoryProductBatches(db)
		_, err := repoPB.CreatePB(productBatches)
		assert.NoError(t, err)
	})
	t.Run("create product_batches, erro product_id not registered", func(t *testing.T) {
		// caso de erro product_id invalid = "product_id is not registered on products"
		db, mock := NewConnectionMock()

		defer db.Close()
		messageError := errors.New("batch_number_UNIQUE is unique, and 111 already registered")

		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreatePB)).
			WithArgs(
				&productBatches.BatchNumber,
				&productBatches.CurrentQuantity,
				&productBatches.CurrentTemperature,
				&productBatches.DueDate,
				&productBatches.InitialQuantity,
				&productBatches.ManufacturingDate,
				&productBatches.ManufacturingHour,
				&productBatches.MinimumTemperature,
				&productBatches.ProductId,
				&productBatches.SectionId,
			).
			WillReturnError(messageError)

		repoPB := productbatch.NewRepositoryProductBatches(db)
		_, err := repoPB.CreatePB(productBatches)
		assert.Error(t, err)
		assert.Equal(t, messageError, err)
	})
	t.Run("create product_batches, erro section_id is not registered", func(t *testing.T) {
		// caso de erro product_id invalid = "section_id is not registered on sections"
		db, mock := NewConnectionMock()

		defer db.Close()
		messageError := errors.New("section_id is not registered on sections")
		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreatePB)).
			WithArgs(
				&productBatches.BatchNumber,
				&productBatches.CurrentQuantity,
				&productBatches.CurrentTemperature,
				&productBatches.DueDate,
				&productBatches.InitialQuantity,
				&productBatches.ManufacturingDate,
				&productBatches.ManufacturingHour,
				&productBatches.MinimumTemperature,
				&productBatches.ProductId,
				&productBatches.SectionId,
			).
			WillReturnError(messageError)

		repoPB := productbatch.NewRepositoryProductBatches(db)
		_, err := repoPB.CreatePB(productBatches)
		assert.Error(t, err)
		assert.Equal(t, messageError, err)
	})
	t.Run("create product_batches, erro, batch_number_unique", func(t *testing.T) {
		// caso de erro product_id invalid = "batch_number_UNIQUE is unique, and 111 already registered"
		db, mock := NewConnectionMock()

		defer db.Close()
		messageError := errors.New("batch_number_UNIQUE is unique, and 111 already registered")
		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreatePB)).WithArgs().WillReturnError(messageError)

		repoPB := productbatch.NewRepositoryProductBatches(db)
		_, err := repoPB.CreatePB(productBatches)
		assert.Error(t, err)
		assert.Equal(t, messageError, err)
	})
	t.Run("create product_batches,erro 0 rows affetcted", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()
		messageError := errors.New("batch_number_UNIQUE is unique, and 111 already registered")
		result := sqlmock.NewErrorResult(messageError)
		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreatePB)).
			WithArgs().
			WillReturnResult(result)

		repoPB := productbatch.NewRepositoryProductBatches(db)
		_, err := repoPB.CreatePB(productBatches)
		assert.Error(t, err)
		assert.Equal(t, messageError, err)

		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreatePB)).
			WithArgs().
			WillReturnResult(sqlmock.NewResult(0, 0))

		_, err = repoPB.CreatePB(productBatches)
		assert.Error(t, err)
		assert.Equal(t, "falha ao registrar um novo products_batches", err.Error())
	})
	t.Run("create product_batches,validação do campo DueDate", func(t *testing.T) {
		db, mock := NewConnectionMock()
		defer db.Close()
		messageError := errors.New("due_date is invalid must use '-' to separate date ex: '000-00-00'")
		productBatches.DueDate = "2020/12/30 12:20:30"
		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreatePB)).
			WithArgs(&productBatches.BatchNumber, &productBatches.CurrentQuantity, &productBatches.CurrentTemperature, &productBatches.DueDate, &productBatches.InitialQuantity, &productBatches.ManufacturingDate, &productBatches.ManufacturingHour, &productBatches.MinimumTemperature, &productBatches.ProductId, &productBatches.SectionId).
			WillReturnError(messageError)

		repoPB := productbatch.NewRepositoryProductBatches(db)
		_, err := repoPB.CreatePB(productBatches)
		assert.Error(t, err)
		assert.Equal(t, messageError, err)
	})
	t.Run("erro ManufacturingDate", func(t *testing.T) {
		db, mock := NewConnectionMock()

		messageError := errors.New("manufacturing_hour is invalid must use ':' to separate the time, ex: '00:00:00'")
		productBatches.ManufacturingHour = "2020-1230 122030"
		productBatches.DueDate = "2020-12-30 12:20:30"
		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreatePB)).
			WithArgs(&productBatches.BatchNumber, &productBatches.CurrentQuantity, &productBatches.CurrentTemperature, &productBatches.DueDate, &productBatches.InitialQuantity, &productBatches.ManufacturingDate, &productBatches.ManufacturingHour, &productBatches.MinimumTemperature, &productBatches.ProductId, &productBatches.SectionId).
			WillReturnError(messageError)

		repoPB := productbatch.NewRepositoryProductBatches(db)
		_, err := repoPB.CreatePB(productBatches)
		assert.Error(t, err)
		assert.Equal(t, messageError, err)
	})

	t.Run("erro no campos Manifactturing_date", func(t *testing.T) {
		db, mock := NewConnectionMock()

		productBatches.ManufacturingDate = "2020.12.30 12:20:30"
		productBatches.ManufacturingHour = "2020-12-30 12:20:30"

		messageError := errors.New("manufacturing_date is invalid must use '-' to separate date ex: '000-00-00'")

		mock.ExpectExec(regexp.QuoteMeta(productbatch.SqlCreatePB)).
			WithArgs(&productBatches.BatchNumber, &productBatches.CurrentQuantity, &productBatches.CurrentTemperature, &productBatches.DueDate, &productBatches.InitialQuantity, &productBatches.ManufacturingDate, &productBatches.ManufacturingHour, &productBatches.MinimumTemperature, &productBatches.ProductId, &productBatches.SectionId).
			WillReturnError(messageError)

		repoPB := productbatch.NewRepositoryProductBatches(db)
		_, err := repoPB.CreatePB(productBatches)
		assert.Error(t, err)
		assert.Equal(t, messageError, err)
	})
}
