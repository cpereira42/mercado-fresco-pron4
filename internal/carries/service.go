package carries

type Service interface {
	GetByIDReport(id int) (Localities, error)
	Create(cid, companyName, address, telephone string, localityID int) (Carries, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// func (s *service) GetByID(id int) (Carries, error) {
// 	carry, err := s.repository.GetByID(id)
// 	if err != nil {
// 		return Carries{}, err
// 	}
// 	return carry, nil
// }

func (s *service) GetByIDReport(id int) (Localities, error) {
	locality, err := s.repository.GetByIDReport(id)
	if err != nil {
		return Localities{}, err
	}
	return locality, nil
}

func (s *service) Create(cid, companyName, address, telephone string, localityID int) (Carries, error) {
	carry, err := s.repository.Create(cid, companyName, address, telephone, localityID)
	if err != nil {
		return Carries{}, err
	}
	return carry, nil
}
