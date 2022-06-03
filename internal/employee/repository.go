package employee

import (
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
)

var employees []Employee
var employee Employee

type Repository interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (Employee, error)
	Create(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error)
	LastID() (int, error)
	Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Employee, error) {
	if err := r.db.Read(&employees); err != nil {
		return []Employee{}, nil
	}
	return employees, nil
}

func (r *repository) LastID() (int, error) {
	if err := r.db.Read(&employees); err != nil {
		return 0, err
	}
	if len(employees) == 0 {
		return 0, nil
	}

	return employees[len(employees)-1].ID, nil
}

func (r *repository) GetByID(id int) (Employee, error) {
	if err := r.db.Read(&employees); err != nil {
		return Employee{}, err
	}
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

func (r *repository) Create(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	if err := r.db.Read(&employees); err != nil {
		return Employee{}, err
	}
	employee = Employee{id, cardNumberID, firstName, lastName, warehouseID}
	employees = append(employees, employee)
	if err := r.db.Write(employees); err != nil {
		return Employee{}, err
	}
	return employee, nil
}
func (r *repository) Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	fmt.Println("id no repository", id)
	if err := r.db.Read(&employees); err != nil {
		return Employee{}, err
	}
	employee = Employee{CardNumberID: cardNumberID, FirstName: firstName, LastName: lastName, WarehouseID: warehouseID}
	updated := false
	for i := range employees {
		if employees[i].ID == id {
			employee.ID = id
			employees[i] = employee
			updated = true
		}
	}
	fmt.Println("id employee no repository", employee.ID)
	if !updated {
		return Employee{}, fmt.Errorf("user with id %d not found", id)
	}
	if err := r.db.Write(employees); err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (r *repository) Delete(id int) error {
	if err := r.db.Read(&employees); err != nil {
		return err
	}
	deleted := false
	var index int
	for i := range employees {
		if employees[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("user with id %d not found", id)
	}
	employees = append(employees[:index], employees[index+1:]...)
	if err := r.db.Write(employees); err != nil {
		return err
	}

	return nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
