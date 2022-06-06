package section

import (
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/fatih/structs"
)

func (r repository) CreateSection(newSection Section) (Section, error) {
	var sectionsList []Section
	if err := r.db.Read(&sectionsList); err != nil {
		return Section{}, err
	}

	for index := range sectionsList {
		if sectionsList[index].SectionNumber == newSection.SectionNumber {
			return newSection, fmt.Errorf("section invalid, section_number field must be unique")
		}
	}

	lastID, _ := r.lastID()
	lastID++

	newSection.Id = lastID
	sectionsList = append(sectionsList, newSection)
	if err := r.db.Write(&sectionsList); err != nil {
		return Section{}, err
	}
	return newSection, nil
}
func (r repository) ListarSectionAll() ([]Section, error) {
	var sectionsList []Section
	if err := r.db.Read(&sectionsList); err != nil {
		return sectionsList, err
	}
	return sectionsList, nil
}
func (r repository) ListarSectionOne(id int) (Section, error) {
	var (
		sectionList []Section
		section     Section
	)
	if err := r.db.Read(&sectionList); err != nil {
		return Section{}, err
	}
	for index := range sectionList {
		if sectionList[index].Id == id {
			section = sectionList[index]
			return section, nil
		}
	}
	return Section{}, fmt.Errorf("Section is not registered")
}

func (r repository) UpdateSection(id int, sectionUp Section) (Section, error) {
	var sectionList []Section
	if err := r.db.Read(&sectionList); err != nil {
		return Section{}, err
	}

	for index := range sectionList {
		if sectionList[index].Id != id && sectionList[index].SectionNumber == sectionUp.SectionNumber {
			return Section{}, fmt.Errorf("this section %d is already registered", sectionUp.SectionNumber)
		}
	}
	var updated, sectionEncontrado = false, false
	strSection := structs.Map(sectionUp)

	field := []string{"SectionNumber", "CurrentTemperature", "MinimumTemperature", "CurrentCapacity",
		"MinimumCapacity", "MaximumCapacity", "WareHouseId", "ProductTypeId"}

	for index := range sectionList {
		strSection2 := structs.Map(sectionList[index])
		for _, value := range field {
			if strSection2["Id"] == id {
				sectionEncontrado = true
				if strSection[value] != 0 && strSection2[value] != strSection[value] {
					updated = true
					strSection2[value] = strSection[value]
				}
			}
		}
		if updated {
			sectionList[index].SectionNumber = strSection2["SectionNumber"].(int)
			sectionList[index].CurrentTemperature = strSection2["CurrentTemperature"].(int)
			sectionList[index].MinimumTemperature = strSection2["MinimumTemperature"].(int)
			sectionList[index].CurrentCapacity = strSection2["CurrentCapacity"].(int)
			sectionList[index].MinimumCapacity = strSection2["MinimumCapacity"].(int)
			sectionList[index].MaximumCapacity = strSection2["MaximumCapacity"].(int)
			sectionList[index].WareHouseId = strSection2["WareHouseId"].(int)
			sectionList[index].ProductTypeId = strSection2["ProductTypeId"].(int)
			sectionUp = sectionList[index]
			sectionUp.Id = sectionList[index].Id

			if err := r.db.Write(&sectionList); err != nil {
				return Section{}, err
			}
			return sectionUp, nil
		}
	}

	if sectionEncontrado {
		sectionUp.Id = id
		return sectionUp, nil
	}
	return Section{}, fmt.Errorf("unable to update section")
}

func (r repository) DeleteSection(id int) error {
	var sectionsList []Section
	if err := r.db.Read(&sectionsList); err != nil {
		return err
	}
	if err := iterateAboutSectionList(r, sectionsList, id); err != nil {
		return err
	}
	return nil
}

func (r repository) lastID() (int, error) {
	var (
		sectionsList  []Section
		erro          error
		totalSections int
	)
	if erro = r.db.Read(&sectionsList); erro != nil {
		return 0, erro
	}
	totalSections = len(sectionsList)
	if totalSections > 0 {
		return sectionsList[totalSections-1].Id, nil
	}
	return 0, nil
}
func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

//
// HELPERS
//
func iterateAboutSectionList(rep repository, sections []Section, id int) error {
	for index := range sections {
		if sections[index].Id == id {
			if len(sections)-1 == index {
				sections = append([]Section{}, sections[:index]...)
			} else {
				sections = append(sections[:index], sections[index+1:]...)
			}
			if err := rep.db.Write(&sections); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("section is not registered")
}
