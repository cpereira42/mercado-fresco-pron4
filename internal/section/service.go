package section

import ( 
	"fmt"
)
 

func (s service) ListarSectionAll() ([]Section, error){ 
	return s.repository.ListarSectionAll()
}


func (s service) ListarSectionOne(id int) (Section, error) { 
	return s.repository.ListarSectionOne(id)	 
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
	return s.repository.DeleteSection(id)
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
