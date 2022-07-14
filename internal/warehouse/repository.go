package warehouse

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	GetAll() ([]Warehouse, error)
	Create(address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_id int) (RequestWarehouseCreate, error)
	Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_id int) (Warehouse, error)
	GetByID(id int) (Warehouse, error)
	Delete(id int) error
}

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

		err := rows.Scan(&w.ID, &w.Address, &w.Telephone, &w.Warehouse_code, &w.Minimum_capacity, &w.Minimum_temperature, &w.Locality_id)
		if err != nil {
			return []Warehouse{}, err
		}
		wr = append(wr, w)
	}

	return wr, nil
}

func (r *repository) Create(address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_id int) (RequestWarehouseCreate, error) {

	w := RequestWarehouseCreate{address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id}

	stmt, err := r.db.Exec(CreateWarehouse, &w.Address, &w.Telephone, &w.Warehouse_code, &w.Minimum_capacity, &w.Minimum_temperature, &w.Locality_id)
	if err != nil {
		return RequestWarehouseCreate{}, err
	}
	rowsAffected, _ := stmt.RowsAffected()
	if rowsAffected == 0 {
		return RequestWarehouseCreate{}, errors.New("no rows affected")
	}
	return w, nil

}

func (r *repository) Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_id int) (Warehouse, error) {
	w := Warehouse{id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id}

	_, err := r.db.Exec(UpdateWarehouse, &w.Address, &w.Telephone, &w.Warehouse_code, &w.Minimum_capacity, &w.Minimum_temperature, &w.Locality_id, id)
	if err != nil {
		return Warehouse{}, util.CheckError(err)
	}
	return w, nil
}

func (r *repository) GetByID(id int) (Warehouse, error) {

	var w Warehouse
	stmt, err := r.db.Prepare(GetId)
	if err != nil {
		return Warehouse{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&w.ID,
		&w.Address,
		&w.Telephone,
		&w.Warehouse_code,
		&w.Minimum_capacity,
		&w.Minimum_temperature,
		&w.Locality_id,
	)

	if err != nil {
		return Warehouse{}, fmt.Errorf("Warehouse %d not found", id)
	}
	return w, nil

}

func (r *repository) Delete(id int) error {

	stmt, err := r.db.Exec(DeleteWarehouse, id)
	if err != nil {
		return util.CheckError(err)
	}
	rowsAffected, _ := stmt.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("ID not Found")
	}

	return nil
}
