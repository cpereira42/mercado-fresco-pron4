package products

import (
	"fmt"

	"github.com/fatih/structs"
)

type Service interface {
	GetAll() ([]Product, error)
	GetId(id int) (Product, error)
	Delete(id int) error
	Create(p RequestProductsCreate) (Product, error)
	Update(id int, p RequestProductsUpdate) (Product, error)
	CheckCode(id int, code string) bool
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetId(id int) (Product, error) {
	ps, err := s.repository.GetId(id)
	if err != nil {
		return Product{}, err
	}
	return ps, nil
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) CheckCode(id int, code string) bool {

	ps, _ := s.repository.GetAll()
	for i := range ps {
		if ps[i].ProductCode == code && ps[i].Id != id {
			return true
		}
	}
	return false
}

func (s *service) Create(p RequestProductsCreate) (Product, error) {

	var prod Product
	if s.CheckCode(0, p.ProductCode) {
		return Product{}, fmt.Errorf("code Product %s already registred", p.ProductCode)
	}

	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++
	prod.Id = lastID
	prod.ProductCode = p.ProductCode
	prod.Description = p.Description
	prod.Width = p.Width
	prod.Length = p.Length
	prod.Height = p.Height
	prod.NetWeight = p.NetWeight
	prod.ExpirationRate = p.ExpirationRate
	prod.RecommendedFreezingTemperature = p.RecommendedFreezingTemperature
	prod.FreezingRate = p.FreezingRate
	prod.ProductTypeId = p.ProductTypeId
	prod.SellerId = p.SellerId
	product, err := s.repository.Create(prod)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Update(id int, p RequestProductsUpdate) (Product, error) {

	list := []string{"ProductCode", "Description", "Width", "Length", "Height", "NetWeight", "ExpirationRate", "RecommendedFreezingTemperature", "FreezingRate", "ProductTypeId", "SellerId"}
	prodUp, err := s.repository.GetId(id)
	if err != nil {
		return Product{}, fmt.Errorf("Product %d not found", id)
	}

	if s.CheckCode(id, p.ProductCode) {
		return Product{}, fmt.Errorf("code Product %s already registred", p.ProductCode)
	}

	m1 := structs.Map(p)
	m2 := structs.Map(prodUp)
	for i := 0; i < len(list); i++ {
		if m2["Id"] == id {
			if m2[list[i]] != m1[list[i]] && m1[list[i]] != "" && m1[list[i]] != nil && m1[list[i]] != 0.0 {
				m2[list[i]] = m1[list[i]]
			}
		}
	}

	prodUp.ProductCode = m2["ProductCode"].(string)
	prodUp.Description = m2["Description"].(string)
	prodUp.Width = m2["Width"].(float64)
	prodUp.Length = m2["Length"].(float64)
	prodUp.Height = m2["Height"].(float64)
	prodUp.NetWeight = m2["NetWeight"].(float64)
	prodUp.ExpirationRate = m2["ExpirationRate"].(float64)
	prodUp.RecommendedFreezingTemperature = m2["RecommendedFreezingTemperature"].(float64)
	prodUp.FreezingRate = m2["FreezingRate"].(float64)
	prodUp.ProductTypeId = m2["ProductTypeId"].(int)
	prodUp.SellerId = m2["SellerId"].(int)
	s.repository.Update(id, prodUp)
	return prodUp, nil
}
