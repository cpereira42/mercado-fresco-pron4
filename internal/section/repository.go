package section

import "fmt"

var sections []Section = []Section{}

var lastId int


type Repository interface {
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int) (Section, error)
	CreateSection(sectionNumber int, currentTemperature int, minimumTemperature int, currentCapacity int, 
		minimumCapacity int, maximumCapacity int, warehouseId int, productTypeId int) (Section, error)
	UpdateSection(id int, sectionNumber int, currentTemperature int, minimumTemperature int, currentCapacity int, 
			minimumCapacity int, maximumCapacity int, warehouseId int, productTypeId int) (Section, error)
	DeleteSection(id int) error
}

type repository struct {} 

func (repository) CreateSection(sectionNumber int, currentTemperature int, minimumTemperature int, currentCapacity int, 
	minimumCapacity int, maximumCapacity int, warehouseId int, productTypeId int) (Section, error) {
		
	lastId ++
	
	sectionObject := Section{
		Id: lastId,
		SectionNumber: sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity: currentCapacity,
		MinimumCapacity: minimumCapacity,
		MaximumCapacity: maximumCapacity,
		ProductTypeId: productTypeId,
	}

	sections = append(sections, sectionObject)
	
	return sectionObject, fmt.Errorf("method not implement")
}


func (repository) ListarSectionAll() ([]Section, error) {
	return sections, nil
}


func (repository) ListarSectionOne(id int) (Section, error) {
	for index := range sections {
		if sections[index].Id == id {
			return sections[index], nil
		} else {
			continue
		}
	}
	return Section{}, fmt.Errorf("Section não esta registrado")
}



func (repository) UpdateSection(id int, sectionNumber int, currentTemperature int, minimumTemperature int, currentCapacity int, 
	minimumCapacity int, maximumCapacity int, warehouseId int, productTypeId int) (Section, error) {
		section := Section{SectionNumber: sectionNumber, CurrentTemperature: currentTemperature, MinimumTemperature: minimumTemperature,
			CurrentCapacity: currentCapacity, MinimumCapacity: minimumCapacity, MaximumCapacity: maximumCapacity, ProductTypeId: productTypeId}
		for index := range sections {
			if sections[index].Id == id {
				sections[index] = section 
				return section, nil
			}
		}
	return Section{}, fmt.Errorf("section não esta registrado")
}


func (repository) DeleteSection(id int) error {
	var sects []Section = []Section{}

	for index := range sections {
		if sections[index].Id == id {
			if len(sections)-1 == index {
				sects = append(sects, sections[:index]... )
			} else {
				sects = append(sections[:index], sections[index+1:]... )
			}
			sections = sects
			return nil
		}
	}
	return fmt.Errorf("method not implement")
}

func NewSection() Repository {
	return &repository{}
}
