package locality

import (
	"database/sql"
	"fmt"
)

type RepositoryLocality interface {
	Create(localityName, provinceName, countryName string) (Locality, error)
	GenerateReportAll() ([]LocalityReport, error)
	GenerateReportById(id int) (LocalityReport, error)
	GetAll() ([]Locality, error)
	// GetId(id int) (Locality, error)
	// Update(id int, localityName, provinceName, coutryName string) (Locality, error)
	// Delete(id int) error
	// CheckLocality(id int) (bool, error)
}

type repositoryLocality struct {
	db *sql.DB
}

func NewRepositoryLocality(db *sql.DB) *repositoryLocality {
	return &repositoryLocality{
		db: db,
	}
}

func (r *repositoryLocality) Create(localityName, provinceName, countryName string) (Locality, error) {
	stmt, err := r.db.Prepare(`INSERT INTO localities 
	(locality_name,
	province_name,
	country_name) 
   	VALUES(?,?,?)`)

	if err != nil {
		return Locality{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(
		localityName,
		provinceName,
		countryName,
	)
	if err != nil {
		return Locality{}, err
	}
	lastID, err := rows.LastInsertId()
	if err != nil {
		return Locality{}, err
	}
	newLocality := Locality{int(lastID), localityName, provinceName, countryName}
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

// func (r *repositoryLocality) GetId(id int) (Locality, error) {
// 	stmt, err := r.db.Prepare("SELECT * FROM localities WHERE id = ?")
// 	if err != nil {
// 		return Locality{}, err
// 	}
// 	defer stmt.Close()

// 	locality := Locality{}

// 	err = stmt.QueryRow(id).Scan(
// 		&locality.Id,
// 		&locality.LocalityName,
// 		&locality.ProvinceName,
// 		&locality.CountryName,
// 	)
// 	if err != nil {
// 		return locality, fmt.Errorf("Locality %d not found", id)
// 	}
// 	return locality, nil
// }

// func (r *repositoryLocality) Update(id int, localityName, provinceName, coutryName string) (Locality, error) {
// 	updatedLocality := Locality{id, localityName, provinceName, coutryName}
// 	stmt, err := r.db.Prepare(`UPDATE localities SET
// 	 	locality_name=?,
// 	  	province_name=?,
// 		country_name=? WHERE id=?`)
// 	if err != nil {
// 		return Locality{}, err
// 	}

// 	defer stmt.Close()

// 	rows, err := stmt.Exec(
// 		localityName,
// 		provinceName,
// 		coutryName,
// 		id)
// 	if err != nil {
// 		return updatedLocality, err
// 	}

// 	totLines, err := rows.RowsAffected()
// 	if err != nil {
// 		return Locality{}, err
// 	}

// 	if totLines == 0 {
// 		return updatedLocality, err
// 	}
// 	return updatedLocality, nil
// }

// func (r *repositoryLocality) Delete(id int) error {
// 	stmt, err := r.db.Prepare("DELETE FROM localities WHERE id=?")
// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	res, err := stmt.Exec(id)
// 	if err != nil {
// 		return err
// 	}
// 	RowsAffected, _ := res.RowsAffected()
// 	if RowsAffected == 0 {
// 		return fmt.Errorf("Locality %d not found", id)
// 	}
// 	return nil
// }
