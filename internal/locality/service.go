package locality

type Service interface {
	Create(localityName, provinceName, countryName string) (Locality, error)
	GenerateReportById(id int) (LocalityReport, error)
	GenerateReportAll() ([]LocalityReport, error)
}

type service struct {
	repository RepositoryLocality
}

func NewService(r RepositoryLocality) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(localityName, provinceName, countryName string) (Locality, error) {
	locality, err := s.repository.Create(localityName, provinceName, countryName)

	if err != nil {
		return Locality{}, err
	}
	return locality, nil
}

func (s *service) GenerateReportAll() ([]LocalityReport, error) {
	report, err := s.repository.GenerateReportAll()
	if err != nil {
		return []LocalityReport{}, err
	}
	return report, nil
}

func (s *service) GenerateReportById(id int) (LocalityReport, error) {
	report, err := s.repository.GenerateReportById(id)
	if err != nil {
		return LocalityReport{}, err
	}
	return report, nil
}
