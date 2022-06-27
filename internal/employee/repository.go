package employee

import (
	"database/sql"
	"fmt"
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

	rows, err := r.db.Query("SELECT * FROM employee")
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
	stmt, err := r.db.Prepare("SELECT * FROM employee WHERE id = ?")
	if err != nil {
		return Employee{}, err
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
		return Employee{}, err
	}

	return employee, nil
}

func (r *repository) Create(cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	employee := Employee{0, cardNumberID, firstName, lastName, warehouseID}

	stmt, err := r.db.Prepare(`INSERT INTO employee
	(card_number_id,
	first_name,
	last_name,
	warehouse_id) VALUES(?,?,?,?)`)
	if err != nil {
		return Employee{}, err
	}
	defer stmt.Close()
	rows, err := stmt.Exec(
		cardNumberID,
		firstName,
		lastName,
		warehouseID,
	)
	RowsAffected, _ := rows.RowsAffected()
	if RowsAffected == 0 {
		return Employee{}, fmt.Errorf("fail to save")
	}

	lastID, err := rows.LastInsertId()
	if err != nil {
		return Employee{}, err
	}

	employee.ID = int(lastID)

	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}
func (r *repository) Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	employee := Employee{id, cardNumberID, firstName, lastName, warehouseID}
	stmt, err := r.db.Prepare(`UPDATE employee SET
	card_number_id=?,
	first_name=?,
	last_name=?,
	warehouse_id=? WHERE id=?`)

	if err != nil {
		return Employee{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Exec(
		cardNumberID,
		firstName,
		lastName,
		warehouseID,
	)
	if err != nil {
		return Employee{}, err
	}
	RowsAffected, _ := rows.RowsAffected()
	if RowsAffected == 0 {
		return Employee{}, fmt.Errorf("fail to update")
	}

	return employee, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM employee WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	RowsAffected, _ := rows.RowsAffected()
	if RowsAffected == 0 {
		return fmt.Errorf("fail to delete")
	}
	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
