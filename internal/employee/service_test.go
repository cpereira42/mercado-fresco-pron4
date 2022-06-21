package employee_test

import (
	"fmt"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee/mocks"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var (
	employee1 = employee.Employee{ID: 1, CardNumberID: "123", FirstName: "Eduardo", LastName: "Araujo", WarehouseID: 1}

	employee2 = employee.Employee{ID: 2, CardNumberID: "1234", FirstName: "Maria", LastName: "Silva", WarehouseID: 2}

	employee3 = employee.Employee{ID: 3, CardNumberID: "12345", FirstName: "Jefferson", LastName: "Filho", WarehouseID: 1}

	newEmployee = employee.Employee{CardNumberID: "123456", FirstName: "Marta", LastName: "Gomes", WarehouseID: 3}

	employee1Update = employee.Employee{ID: 1, CardNumberID: "123", FirstName: "Gustavo", LastName: "Junior", WarehouseID: 1}

	employeeUpdateSameCardNumberID = employee.Employee{ID: 2, CardNumberID: "123", FirstName: "Gustavo", LastName: "Junior", WarehouseID: 1}
)

func TestServiceGetAll(t *testing.T) {
	employees := []employee.Employee{employee1, employee2, employee3}
	t.Run("If GetAll is ok, it should return a list of employees", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetAll").Return(employees, nil)

		service := employee.NewService(repo)
		employeesService, err := service.GetAll()

		assert.Nil(t, err)
		assert.True(t, len(employeesService) == 3)
		assert.Equal(t, employees, employeesService)

	})

	t.Run("If GetAll has an error, it should return an error", func(t *testing.T) {
		errorMsg := fmt.Errorf("Failed to get all employees")
		repo := &mocks.Repository{}
		repo.On("GetAll").Return([]employee.Employee{}, errorMsg).Once()

		service := employee.NewService(repo)
		employeesService, err := service.GetAll()

		assert.True(t, len(employeesService) == 0)
		assert.Equal(t, fmt.Errorf("Failed to get all employees"), err)
		repo.AssertExpectations(t)

	})
}

func TestServiceGetByID(t *testing.T) {
	employees := []employee.Employee{employee1, employee2, employee3}
	t.Run("If GetByID is success, it should return a employee",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("GetByID", tmock.AnythingOfType("int")).Return(employee1, nil).Once()

			service := employee.NewService(repo)
			employeeByID, err := service.GetByID(1)

			assert.NoError(t, err)
			assert.Equal(t, employee1, employeeByID)

		})
	t.Run("If GetByID has an error to get all employees, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("Failed to get all employees in the GetByID ")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return([]employee.Employee{}, errorMsg).Once()

			service := employee.NewService(repo)
			_, err := service.GetByID(1)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())

		})
	t.Run("If the id in the GetByID does not exists, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("user with id 10 not found")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("GetByID", tmock.AnythingOfType("int")).Return(employee.Employee{}, errorMsg).Once()

			service := employee.NewService(repo)
			_, err := service.GetByID(10)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())

		})

	t.Run("If GetByID has an error, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("error to GetByID")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("GetByID", tmock.AnythingOfType("int")).Return(employee.Employee{}, errorMsg).Once()

			service := employee.NewService(repo)
			_, err := service.GetByID(1)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())
		})
}

func TestServiceCreate(t *testing.T) {
	employees := []employee.Employee{employee1, employee2, employee3}

	t.Run("If Create is success, it should return a employee",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("LastID").Return(3, nil).Once()
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
			).Return(newEmployee, nil)

			service := employee.NewService(repo)
			employeeCreated, err := service.Create(newEmployee.CardNumberID, newEmployee.FirstName, newEmployee.LastName, newEmployee.WarehouseID)

			assert.NoError(t, err)
			assert.Equal(t, newEmployee, employeeCreated)

		})
	t.Run("If Create has an error to get last id, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("Failed to get last id in the Create")
			repo := &mocks.Repository{}
			repo.On("LastID").Return(0, errorMsg).Once()

			service := employee.NewService(repo)
			_, err := service.Create(newEmployee.CardNumberID, newEmployee.FirstName, newEmployee.LastName, newEmployee.WarehouseID)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())

		})
	t.Run("If Create has an error to get all employees, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("Failed to get all employees in the Create")
			repo := &mocks.Repository{}
			repo.On("LastID").Return(3, nil).Once()
			repo.On("GetAll").Return([]employee.Employee{}, errorMsg).Once()

			service := employee.NewService(repo)
			_, err := service.Create(newEmployee.CardNumberID, newEmployee.FirstName, newEmployee.LastName, newEmployee.WarehouseID)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())

		})
	t.Run("If the CardNumberID  in the Create exists, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("user with this card number id 123 exists")
			repo := &mocks.Repository{}
			repo.On("LastID").Return(3, nil).Once()
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
			).Return(employee.Employee{}, errorMsg)

			service := employee.NewService(repo)
			_, err := service.Create(employeeUpdateSameCardNumberID.CardNumberID, employeeUpdateSameCardNumberID.FirstName, employeeUpdateSameCardNumberID.LastName, employeeUpdateSameCardNumberID.WarehouseID)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())

		})

	t.Run("If Create has an error, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("error to Create")
			repo := &mocks.Repository{}
			repo.On("LastID").Return(3, nil).Once()
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("GetByID", tmock.AnythingOfType("int")).Return(employee.Employee{}, errorMsg).Once()
			repo.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
			).Return(employee.Employee{}, errorMsg)

			service := employee.NewService(repo)
			_, err := service.Create(newEmployee.CardNumberID, newEmployee.FirstName, newEmployee.LastName, newEmployee.WarehouseID)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())
		})
}

func TestServiceUpdate(t *testing.T) {
	employees := []employee.Employee{employee1, employee2, employee3}

	t.Run("If Update is success, it should return an updated employee",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
			).Return(employee1Update, nil)

			service := employee.NewService(repo)
			employeeUpdated, err := service.Update(employee1Update.ID, employee1Update.CardNumberID, employee1Update.FirstName, employee1Update.LastName, employee1Update.WarehouseID)

			assert.NoError(t, err)
			assert.Equal(t, employee1Update, employeeUpdated)

		})

	t.Run("If Update has an error to get all employees, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("Failed to get all employees in the Update")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return([]employee.Employee{}, errorMsg).Once()

			service := employee.NewService(repo)
			_, err := service.Update(employee1Update.ID, employee1Update.CardNumberID, employee1Update.FirstName, employee1Update.LastName, employee1Update.WarehouseID)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())

		})

	t.Run("If the id in the Update does not exists, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("user with id 10 not found")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
			).Return(employee.Employee{}, errorMsg)

			service := employee.NewService(repo)
			_, err := service.Update(10, "", "", "", 0)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())

		})
	t.Run("If the CardNumberID  in the Update exists, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("user with this card number id 123 exists")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
			).Return(employee.Employee{}, errorMsg)

			service := employee.NewService(repo)
			_, err := service.Update(employeeUpdateSameCardNumberID.ID, employeeUpdateSameCardNumberID.CardNumberID, employeeUpdateSameCardNumberID.FirstName, employeeUpdateSameCardNumberID.LastName, employeeUpdateSameCardNumberID.WarehouseID)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())

		})
	t.Run(
		"If CardNumberID is empty, should return the same CardNumberID of the employee",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
			).Return(employee1Update, nil)

			service := employee.NewService(repo)
			employeeUpdated, err := service.Update(employee1Update.ID, "", employee1Update.FirstName, employee1Update.LastName, employee1Update.WarehouseID)

			assert.NoError(t, err)
			assert.Equal(t, employee1Update, employeeUpdated)
		})
	t.Run(
		"If FirstName is empty, should return the same FirstName of the employee",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
			).Return(employee1Update, nil)

			service := employee.NewService(repo)
			employeeUpdated, err := service.Update(employee1Update.ID, employee1.CardNumberID, "", employee1Update.LastName, employee1Update.WarehouseID)

			assert.NoError(t, err)
			assert.Equal(t, employee1Update, employeeUpdated)
		})
	t.Run(
		"If LastName is empty, should return the same LastName of the employee",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
			).Return(employee1Update, nil)

			service := employee.NewService(repo)
			employeeUpdated, err := service.Update(employee1Update.ID, employee1.CardNumberID, employee1.FirstName, "", employee1Update.WarehouseID)

			assert.NoError(t, err)
			assert.Equal(t, employee1Update, employeeUpdated)
		})
	t.Run(
		"If WarehouseID is empty, should return the same WarehouseID of the employee",
		func(t *testing.T) {
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
			).Return(employee1Update, nil)

			service := employee.NewService(repo)
			employeeUpdated, err := service.Update(employee1Update.ID, employee1.CardNumberID, employee1.FirstName, employee1.LastName, 0)

			assert.NoError(t, err)
			assert.Equal(t, employee1Update, employeeUpdated)
		})

	t.Run("If Update has an error, it should return an error",
		func(t *testing.T) {
			errorMsg := fmt.Errorf("error to Update")
			repo := &mocks.Repository{}
			repo.On("GetAll").Return(employees, nil).Once()
			repo.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("int"),
			).Return(employee.Employee{}, errorMsg)

			service := employee.NewService(repo)
			_, err := service.Update(employee1Update.ID, employee1Update.CardNumberID, employee1Update.FirstName, employee1Update.LastName, employee1Update.WarehouseID)

			assert.Error(t, err)
			assert.EqualError(t, err, errorMsg.Error())
		})
}
