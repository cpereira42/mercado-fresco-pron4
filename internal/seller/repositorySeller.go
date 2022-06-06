package seller

import (
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
)

type repositorySeller struct {
	db store.Store
}

var ps []Seller

func NewRepositorySeller(db store.Store) *repositorySeller {
	return &repositorySeller{
		db: db,
	}
}

func (r *repositorySeller) LastID() (int, error) {
	var ps []Seller
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}

func (r *repositorySeller) Create(id, cid int, company, adress, telephone string) (Seller, error) {
	var ps []Seller
	if err := r.db.Read(&ps); err != nil {
		return Seller{}, err
	}
	p := Seller{id, cid, company, adress, telephone}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Seller{}, err
	}
	return p, nil
}

func (r *repositorySeller) GetAll() ([]Seller, error) {
	var ps []Seller
	r.db.Read(&ps)
	return ps, nil
}

func (r *repositorySeller) GetId(id int) (Seller, error) {
	var ps []Seller
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].Id == id {
			return ps[i], nil
		}
	}
	return Seller{}, fmt.Errorf("vendedor %d não encontrado", id)
}

func (r *repositorySeller) Update(id, cid int, company, adress, telephone string) (Seller, error) {
	var ps []Seller
	r.db.Read(&ps)
	seller := Seller{id, cid, company, adress, telephone}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			seller.Id = id
			ps[i] = seller
			updated = true
		}
	}
	if !updated {
		return Seller{}, fmt.Errorf("produto %d não encontrado", id)
	}
	if err := r.db.Write(ps); err != nil {
		return Seller{}, err
	}
	return seller, nil
}

func (r *repositorySeller) Delete(id int) error {
	var ps []Seller
	r.db.Read(&ps)
	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Vendedor %d nao encontrado", id)
	}

	ps = append(ps[:index], ps[index+1:]...)
	if err := r.db.Write(ps); err != nil {
		return err
	}
	return nil
}
