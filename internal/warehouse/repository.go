package warehouse

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	GetAll() ([]Warehouse, error)
	Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error)
	LastID() (int, error)
	Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error)
	GetByID(id int) (Warehouse, error)
	Delete(id int) error
}

const (
	GetAll          = "SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature FROM warehouse"
	LastID          = "SELECT MAX(id) FROM warehouse"
	CreateWarehouse = "INSERT INTO warehouse (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)"
	UpdateWarehouse = "UPDATE warehouse SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?"
	GetId           = "SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature FROM warehouse WHERE id=?"
	DeleteWarehouse = "DELETE FROM warehouse WHERE id=?"
)

//var wr []Warehouse
var w Warehouse

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db}
}

func (r *repository) GetAll() ([]Warehouse, error) {

	var wr []Warehouse

	rows, err := r.db.Query(GetAll)
	if err != nil {
		return []Warehouse{}, err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&w.ID, &w.Address, &w.Telephone, &w.Warehouse_code, &w.Minimum_capacity, &w.Minimum_temperature)
		if err != nil {
			return []Warehouse{}, err
		}
		wr = append(wr, w)
	}

	return wr, nil
}

func (r *repository) LastID() (int, error) {
	var maxCount int

	row := r.db.QueryRow(LastID)
	err := row.Scan(&maxCount)
	if err != nil {
		return 0, err
	}
	return maxCount, nil
}

func (r *repository) Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {

	w = Warehouse{id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature}

	stmt, err := r.db.Prepare(CreateWarehouse)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var result sql.Result

	result, err = stmt.Exec(&w.Address, &w.Telephone, &w.Warehouse_code, &w.Minimum_capacity, &w.Minimum_temperature)
	if err != nil {
		return Warehouse{}, err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return Warehouse{}, err
	}
	w.ID = int(insertedId)
	return w, nil

}

func (r *repository) Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {
	w := Warehouse{id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature}

	stmt, err := r.db.Prepare(UpdateWarehouse)
	if err != nil {
		return Warehouse{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(&w.Address, &w.Telephone, &w.Warehouse_code, &w.Minimum_capacity, &w.Minimum_temperature, &w.ID)
	if err != nil {
		return Warehouse{}, err
	}

	lines, err := rows.RowsAffected()
	if err != nil {
		return Warehouse{}, err
	}
	if lines == 0 {
		return Warehouse{}, sql.ErrNoRows
	}

	return w, nil
}

func (r *repository) GetByID(id int) (Warehouse, error) {

	stmt, err := r.db.Prepare(GetId)
	if err != nil {
		return Warehouse{}, err
	}
	defer stmt.Close()
	res, err := r.db.Query(GetId, id)

	if err != nil {
		return Warehouse{}, err
	}
	defer stmt.Close()

	for res.Next() {
		if err := res.Scan(&w.ID, &w.Address, &w.Telephone, &w.Warehouse_code, &w.Minimum_capacity, &w.Minimum_temperature); err != nil {
			return Warehouse{}, err
		}
	}
	return w, nil

}

func (r *repository) Delete(id int) error {

	stmt, err := r.db.Prepare(DeleteWarehouse)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := r.db.Exec(DeleteWarehouse, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
