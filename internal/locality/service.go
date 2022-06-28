package locality

type Service interface {
	Create(localityName, provinceName, countryName string) (Locality, error)
	GenerateReport() (LocalityReport, error)
}

// type service struct {
// 	repository RepositoryLocality
// }

// func NewService(r RepositoryLocality) Service {
// 	return &service{
// 		repository: r,
// 	}
// }

// func (s *service) Create(cid, company, address, telephone string, localityId int) (Seller, error) {
// 	checkCid, err := s.CheckCid(cid)
// 	if err != nil {
// 		return Seller{}, err
// 	}
// 	if !checkCid {
// 		return Seller{}, errors.New("Cid already registered")
// 	}

// 	checkLocality, err := s.repository.CheckLocality(localityId)
// 	if err != nil {
// 		return Seller{}, err
// 	}
// 	if !checkLocality {
// 		return Seller{}, errors.New("Locality not found")
// 	}

// 	seller, err := s.repository.Create(cid, company, address, telephone, localityId)

// 	if err != nil {
// 		return Seller{}, err
// 	}

// 	return seller, nil

// }
