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
	return Product{}, fmt.Errorf("Product %d not found", id)
}

func (r *repositoryProduct) CheckCode(code string) error {
	var ps []Product
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].Product_code == code {
			return fmt.Errorf("Product Code %s already registered", code)
		}
	}
	return nil
}

func (r *repositoryProduct) Delete(id int) error {
	var index int
	var ps []Product
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].Id == id {
			index = i
			ps = append(ps[:index], ps[index+1:]...)
			if err := r.db.Write(ps); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("Product %d not found", id)
}

func (r *repositoryProduct) Create(p Product) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repositoryProduct) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].Id, nil
}

func (r *repositoryProduct) Update(id int, p Product) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			if err := r.db.Write(ps); err != nil {
				return Product{}, err
			}
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("Product %d not found", id)
}
