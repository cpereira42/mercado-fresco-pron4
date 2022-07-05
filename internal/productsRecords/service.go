package productsRecords

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
