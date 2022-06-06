package section

import (
	"fmt"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
) 

func (r repository) CreateSection(newSection Section) (Section, error) {	
	var sectionsList 	[]Section
	if err := r.db.Read(&sectionsList); err != nil {
		return Section{}, err
	}

	for index := range sectionsList {
		if sectionsList[index].SectionNumber == newSection.SectionNumber {
			return newSection, fmt.Errorf("section inlidada, o campo section_number deve ser único")
		}
	}

	lastID, _ := r.lastID()	
	lastID ++

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
		section Section
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
 	return Section{}, fmt.Errorf("Section não esta registrado")
}
func (r repository) UpdateSection(id int, sectionUp Section) (Section, error) {
		var sectionList []Section
		if err := r.db.Read(&sectionList); err != nil {
			return Section{}, err
		}
		
		for index := range sectionList {
			if sectionList[index].Id == id { 
				sectionUp.Id = sectionList[index].Id
				sectionList[index] = sectionUp
				if err := r.db.Write(&sectionList); err != nil {
					return Section{}, err
				}
				return sectionUp, nil				
			}
		}
	return Section{}, fmt.Errorf("section não esta registrado")
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
func (r repository) ModifyParcial(id int, section *ModifyParcial) (*ModifyParcial, error) {
	var sections []Section
	if err := r.db.Read(&sections); err != nil {
	return &ModifyParcial{}, err
	} 
	upSection, err := iterateAboutSectionListModify(r, sections, section, id)	
	if err != nil {
		return upSection, err
	}	
	return upSection, nil
} 
func (r repository) lastID() (int, error) {
	var (
		sectionsList 	[]Section
		erro 			error
		totalSections 	int
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
	return &repository{ db: db}
}

// 
// HELPERS
//
func iterateAboutSectionList(rep repository, sections []Section, id int) (error) {
	for index := range sections {
		if sections[index].Id == id {
			if len(sections)-1 == index {
				sections = append([]Section{}, sections[:index]... )
			} else {
				sections = append(sections[:index], sections[index+1:]... )
			}
			if err := rep.db.Write(&sections); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("section não esta registrado")
}
func iterateAboutSectionListModify(rep repository, sections []Section, objeto *ModifyParcial, 
	objetoId int) (*ModifyParcial, error) {
	for index := range sections {
		if sections[index].Id == objetoId { 
			sections[index].SectionNumber		= objeto.SectionNumber
			sections[index].CurrentTemperature	= objeto.CurrentTemperature
			sections[index].WareHouseId			= objeto.WareHouseId
			sections[index].CurrentCapacity		= objeto.CurrentCapacity
			sections[index].ProductTypeId		= objeto.ProductTypeId
			if err := rep.db.Write(&sections); err != nil {
				return &ModifyParcial{}, err
			}
 			return objeto, nil
		}
	}
	return &ModifyParcial{}, fmt.Errorf("produto não esta registrado")	
}