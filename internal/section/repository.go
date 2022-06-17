package section

import (
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/fatih/structs"
)

func (r *repository) CreateSection(newSection Section) (Section, error) {
	var sectionsList []Section
	if err := r.db.Read(&sectionsList); err != nil {
		return Section{}, err
	} 
	lastID, _ := r.LastID()
	lastID ++

	newSection.Id = lastID
	sectionsList = append(sectionsList, newSection)		
	if err := r.db.Write(sectionsList); err != nil {
		return Section{}, err
	}
	return omitFieldId(newSection), nil
}
func (r *repository) ListarSectionAll() ([]Section, error) {
	var sectionsList []Section
	if err := r.db.Read(&sectionsList); err != nil {
		return sectionsList, err
	}
	return sectionsList, nil
}
func (r *repository) ListarSectionOne(id int) (Section, error) {
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

func (r *repository) UpdateSection(id int, sectionUp Section) (Section, error) {
	var (
		sectionList []Section
		section Section = sectionUp
	)
	if err := r.db.Read(&sectionList); err != nil {
		return Section{}, err
	} 
	var updated, sectionEncontrado = false, false
	strSection := structs.Map(sectionUp)

	field := []string{"SectionNumber", "CurrentTemperature", "MinimumTemperature", "CurrentCapacity",
		"MinimumCapacity", "MaximumCapacity", "WarehouseId", "ProductTypeId"}

		for index := range sectionList {
		strSection2 := structs.Map(sectionList[index])
		for _, value := range field {
			if strSection2["Id"] == id {
				sectionEncontrado = true
				section = sectionList[index]
				if strSection[value] != 0 && strSection2[value] != strSection[value] {
					updated = true
					strSection2[value] = strSection[value]
				}
			}
		}
		if updated {
			sectionUp.SectionNumber = strSection2["SectionNumber"].(int)
			sectionUp.CurrentTemperature = strSection2["CurrentTemperature"].(int)
			sectionUp.MinimumTemperature = strSection2["MinimumTemperature"].(int)
			sectionUp.CurrentCapacity = strSection2["CurrentCapacity"].(int)
			sectionUp.MinimumCapacity = strSection2["MinimumCapacity"].(int)
			sectionUp.MaximumCapacity = strSection2["MaximumCapacity"].(int)
			sectionUp.WarehouseId = strSection2["WarehouseId"].(int)
			sectionUp.ProductTypeId = strSection2["ProductTypeId"].(int)
			sectionUp.Id = sectionList[index].Id
			sectionList[index] = sectionUp
			if err := r.db.Write(sectionList); err != nil {
				return Section{}, err
			}
			return omitFieldId(sectionUp), nil
		}
	}
	
	if sectionEncontrado {
		return omitFieldId(section), nil
	}
	return Section{}, fmt.Errorf("unable to update section")
}
func omitFieldId(section Section) Section {
	section.Id = 0
	return section
}

func (r *repository) DeleteSection(id int) error {
	var sectionsList []Section
	if err := r.db.Read(&sectionsList); err != nil {
		return err
	}
	if err := iterateAboutSectionList(r, sectionsList, id); err != nil {
		return err
	}
	return nil
}

func (r *repository) LastID() (int, error) {
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
func iterateAboutSectionList(rep *repository, sections []Section, id int) error {
	for index := range sections {
		if sections[index].Id == id {
			if len(sections)-1 == index {
				sections = append([]Section{}, sections[:index]...)
			} else {
				sections = append(sections[:index], sections[index+1:]...)
			}
			if err := rep.db.Write(sections); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("section is not registered")
}
