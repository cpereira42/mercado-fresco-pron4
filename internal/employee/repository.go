package employee

import (
	"database/sql"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
)

type Repository interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (Employee, error)
	Create(cardNumberID, firstName, lastName string, warehouseID int) (Employee, error)
	Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll() ([]Employee, error) {
	var employees []Employee

	rows, err := r.db.Query(GET_ALL_EMPLOYEES)
	if err != nil {
		return employees, err
	}

	defer rows.Close()

	for rows.Next() {
		var employee Employee

		err := rows.Scan(
			&employee.ID,
			&employee.CardNumberID,
			&employee.FirstName,
			&employee.LastName,
			&employee.WarehouseID,
		)
		if err != nil {
			return employees, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (r *repository) GetByID(id int) (Employee, error) {
	stmt, err := r.db.Prepare(GET_EMPLOYEE_BY_ID)
	if err != nil {
		return Employee{}, fmt.Errorf(FAIL_TO_PREPARE_QUERY)
	}
	defer stmt.Close()

	var employee Employee

	err = stmt.QueryRow(id).Scan(
		&employee.ID,
		&employee.CardNumberID,
		&employee.FirstName,
		&employee.LastName,
		&employee.WarehouseID,
	)

	if err != nil {
		return Employee{}, fmt.Errorf(EMPLOYEE_NOT_FOUND)
	}

	return employee, nil
}

func (r *repository) Create(cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	employee := Employee{CardNumberID: cardNumberID, FirstName: firstName, LastName: lastName, WarehouseID: warehouseID}

	stmt, err := r.db.Exec(CREATE_EMPLOYEE, cardNumberID, firstName, lastName,
		warehouseID)
	if err != nil {
		return Employee{}, util.CheckError(err)
	}

	RowsAffected, _ := stmt.RowsAffected()
	if RowsAffected == 0 {
		return Employee{}, fmt.Errorf(FAIL_TO_SAVE)
	}

	lastID, _ := stmt.LastInsertId()

	employee.ID = int(lastID)

	return employee, nil
}
func (r *repository) Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	employee := Employee{id, cardNumberID, firstName, lastName, warehouseID}

	_, err := r.db.Exec(UPDATE_EMPLOYEE, cardNumberID, firstName, lastName,
		warehouseID, id)

	if err != nil {
		return Employee{}, util.CheckError(err)
	}

	return employee, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Exec(DELETE_EMPLOYEE, id)
	if err != nil {
		return util.CheckError(err)
	}
	RowsAffected, _ := stmt.RowsAffected()
	if RowsAffected == 0 {
		return fmt.Errorf(EMPLOYEE_NOT_FOUND)
	}
	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
