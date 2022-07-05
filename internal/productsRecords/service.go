package productsRecords

import (
	"time"
)

type Service interface {
	GetIdRecords(id int) (ReturnProductRecords, error)
	GetAllRecords() ([]ReturnProductRecords, error)
	Create(p RequestProductRecordsCreate) (ProductRecords, error)
}

type service struct {
	repositoryRecords Repository
}

func NewService(r Repository) Service {
	return &service{
		repositoryRecords: r,
	}
}

func (s *service) GetIdRecords(id int) (ReturnProductRecords, error) {
	ps, err := s.repositoryRecords.GetIdRecords(id)
	if err != nil {
		return ReturnProductRecords{}, err
	}
	return ps, nil
}

func (s *service) GetAllRecords() ([]ReturnProductRecords, error) {
	ps, err := s.repositoryRecords.GetAllRecords()
	if err != nil {
		return []ReturnProductRecords{}, err
	}
	return ps, nil
}

func (s *service) Create(p RequestProductRecordsCreate) (ProductRecords, error) {
	var prod ProductRecords
	currentTime := time.Now()

	theTime := time.Date(currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
		100,
		time.Local).Format("2006-01-02 15:04:05")

	prod.LastUpdateDate = theTime
	prod.PurchasePrice = p.PurchasePrice
	prod.SalePrice = p.SalePrice
	prod.ProductId = p.ProductId
	product, err := s.repositoryRecords.Create(prod)
	if err != nil {
		return ProductRecords{}, err
	}

	return product, nil
}
