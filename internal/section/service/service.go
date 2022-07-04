package sectionService

import (
	"errors"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section/entites"
	// "github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	// main "github.com/cpereira42/mercado-fresco-pron4/cmd/server"
)

type service struct {
	repository entites.Repository
}

func (s service) ListarSectionAll() ([]entites.Section, error) {
	sectionList, err := s.repository.ListarSectionAll()
	if err != nil {
		return []entites.Section{}, errors.New("não há sections registrado")
	}
	return sectionList, nil
}

func (s service) ListarSectionOne(id int) (entites.Section, error) {
	sect, err := s.repository.ListarSectionOne(id)
	if err != nil {
		return entites.Section{}, errors.New("sections not found")
	}
	return sect, nil
}

func (s service) CreateSection(newSection entites.SectionRequestCreate) (entites.Section, error) {
	//dbWarehouse := store.New("file", "./internal/repositories/warehouse.json")
	//repoWarehouse := warehouse.NewRepository(dbWarehouse)
	//serviceWarehouse := warehouse.NewService(repoWarehouse)

	//_, err := serviceWarehouse.GetByID(newSection.WarehouseId)
	//if err != nil {
	//	return entites.Section{}, err
	//}

	var sectionList []entites.Section

	sectionList, err := s.repository.ListarSectionAll()
	if err != nil {
		return entites.Section{}, err
	}
	for index := range sectionList {
		if sectionList[index].SectionNumber == newSection.SectionNumber {
			return entites.Section{}, fmt.Errorf("section invalid, section_number field must be unique")
		}
	}
	var sec entites.Section = factorySection(newSection)

	return s.repository.CreateSection(sec)
}

func (s service) UpdateSection(id int, sectionUp entites.SectionRequestUpdate) (entites.Section, error) {
	_, err := s.ListarSectionOne(id)
	if err != nil {
		return entites.Section{}, err
	}

	var sectionList []entites.Section
	sectionList, err = s.repository.ListarSectionAll()
	if err != nil {
		return entites.Section{}, err
	}

	for index := range sectionList {
		if sectionList[index].Id != id && sectionList[index].SectionNumber == sectionUp.SectionNumber {
			return entites.Section{}, fmt.Errorf("this section_number %v is already registered", sectionUp.SectionNumber)
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

func NewService(repository entites.Repository) entites.Service {
	return &service{repository: repository}
}

func factorySection(sectionRequest entites.SectionRequestCreate) entites.Section {
	return entites.Section{
		SectionNumber: sectionRequest.SectionNumber, CurrentTemperature: sectionRequest.CurrentTemperature,
		MinimumTemperature: sectionRequest.MinimumTemperature, CurrentCapacity: sectionRequest.CurrentCapacity,
		MinimumCapacity: sectionRequest.MinimumCapacity, MaximumCapacity: sectionRequest.MaximumCapacity,
		WarehouseId: sectionRequest.WarehouseId, ProductTypeId: sectionRequest.ProductTypeId,
	}
}

func factorySectionUpdate(sectionRequest entites.SectionRequestUpdate) entites.Section {
	return entites.Section{
		SectionNumber: sectionRequest.SectionNumber, CurrentTemperature: sectionRequest.CurrentTemperature,
		MinimumTemperature: sectionRequest.MinimumTemperature, CurrentCapacity: sectionRequest.CurrentCapacity,
		MinimumCapacity: sectionRequest.MinimumCapacity, MaximumCapacity: sectionRequest.MaximumCapacity,
		WarehouseId: sectionRequest.WarehouseId, ProductTypeId: sectionRequest.ProductTypeId,
	}
}
