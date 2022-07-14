package section_test

import (
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewConnectionMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("erro ao fazer conexão do mock, %s", err.Error())
	}
	return db, mock
}

func TestRepositoryCreateSection(t *testing.T) {
	mockSection := &section.Section{
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
	SqlCreateSectionTest := `INSERT INTO mercadofresco.sections (section_number,current_capacity,current_temperature,
		maximum_capacity,minimum_capacity,minimum_temperature,product_type_id,warehouse_id) VALUES (?,?,?,?,?,?,?,?)`

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

		sectionRepo := section.NewRepository(db)

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

		sectionRepo := section.NewRepository(db)

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

		sectionRepo := section.NewRepository(db)

		newSection, err := sectionRepo.CreateSection(*mockSection)

		assert.Error(t, err)
		// assert.Equal(t, expectError, err)

		sectionNumberfield := 0
		assert.Equal(t, sectionNumberfield, newSection.SectionNumber)
	})
}

func TestRepositoryUpdateSection(t *testing.T) {
	mockSection := &section.Section{
		Id:                 1,
		SectionNumber:      22,
		CurrentCapacity:    135,
		CurrentTemperature: 79845,
		MaximumCapacity:    456,
		MinimumCapacity:    23,
		MinimumTemperature: 4,
		ProductTypeId:      1,
		WarehouseId:        1,
	}
	t.Run("[success], Update", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(section.SqlUpdateSection)).
			WithArgs(
				&mockSection.SectionNumber, &mockSection.CurrentCapacity, &mockSection.CurrentTemperature,
				&mockSection.MaximumCapacity, &mockSection.MinimumCapacity, &mockSection.MinimumTemperature,
				&mockSection.ProductTypeId, &mockSection.WarehouseId, &mockSection.Id,
			).WillReturnResult(sqlmock.NewResult(0, 1))

		repoSection := section.NewRepository(db)

		sectionUpdate, err := repoSection.UpdateSection(*mockSection)
		assert.NoError(t, err)
		assert.ObjectsAreEqual(sectionUpdate, mockSection)
	})
	t.Run("[error]: section_number_UNIQUE", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		messageError := errors.New("section_number_UNIQUE is unique, and 22 already registered")

		mock.ExpectExec(regexp.QuoteMeta(section.SqlUpdateSection)).
			WithArgs(
				&mockSection.SectionNumber,
				&mockSection.CurrentCapacity,
				&mockSection.CurrentTemperature,
				&mockSection.MaximumCapacity,
				&mockSection.MinimumCapacity,
				&mockSection.MinimumTemperature,
				&mockSection.ProductTypeId,
				&mockSection.WarehouseId,
				&mockSection.Id,
			).WillReturnError(messageError)

		repoSection := section.NewRepository(db)

		_, err := repoSection.UpdateSection(*mockSection)

		assert.Error(t, err)
		assert.Equal(t, messageError.Error(), err.Error())
	})
	t.Run("[error]: warehouse_id", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		messageError := errors.New("warehouse_id is not registered on warehouse")

		mock.ExpectExec(regexp.QuoteMeta(section.SqlUpdateSection)).
			WithArgs(
				&mockSection.SectionNumber,
				&mockSection.CurrentCapacity,
				&mockSection.CurrentTemperature,
				&mockSection.MaximumCapacity,
				&mockSection.MinimumCapacity,
				&mockSection.MinimumTemperature,
				&mockSection.ProductTypeId,
				&mockSection.WarehouseId,
				&mockSection.Id,
			).WillReturnError(messageError)

		repoSection := section.NewRepository(db)

		_, err := repoSection.UpdateSection(*mockSection)

		assert.Error(t, err)
		assert.Equal(t, messageError.Error(), err.Error())
	})
	t.Run("[error]: product_type_id", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		messageError := errors.New("product_type_id is not registered on products_types")

		mock.ExpectExec(regexp.QuoteMeta(section.SqlUpdateSection)).
			WithArgs(
				&mockSection.SectionNumber,
				&mockSection.CurrentCapacity,
				&mockSection.CurrentTemperature,
				&mockSection.MaximumCapacity,
				&mockSection.MinimumCapacity,
				&mockSection.MinimumTemperature,
				&mockSection.ProductTypeId,
				&mockSection.WarehouseId,
				&mockSection.Id,
			).WillReturnError(messageError)

		repoSection := section.NewRepository(db)

		_, err := repoSection.UpdateSection(*mockSection)

		assert.Error(t, err)
		assert.Equal(t, messageError.Error(), err.Error())
	})
	t.Run("[error]: section is not found", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		expectError := errors.New("section is not found")

		mock.ExpectExec(regexp.QuoteMeta(section.SqlUpdateSection)).
			WithArgs(
				&mockSection.SectionNumber,
				&mockSection.CurrentCapacity,
				&mockSection.CurrentTemperature,
				&mockSection.MaximumCapacity,
				&mockSection.MinimumCapacity,
				&mockSection.MinimumTemperature,
				&mockSection.ProductTypeId,
				&mockSection.WarehouseId,
				&mockSection.Id,
			).WillReturnResult(sqlmock.NewErrorResult(expectError))

		repoSection := section.NewRepository(db)

		_, err := repoSection.UpdateSection(*mockSection)

		assert.Error(t, err)
		assert.Equal(t, expectError, err)
	})
	t.Run("[error]: section is not modifycation", func(t *testing.T) {
		db, mock := NewConnectionMock()

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(section.SqlUpdateSection)).
			WithArgs(
				&mockSection.SectionNumber,
				&mockSection.CurrentCapacity,
				&mockSection.CurrentTemperature,
				&mockSection.MaximumCapacity,
				&mockSection.MinimumCapacity,
				&mockSection.MinimumTemperature,
				&mockSection.ProductTypeId,
				&mockSection.WarehouseId,
				&mockSection.Id,
			).WillReturnResult(sqlmock.NewResult(0, 0))

		repoSection := section.NewRepository(db)

		_, err := repoSection.UpdateSection(*mockSection)

		assert.Error(t, err)
		assert.Equal(t, section.ErrorNotModify, err)
	})
}

func TestRepositoryListarSectionAll(t *testing.T) {
	db, mock := NewConnectionMock()

	rows := sqlmock.NewRows([]string{
		"Id",
		"SectionNumber",
		"CurrentCapacity",
		"CurrentTemperature",
		"MaximumCapacity",
		"MinimumCapacity",
		"MinimumTemperature",
		"ProductTypeId",
		"WarehouseId",
	}).
		AddRow(SectionList[0].Id,
			SectionList[0].SectionNumber,
			SectionList[0].CurrentCapacity,
			SectionList[0].CurrentTemperature,
			SectionList[0].MaximumCapacity,
			SectionList[0].MinimumCapacity,
			SectionList[0].MinimumTemperature,
			SectionList[0].ProductTypeId,
			SectionList[0].WarehouseId).
		AddRow(SectionList[1].Id,
			SectionList[1].SectionNumber,
			SectionList[1].CurrentCapacity,
			SectionList[1].CurrentTemperature,
			SectionList[1].MaximumCapacity,
			SectionList[1].MinimumCapacity,
			SectionList[1].MinimumTemperature,
			SectionList[1].ProductTypeId,
			SectionList[1].WarehouseId)

	t.Run("lista sections, sucesso", func(t *testing.T) {
		mock.ExpectQuery(section.SqlSelect).WillReturnRows(rows)
		repoSection := section.NewRepository(db)
		sections, err := repoSection.ListarSectionAll()
		assert.Nil(t, err)
		assert.Equal(t, SectionList[0].SectionNumber, sections[0].SectionNumber)
		assert.True(t, len(sections) == 2)
	})
	t.Run("lista sections, error section not registered", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"Id",
			"SectionNumber",
			"CurrentCapacity",
			"CurrentTemperature",
			"MaximumCapacity",
			"MinimumCapacity",
			"MinimumTemperature",
			"ProductTypeId",
			"WarehouseId",
		}).AddRow(0, 0, 0, 0, 0, 0, 0, 0, 0)

		mock.ExpectQuery(section.SqlSelect).WillReturnRows(rows).WillReturnError(section.ErrorFalhaInListAll)
		repoSection := section.NewRepository(db)
		sections, err := repoSection.ListarSectionAll()
		assert.Error(t, err)
		assert.Equal(t, section.ErrorFalhaInListAll, err)
		assert.True(t, len(sections) == 0)
	})
	t.Run("lista sections, error serializer fields of section", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"Id",
			"SectionNumber",
			"CurrentCapacity",
			"CurrentTemperature",
			"MaximumCapacity",
			"MinimumCapacity",
			"MinimumTemperature",
			"ProductTypeId",
			"WarehouseId",
		}).AddRow("", "", "", "", "", "", "", "", "")

		mock.ExpectQuery(section.SqlSelect).WillReturnRows(rows)
		repoSection := section.NewRepository(db)
		sections, err := repoSection.ListarSectionAll()
		assert.Error(t, err)
		assert.Equal(t, section.ErrorFalhaInserializerFields, err)
		assert.True(t, len(sections) == 0)
	})
}

func TestRepositoryListarSectionOne(t *testing.T) {
	t.Run("lista section, sucesso", func(t *testing.T) {
		db, mock := NewConnectionMock()

		sectionOne := section.Section{
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

		mock.ExpectQuery(section.SqlSelectByID).WithArgs(1).WillReturnRows(rows)

		repoSection := section.NewRepository(db)
		sec, err := repoSection.ListarSectionOne(1)
		assert.Nil(t, err)
		assert.ObjectsAreEqual(sectionOne, sec)
	})
	t.Run("lista section, error not found", func(t *testing.T) {
		db, mock := NewConnectionMock()

		rows := sqlmock.NewRows([]string{
			"Id", "SectionNumber", "CurrentCapacity", "CurrentTemperature", "MaximumCapacity", "MinimumCapacity", "MinimumTemperature", "ProductTypeId", "WarehouseId",
		})
		// .AddRow(sectionOne.Id, sectionOne.SectionNumber, sectionOne.CurrentCapacity, sectionOne.CurrentTemperature, sectionOne.MaximumCapacity, sectionOne.MinimumCapacity, sectionOne.MinimumTemperature, sectionOne.ProductTypeId, sectionOne.WarehouseId)

		mock.ExpectQuery(section.SqlSelectByID).WithArgs(10).WillReturnRows(rows)

		repoSection := section.NewRepository(db)
		sec, err := repoSection.ListarSectionOne(10)

		expectErrr := errors.New("section is not found")

		assert.Error(t, err)
		assert.Equal(t, expectErrr, err)
		assert.Empty(t, sec)
		assert.ObjectsAreEqual(section.Section{}, sec)
	})
}

func TestRepositoryDeleteSection(t *testing.T) {
	t.Run("[error]: delete section, error cannt be removed", func(t *testing.T) {
		db, mock := NewConnectionMock()

		paramId := 2

		expError := errors.New("this section cannot be removed")

		mock.ExpectExec(regexp.QuoteMeta(section.SqlDeleteSection)).
			WithArgs(paramId).WillReturnError(expError)

		repoSection := section.NewRepository(db)

		err := repoSection.DeleteSection(int64(paramId))

		assert.Error(t, err)

		assert.Equal(t, section.ErrorKeyTableSectionId, err)
	})
	t.Run("[error]: delete section, error section not found", func(t *testing.T) {
		db, mock := NewConnectionMock()

		paramId := 2

		mock.ExpectExec(regexp.QuoteMeta(section.SqlDeleteSection)).
			WithArgs(paramId).WillReturnResult(sqlmock.NewResult(0, 0)).WillReturnError(nil)

		repoSection := section.NewRepository(db)

		err := repoSection.DeleteSection(int64(paramId))

		assert.Equal(t, section.ErrorNotFound, err)
	})
	t.Run("[success]: delete section, sucesso", func(t *testing.T) {
		db, mock := NewConnectionMock()

		paramId := 20

		mock.ExpectExec(regexp.QuoteMeta(section.SqlDeleteSection)).
			WithArgs(paramId).WillReturnResult(sqlmock.NewResult(0, 1)).WillReturnError(nil)

		repoSection := section.NewRepository(db)

		err := repoSection.DeleteSection(int64(paramId))

		assert.NoError(t, err)
	})
}
