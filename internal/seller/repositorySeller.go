package seller

import (
	"database/sql"
	"fmt"
)

type RepositorySeller interface {
	GetAll() ([]Seller, error)
	GetId(id int) (Seller, error)
	Create(cid int, company, adress, telephone string) (Seller, error)
	Update(id, cid int, company, adress, telephone string) (Seller, error)
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

func (r *repositorySeller) Create(cid int, company, adress, telephone string) (Seller, error) {
	stmt, err := r.db.Prepare(`INSERT INTO sellers 
	(cid,
	company,
	adress,
   	telephone) 
   	VALUES(?,?,?,?)`)

	if err != nil {
		return Seller{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(
		cid,
		company,
		adress,
		telephone,
	)
	if err != nil {
		return Seller{}, err
	}
	lastID, err := rows.LastInsertId()
	if err != nil {
		return Seller{}, err
	}
	newSeller := Seller{int(lastID), cid, company, adress, telephone}
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
			&seller.Adress,
			&seller.Telephone,
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
		&seller.Adress,
		&seller.Telephone,
	)
	if err != nil {
		return seller, fmt.Errorf("Seller %d not found", id)
	}
	return seller, nil
}

func (r *repositorySeller) Update(id, cid int, company, adress, telephone string) (Seller, error) {
	stmt, err := r.db.Prepare(`UPDATE sellers SET 
	 	cid=?,
	  	company=?,
		adress=?,
		telephone=?,
		WHERE id=?`)
	if err != nil {
		return Seller{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Exec(
		cid,
		company,
		adress,
		telephone,
		id)
	if err != nil {
		return Seller{}, fmt.Errorf("Seller %d not found", id)
	}

	totLines, err := rows.RowsAffected()
	if err != nil {
		return Seller{}, err
	}

	if totLines == 0 {
		return Seller{}, fmt.Errorf("Seller %d not found", id)
	}
	updatedSeller := Seller{id, cid, company, adress, telephone}
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
