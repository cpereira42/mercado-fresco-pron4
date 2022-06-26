package employee

import (
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
)

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
	var employees []Employee
	if err := r.db.Read(&employees); err != nil {
		return []Employee{}, nil
	}
	return employees, nil
}

func (r *repository) LastID() (int, error) {
	var employees []Employee
	if err := r.db.Read(&employees); err != nil {
		return 0, err
	}
	if len(employees) == 0 {
		return 0, nil
	}

	return employees[len(employees)-1].ID, nil
}

func (r *repository) GetByID(id int) (Employee, error) {
	var employees []Employee
	if err := r.db.Read(&employees); err != nil {
		return Employee{}, nil
	}
	var employee Employee
	for i := range employees {
		if employees[i].ID == id {
			employee = employees[i]
		}
	}
	return employee, nil
}

func (r *repository) Create(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	var employees []Employee
	if err := r.db.Read(&employees); err != nil {
		return Employee{}, err
	}
	employee := Employee{id, cardNumberID, firstName, lastName, warehouseID}

	employees = append(employees, employee)
	if err := r.db.Write(employees); err != nil {
		return Employee{}, err
	}
	return employee, nil
}
func (r *repository) Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	var employees []Employee
	if err := r.db.Read(&employees); err != nil {
		return Employee{}, err
	}

	employee := Employee{id, cardNumberID, firstName, lastName, warehouseID}

	for i := range employees {
		if employees[i].ID == id {
			employees[i] = employee
		}
	}

	if err := r.db.Write(employees); err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (r *repository) Delete(id int) error {
	var employees []Employee
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
