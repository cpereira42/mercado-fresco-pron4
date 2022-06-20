package employee_test

import (
	"fmt"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	employee1 = employee.Employee{ID: 1, CardNumberID: "123", FirstName: "Eduardo", LastName: "Araujo", WarehouseID: 1}

	employee2 = employee.Employee{ID: 2, CardNumberID: "1234", FirstName: "Maria", LastName: "Silva", WarehouseID: 2}

	employee3 = employee.Employee{ID: 3, CardNumberID: "12345", FirstName: "Jefferson", LastName: "Filho", WarehouseID: 1}

	newEmployee = employee.Employee{CardNumberID: "123456", FirstName: "Marta", LastName: "Gomes", WarehouseID: 3}

	employee1Update = employee.Employee{ID: 1, CardNumberID: "123", FirstName: "Gustavo", LastName: "Junior", WarehouseID: 1}

	employeeUpdateSameCardNumberID = employee.Employee{ID: 1, CardNumberID: "1235", FirstName: "Gustavo", LastName: "Junior", WarehouseID: 1}
)

func Test_RepositoryGetAll(t *testing.T) {
	employees := []employee.Employee{employee1, employee2, employee3}
	t.Run("Get all employees Ok", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetAll").Return(employees, nil)

		service := employee.NewService(repo)
		employeesService, err := service.GetAll()

		assert.Nil(t, err)
		assert.True(t, len(employeesService) == 3)
		assert.Equal(t, employees, employeesService)

	})

	t.Run("Get all employees error", func(t *testing.T) {
		errorMsg := fmt.Errorf("Cannot get all")
		repo := &mocks.Repository{}
		repo.On("GetAll").Return([]employee.Employee{}, errorMsg).Once()

		service := employee.NewService(repo)
		_, err := service.GetAll()

		assert.Error(t, err)
		assert.Equal(t, err, errorMsg)

	})
}
