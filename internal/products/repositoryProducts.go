package products

import (
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/fatih/structs"
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

// verificar o checkcode
func (r *repositoryProduct) CheckCode(code string) error {
	var ps []Product
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].Product_code == code {
			return fmt.Errorf("codigo do produto %s já cadastrado", code)
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
	list := []string{"Product_code", "Description", "Width", "Length", "Height", "NetHeight", "ExpirationRate", "RecommendedFreezingTemperature", "FreezingRate", "ProductTypeId", "SellerId"}

	m1 := structs.Map(p)
	for z := 0; z < len(ps); z++ {
		update := false
		m2 := structs.Map(ps[z])
		for i := 0; i < len(list); i++ {
			if m2["Id"] == id {
				if m2[list[i]] != m1[list[i]] && m1[list[i]] != "" && m1[list[i]] != 0.0 {
					update = true
					m2[list[i]] = m1[list[i]]
				}
			}
		}
		if update {
			ps[z].Product_code = m2["Product_code"].(string)
			ps[z].Description = m2["Description"].(string)
			ps[z].Width = m2["Width"].(float64)
			ps[z].Length = m2["Length"].(float64)
			ps[z].Height = m2["Height"].(float64)
			ps[z].NetWeight = m2["NetWeight"].(float64)
			ps[z].ExpirationRate = m2["ExpirationRate"].(string)
			ps[z].RecommendedFreezingTemperature = m2["RecommendedFreezingTemperature"].(float64)
			ps[z].FreezingRate = m2["FreezingRate"].(float64)
			ps[z].ProductType_Id = m2["ProductType_Id"].(int)
			ps[z].SellerId = m2["SellerId"].(int)
			p := ps[z]
			if err := r.db.Write(ps); err != nil {
				return Product{}, err
			}
			return p, nil
		}

	}

	/*
		for i, p := range ps {
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
		}*/
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}
