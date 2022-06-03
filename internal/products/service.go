package products

type Service interface {
	GetAll() ([]Product, error)
	GetId(id int) (Product, error)
	Delete(id int) error
	Store(p Product) (Product, error)
	Update(id int, p Product) (Product, error)
	UpdatePatch(id int, p Product) (Product, error)
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

func (s *service) Store(p Product) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	lastID++
	p.Id = lastID
	product, err := s.repository.Store(p)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Update(id int, p Product) (Product, error) {
	return s.repository.Update(id, p)
}

func (s *service) UpdatePatch(id int, p Product) (Product, error) {
	return s.repository.UpdatePatch(id, p)
}
