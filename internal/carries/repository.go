package carries

import (
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	GetByIDReport(id int) (Localities, error)
	Create(cid, companyName, address, telephone string, localityID int) (Carries, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	CreateCarrie  = `INSERT INTO carries (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)`
	GetByID       = `SELECT id, locality_id, cid, company_name, address, telephone FROM carries WHERE id=?`
	GetByIDReport = `SELECT locality_id, locality_name, COUNT(carries.locality_id)
	FROM localities
	INNER JOIN carries ON localities.id = carries.locality_id
	WHERE localities.id = ?`
)

// func (r *repository) GetByID(id int) (Carries, error) {
// 	row := r.db.QueryRow("SELECT * from carries WHERE id = ?", id)

// 	var carry Carries

// 	err := row.Scan(&carry.ID, &carry.Cid, &carry.CompanyName, &carry.Address, &carry.Telephone, &carry.LocalityID)

// 	if errors.Is(err, sql.ErrNoRows) {
// 		return carry, fmt.Errorf("id %d not found", id)
// 	}
// 	if err != nil {
// 		return carry, err
// 	}
// 	return carry, nil
// }
func (r *repository) GetByIDReport(id int) (Localities, error) {
	row := r.db.QueryRow(GetByIDReport, id)

	var locality Localities

	err := row.Scan(&locality.LocalityID, &locality.LocalityName, &locality.Count)

	if errors.Is(err, sql.ErrNoRows) {
		return locality, fmt.Errorf("id %d not found", id)
	}
	if err != nil {
		return locality, err
	}
	return locality, nil

}
func (r *repository) Create(cid, company_name, address, telephone string, localityID int) (Carries, error) {
	newCarry := Carries{Cid: cid, CompanyName: company_name, Address: address, Telephone: telephone, LocalityID: localityID}

	result, err := r.db.Exec(CreateCarrie, &newCarry.Cid, &newCarry.CompanyName, &newCarry.Address, &newCarry.Telephone, &newCarry.LocalityID)
	if err != nil {
		return newCarry, err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return newCarry, err
	}
	if err != nil {
		return newCarry, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return newCarry, err
	}
	newCarry.ID = int(lastID)

	return newCarry, nil
}
