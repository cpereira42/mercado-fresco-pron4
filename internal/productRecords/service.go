package productsRecords

import (
	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
)

type Service interface {
	GetId(id int) (ProductRecords, error)
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

func (s *service) GetId(id int) (ProductRecords, error) {
	ps, err := s.repositoryRecords.GetId(id)
	if err != nil {
		return ProductRecords{}, err
	}
	return ps, nil
}

/*func (s *service) CheckCode(id int, code string) bool {

	ps, _ := s.rep.GetAll()
	for i := range ps {
		if ps[i].ProductCode == code && ps[i].Id != id {
			return true
		}
	}
	return false
}*/

func (s *service) Create(p RequestProductRecordsCreate) (ProductRecords, error) {
	var prod ProductRecords

	_, err := s.repositoryProducts.GetId(p.productTypeId)
	if err != nil {
		return ProductRecords{}, err
	}

	/*if s.CheckCode(0, p.ProductCode) {
		return Product{}, fmt.Errorf("code Product %s already registred", p.ProductCode)
	}*/

	prod.LastUpdateDate = p.LastUpdateDate
	prod.PurchasePrice = p.PurchasePrice
	prod.SalePrice = p.SalePrice
	prod.productTypeId = p.productTypeId
	product, err := s.repositoryRecords.Create(prod)
	if err != nil {
		return ProductRecords{}, err
	}

	return product, nil
}
