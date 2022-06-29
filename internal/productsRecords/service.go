package productsRecords

import (
	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
)

type Service interface {
	GetIdRecords(id int) (ReturnProductRecords, error)
	Create(p RequestProductRecordsCreate) (ProductRecords, error)
}

type service struct {
	repositoryRecords  Repository
	repositoryProducts products.Repository
}

func NewService(r Repository, repositoryProducts products.Repository) Service {
	return &service{
		repositoryRecords:  r,
		repositoryProducts: repositoryProducts,
	}
}

func (s *service) GetIdRecords(id int) (ReturnProductRecords, error) {
	ps, err := s.repositoryRecords.GetIdRecords(id)
	if err != nil {
		return ReturnProductRecords{}, err
	}
	return ps, nil
}

func (s *service) Create(p RequestProductRecordsCreate) (ProductRecords, error) {
	var prod ProductRecords

	_, err := s.repositoryProducts.GetId(p.ProductId)
	if err != nil {
		return ProductRecords{}, err
	}
	prod.LastUpdateDate = p.LastUpdateDate
	prod.PurchasePrice = p.PurchasePrice
	prod.SalePrice = p.SalePrice
	prod.ProductId = p.ProductId
	product, err := s.repositoryRecords.Create(prod)
	if err != nil {
		return ProductRecords{}, err
	}

	return product, nil
}
