package section

import (
	"errors"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
)

type service struct {
	repository          Repository
	repositoryWarehouse warehouse.Repository
}

func NewService(repository Repository, repositoryWarehouse warehouse.Repository) Service {
	return &service{repository: repository, repositoryWarehouse: repositoryWarehouse}
}

func (s service) ListarSectionAll() ([]Section, error) {
	sectionList, err := s.repository.ListarSectionAll()
	if err != nil {
		return []Section{}, err
	}
	return sectionList, nil
}

func (s service) ListarSectionOne(id int64) (Section, error) {
	sect, err := s.repository.ListarSectionOne(id)
	if err != nil {
		return Section{}, errors.New("sections not found")
	}
	return sect, nil
}

func (s service) CreateSection(newSection SectionRequestCreate) (SectionRequestCreate, error) {
	if _, err := s.repository.getProductTypes(newSection.ProductTypeId); err != nil {
		return newSection, err
	}

	if _, err := s.repositoryWarehouse.GetByID(int(newSection.WarehouseId)); err != nil {
		return newSection, err
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
	var sec Section = factorySection(newSection)

	if _, err = s.repository.CreateSection(sec); err != nil {
		return newSection, err
	}

	return newSection, nil
}

func (s service) UpdateSection(id int64, sectionUp SectionRequestUpdate) (Section, error) {
	if _, err := s.repository.getWarehouse(int64(sectionUp.WarehouseId)); err != nil && sectionUp.WarehouseId != 0 {
		return Section{}, err
	}

	if _, err := s.repository.getProductTypes(int64(sectionUp.ProductTypeId)); err != nil && sectionUp.ProductTypeId != 0 {
		return Section{}, errors.New("product_type_id invalid")
	}

	var sectionList []Section
	sectionList, err := s.ListarSectionAll()
	if err != nil {
		return Section{}, err
	}

	for index := range sectionList {
		if sectionList[index].Id != id && sectionList[index].SectionNumber == sectionUp.SectionNumber {
			return Section{}, fmt.Errorf("this section_number %v is already registered", sectionUp.SectionNumber)
		}
	}
	sec, _ := s.ListarSectionOne(id)
	
	object := factorySectionUpdate(&s, &sectionUp, &sec)

	if _, err := s.repository.UpdateSection(object); err != nil {
		return Section{}, err
	}

	return object, nil
}

func (s service) DeleteSection(id int64) error {
	if _, err := s.ListarSectionOne(id); err != nil {
		return err
	}

	err := s.repository.DeleteSection(id)
	if err != nil {
		return err
	}
	return nil
}

func factorySection(sectionRequest SectionRequestCreate) Section {
	return Section{
		SectionNumber:      sectionRequest.SectionNumber,
		CurrentCapacity:    sectionRequest.CurrentCapacity,
		CurrentTemperature: sectionRequest.CurrentTemperature,
		MaximumCapacity:    sectionRequest.MaximumCapacity,
		MinimumCapacity:    sectionRequest.MinimumCapacity,
		MinimumTemperature: sectionRequest.MinimumTemperature,
		WarehouseId:        sectionRequest.WarehouseId,
		ProductTypeId:      sectionRequest.ProductTypeId,
	}
}

func factorySectionUpdate(s *service, sectionUp *SectionRequestUpdate, sec *Section) Section {
	if sec.SectionNumber != sectionUp.SectionNumber && sectionUp.SectionNumber != 0 {
		sec.SectionNumber = sectionUp.SectionNumber
	}
	if sec.CurrentCapacity != sectionUp.CurrentCapacity && sectionUp.CurrentCapacity != 0 {
		sec.CurrentCapacity = sectionUp.CurrentCapacity
	}
	if sec.CurrentTemperature != sectionUp.CurrentTemperature && sectionUp.CurrentTemperature != 0 {
		sec.CurrentTemperature = sectionUp.CurrentTemperature
	}
	if sec.MaximumCapacity != sectionUp.MaximumCapacity && sectionUp.MaximumCapacity != 0 {
		sec.MaximumCapacity = sectionUp.MaximumCapacity
	}
	if sec.MinimumCapacity != sectionUp.MinimumCapacity && sectionUp.MinimumCapacity != 0 {
		sec.MinimumCapacity = sectionUp.MinimumCapacity
	}
	if sec.MinimumTemperature != sectionUp.MinimumTemperature && sectionUp.MinimumTemperature != 0 {
		sec.MinimumTemperature = sectionUp.MinimumTemperature
	}
	if sec.WarehouseId != sectionUp.WarehouseId && sectionUp.WarehouseId != 0 {
		sec.WarehouseId = sectionUp.WarehouseId
	}
	if sec.ProductTypeId != sectionUp.ProductTypeId && sectionUp.ProductTypeId != 0 {
		sec.ProductTypeId = sectionUp.ProductTypeId
	}
	return *sec
}
