package buyer

import (
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
)

var buyers []Buyer
var buyer Buyer

type Repository interface {
	GetAll() ([]Buyer, error)
	GetId(id int) (Buyer, error)
	Create(id int, card_number_ID, first_name, last_name string) (Buyer, error)
	LastID() (int, error)
	Update(id int, card_number_ID, first_name, last_name string) (Buyer, error)
	Delete(id int) error
}

type repositoryBuyer struct {
	db store.Store
}

func (r *repositoryBuyer) GetAll() ([]Buyer, error) {
	if err := r.db.Read(&buyers); err != nil {
		return []Buyer{}, nil
	}
	return buyers, nil
}

func (r *repositoryBuyer) LastID() (int, error) {
	var buyers []Buyer
	if err := r.db.Read(&buyers); err != nil {
		return 0, err
	}
	if len(buyers) == 0 {
		return 0, nil
	}

	return buyers[len(buyers)-1].ID, nil
}

func (r *repositoryBuyer) GetId(id int) (Buyer, error) {
	if err := r.db.Read(&buyers); err != nil {
		return Buyer{}, err
	}
	exists := false
	for i := range buyers {
		if buyers[i].ID == id {
			buyer = buyers[i]
			exists = true
		}
	}
	if !exists {
		return Buyer{}, fmt.Errorf("buyer %d not found", id)
	}
	return buyer, nil
}

func (r *repositoryBuyer) Create(id int, card_number_ID, first_name, last_name string) (Buyer, error) {
	var buyers []Buyer
	if err := r.db.Read(&buyers); err != nil {
		return Buyer{}, err
	}
	buyer = Buyer{id, card_number_ID, first_name, last_name}
	exists := false
	for i := range buyers {
		if buyers[i].Card_number_ID == card_number_ID {
			exists = true
		}
	}

	if exists {
		return Buyer{}, fmt.Errorf("a buyer with id %s, already exists", card_number_ID)
	}
	buyers = append(buyers, buyer)
	if err := r.db.Write(buyers); err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func (r *repositoryBuyer) Update(id int, card_number_ID, first_name, last_name string) (Buyer, error) {

	if err := r.db.Read(&buyers); err != nil {
		return Buyer{}, err
	}
	buyer = Buyer{Card_number_ID: card_number_ID, First_name: first_name, Last_name: last_name}
	updated := false
	for i := range buyers {
		if buyers[i].ID == id {
			buyer.ID = id
			buyers[i] = buyer
			updated = true
		}
	}

	if !updated {
		return Buyer{}, fmt.Errorf("buyer %d not found", id)
	}
	if err := r.db.Write(buyers); err != nil {
		return Buyer{}, err
	}

	return buyer, nil
}

func (r *repositoryBuyer) Delete(id int) error {
	if err := r.db.Read(&buyers); err != nil {
		return err
	}
	deleted := false
	var index int
	for i := range buyers {
		if buyers[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("buyer %d not found", id)
	}
	buyers = append(buyers[:index], buyers[index+1:]...)
	if err := r.db.Write(buyers); err != nil {
		return err
	}

	return nil
}

func NewRepository(db store.Store) Repository {
	return &repositoryBuyer{
		db: db,
	}
}
