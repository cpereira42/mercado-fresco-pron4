package section

import ( 
	"fmt"

	"github.com/fatih/structs"
)
 
type Service interface {
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int) (Section, error)
	CreateSection(newSection Section) (Section, error)
	UpdateSection(id int, sectionUp Section) (Section, error)
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


func (s service) CreateSection(newSection Section) (Section, error) {
	fields := []string{"SectionNumber", "CurrentTemperature", "MinimumTemperature", "CurrentCapacity", 
		"MinimumCapacity", "MaximumCapacity", "WareHouseId", "ProductTypeId"}

	mapSection := structs.Map(newSection)
	for _, value := range fields {
		if mapSection[value] == 0 {
			return Section{}, fmt.Errorf("o campo %s é obrigatório", value)
		}
	} 

	return s.repository.CreateSection(newSection)
}

func (s service) UpdateSection(id int, sectionUp Section) (Section, error) { 
	return s.repository.UpdateSection(id, sectionUp)	
}

func (s service) DeleteSection(id int) error {
	return s.repository.DeleteSection(id)
}

func NewService(repository Repository) Service {
	return &service{ repository: repository }
}
