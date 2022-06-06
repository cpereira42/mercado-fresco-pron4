package products

import (
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
)

type repositoryProduct struct {
	db store.Store
}

func NewRepositoryProducts(db store.Store) Repository {
	return &repositoryProduct{
		db: db,
	}
}

func (r *repositoryProduct) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *repositoryProduct) GetId(id int) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].Id == id {
			return ps[i], nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}
