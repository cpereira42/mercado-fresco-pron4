package warehouse_test

import (
	"fmt"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse/mocks"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var (
	warehouse1 = warehouse.Warehouse{ID: 1, Address: "Rua 1", Telephone: "11111-1111", Warehouse_code: "W1", Minimum_capacity: 10, Minimum_temperature: 20}
	warehouse2 = warehouse.Warehouse{ID: 2, Address: "Rua 2", Telephone: "22222-2222", Warehouse_code: "W2", Minimum_capacity: 20, Minimum_temperature: 30}
	warehouse3 = warehouse.Warehouse{ID: 3, Address: "Rua 3", Telephone: "33333-3333", Warehouse_code: "W3", Minimum_capacity: 30, Minimum_temperature: 40}
)

func TestServiceCreate(t *testing.T) {

	warehouseListSucess := []warehouse.Warehouse{warehouse1, warehouse2}
	warehouseListError := []warehouse.Warehouse{warehouse1}

	t.Run(
		"If all fields are valid, should return a warehouse",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(warehouseListSucess, nil).Once()
			repo.On("LastID").Return(2, nil).Once()
			repo.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse3, nil)
			service := warehouse.NewService(repo)
			newWarehouse, err := service.Create(warehouse3.Address, warehouse3.Telephone, warehouse3.Warehouse_code, warehouse3.Minimum_capacity, warehouse3.Minimum_temperature)
			assert.NoError(t, err)
			assert.Equal(t, "Rua 3", newWarehouse.Address)
			assert.Equal(t, "33333-3333", newWarehouse.Telephone)
			assert.Equal(t, "W3", newWarehouse.Warehouse_code)
			assert.Equal(t, 30, newWarehouse.Minimum_capacity)
			assert.Equal(t, 40, newWarehouse.Minimum_temperature)
			repo.AssertExpectations(t)

		})
	t.Run(
		"If Warehouse_code is already exists, should return an error",
		func(t *testing.T) {
			errorMsgWarehouseCodeAlreadyExists := "Warehouse already exists"
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(warehouseListError, nil).Once()
			repo.On("LastID").Return(1, nil).Once()
			repo.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse.Warehouse{}, errorMsgWarehouseCodeAlreadyExists)

			service := warehouse.NewService(repo)

			newWarehouse, err := service.Create(warehouse1.Address, warehouse1.Telephone, warehouse1.Warehouse_code, warehouse1.Minimum_capacity, warehouse1.Minimum_temperature)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsgWarehouseCodeAlreadyExists)
			assert.ObjectsAreEqual(warehouse.Warehouse{}, newWarehouse)

		})
	t.Run(
		"If create cannot save on DB, should return an error - GetAll",
		func(t *testing.T) {
			errorMsgCannotSaveOnDB := fmt.Errorf("Cannot save on DB")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return([]warehouse.Warehouse{}, errorMsgCannotSaveOnDB).Once()
			repo.On("LastID").Return(1, nil).Once()
			repo.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse.Warehouse{}, nil)

			service := warehouse.NewService(repo)

			newWarehouse, err := service.Create(warehouse3.Address, warehouse3.Telephone, warehouse3.Warehouse_code, warehouse3.Minimum_capacity, warehouse3.Minimum_temperature)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsgCannotSaveOnDB.Error())
			assert.ObjectsAreEqual([]warehouse.Warehouse{}, newWarehouse)

		})
	t.Run(
		"If create cannot save on DB, should return an error - LastIdError",
		func(t *testing.T) {
			errorMsgCannotSaveOnDB := fmt.Errorf("Cannot save on DB")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(warehouseListSucess, nil).Once()
			repo.On("LastID").Return(0, errorMsgCannotSaveOnDB).Once()
			repo.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse.Warehouse{}, nil)

			service := warehouse.NewService(repo)

			newWarehouse, err := service.Create(warehouse3.Address, warehouse3.Telephone, warehouse3.Warehouse_code, warehouse3.Minimum_capacity, warehouse3.Minimum_temperature)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsgCannotSaveOnDB.Error())
			assert.ObjectsAreEqual([]warehouse.Warehouse{}, newWarehouse)

		})
	t.Run(
		"If create cannot save on DB, should return an error - Create",
		func(t *testing.T) {
			errorMsgCannotSaveOnDB := fmt.Errorf("Cannot save on DB")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(warehouseListSucess, nil).Once()
			repo.On("LastID").Return(1, nil).Once()
			repo.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int")).
				Return(warehouse.Warehouse{}, errorMsgCannotSaveOnDB)

			service := warehouse.NewService(repo)

			newWarehouse, err := service.Create(warehouse3.Address, warehouse3.Telephone, warehouse3.Warehouse_code, warehouse3.Minimum_capacity, warehouse3.Minimum_temperature)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsgCannotSaveOnDB.Error())
			assert.ObjectsAreEqual([]warehouse.Warehouse{}, newWarehouse)

		})

}

func TestServiceGetAll(t *testing.T) {

	warehouseListSucess := []warehouse.Warehouse{warehouse1, warehouse2}
	//warehouseListError := []warehouse.Warehouse{warehouse1}

	t.Run(
		"If GetAll is success, should return a list of warehouses",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(warehouseListSucess, nil).Once()
			service := warehouse.NewService(repo)
			warehouseAllList, err := service.GetAll()

			assert.NoError(t, err)
			assert.Equal(t, warehouseAllList, warehouseListSucess)
		})
	t.Run(
		"If GetAll is error, should return an error",
		func(t *testing.T) {
			errorMsgCannotGetAll := fmt.Errorf("Cannot get all")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return([]warehouse.Warehouse{}, errorMsgCannotGetAll).Once()
			service := warehouse.NewService(repo)
			_, err := service.GetAll()

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsgCannotGetAll.Error())

		})

}

func TestServiceGetById(t *testing.T) {

	warehouseListSucess := []warehouse.Warehouse{warehouse1, warehouse2}

	t.Run(
		"If GetById is success, should return a warehouse",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(warehouseListSucess, nil).Once()
			repo.On("GetByID", tmock.AnythingOfType("int")).Return(warehouse1, nil).Once()
			service := warehouse.NewService(repo)
			warehouseById, err := service.GetByID(1)

			assert.NoError(t, err)
			assert.Equal(t, warehouse1, warehouseById)
		})
	t.Run(
		"If GetByID has an error use GetAll, should return an error",
		func(t *testing.T) {
			errorMsgCannotGetAll := fmt.Errorf("Cannot get all")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return([]warehouse.Warehouse{}, errorMsgCannotGetAll).Once()
			service := warehouse.NewService(repo)
			_, err := service.GetByID(1)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsgCannotGetAll.Error())

		})

	t.Run(
		"If GetById dont found a ID, should return an error",
		func(t *testing.T) {
			errorMsgCannotGetById := fmt.Errorf("Warehouse not found")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(warehouseListSucess, nil).Once()
			repo.On("GetByID", tmock.AnythingOfType("int")).Return(warehouse.Warehouse{}, errorMsgCannotGetById).Once()
			service := warehouse.NewService(repo)
			_, err := service.GetByID(4)
			assert.Error(t, err)
			assert.EqualError(t, err, errorMsgCannotGetById.Error())

		})
	t.Run(
		"If GetById dont found a ID on DB, should return an error",
		func(t *testing.T) {
			errorMsgCannotGetById := fmt.Errorf("cannot get on bd")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(warehouseListSucess, nil).Once()
			repo.On("GetByID", tmock.AnythingOfType("int")).Return(warehouse.Warehouse{}, errorMsgCannotGetById).Once()
			service := warehouse.NewService(repo)
			_, err := service.GetByID(1)
			assert.Error(t, err)
			assert.EqualError(t, err, errorMsgCannotGetById.Error())

		})

}
