package employee_test

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/stretchr/testify/assert"
)

var (
	emply1DB       = employee.Employee{1, "123", "Eduardo", "Araujo", 1}
	emply2DB       = employee.Employee{2, "1234", "Edu", "Araujo", 1}
	emply2DBUpdate = employee.Employee{2, "12345", "Eduu", "Araujo", 1}
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	employees := []employee.Employee{emply1DB, emply2DB}
	employeesRepo := employee.NewRepository(db)

	t.Run("if GetAll is everything ok in the scan, it should return an array of employees", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"ID",
			"CardNumberID",
			"FirstName",
			"LastName",
			"WarehouseID",
		}).AddRow(
			employees[0].ID,
			employees[0].CardNumberID,
			employees[0].FirstName,
			employees[0].LastName,
			employees[0].WarehouseID,
		).AddRow(
			employees[1].ID,
			employees[1].CardNumberID,
			employees[1].FirstName,
			employees[1].LastName,
			employees[1].WarehouseID,
		)
		mock.ExpectQuery(employee.GET_ALL_EMPLOYEES).WillReturnRows(rows)
		result, err := employeesRepo.GetAll()
		assert.NoError(t, err)

		assert.Equal(t, employees[0].CardNumberID, result[0].CardNumberID)
		assert.Equal(t, employees[1].CardNumberID, result[1].CardNumberID)
	})

	t.Run("if there is an error to scan in the GetAll, it should return an error message", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"ID",
			"CardNumberID",
			"FirstName",
			"LastName",
			"WarehouseID",
		}).AddRow("", "", "", "", "")

		mock.ExpectQuery(employee.GET_ALL_EMPLOYEES).WillReturnRows(rows)
		_, err = employeesRepo.GetAll()

		assert.Error(t, err)
	})

	t.Run("if there is an error in the GetAll select, it should return an error", func(t *testing.T) {
		_, err = employeesRepo.GetAll()
		assert.Error(t, err)
		mock.ExpectQuery(employee.GET_ALL_EMPLOYEES).WillReturnError(sql.ErrNoRows)
	})
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	employees := []employee.Employee{emply1DB, emply2DB}
	employeesRepo := employee.NewRepository(db)

	t.Run("if GetByID is OK, it should return an employee", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"ID",
			"CardNumberID",
			"FirstName",
			"LastName",
			"WarehouseID",
		}).AddRow(
			emply1DB.ID,
			emply1DB.CardNumberID,
			emply1DB.FirstName,
			emply1DB.LastName,
			emply1DB.WarehouseID,
		)

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(employee.GET_EMPLOYEE_BY_ID))
		stmt.ExpectQuery().WithArgs(1).WillReturnRows(rows)

		result, _ := employeesRepo.GetByID(1)
		assert.NoError(t, err)
		assert.Equal(t, employees[0].CardNumberID, result.CardNumberID)
	})
	t.Run("if there is an error to Prepare in the GetByID, it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(employee.GET_EMPLOYEE_BY_ID).WithArgs(1).WillReturnError(fmt.Errorf(employee.FAIL_TO_PREPARE_QUERY))
		_, err = employeesRepo.GetByID(1)

		assert.Equal(t, fmt.Errorf(employee.FAIL_TO_PREPARE_QUERY), err)
	})

	t.Run("if there is an error to find an employee in the GetByID, it should return employee not found", func(t *testing.T) {
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(employee.GET_EMPLOYEE_BY_ID))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf(employee.EMPLOYEE_NOT_FOUND))

		_, err := employeesRepo.GetByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf(employee.EMPLOYEE_NOT_FOUND), err)

	})
}

func TestCreate(t *testing.T) {

	t.Run("if Create is OK, it should return an employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeesRepo := employee.NewRepository(db)
		mock.ExpectExec(regexp.QuoteMeta(employee.CREATE_EMPLOYEE)).WithArgs(
			emply2DB.CardNumberID,
			emply2DB.FirstName,
			emply2DB.LastName,
			emply2DB.WarehouseID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		emp, err := employeesRepo.Create(emply2DB.CardNumberID, emply2DB.FirstName, emply2DB.LastName, emply2DB.WarehouseID)

		assert.NoError(t, err)

		expectedCardNumberID := "1234"

		assert.Equal(t, expectedCardNumberID, emp.CardNumberID)
	})

	t.Run("if there is an error to Exec in the Create, it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		employeesRepo := employee.NewRepository(db)
		rows := sqlmock.NewRows([]string{
			"ID",
			"CardNumberID",
			"FirstName",
			"LastName",
			"WarehouseID",
		}).AddRow("", "", "", "", "")

		mock.ExpectQuery(employee.CREATE_EMPLOYEE).WillReturnRows(rows)
		_, err = employeesRepo.Create(emply2DB.CardNumberID, emply2DB.FirstName, emply2DB.LastName, emply2DB.WarehouseID)

		assert.Error(t, err)
	})
	t.Run("if no rows affected in the Create, it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(employee.CREATE_EMPLOYEE)).WithArgs(
			emply2DB.CardNumberID,
			emply2DB.FirstName,
			emply2DB.LastName,
			emply2DB.WarehouseID,
		).WillReturnResult(sqlmock.NewResult(0, 0))
		employeesRepo := employee.NewRepository(db)
		_, err = employeesRepo.Create(emply2DB.CardNumberID, emply2DB.FirstName, emply2DB.LastName, emply2DB.WarehouseID)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf(employee.FAIL_TO_SAVE), err)
	})
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	employeesRepo := employee.NewRepository(db)

	t.Run("if Update is OK, it should return an employee", func(t *testing.T) {

		mock.ExpectExec(regexp.QuoteMeta(employee.UPDATE_EMPLOYEE)).WithArgs(
			emply2DBUpdate.CardNumberID,
			emply2DBUpdate.FirstName,
			emply2DBUpdate.LastName,
			emply2DBUpdate.WarehouseID,
			emply2DBUpdate.ID,
		).WillReturnResult(driver.RowsAffected(1))

		emp, err := employeesRepo.Update(emply2DBUpdate.ID, emply2DBUpdate.CardNumberID, emply2DBUpdate.FirstName, emply2DBUpdate.LastName, emply2DBUpdate.WarehouseID)

		assert.NoError(t, err)

		assert.Equal(t, emply2DBUpdate.CardNumberID, emp.CardNumberID)
		assert.Equal(t, emply2DBUpdate.FirstName, emp.FirstName)
	})
	t.Run("if Update has an error, it should return an error", func(t *testing.T) {

		mock.ExpectExec(regexp.QuoteMeta(employee.UPDATE_EMPLOYEE)).WillReturnError(fmt.Errorf(employee.FAIL_TO_UPDATE))

		_, err := employeesRepo.Update(emply2DBUpdate.ID, emply2DBUpdate.CardNumberID, emply2DBUpdate.FirstName, emply2DBUpdate.LastName, emply2DBUpdate.WarehouseID)

		// assert.Error(t, err)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf(employee.FAIL_TO_UPDATE), err)
	})
}
