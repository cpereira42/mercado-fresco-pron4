package employee

import "fmt"

var employees []Employee = []Employee{}

var lastID int

type Repository interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (Employee, error)
	Create(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error)
	LastID() (int, error)
	Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error)
	Delete(id int) error
}

type repository struct{}

func (repository) GetAll() ([]Employee, error) {
	return employees, nil
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) GetByID(id int) (Employee, error) {
	var employee Employee
	exists := false
	for i := range employees {
		if employees[i].ID == id {
			employee = employees[i]
			exists = true
		}
	}
	if !exists {
		return Employee{}, fmt.Errorf("User with id %d not found", id)
	}
	return employee, nil
}

func (repository) Create(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	employee := Employee{id, cardNumberID, firstName, lastName, warehouseID}
	employees = append(employees, employee)
	lastID = employee.ID

	return employee, nil
}
func (repository) Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	employee := Employee{CardNumberID: cardNumberID, FirstName: firstName, LastName: lastName, WarehouseID: warehouseID}
	updated := false
	for i := range employees {
		if employees[i].ID == id {
			employees[i].ID = id
			employees[i] = employee
			updated = true
		}
	}
	if !updated {
		return Employee{}, fmt.Errorf("User with id %d not found", id)
	}
	return employee, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range employees {
		if employees[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("User with id %d not found", id)
	}
	employees = append(employees[:index], employees[index+1:]...)
	return nil
}

func NewRepository() Repository {
	return &repository{}
}
