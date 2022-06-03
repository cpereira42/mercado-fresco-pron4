package seller

import (
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
)

type repositorySeller struct {
	db store.Store
}

func NewRepositorySeller(db store.Store) Repository {
	return &repositorySeller{
		db: db,
	}
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
