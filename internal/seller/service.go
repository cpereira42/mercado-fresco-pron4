package seller

import "errors"

type Service interface {
	GetAll() ([]Seller, error)
	GetId(id int) (Seller, error)
	Create(cid int, company, adress, telephone string) (Seller, error)
	CheckCid(cid int) (bool, error)
	Update(id, cid int, company, adress, telephone string) (Seller, error)
	Delete(id int) error
}

type service struct {
	repository RepositorySeller
}

func NewService(r RepositorySeller) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(cid int, company, adress, telephone string) (Seller, error) {
	checkCid, err1 := s.CheckCid(cid)
	if err1 != nil {
		return Seller{}, err1
	}
	if !checkCid {
		return Seller{}, errors.New("Cid already registered")
	}

	product, err3 := s.repository.Create(cid, company, adress, telephone)

	if err3 != nil {
		return Seller{}, err3
	}

	return product, nil

}

func (s service) CheckCid(cid int) (bool, error) {
	sellers, err := s.repository.GetAll()
	if err != nil {
		return false, err
	}
	for _, seller := range sellers {
		if seller.Cid == cid {
			return false, nil
		}
	}
	return true, nil
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

func (s *service) Update(id, cid int, company, adress, telephone string) (Seller, error) {
	seller, err := s.repository.Update(id, cid, company, adress, telephone)
	if err != nil {
		return Seller{}, err
	}
	return seller, err
}

func (s *service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return err
	}
	return nil
}
