package employee

import "fmt"

var emps []Employee = []Employee{}

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
	return emps, nil
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) GetByID(id int) (Employee, error) {
	var emp Employee
	exists := false
	for i := range emps {
		if emps[i].ID == id {
			emp = emps[i]
			exists = true
		}
	}
	if !exists {
		return Employee{}, fmt.Errorf("User with id %d not found", id)
	}
	return emp, nil
}

func (repository) Create(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	emp := Employee{id, cardNumberID, firstName, lastName, warehouseID}
	emps = append(emps, emp)
	lastID = emp.ID

	return emp, nil
}
func (repository) Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	emp := Employee{CardNumberID: cardNumberID, FirstName: firstName, LastName: lastName, WarehouseID: warehouseID}
	updated := false
	for i := range emps {
		if emps[i].ID == id {
			emps[i].ID = id
			emps[i] = emp
			updated = true
		}
	}
	if !updated {
		return Employee{}, fmt.Errorf("User with id %d not found", id)
	}
	return emp, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range emps {
		if emps[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("User with id %d not found", id)
	}
	emps = append(emps[:index], emps[index+1:]...)
	return nil
}

func NewRepository() Repository {
	return &repository{}
}
