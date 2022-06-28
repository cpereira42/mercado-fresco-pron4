package section

import (
	"errors"
	"fmt"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
)

type service struct {
	repository Repository
	repositoryWarehouse warehouse.Repository
}

func NewService(repository Repository, repositoryWarehouse warehouse.Repository) Service {
	return &service{repository: repository, repositoryWarehouse: repositoryWarehouse}
}

func (s service) ListarSectionAll() ([]Section, error)  {
	sectionList, err := s.repository.ListarSectionAll()
	if err != nil {
		return []Section{}, err
	}
	return sectionList, nil
}

func (s service) ListarSectionOne(id int) (Section, error) {
	sect, err := s.repository.ListarSectionOne(id)
	if err != nil {
		return Section{}, errors.New("sections not found")
	}
	return sect, nil
}

func (s service) CreateSection(newSection SectionRequestCreate) (Section, error) {
	_, err := s.repositoryWarehouse.GetByID(newSection.WarehouseId)
	if err != nil {
		return Section{}, err
	}

	var sectionList []Section

	sectionList, err = s.repository.ListarSectionAll()
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
	_, err := s.ListarSectionOne(id)
	if err != nil {
		return Section{}, err
	}

	var sectionList []Section
	sectionList, err = s.repository.ListarSectionAll()
	if err != nil {
		return Section{}, err
	}

	for index := range sectionList {
		if sectionList[index].Id != id && sectionList[index].SectionNumber == sectionUp.SectionNumber {
			return Section{}, fmt.Errorf("this section_number %v is already registered", sectionUp.SectionNumber)
		}
	}
	newSectionRequest := factorySectionUpdate(sectionUp)
	return s.repository.UpdateSection(id, newSectionRequest)
}

func (s service) DeleteSection(id int) error {
	if _, err := s.repository.ListarSectionOne(id); err != nil {
		return errors.New("section not this registered")
	}
	

	err := s.repository.DeleteSection(id)
	if err != nil {
		return err
	}
	return nil
}
 
func factorySection(sectionRequest SectionRequestCreate) Section {
	return Section{
		SectionNumber: sectionRequest.SectionNumber, CurrentTemperature: sectionRequest.CurrentTemperature,
		MinimumTemperature: sectionRequest.MinimumTemperature, CurrentCapacity: sectionRequest.CurrentCapacity,
		MinimumCapacity: sectionRequest.MinimumCapacity, MaximumCapacity: sectionRequest.MaximumCapacity,
		WarehouseId: sectionRequest.WarehouseId, ProductTypeId: sectionRequest.ProductTypeId,
	}
}

func factorySectionUpdate(sectionRequest SectionRequestUpdate) Section {
	return Section{
		SectionNumber: sectionRequest.SectionNumber, CurrentTemperature: sectionRequest.CurrentTemperature,
		MinimumTemperature: sectionRequest.MinimumTemperature, CurrentCapacity: sectionRequest.CurrentCapacity,
		MinimumCapacity: sectionRequest.MinimumCapacity, MaximumCapacity: sectionRequest.MaximumCapacity,
		WarehouseId: sectionRequest.WarehouseId, ProductTypeId: sectionRequest.ProductTypeId,
	}
}
