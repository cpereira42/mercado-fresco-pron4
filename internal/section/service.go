package section
 
type Service interface {
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int) (Section, error)
	CreateSection(newSection Section) (Section, error)
	UpdateSection(id int, sectionUp Section) (Section, error)
	DeleteSection(id int) error
	ModifyParcial(id int, section ModifyParcial) (ModifyParcial, error)
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


func (s service) CreateSection(newSection Section) (Section, error) {
	return s.repository.CreateSection(newSection)
}

 
func (s service) UpdateSection(id int, sectionUp Section) (Section, error) {
	return s.repository.UpdateSection(id, sectionUp)	
}


func (s service) DeleteSection(id int) error {
	return s.repository.DeleteSection(id)
}


func (s service) ModifyParcial(id int, section ModifyParcial) (ModifyParcial, error) {
	return s.repository.ModifyParcial(id, section)
}


func NewService(repository Repository) Service {
	return &service{ repository: repository }
}
