package seller

type Service interface {
	GetAll() ([]Seller, error)
	GetId(id int) (Seller, error)
	Create(cid int, company, adress, telephone string) (Seller, error)
	CheckCid(cid int) bool
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

func (s *service) Create(cid int, company, adress, telephone string) (Seller, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Seller{}, err
	}

	lastID++

	product, err := s.repository.Create(lastID, cid, company, adress, telephone)

	if err != nil {
		return Seller{}, err
	}

	return product, nil

}

func (s *service) CheckCid(cid int) bool {
	sellers, err := s.repository.GetAll()
	if err != nil {
		return false
	}
	for _, seller := range sellers {
		if seller.Cid == cid {
			return false
		}
	}
	return true
}

func (s *service) GetAll() ([]Seller, error) {
	sellers, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return sellers, nil
}

func (s *service) GetId(id int) (Seller, error) {
	ps, err := s.repository.GetId(id)
	if err != nil {
		return Seller{}, err
	}
	return ps, nil
}
