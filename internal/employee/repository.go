package employee

import (
	"database/sql"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
)

type Repository interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (Employee, error)
	Create(cardNumberID, firstName, lastName string, warehouseID int) (Employee, error)
	Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error)
	GetByIDWarehouse(id int) (warehouse.Warehouse, error)
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
func (r *repository) GetByIDWarehouse(id int) (warehouse.Warehouse, error) {
	stmt, err := r.db.Prepare("SELECT * FROM warehouse WHERE id = ?")
	if err != nil {
		return warehouse.Warehouse{}, err
	}
	defer stmt.Close()

	var warehouse warehouse.Warehouse

	err = stmt.QueryRow(id).Scan(
		&warehouse.ID,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.Warehouse_code,
	)

	if err != nil {
		return warehouse, err
	}

	return warehouse, nil
}

func (r *repository) Create(cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	employee := Employee{CardNumberID: cardNumberID, FirstName: firstName, LastName: lastName, WarehouseID: warehouseID}

	stmt, err := r.db.Exec(`INSERT INTO employee
	(card_number_id,
		first_name,
		last_name,
		warehouse_id) VALUES(?,?,?,?)`, cardNumberID,
		firstName,
		lastName,
		warehouseID)
	if err != nil {
		return Employee{}, err
	}

	RowsAffected, _ := stmt.RowsAffected()
	if RowsAffected == 0 {
		return Employee{}, fmt.Errorf("fail to save")
	}

	lastID, err := stmt.LastInsertId()
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

	stmt, err := r.db.Exec(`UPDATE employee SET
	card_number_id=?,
	first_name=?,
	last_name=?,
	warehouse_id=? WHERE id=?`,
		cardNumberID,
		firstName,
		lastName,
		warehouseID,
		id)

	if err != nil {
		return Employee{}, err
	}

	RowsAffected, _ := stmt.RowsAffected()
	if RowsAffected == 0 {
		return Employee{}, fmt.Errorf("fail to update")
	}

	return employee, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Exec("DELETE FROM employee WHERE id=?", id)
	if err != nil {
		return err
	}
	RowsAffected, _ := stmt.RowsAffected()
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
