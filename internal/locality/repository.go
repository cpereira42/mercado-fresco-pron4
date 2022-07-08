package locality

import (
	"database/sql"
	"fmt"
	"log"
)

type RepositoryLocality interface {
	Create(id int, localityName, provinceName, countryName string) (Locality, error)
	GenerateReportAll() ([]LocalityReport, error)
	GenerateReportById(id int) (LocalityReport, error)
	GetAll() ([]Locality, error)
}

type repositoryLocality struct {
	db *sql.DB
}

func NewRepositoryLocality(db *sql.DB) *repositoryLocality {
	return &repositoryLocality{
		db: db,
	}
}

func (r *repositoryLocality) Create(id int, localityName, provinceName, countryName string) (Locality, error) {
	stmt, err := r.db.Prepare(`INSERT INTO localities 
	(id,
	locality_name,
	province_name,
	country_name) 
   	VALUES(?,?,?,?)`)

	if err != nil {
		return Locality{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(
		id,
		localityName,
		provinceName,
		countryName,
	)
	if err != nil {
		return Locality{}, fmt.Errorf("Locality already registered")
	}

	log.Println(rows.RowsAffected())
	// lastID, err := rows.LastInsertId()
	// if err != nil {
	// 	return Locality{}, err
	// }
	newLocality := Locality{id, localityName, provinceName, countryName}
	return newLocality, nil
}

func (r *repositoryLocality) GenerateReportAll() ([]LocalityReport, error) {
	var localityList []LocalityReport
	rows, err := r.db.Query(`SELECT localities.id, localities.locality_name, COUNT(sellers.locality_id) 
	FROM mercadofresco.localities
	INNER JOIN sellers ON localities.id = sellers.locality_id
	GROUP BY locality_id;`)
	if err != nil {
		return localityList, err
	}
	defer rows.Close()

	for rows.Next() {
		locality := LocalityReport{}

		err := rows.Scan(
			&locality.LocalityId,
			&locality.LocalityName,
			&locality.SellersCount,
		)
		if err != nil {
			return localityList, err
		}
		localityList = append(localityList, locality)
	}

	return localityList, nil
}

func (r *repositoryLocality) GenerateReportById(id int) (LocalityReport, error) {
	stmt, err := r.db.Prepare(`SELECT localities.id, localities.locality_name, COUNT(sellers.locality_id) 
	FROM mercadofresco.localities
	INNER JOIN sellers ON localities.id = sellers.locality_id
	WHERE localities.id = ?
	GROUP BY locality_id;`)
	if err != nil {
		return LocalityReport{}, err
	}

	defer stmt.Close()

	locality := LocalityReport{}

	err = stmt.QueryRow(id).Scan(
		&locality.LocalityId,
		&locality.LocalityName,
		&locality.SellersCount,
	)
	if err != nil {
		return locality, fmt.Errorf("Locality %d not found", id)
	}

	return locality, nil
}

func (r *repositoryLocality) GetAll() ([]Locality, error) {
	var localityList []Locality
	rows, err := r.db.Query("SELECT * FROM localities")
	if err != nil {
		return localityList, err
	}
	defer rows.Close()

	for rows.Next() {
		locality := Locality{}

		err := rows.Scan(
			&locality.Id,
			&locality.LocalityName,
			&locality.ProvinceName,
			&locality.CountryName,
		)
		if err != nil {
			return localityList, err
		}
		localityList = append(localityList, locality)
	}

	return localityList, nil
}
