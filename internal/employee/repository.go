package employee

import (
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
		return Employee{}, nil
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

	employee := Employee{id, cardNumberID, firstName, lastName, warehouseID}
	if err := r.db.Write(employees); err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (r *repository) Delete(id int) error {
	if err := r.db.Read(&employees); err != nil {
		return err
	}

	employees = append(employees[:id], employees[id+1:]...)
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
