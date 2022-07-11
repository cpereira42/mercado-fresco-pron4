package carries

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
)

type Repository interface {
	GetByIDReport(id int) (Localities, error)
	Create(cid, companyName, address, telephone string, localityID int) (Carries, error)
	GetAllReport() ([]Localities, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllReport() ([]Localities, error) {
	var localities []Localities
	rows, err := r.db.Query(GetAllReport)
	if err != nil {
		return localities, err
	}
	defer rows.Close()
	for rows.Next() {
		var locality Localities
		err := rows.Scan(&locality.LocalityID, &locality.LocalityName, &locality.Count)
		if err != nil {
			return localities, err
		}
		localities = append(localities, locality)
	}
	return localities, nil

}
func (r *repository) GetByIDReport(id int) (Localities, error) {
	var locality Localities
	row := r.db.QueryRow(GetByIDReport, id)

	err := row.Scan(&locality.LocalityID, &locality.LocalityName, &locality.Count)

	if errors.Is(err, sql.ErrNoRows) {
		return locality, fmt.Errorf(failedIdNotFound)
	}
	if err != nil {
		return locality, util.CheckError(err)
	}
	return locality, nil

}
func (r *repository) Create(cid, company_name, address, telephone string, localityID int) (Carries, error) {
	newCarry := Carries{Cid: cid, CompanyName: company_name, Address: address, Telephone: telephone, LocalityID: localityID}

	result, err := r.db.Exec(CreateCarry, &newCarry.Cid, &newCarry.CompanyName, &newCarry.Address, &newCarry.Telephone, &newCarry.LocalityID)
	if err != nil {
		return Carries{}, util.CheckError(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Carries{}, fmt.Errorf("failed to create carry")
	}

	return newCarry, nil
}
