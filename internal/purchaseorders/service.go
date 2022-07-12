package purchaseorders

type Service interface {
	GetById(id int) (Purchase, error)
	Create(order_number, order_date, tracking_code string, buyer_id, product_record_id, order_status_id int) (Purchase, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetById(id int) (Purchase, error) {
	purchase, err := s.repository.GetById(id)
	if err != nil {
		return Purchase{}, err
	}
	return purchase, nil
}

func (s *service) Create(order_number,
	order_date,
	tracking_code string,
	buyer_id,
	product_record_id,
	order_status_id int) (Purchase, error) {
	purchase, err := s.repository.Create(
		order_number,
		order_date,
		tracking_code,
		buyer_id,
		product_record_id,
		order_status_id,
	)
	if err != nil {
		return Purchase{}, err
	}

	return purchase, nil
}
