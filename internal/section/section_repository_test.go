package section

import (
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

var sectionList = []Section{
	{
		Id:                 1,
		SectionNumber:      1,
		CurrentCapacity:    1,
		CurrentTemperature: 1,
		MaximumCapacity:    1,
		MinimumCapacity:    1,
		MinimumTemperature: 1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}, {
		Id:                 2,
		SectionNumber:      3,
		CurrentTemperature: 79845,
		MinimumTemperature: 4,
		CurrentCapacity:    135,
		MinimumCapacity:    23,
		MaximumCapacity:    456,
		WarehouseId:        1,
		ProductTypeId:      1,
	},
}

func NewConnectionMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("erro ao fazer conexão do mock, %s", err.Error())
	}
	return db, mock
}

func TestRepositoryCreateSection(t *testing.T) {
	mockSection := &Section{
		Id:                 1,
		SectionNumber:      3,
		CurrentTemperature: 79845,
		MinimumTemperature: 4,
		CurrentCapacity:    135,
		MinimumCapacity:    23,
		MaximumCapacity:    456,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	SqlCreateSectionTest := `INSERT INTO mercadofresco.sections (section_number,current_capacity,current_temperature,maximum_capacity,minimum_capacity,minimum_temperature,product_type_id,warehouse_id) VALUES (?,?,?,?,?,?,?,?)`

	t.Run("sucesso", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(SqlCreateSectionTest)).
			WithArgs(
				mockSection.SectionNumber,
				mockSection.CurrentCapacity,
				mockSection.CurrentTemperature,
				mockSection.MaximumCapacity,
				mockSection.MinimumCapacity,
				mockSection.MinimumTemperature,
				mockSection.ProductTypeId,
				mockSection.WarehouseId,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		sectionRepo := NewRepository(db)

		newSection, err := sectionRepo.CreateSection(*mockSection)

		assert.NoError(t, err)
		assert.Equal(t, nil, err)
		assert.Equal(t, newSection.SectionNumber, mockSection.SectionNumber)
		assert.Equal(t, newSection.CurrentCapacity, mockSection.CurrentCapacity)
		assert.Equal(t, newSection.CurrentTemperature, mockSection.CurrentTemperature)
		assert.Equal(t, newSection.MaximumCapacity, mockSection.MaximumCapacity)
		assert.Equal(t, newSection.MinimumCapacity, mockSection.MinimumCapacity)
		assert.Equal(t, newSection.ProductTypeId, mockSection.ProductTypeId)
		assert.Equal(t, newSection.WarehouseId, mockSection.WarehouseId)
	})
	t.Run("create section com warehouse_id invalido", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		mockSection.WarehouseId = 2

		expectError := errors.New("product_type_id id not found")

		mock.ExpectExec(regexp.QuoteMeta(SqlCreateSectionTest)).
			WithArgs(
				mockSection.SectionNumber,
				mockSection.CurrentCapacity,
				mockSection.CurrentTemperature,
				mockSection.MaximumCapacity,
				mockSection.MinimumCapacity,
				mockSection.MinimumTemperature,
				mockSection.ProductTypeId,
				mockSection.WarehouseId,
			).
			WillReturnResult(sqlmock.NewErrorResult(expectError))

		sectionRepo := NewRepository(db)

		newSection, err := sectionRepo.CreateSection(*mockSection)

		assert.Error(t, err)
		assert.Equal(t, expectError, err)
		assert.Equal(t, 0, newSection.SectionNumber)
	})
	t.Run("create section com campos invalido", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		mockSection.WarehouseId = 2
		expectError := errors.New("inserção de section falho, campos invalidos")
		mock.ExpectExec(regexp.QuoteMeta(SqlCreateSectionTest)).
			WithArgs(
				mockSection.SectionNumber,
				mockSection.CurrentCapacity,
				mockSection.CurrentTemperature,
				mockSection.MaximumCapacity,
				// mockSection.MinimumCapacity,
				mockSection.MinimumTemperature,
				mockSection.ProductTypeId,
				mockSection.WarehouseId,
			).
			WillReturnResult(sqlmock.NewErrorResult(expectError))

		sectionRepo := NewRepository(db)

		newSection, err := sectionRepo.CreateSection(*mockSection)

		assert.Error(t, err)
		assert.Equal(t, expectError, err)
		assert.Equal(t, 0, newSection.SectionNumber)
	})
	t.Run("create section com campo product_type_id invalido", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()
		mockSection.ProductTypeId = 2
		expectError := errors.New("o campo product_type_id não é valido")
		mock.ExpectExec(regexp.QuoteMeta(SqlCreateSectionTest)).
			WithArgs(
				mockSection.SectionNumber,
				mockSection.CurrentCapacity,
				mockSection.CurrentTemperature,
				mockSection.MaximumCapacity,
				mockSection.MinimumCapacity,
				mockSection.MinimumTemperature,
				mockSection.ProductTypeId,
				mockSection.WarehouseId,
			).
			WillReturnResult(sqlmock.NewErrorResult(expectError))

		sectionRepo := NewRepository(db)

		newSection, err := sectionRepo.CreateSection(*mockSection)

		assert.Error(t, err)
		assert.Equal(t, expectError, err)
		assert.Equal(t, 0, newSection.SectionNumber)
	})
}

func TestRepositoryUpdateSection(t *testing.T) {
	mockSection := &Section{
		Id:                 1,
		SectionNumber:      3,
		CurrentCapacity:    135,
		CurrentTemperature: 79845,
		MaximumCapacity:    456,
		MinimumCapacity:    23,
		MinimumTemperature: 4,
		ProductTypeId:      1,
		WarehouseId:        1,
	}

	t.Run("sucesso, Update", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta(SqlUpdateSection)).ExpectExec().
			WithArgs(
				mockSection.SectionNumber,
				mockSection.CurrentCapacity,
				mockSection.CurrentTemperature,
				mockSection.MaximumCapacity,
				mockSection.MinimumCapacity,
				mockSection.MinimumTemperature,
				mockSection.ProductTypeId,
				mockSection.WarehouseId,
				mockSection.Id,
			).WillReturnResult(sqlmock.NewResult(0, 1))

		repoSection := NewRepository(db)

		sectionUpdate, err := repoSection.UpdateSection(*mockSection)
		assert.NoError(t, err)
		assert.ObjectsAreEqual(sectionUpdate, mockSection)
	})
	t.Run("sucesso, Update, error ordem dos campos incorreta", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta(SqlUpdateSection)).ExpectExec().
			WithArgs(
				mockSection.CurrentCapacity,
				mockSection.SectionNumber,
				mockSection.MaximumCapacity,
				mockSection.CurrentTemperature,
				mockSection.MinimumTemperature,
				mockSection.MinimumCapacity,
				mockSection.ProductTypeId,
				mockSection.WarehouseId,
				mockSection.Id,
			).WillReturnResult(sqlmock.NewResult(0, 0))

		repoSection := NewRepository(db)

		_, err := repoSection.UpdateSection(*mockSection)

		expectError := errors.New("falha ao atualizar o section")

		assert.Error(t, err)
		assert.Equal(t, expectError, err)
	})
	t.Run("sucesso, Update, error ordem dos campos incorreta", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta(SqlUpdateSection)).ExpectExec().
			WithArgs(
				mockSection.CurrentCapacity,
				mockSection.SectionNumber,
				mockSection.MaximumCapacity,
				mockSection.CurrentTemperature,
				mockSection.MinimumTemperature,
				mockSection.MinimumCapacity,
				mockSection.ProductTypeId,
				mockSection.WarehouseId,
				mockSection.Id,
			).WillReturnResult(sqlmock.NewResult(0, 0))

		repoSection := NewRepository(db)

		_, err := repoSection.UpdateSection(*mockSection)

		expectError := errors.New("falha ao atualizar o section")

		assert.Error(t, err)
		assert.Equal(t, expectError, err)
	})
	t.Run("sucesso, Update, error na query sql", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		SqlUpdateSection := `
			UPDATE mercadofresco.section 
			SET section_number=?,current_capacity=?,current_temperature=?,maximum_capacity=?,minimum_capacity=?,
			minimum_temperature=?,product_type_id=?, warehouse_id=? WHERE id=?`

		mock.ExpectPrepare(regexp.QuoteMeta(SqlUpdateSection)).ExpectExec().
			WithArgs(
				mockSection.CurrentCapacity,
				mockSection.SectionNumber,
				mockSection.MaximumCapacity,
				mockSection.CurrentTemperature,
				mockSection.MinimumTemperature,
				mockSection.MinimumCapacity,
				mockSection.ProductTypeId,
				mockSection.WarehouseId,
				mockSection.Id,
			).WillReturnResult(sqlmock.NewResult(0, 0))

		repoSection := NewRepository(db)

		_, err := repoSection.UpdateSection(*mockSection)

		expectError := errors.New("falha ao executar query sql")

		assert.Error(t, err)
		assert.Equal(t, expectError, err)
	})

	t.Run("sucesso, Update, error rows affected 0", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta(SqlUpdateSection)).ExpectExec().
			WithArgs(
				mockSection.SectionNumber,
				mockSection.CurrentCapacity,
				mockSection.CurrentTemperature,
				mockSection.MaximumCapacity,
				mockSection.MinimumCapacity,
				mockSection.MinimumTemperature,
				mockSection.ProductTypeId,
				mockSection.WarehouseId,
				mockSection.Id,
			).WillReturnResult(sqlmock.NewResult(0, 0))

		repoSection := NewRepository(db)

		sectionExpct, err := repoSection.UpdateSection(*mockSection)

		expectError := errors.New("section não atualizada")
		assert.Error(t, err)
		assert.Equal(t, expectError, err)
		assert.ObjectsAreEqual(Section{}, sectionExpct)
	})
}

func TestRepositoryListarSectionAll(t *testing.T) {
	t.Run("lista sections, sucesso", func(t *testing.T) {
		db, mock := NewConnectionMock()

		mockSection := Section{
			Id:                 1,
			SectionNumber:      3,
			CurrentTemperature: 79845,
			MinimumTemperature: 4,
			CurrentCapacity:    135,
			MinimumCapacity:    23,
			MaximumCapacity:    456,
			WarehouseId:        1,
			ProductTypeId:      1,
		}

		rows := sqlmock.NewRows([]string{"Id", "SectionNumber", "CurrentCapacity", "CurrentTemperature", "MaximumCapacity", "MinimumCapacity", "MinimumTemperature", "ProductTypeId", "WarehouseId"}).
			AddRow(sectionList[0].Id, sectionList[0].SectionNumber, sectionList[0].CurrentCapacity, sectionList[0].CurrentTemperature, sectionList[0].MaximumCapacity, sectionList[0].MinimumCapacity, sectionList[0].MinimumTemperature, sectionList[0].ProductTypeId, sectionList[0].WarehouseId).
			AddRow(sectionList[1].Id, sectionList[1].SectionNumber, sectionList[1].CurrentCapacity, sectionList[1].CurrentTemperature, sectionList[1].MaximumCapacity, sectionList[1].MinimumCapacity, sectionList[1].MinimumTemperature, sectionList[1].ProductTypeId, sectionList[1].WarehouseId)

		mock.ExpectQuery(SqlSelect).WillReturnRows(rows)

		repoSection := NewRepository(db)

		sections, err := repoSection.ListarSectionAll()

		assert.Nil(t, err)
		assert.Equal(t, mockSection.SectionNumber, sections[1].SectionNumber)
		assert.True(t, len(sections) == 2)
	})
	t.Run("lista sections, error", func(t *testing.T) {
		db, mock := NewConnectionMock()

		// mockSection := Section{}

		rows := sqlmock.NewRows([]string{"Id", "SectionNumber", "CurrentCapacity", "CurrentTemperature", "MaximumCapacity", "MinimumCapacity", "MinimumTemperature", "ProductTypeId", "WarehouseId"})
		// .
		// AddRow(sectionList[0].Id, sectionList[0].SectionNumber, sectionList[0].CurrentCapacity, sectionList[0].CurrentTemperature, sectionList[0].MaximumCapacity, sectionList[0].MinimumCapacity, sectionList[0].MinimumTemperature, sectionList[0].ProductTypeId, sectionList[0].WarehouseId).
		// AddRow(sectionList[1].Id, sectionList[1].SectionNumber, sectionList[1].CurrentCapacity, sectionList[1].CurrentTemperature, sectionList[1].MaximumCapacity, sectionList[1].MinimumCapacity, sectionList[1].MinimumTemperature, sectionList[1].ProductTypeId, sectionList[1].WarehouseId)

		SqlSelect := `
			SELECT i123d,section_number,current_capacity,current_temperature,maximum_capacity,minimum_capacity,minimum_temperature,product_type_id,warehouse_id
			FROM mercadofresco.sections`

		mock.ExpectQuery(SqlSelect).WillReturnRows(rows)

		repoSection := NewRepository(db)

		sections, err := repoSection.ListarSectionAll()

		expectError := errors.New("sections not this registered")

		assert.Error(t, err)
		assert.Equal(t, expectError, err)
		assert.True(t, len(sections) == 0)
	})
}

func TestRepositoryListarSectionOne(t *testing.T) {
	t.Run("lista section, sucesso", func(t *testing.T) {
		db, mock := NewConnectionMock()

		sectionOne := Section{
			Id:                 1,
			SectionNumber:      1,
			CurrentCapacity:    1,
			CurrentTemperature: 1,
			MaximumCapacity:    1,
			MinimumCapacity:    1,
			MinimumTemperature: 1,
			WarehouseId:        1,
			ProductTypeId:      1,
		}

		rows := sqlmock.NewRows([]string{
			"Id", "SectionNumber", "CurrentCapacity", "CurrentTemperature", "MaximumCapacity", "MinimumCapacity", "MinimumTemperature", "ProductTypeId", "WarehouseId",
		}).AddRow(sectionOne.Id, sectionOne.SectionNumber, sectionOne.CurrentCapacity, sectionOne.CurrentTemperature, sectionOne.MaximumCapacity, sectionOne.MinimumCapacity, sectionOne.MinimumTemperature, sectionOne.ProductTypeId, sectionOne.WarehouseId)

		mock.ExpectQuery(SqlSelectByID).WithArgs(1).WillReturnRows(rows)

		repoSection := NewRepository(db)
		sec, err := repoSection.ListarSectionOne(1)
		assert.Nil(t, err)
		assert.ObjectsAreEqual(sectionOne, sec)
	})
	t.Run("lista section, error", func(t *testing.T) {
		db, mock := NewConnectionMock()

		rows := sqlmock.NewRows([]string{
			"Id", "SectionNumber", "CurrentCapacity", "CurrentTemperature", "MaximumCapacity", "MinimumCapacity", "MinimumTemperature", "ProductTypeId", "WarehouseId",
		})
		// .AddRow(sectionOne.Id, sectionOne.SectionNumber, sectionOne.CurrentCapacity, sectionOne.CurrentTemperature, sectionOne.MaximumCapacity, sectionOne.MinimumCapacity, sectionOne.MinimumTemperature, sectionOne.ProductTypeId, sectionOne.WarehouseId)

		mock.ExpectQuery(SqlSelectByID).WithArgs(10).WillReturnRows(rows)

		repoSection := NewRepository(db)
		sec, err := repoSection.ListarSectionOne(10)
		assert.Error(t, err)
		assert.Empty(t, sec)
	})
}

func TestRepositoryDeleteSection(t *testing.T) {
	t.Run("delete section, error", func(t *testing.T) {
		db, mock := NewConnectionMock()

		paramId := 1

		expError := errors.New("falha ao remove o section")

		mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM mercadofresco.sections WHERE id = ?`)).
			WithArgs(1).WillReturnError(expError)

		repoSection := NewRepository(db)

		err := repoSection.DeleteSection(int64(paramId))

		assert.Error(t, err)

		assert.Equal(t, expError, err)
	})
}

func TestRepositorygetWarehouse(t *testing.T) {
	t.Run("", func(t *testing.T) {
	})
}
