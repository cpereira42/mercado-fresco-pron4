package section

import (
	"errors"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/fatih/structs"
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

func (s service) CreateSection(newSection SectionRequestCreate) (Section, error) {
	_, err := s.repositoryWarehouse.GetByID(int(newSection.WarehouseId))
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

func (s service) UpdateSection(id int64, sectionUp SectionRequestUpdate) (Section, error) {
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
	newSectionRequest := factorySectionUpdate(&s, id, sectionUp)

	return s.repository.UpdateSection(newSectionRequest)
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

func factorySectionUpdate(s *service, id int64, sectionRequest SectionRequestUpdate) Section {
	fieldSection := []string{
		"SectionNumber",
		"CurrentCapacity",
		"CurrentTemperature",
		"MaximumCapacity",
		"MinimumCapacity",
		"MinimumTemperature",
		"WarehouseId",
		"ProductTypeId",
	}
	sec, _ := s.ListarSectionOne(id)
	sectionMapCurrent := structs.Map(sec)
	sectionMapModify := structs.Map(sectionRequest)

	for _, value := range fieldSection {
		if sectionMapModify[value] != 0 && sectionMapModify[value] != sectionMapCurrent[value] {
			switch value {
			case "WarehouseId":
				fmt.Println(value)
				sectionMapCurrent[value] = sectionMapModify[value]
			case "ProductTypeId":
				fmt.Println(value)
				sectionMapCurrent[value] = sectionMapModify[value]
			default:
				sectionMapCurrent[value] = sectionMapModify[value]
			}
		}
	}
	objetoSec := Section{
		Id:                 id,
		SectionNumber:      sectionMapCurrent["SectionNumber"].(int),
		CurrentCapacity:    sectionMapCurrent["CurrentCapacity"].(int),
		CurrentTemperature: sectionMapCurrent["CurrentTemperature"].(int),
		MaximumCapacity:    sectionMapCurrent["MaximumCapacity"].(int),
		MinimumCapacity:    sectionMapCurrent["MinimumCapacity"].(int),
		MinimumTemperature: sectionMapCurrent["MinimumTemperature"].(int),
		WarehouseId:        sectionMapCurrent["WarehouseId"].(int64),
		ProductTypeId:      sectionMapCurrent["ProductTypeId"].(int64),
	}
	return objetoSec
}
