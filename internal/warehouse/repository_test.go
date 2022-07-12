package warehouse_test

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"regexp"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	warehouseOne        = warehouse.Warehouse{ID: 1, Address: "Rua 1", Telephone: "11111-1111", Warehouse_code: "W1", Minimum_capacity: 10, Minimum_temperature: 20, Locality_id: 1}
	warehouseTwo        = warehouse.Warehouse{ID: 2, Address: "Rua 2", Telephone: "22222-2222", Warehouse_code: "W2", Minimum_capacity: 20, Minimum_temperature: 30, Locality_id: 1}
	warehouseThree      = warehouse.Warehouse{ID: 3, Address: "Rua 3", Telephone: "33333-3333", Warehouse_code: "W3", Minimum_capacity: 30, Minimum_temperature: 40, Locality_id: 1}
	warehouseOneUpdated = warehouse.Warehouse{ID: 1, Address: "Rua 4", Telephone: "11111-1111", Warehouse_code: "W1", Minimum_capacity: 10, Minimum_temperature: 20, Locality_id: 1}
	warehouses          = []warehouse.Warehouse{warehouse1, warehouse2}
)

func TestGetAll(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	warehousesRepo := warehouse.NewRepository(db)

	t.Run("GetAll - OK -", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"ID",
			"Address",
			"Telephone",
			"Warehouse_code",
			"Minimum_capacity",
			"Minimum_temperature",
			"Locality_id",
		}).AddRow(
			warehouses[0].ID,
			warehouses[0].Address,
			warehouses[0].Telephone,
			warehouses[0].Warehouse_code,
			warehouses[0].Minimum_capacity,
			warehouses[0].Minimum_temperature,
			warehouses[0].Locality_id,
		).AddRow(
			warehouses[1].ID,
			warehouses[1].Address,
			warehouses[1].Telephone,
			warehouses[1].Warehouse_code,
			warehouses[1].Minimum_capacity,
			warehouses[1].Minimum_temperature,
			warehouses[1].Locality_id,
		)

		mock.ExpectQuery(regexp.QuoteMeta(warehouse.GetAll)).WillReturnRows(rows)

		result, err := warehousesRepo.GetAll()

		assert.NoError(t, err)
		assert.ObjectsAreEqualValues(result[0], warehouses[0])
		assert.ObjectsAreEqualValues(result[1], warehouses[1])

	})
	t.Run("GetAll - Fail to Scan", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"ID",
			"Address",
			"Telephone",
			"Warehouse_code",
			"Minimum_capacity",
			"Minimum_temperature",
			"Locality_id",
		}).AddRow("", "", "", "", "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(warehouse.GetAll)).WillReturnRows(rows)

		_, err = warehousesRepo.GetAll()
		assert.Error(t, err)
	})
	t.Run("GetAll - Fail to Read", func(t *testing.T) {
		_, err = warehousesRepo.GetAll()
		assert.Error(t, err)
		mock.ExpectQuery(regexp.QuoteMeta(warehouse.GetAll)).WillReturnError(sql.ErrNoRows)
	})
}

func TestGetId(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	warehousesRepo := warehouse.NewRepository(db)
	rows := sqlmock.NewRows([]string{
		"ID",
		"Address",
		"Telephone",
		"Warehouse_code",
		"Minimum_capacity",
		"Minimum_temperature",
		"Locality_id",
	}).AddRow(
		warehouses[0].ID,
		warehouses[0].Address,
		warehouses[0].Telephone,
		warehouses[0].Warehouse_code,
		warehouses[0].Minimum_capacity,
		warehouses[0].Minimum_temperature,
		warehouses[0].Locality_id,
	)

	t.Run("GetByID - OK", func(t *testing.T) {

		stmt := mock.ExpectPrepare(regexp.QuoteMeta(warehouse.GetId))
		stmt.ExpectQuery().WithArgs(1).WillReturnRows(rows)
		result, _ := warehousesRepo.GetByID(1)
		assert.NoError(t, err)

		assert.ObjectsAreEqual(warehouses[0], result)
	})
	t.Run("Get ID - Fail prepar query", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		warehousesRepo := warehouse.NewRepository(db)
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(warehouse.GetId)).WillReturnError(fmt.Errorf("Fail to prepare query"))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf("Fail to prepare query"))

		_, err = warehousesRepo.GetByID(1)
		assert.Equal(t, fmt.Errorf("Fail to prepare query"), err)

	})
	t.Run("GetByID - Warehouse ID not found", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		warehousesRepo := warehouse.NewRepository(db)
		stmt := mock.ExpectPrepare(regexp.QuoteMeta(warehouse.GetId))
		stmt.ExpectQuery().WithArgs(1).WillReturnError(fmt.Errorf("Warehouse 1 not found"))

		_, err = warehousesRepo.GetByID(1)
		assert.Equal(t, fmt.Errorf("Warehouse 1 not found"), err)

	})
}

func TestCreate(t *testing.T) {
	t.Run("Create - OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		warehousesRepo := warehouse.NewRepository(db)

		mock.ExpectExec(regexp.QuoteMeta(warehouse.CreateWarehouse)).WithArgs(
			warehouse1.Address, warehouse1.Telephone, warehouse1.Warehouse_code, warehouse1.Minimum_capacity, warehouse1.Minimum_temperature, warehouse1.Locality_id).WillReturnResult(sqlmock.NewResult(0, 1))
		_, err = warehousesRepo.Create(warehouse1.Address, warehouse1.Telephone, warehouse1.Warehouse_code, warehouse1.Minimum_capacity, warehouse1.Minimum_temperature, warehouse1.Locality_id)
		assert.NoError(t, err)
	})
	t.Run("Create Fail", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		warehousesRepo := warehouse.NewRepository(db)

		rows := sqlmock.NewRows([]string{
			"Address",
			"Telephone",
			"Warehouse_code",
			"Minimum_capacity",
			"Minimum_temperature",
			"Locality_id",
		}).AddRow("", "", "", "", "", "")
		mock.ExpectQuery(regexp.QuoteMeta(warehouse.CreateWarehouse)).WillReturnRows(rows)
		_, err = warehousesRepo.Create(warehouse1.Address, warehouse1.Telephone, warehouse1.Warehouse_code, warehouse1.Minimum_capacity, warehouse1.Minimum_temperature, warehouse1.Locality_id)
		assert.Error(t, err)
	})

	t.Run("Create - Rows not affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		warehousesRepo := warehouse.NewRepository(db)
		mock.ExpectExec(regexp.QuoteMeta(warehouse.CreateWarehouse)).WithArgs(
			warehouseOne.Address,
			warehouseOne.Telephone,
			warehouseOne.Warehouse_code,
			warehouseOne.Minimum_capacity,
			warehouseOne.Minimum_temperature,
			warehouseOne.Locality_id).WillReturnResult(sqlmock.NewResult(0, 0))

		_, err = warehousesRepo.Create(warehouseOne.Address, warehouseOne.Telephone, warehouseOne.Warehouse_code, warehouseOne.Minimum_capacity, warehouseOne.Minimum_temperature, warehouseOne.Locality_id)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("no rows affected"), err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Udpdate - OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		warehousesRepo := warehouse.NewRepository(db)

		mock.ExpectExec(regexp.QuoteMeta(warehouse.UpdateWarehouse)).WithArgs(
			warehouseOneUpdated.Address,
			warehouseOneUpdated.Telephone,
			warehouseOneUpdated.Warehouse_code,
			warehouseOneUpdated.Minimum_capacity,
			warehouseOneUpdated.Minimum_temperature,
			warehouseOneUpdated.Locality_id,
			warehouseOneUpdated.ID).WillReturnResult(sqlmock.NewResult(0, 1))
		_, err = warehousesRepo.Update(warehouseOneUpdated.ID, warehouseOneUpdated.Address, warehouseOneUpdated.Telephone, warehouseOneUpdated.Warehouse_code, warehouseOneUpdated.Minimum_capacity, warehouseOneUpdated.Minimum_temperature, warehouseOneUpdated.Locality_id)
		assert.NoError(t, err)
	})

	t.Run("Update - Fail on query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		warehousesRepo := warehouse.NewRepository(db)
		mock.ExpectExec(regexp.QuoteMeta(warehouse.UpdateWarehouse)).WillReturnError(fmt.Errorf("Fail on query"))

		_, err = warehousesRepo.Update(warehouseOneUpdated.ID, warehouseOneUpdated.Address, warehouseOneUpdated.Telephone, warehouseOneUpdated.Warehouse_code, warehouseOneUpdated.Minimum_capacity, warehouseOneUpdated.Minimum_temperature, warehouseOneUpdated.Locality_id)

		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("Fail on query"), err)
	})
}
func TestDelete(t *testing.T) {
	t.Run("Delete OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		warehousesRepo := warehouse.NewRepository(db)

		mock.ExpectExec(regexp.QuoteMeta(warehouse.DeleteWarehouse)).WithArgs(warehouseOne.ID).WillReturnResult(driver.RowsAffected(1))

		err = warehousesRepo.Delete(warehouseOne.ID)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
	t.Run("if there is an error in the exec of Delete, it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		warehousesRepo := warehouse.NewRepository(db)

		mock.ExpectExec(regexp.QuoteMeta(warehouse.DeleteWarehouse)).WithArgs(warehouse1.ID).WillReturnError(fmt.Errorf("error"))

		err = warehousesRepo.Delete(warehouse1.ID)

		assert.Error(t, err)

		assert.NotNil(t, err)
	})
	t.Run("if there is no rows affected in Delete, it should return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		warehousesRepo := warehouse.NewRepository(db)

		mock.ExpectExec(regexp.QuoteMeta(warehouse.DeleteWarehouse)).WithArgs(warehouse1.ID).WillReturnResult(driver.ResultNoRows)

		err = warehousesRepo.Delete(warehouse1.ID)

		assert.Error(t, err)

		assert.NotNil(t, err)
	})
}
