package buyer

import (
	"database/sql"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
)

type Repository interface {
	GetAll() ([]Buyer, error)
	GetId(id int) (Buyer, error)
	Create(card_number_ID, first_name, last_name string) (Buyer, error)
	Update(id int, card_number_ID, first_name, last_name string) (Buyer, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]Buyer, error) {
	var buyers []Buyer
	rows, err := r.db.Query(GET_ALL_BUYERS)
	if err != nil {
		return buyers, err
	}

	defer rows.Close()

	for rows.Next() {
		var buyer Buyer

		err := rows.Scan(
			&buyer.ID,
			&buyer.Card_number_ID,
			&buyer.First_name,
			&buyer.Last_name,
		)
		if err != nil {
			return buyers, err
		}
		buyers = append(buyers, buyer)
	}
	return buyers, nil
}

func (r *repository) GetId(id int) (Buyer, error) {
	row := r.db.QueryRow(GET_BUYER_BY_ID)

	buyer := Buyer{}

	err := row.Scan(
		&buyer.ID,
		&buyer.Card_number_ID,
		&buyer.First_name,
		&buyer.Last_name,
	)

	if err != nil {
		return Buyer{}, fmt.Errorf("Fail to get id")
	}

	return buyer, nil
}

func (r *repository) Create(card_number_ID, first_name, last_name string) (Buyer, error) {
	buyer := Buyer{Card_number_ID: card_number_ID, First_name: first_name, Last_name: last_name}

	stmt, err := r.db.Exec(CREATE_BUYER, card_number_ID, first_name, last_name)

	if err != nil {
		return Buyer{}, util.CheckError(err)
	}

	rowsAffected, _ := stmt.RowsAffected()
	if rowsAffected == 0 {
		return Buyer{}, fmt.Errorf("Fail to create")
	}

	lastId, _ := stmt.LastInsertId()

	buyer.ID = int(lastId)

	return buyer, nil
}

func (r *repository) Update(id int, card_number_ID, first_name, last_name string) (Buyer, error) {
	buyer := Buyer{id, card_number_ID, first_name, last_name}

	_, err := r.db.Exec(UPDATE_BUYER, id, card_number_ID, first_name, last_name)

	if err != nil {
		return Buyer{}, util.CheckError(err)
	}

	return buyer, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Exec(DELETE_BUYER, id)
	if err != nil {
		return util.CheckError(err)
	}

	rowsAffected, _ := stmt.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Buyer not found")
	}
	return nil
}
