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
	return fmt.Errorf("produto não encontrado")
}

func (r *repositoryProduct) Store(p Product) (Product, error) {
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
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

func (r *repositoryProduct) UpdatePatch(id int, p Product) (Product, error) {
	var ps []Product
	r.db.Read(&ps)

	for i := range ps {
		if ps[i].Id == id {
			if p.Product_code != "" {
				ps[i].Product_code = p.Product_code
			}
			if p.Description != "" {
				ps[i].Description = p.Description
			}
			if p.Width != 0 {
				ps[i].Width = p.Width
			}
			if p.Height != 0 {
				ps[i].Height = p.Height
			}
			if p.Length != 0 {
				ps[i].Length = p.Length
			}
			if p.NetHeight != 0 {
				ps[i].NetHeight = p.NetHeight
			}
			if p.ExpirationRate != "" {
				ps[i].ExpirationRate = p.ExpirationRate
			}
			if p.RecommendedFreezingTemperature != 0 {
				ps[i].RecommendedFreezingTemperature = p.RecommendedFreezingTemperature
			}
			if p.FreezingRate != 0 {
				ps[i].FreezingRate = p.FreezingRate
			}
			if p.ProductType_Id != 0 {
				ps[i].ProductType_Id = p.ProductType_Id
			}
			if p.SellerId != 0 {
				ps[i].SellerId = p.SellerId
			}
			p := ps[i]
			if err := r.db.Write(ps); err != nil {
				return Product{}, err
			}
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}
