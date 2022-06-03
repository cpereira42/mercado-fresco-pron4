package section
 
type Service interface {
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int) (Section, error)
	CreateSection(sectionNumber int, currentTemperature int, minimumTemperature int, currentCapacity int, 
		minimumCapacity int, maximumCapacity int, warehouseId int, productTypeId int) (Section, error)
	UpdateSection(id int, sectionNumber int, currentTemperature int, minimumTemperature int, currentCapacity int,
		minimumCapacity int, maximumCapacity int, warehouseId int, productTypeId int) (Section, error)
	DeleteSection(id int) error

}

type service struct {
	repository Repository
}


func (s service) ListarSectionAll() ([]Section, error){ 
	return s.repository.ListarSectionAll()
}


func (s service) ListarSectionOne(id int) (Section, error) {
	return s.repository.ListarSectionOne(id)	 
}


func (s service) CreateSection(sectionNumber int, currentTemperature int, minimumTemperature int, currentCapacity int, 
	minimumCapacity int, maximumCapacity int, warehouseId int, productTypeId int) (Section, error) {
		
	newSection, err := s.repository.CreateSection(sectionNumber, currentTemperature, minimumTemperature, currentCapacity, 
			minimumCapacity, maximumCapacity, warehouseId, productTypeId)
	if err != nil {
		return Section{}, err
	}		
	return newSection, nil
}


func (s service) UpdateSection(id int, sectionNumber int, currentTemperature int, minimumTemperature int, currentCapacity int,
		minimumCapacity int, maximumCapacity int, warehouseId int, productTypeId int) (Section, error) {
		
		upSection,err := s.repository.UpdateSection(id, sectionNumber, currentTemperature, minimumTemperature, 
			currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId)
		if err != nil {
			return Section{}, err
		}
		return upSection, nil
}
func (s service) DeleteSection(id int) error {
	if err := s.repository.DeleteSection(id); err != nil {
		return err
	}
	return nil
}

func NewService() Service {
	return &service{}
}



















