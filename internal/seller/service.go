package seller

type Service interface {
	GetAll() ([]Seller, error)
	GetId(id int) (Seller, error)
	Create(cid, company, address, telephone string, localityId int) (Seller, error)
	// CheckCid(cid string) (bool, error)
	Update(id int, cid, company, address, telephone string, localityId int) (Seller, error)
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

func (s *service) Create(cid, company, address, telephone string, localityId int) (Seller, error) {
	seller, err := s.repository.Create(cid, company, address, telephone, localityId)

	if err != nil {
		return Seller{}, err
	}

	return seller, nil

}

// func (s service) CheckCid(cid string) (bool, error) {
// 	sellers, err := s.repository.GetAll()
// 	if err != nil {
// 		return false, err
// 	}
// 	for _, seller := range sellers {
// 		if seller.Cid == cid {
// 			return false, nil
// 		}
// 	}
// 	return true, nil
// }

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

func (s *service) Update(id int, cid, company, address, telephone string, localityId int) (Seller, error) {
	seller, err := s.repository.GetId(id)
	if err != nil {
		return Seller{}, err
	}
	sellerToUpdate := Seller{}
	if cid == "" {
		sellerToUpdate.Cid = seller.Cid
	} else {
		sellerToUpdate.Cid = cid
	}
	if company == "" {
		sellerToUpdate.CompanyName = seller.CompanyName
	} else {
		sellerToUpdate.CompanyName = company
	}
	if address == "" {
		sellerToUpdate.Address = seller.Address
	} else {
		sellerToUpdate.Address = address
	}
	if telephone == "" {
		sellerToUpdate.Telephone = seller.Telephone
	} else {
		sellerToUpdate.Telephone = telephone
	}
	if localityId == 0 {
		sellerToUpdate.LocalityId = seller.LocalityId
	} else {
		sellerToUpdate.LocalityId = localityId
	}
	updatedSeller, err := s.repository.Update(id, sellerToUpdate.Cid, sellerToUpdate.CompanyName, sellerToUpdate.Address, sellerToUpdate.Telephone, sellerToUpdate.LocalityId)
	if err != nil {
		return Seller{}, err
	}
	return updatedSeller, nil
}

func (s *service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return err
	}
	return nil
}
