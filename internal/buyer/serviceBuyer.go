package buyer

//import "github.com/gin-gonic/gin"

type Service interface {
	GetAll() ([]Buyer, error)
	GetId(id int) (Buyer, error)
	Create(card_number_ID, first_name, last_name string) (Buyer, error)
	Update(id int, card_number_ID, first_name, last_name string) (Buyer, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Buyer, error) {
	buyers, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return buyers, nil
}

func (s *service) GetId(id int) (Buyer, error) {
	buyer, err := s.repository.GetId(id)
	if err != nil {
		return Buyer{}, err
	}
	return buyer, nil
}

func (s *service) Create(card_number_ID, first_name, last_name string) (Buyer, error) {
	buyer, err := s.repository.Create(card_number_ID, first_name, last_name)
	if err != nil {
		return Buyer{}, err
	}

	return buyer, nil
}

func (s *service) Update(id int, card_number_ID, first_name, last_name string) (Buyer, error) {
	buyer, err := s.repository.Update(id, card_number_ID, first_name, last_name)
	if err != nil {
		return Buyer{}, err
	}

	return buyer, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
