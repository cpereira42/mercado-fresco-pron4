package seller

type Service interface {
	GetAll() ([]Seller, error)
	GetId(id int) (Seller, error)
	//Store(name, tipo string, count int, price float64) (Seller, error)
	//Update(id int, name, tipo string, count int, price float64) (Seller, error)
	//UpdateName(id int, name string) (Seller, error)
	//Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Seller, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetId(id int) (Seller, error) {
	ps, err := s.repository.GetId(id)
	if err != nil {
		return Seller{}, err
	}
	return ps, nil
}
