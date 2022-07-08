package inbound_orders_test

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	inboundOrders "github.com/cpereira42/mercado-fresco-pron4/internal/inbound_orders"
	"github.com/stretchr/testify/assert"
)

var (
	inboundOrderDB1       = inboundOrders.ReportInboundOrders{1, "1", "mercado", "livre", 1, 5}
	inboundOrderDB2       = inboundOrders.ReportInboundOrders{2, "121", "edu", "araujo", 2, 4}
	inboundOrderCreate    = inboundOrders.InboundOrdersCreate{"Order2", 3, 1, 1}
	inboundOrdersResponse = inboundOrders.InboundOrdersResponse{"2022-07-08 10:48:22", "Order2", 3, 1, 1}
	createdAt             = "2022-07-08 10:48:22"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	inboundOrdersReport := []inboundOrders.ReportInboundOrders{inboundOrderDB1, inboundOrderDB2}
	inboundRepo := inboundOrders.NewRepository(db)

	t.Run("if GetAll is everything ok in the scan, it should return an array of employees reports", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"ID",
			"CardNumberID",
			"FirstName",
			"LastName",
			"WarehouseID",
			"InboundOrdersCount",
		}).AddRow(
			inboundOrdersReport[0].ID,
			inboundOrdersReport[0].CardNumberID,
			inboundOrdersReport[0].FirstName,
			inboundOrdersReport[0].LastName,
			inboundOrdersReport[0].WarehouseID,
			inboundOrdersReport[0].InboundOrdersCount,
		).AddRow(
			inboundOrdersReport[1].ID,
			inboundOrdersReport[1].CardNumberID,
			inboundOrdersReport[1].FirstName,
			inboundOrdersReport[1].LastName,
			inboundOrdersReport[1].WarehouseID,
			inboundOrdersReport[1].InboundOrdersCount,
		)
		mock.ExpectQuery(regexp.QuoteMeta(inboundOrders.GET_ALL_REPORT_INBOUND_ORDERS)).WillReturnRows(rows)
		result, err := inboundRepo.GetAll()
		assert.NoError(t, err)

		assert.ObjectsAreEqualValues(inboundOrdersReport, result)
		assert.ObjectsAreEqual(inboundOrdersReport, result)
		assert.Equal(t, inboundOrdersReport, result)
	})

	t.Run("if there is an error to scan in the GetAll, it should return an error message", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"ID",
			"CardNumberID",
			"FirstName",
			"LastName",
			"WarehouseID",
			"InboundOrdersCount",
		}).AddRow("", "", "", "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(inboundOrders.GET_ALL_REPORT_INBOUND_ORDERS)).WillReturnRows(rows)
		_, err = inboundRepo.GetAll()

		assert.Error(t, err)

	})

	t.Run("if there is an error in the GetAll select, it should return an error", func(t *testing.T) {
		_, err = inboundRepo.GetAll()
		mock.ExpectQuery(regexp.QuoteMeta(inboundOrders.GET_ALL_REPORT_INBOUND_ORDERS)).WillReturnError(sql.ErrNoRows)
		assert.Error(t, err)
	})
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	inboundRepo := inboundOrders.NewRepository(db)

	t.Run("if GetByID is OK, it should return an employee report", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"ID",
			"CardNumberID",
			"FirstName",
			"LastName",
			"WarehouseID",
			"InboundOrdersCount",
		}).AddRow(
			inboundOrderDB1.ID,
			inboundOrderDB1.CardNumberID,
			inboundOrderDB1.FirstName,
			inboundOrderDB1.LastName,
			inboundOrderDB1.WarehouseID,
			inboundOrderDB1.InboundOrdersCount,
		)
		stmt := mock.ExpectQuery(regexp.QuoteMeta(inboundOrders.GET_REPORT_INBOUND_ORDER_BY_ID))

		stmt.WithArgs(1).WillReturnRows(rows)

		result, err := inboundRepo.GetByID(1)
		assert.NoError(t, err)
		assert.ObjectsAreEqualValues(inboundOrderDB1, result)
		assert.Equal(t, inboundOrderDB1, result)
	})

	t.Run("if there is an error to find an employee in the GetByID, it should return employee not found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"ID",
			"CardNumberID",
			"FirstName",
			"LastName",
			"WarehouseID",
			"InboundOrdersCount",
		}).AddRow("", "", "", "", "", "")

		stmt := mock.ExpectQuery(regexp.QuoteMeta(inboundOrders.GET_REPORT_INBOUND_ORDER_BY_ID))

		stmt.WithArgs(1).WillReturnRows(rows)

		_, err := inboundRepo.GetByID(1)
		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf(inboundOrders.EMPLOYEE_NOT_FOUND), err)
	})

}

func TestCreate(t *testing.T) {
	t.Run("if Create is OK, it should return an inbound order", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		inboundRepo := inboundOrders.NewRepository(db)
		mock.ExpectExec(regexp.QuoteMeta(inboundOrders.CREATE_INBOUND_ORDERS)).WithArgs(
			inboundOrdersResponse.OrderDate,
			inboundOrdersResponse.OrderNumber,
			inboundOrdersResponse.EmployeeID,
			inboundOrdersResponse.ProductBatchID,
			inboundOrdersResponse.WarehouseID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		result, err := inboundRepo.Create(createdAt, inboundOrderCreate)

		assert.NoError(t, err)
		assert.ObjectsAreEqualValues(inboundOrdersResponse, result)
		assert.Equal(t, inboundOrdersResponse, result)
	})
	t.Run("if there is an error to Exec in the Create, it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		inboundRepo := inboundOrders.NewRepository(db)
		rows := sqlmock.NewRows([]string{
			"ID",
			"OrderDate",
			"OrderNumber",
			"EmployeeID",
			"ProductBatchID",
			"WarehouseID",
		}).AddRow("", "", "", "", "", "")

		mock.ExpectQuery(inboundOrders.CREATE_INBOUND_ORDERS).WillReturnRows(rows)
		_, err = inboundRepo.Create(createdAt, inboundOrderCreate)

		assert.Error(t, err)
	})

	t.Run("if no rows affected in the Create, it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		inboundRepo := inboundOrders.NewRepository(db)

		mock.ExpectExec(regexp.QuoteMeta(inboundOrders.CREATE_INBOUND_ORDERS)).WithArgs(
			inboundOrdersResponse.OrderDate,
			inboundOrdersResponse.OrderNumber,
			inboundOrdersResponse.EmployeeID,
			inboundOrdersResponse.ProductBatchID,
			inboundOrdersResponse.WarehouseID,
		).WillReturnResult(sqlmock.NewResult(0, 0))

		_, err = inboundRepo.Create(createdAt, inboundOrderCreate)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf(inboundOrders.FAIL_TO_SAVE), err)
	})
}
