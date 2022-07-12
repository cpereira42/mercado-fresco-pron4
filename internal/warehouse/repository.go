package warehouse

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	GetAll() ([]Warehouse, error)
	Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_id int) (Warehouse, error)
	//LastID() (int, error)
	Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_id int) (Warehouse, error)
	GetByID(id int) (Warehouse, error)
	Delete(id int) error
	CheckLocality(id int) (bool, error)
}

const (
	GetAll = `SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id FROM warehouse`
	//LastID          = "SELECT MAX(id) FROM warehouse"
	CreateWarehouse = `INSERT INTO warehouse (address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id) VALUES (?, ?, ?, ?, ?, ?)`
	UpdateWarehouse = `UPDATE warehouse SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=?, locality_id=? WHERE id=?`
	GetId           = `SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id FROM warehouse WHERE id=?`
	DeleteWarehouse = `DELETE FROM warehouse WHERE id=?`
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

		err := rows.Scan(&w.ID, &w.Address, &w.Telephone, &w.Warehouse_code, &w.Minimum_capacity, &w.Minimum_temperature, &w.Locality_id)
		if err != nil {
			return []Warehouse{}, err
		}
		wr = append(wr, w)
	}

	return wr, nil
}

func (r *repository) Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_id int) (Warehouse, error) {

	w = Warehouse{id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id}

	stmt, err := r.db.Exec(CreateWarehouse, &w.Address, &w.Telephone, &w.Warehouse_code, &w.Minimum_capacity, &w.Minimum_temperature, locality_id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := stmt.RowsAffected()
	if err != nil {
		return Warehouse{}, err
	}
	if rowsAffected == 0 {
		return Warehouse{}, errors.New("no rows affected")
	}
	insertedId, err := stmt.LastInsertId()
	if err != nil {
		return Warehouse{}, err
	}
	w.ID = int(insertedId)
	return w, nil

}

func (r *repository) Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_id int) (Warehouse, error) {
	w := Warehouse{id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id}

	stmt, err := r.db.Exec(UpdateWarehouse, &w.Address, &w.Telephone, &w.Warehouse_code, &w.Minimum_capacity, &w.Minimum_temperature, &w.Locality_id, id)
	if err != nil {
		return Warehouse{}, err
	}
	rowsAffected, err := stmt.RowsAffected()
	if err != nil {
		return Warehouse{}, err
	}
	if rowsAffected == 0 {
		return Warehouse{}, errors.New("no rows affected")
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

	stmt, err := r.db.Prepare(DeleteWarehouse)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := r.db.Exec(DeleteWarehouse, id)
	if err != nil {
		fmt.Println(err)
		return util.CheckError(err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("ID not Found")
	}

	return nil
}

func (r *repository) CheckLocality(id int) (bool, error) {

	type Locality struct {
		Id           int
		LocalityName string
		ProvinceName string
		CountryName  string
	}

	stmt, err := r.db.Prepare(`SELECT id, locality_name, province_name, country_name FROM localities WHERE id = ?`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var locality Locality

	err = stmt.QueryRow(id).Scan(&locality.Id, &locality.LocalityName, &locality.ProvinceName, &locality.CountryName)
	if err != nil {
		return false, fmt.Errorf("locality %d not found", id)
	}
	return true, nil
}
