package section

import (
	"errors"
	"fmt"
)
 

func (s service) ListarSectionAll() ([]Section, error){
	sectionList, err := s.repository.ListarSectionAll()
	if err != nil {
		return []Section{}, errors.New("não há sections registrado")
	}
	return sectionList, nil
}


func (s service) ListarSectionOne(id int) (Section, error) { 
	sect, err := s.repository.ListarSectionOne(id)	 
	if err != nil {
		return Section{}, err
	}
	return sect, nil
}


func (s service) CreateSection(newSection SectionRequestCreate) (Section, error) {
	var sectionList []Section 
	
	sectionList, err := s.repository.ListarSectionAll()
	if err != nil {
		return Section{}, err
	}
	for index := range sectionList {
		if sectionList[index].SectionNumber == newSection.SectionNumber {
			return Section{}, fmt.Errorf("section invalid, section_number field must be unique")
		}
	} 
	var sec Section = factorySection(newSection)

	return s.repository.CreateSection(sec)
}

func (s service) UpdateSection(id int, sectionUp SectionRequestUpdate) (Section, error) { 
	var sectionList []Section 
	sectionList, err := s.repository.ListarSectionAll()
	if err != nil {
		return Section{}, err
	}
	
	for index := range sectionList {
		if sectionList[index].Id != id && sectionList[index].SectionNumber  == sectionUp.SectionNumber {
			return Section{}, fmt.Errorf("this section_number %v is already registered", sectionUp.SectionNumber)
		}
	}
	newSectionRequest := factorySectionUpdate(sectionUp)
	return s.repository.UpdateSection(id, newSectionRequest)
}

func (s service) DeleteSection(id int) error {	
	err := s.repository.DeleteSection(id)
	if err != nil {
		return err
	}
	return nil
}

func NewService(repository Repository) Service { 
	return &service{ repository: repository }
}

func factorySection(sectionRequest SectionRequestCreate) Section {
	return Section{ 
		SectionNumber: sectionRequest.SectionNumber,CurrentTemperature: sectionRequest.CurrentTemperature,
		MinimumTemperature: sectionRequest.MinimumTemperature,CurrentCapacity: sectionRequest.CurrentCapacity, 
		MinimumCapacity: sectionRequest.MinimumCapacity, MaximumCapacity: sectionRequest.MaximumCapacity, 
		WarehouseId: sectionRequest.WarehouseId,ProductTypeId: sectionRequest.ProductTypeId,
	}
}
func factorySectionUpdate(sectionRequest SectionRequestUpdate) Section {
	return Section{ 
		SectionNumber: sectionRequest.SectionNumber,CurrentTemperature: sectionRequest.CurrentTemperature,
		MinimumTemperature: sectionRequest.MinimumTemperature,CurrentCapacity: sectionRequest.CurrentCapacity, 
		MinimumCapacity: sectionRequest.MinimumCapacity, MaximumCapacity: sectionRequest.MaximumCapacity, 
		WarehouseId: sectionRequest.WarehouseId,ProductTypeId: sectionRequest.ProductTypeId,
	}
}
