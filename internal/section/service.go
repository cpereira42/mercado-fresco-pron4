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
			return Section{}, fmt.Errorf("field %s is required", value)
		}
	}
	var sectionList []Section 
	sectionList, err := s.repository.ListarSectionAll()
	if err != nil {
		return newSection, err
	}
	for index := range sectionList {
		if sectionList[index].SectionNumber == newSection.SectionNumber {
			return newSection, fmt.Errorf("section invalid, section_number field must be unique")
		}
	}
	return s.repository.CreateSection(newSection)
}

func (s service) UpdateSection(id int, sectionUp Section) (Section, error) { 
	var sectionList []Section 
	sectionList, err := s.repository.ListarSectionAll()
	if err != nil {
		return sectionUp, err
	}
	
	for index := range sectionList {
		if sectionList[index].Id != id && sectionList[index].SectionNumber  == sectionUp.SectionNumber {
			return Section{}, fmt.Errorf("this section %d is already registered", sectionUp.SectionNumber)
		}
	}

	return s.repository.UpdateSection(id, sectionUp)	
}

func (s service) DeleteSection(id int) error {
	return s.repository.DeleteSection(id)
}

func NewService(repository Repository) Service {
	return &service{ repository: repository }
}
