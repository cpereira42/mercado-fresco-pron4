package buyer

type Service interface {
	GetAll() ([]Buyer, error)
	GetId(id int) (Buyer, error)
	Create(card_number_ID, first_name, last_name string) (Buyer, error)
	Update(id int, card_number_ID, first_name, last_name string) (Buyer, error)
	Delete(id int) error
}

type service struct {
	repositoryBuyer Repository
}

func NewService(r Repository) Service {
	return &service{
		repositoryBuyer: r,
	}
}

func (s service) GetAll() ([]Buyer, error) {
	buyers, err := s.repositoryBuyer.GetAll()
	if err != nil {
		return nil, err
	}
	return buyers, nil
}

func (s service) GetId(id int) (Buyer, error) {
	buyer, err := s.repositoryBuyer.GetId(id)
	if err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func (s service) Create(card_number_ID, first_name, last_name string) (Buyer, error) {
	lastId, err := s.repositoryBuyer.LastID()
	if err != nil {
		return Buyer{}, err
	}
	lastId++
	buyer, err := s.repositoryBuyer.Create(lastId, card_number_ID, first_name, last_name)

	if err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func (s service) Update(id int, card_number_ID, first_name, last_name string) (Buyer, error) {
	buyer, err := s.repositoryBuyer.Update(id, card_number_ID, first_name, last_name)
	if err != nil {
		return Buyer{}, err
	}

	return buyer, nil
}

func (s service) Delete(id int) error {
	err := s.repositoryBuyer.Delete(id)
	if err != nil {
		return err
	}
	return err
}