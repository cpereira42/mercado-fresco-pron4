package seller

import (
	"database/sql"
	"fmt"
	"strings"
)

type RepositorySeller interface {
	GetAll() ([]Seller, error)
	GetId(id int) (Seller, error)
	Create(cid, company, address, telephone string, localityId int) (Seller, error)
	Update(id int, cid, company, adress, telephone string, localityId int) (Seller, error)
	Delete(id int) error
}

type repositorySeller struct {
	db *sql.DB
}

func NewRepositorySeller(db *sql.DB) *repositorySeller {
	return &repositorySeller{
		db: db,
	}
}

func (r *repositorySeller) Create(cid, company, address, telephone string, localityId int) (Seller, error) {
	stmt, err := r.db.Prepare(`INSERT INTO sellers 
	(cid,
	company_name,
	address,
   	telephone,
	locality_id) 
   	VALUES(?,?,?,?,?)`)

	if err != nil {
		return Seller{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(
		cid,
		company,
		address,
		telephone,
		localityId,
	)
	if err != nil {
		handledError := handleSQLError(err)
		return Seller{}, handledError
	}
	lastID, err := rows.LastInsertId()
	if err != nil {
		return Seller{}, err
	}
	newSeller := Seller{int(lastID), cid, company, address, telephone, localityId}
	return newSeller, nil
}

func (r *repositorySeller) GetAll() ([]Seller, error) {
	var sellerList []Seller
	rows, err := r.db.Query("SELECT * FROM sellers")
	if err != nil {
		return sellerList, err
	}
	defer rows.Close()

	for rows.Next() {
		seller := Seller{}

		err := rows.Scan(
			&seller.Id,
			&seller.Cid,
			&seller.CompanyName,
			&seller.Address,
			&seller.Telephone,
			&seller.LocalityId,
		)
		if err != nil {
			return sellerList, err
		}
		sellerList = append(sellerList, seller)
	}

	return sellerList, nil
}

func (r *repositorySeller) GetId(id int) (Seller, error) {
	stmt, err := r.db.Prepare("SELECT * FROM sellers WHERE id = ?")
	if err != nil {
		return Seller{}, err
	}
	defer stmt.Close()

	seller := Seller{}

	err = stmt.QueryRow(id).Scan(
		&seller.Id,
		&seller.Cid,
		&seller.CompanyName,
		&seller.Address,
		&seller.Telephone,
		&seller.LocalityId,
	)
	if err != nil {
		return seller, fmt.Errorf("Seller %d not found", id)
	}
	return seller, nil
}

func (r *repositorySeller) Update(id int, cid, company, address, telephone string, localityId int) (Seller, error) {
	updatedSeller := Seller{id, cid, company, address, telephone, localityId}
	stmt, err := r.db.Prepare(`UPDATE sellers SET 
	 	cid=?,
	  	company_name=?,
		address=?,
		telephone=?,
		locality_id=? WHERE id=?`)
	if err != nil {
		return Seller{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Exec(
		cid,
		company,
		address,
		telephone,
		localityId,
		id)
	if err != nil {
		handledError := handleSQLError(err)
		return updatedSeller, handledError
	}

	totLines, err := rows.RowsAffected()
	if err != nil {
		return Seller{}, err
	}

	if totLines == 0 {
		return updatedSeller, err
	}
	return updatedSeller, nil
}

func (r *repositorySeller) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM sellers WHERE id=?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return fmt.Errorf("Seller %d not found", id)
	}
	return nil
}

func handleSQLError(sqlError error) error {
	switch {
	case strings.Contains(sqlError.Error(), "Cannot add or update a child row"):
		return fmt.Errorf("Locality id not found")
	case strings.Contains(sqlError.Error(), "Duplicate entry"):
		return fmt.Errorf("Cid already registered")
	}
	return nil
}
