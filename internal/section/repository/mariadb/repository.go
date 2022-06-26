package mariadb

import (
	"database/sql"
	"errors"
	"log"

	entites "github.com/cpereira42/mercado-fresco-pron4/internal/section/entites"
	// "github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	// "github.com/fatih/structs"
)

type repository struct {
	db *sql.DB
}

func (r *repository) CreateSection(newSection entites.Section) (entites.Section, error) {
	 
	stmt, err := r.db.Prepare(`INSERT INTO sections 
	 	(section_number,
	  	current_temperature,
		minimum_temperature,
		current_capacity,
		minimum_capacity,
		maximum_capacity,
		warehouse_id,
		product_type_id) VALUES(?,?,?,?,?,?,?,?)`)
		
	if err != nil {
		return entites.Section{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Exec(
		newSection.SectionNumber,
		newSection.CurrentTemperature,
		newSection.MinimumTemperature,
		newSection.CurrentCapacity,
		newSection.MinimumCapacity,
		newSection.MaximumCapacity,
		newSection.WarehouseId,
		newSection.ProductTypeId,
	)

	if err != nil {
		return entites.Section{}, err
	}


	return omitFieldId(newSection), nil
}

func (r *repository) ListarSectionAll() ([]entites.Section, error) {
	var sectionList []entites.Section
	// log.Println("1")

	rows, err := r.db.Query("SELECT * FROM sections")
	if err != nil {
		log.Println(err)
		return sectionList, err
	}

	defer rows.Close()

	for rows.Next() {
		var section entites.Section

		err := rows.Scan(
			&section.Id,
			&section.SectionNumber,
			&section.CurrentTemperature,
			&section.MinimumTemperature,
			&section.CurrentCapacity,
			&section.MinimumCapacity,
			&section.MaximumCapacity,
			&section.WarehouseId,
			&section.ProductTypeId,
		)
		if err != nil {
			return sectionList, err
		}
		sectionList = append(sectionList, section)
	}

	return sectionList, nil
}

func (r *repository) ListarSectionOne(id int) (entites.Section, error) {
	stmt, err := r.db.Prepare("SELECT * FROM sections WHERE id = ?")
	if err != nil {
		return entites.Section{}, err
	}
	defer stmt.Close()

	var section entites.Section

	err = stmt.QueryRow(id).Scan(
		&section.Id,
		&section.SectionNumber,
		&section.CurrentTemperature,
		&section.MinimumTemperature,
		&section.CurrentCapacity,
		&section.MinimumCapacity,
		&section.MaximumCapacity,
		&section.WarehouseId,
		&section.ProductTypeId)
	if err != nil {
		return section, err
	}
	return section, nil
}

func (r *repository) UpdateSection(id int, sectionUp entites.Section) (entites.Section, error) {
	stmt, err := r.db.Prepare(`UPDATE sections SET 
	 	section_number=?,
	  	current_temperature=?,
		minimum_temperature=?,
		current_capacity=?,
		minimum_capacity=?,
		maximum_capacity=?,
		warehouse_id=?,
		product_type_id=? WHERE id=?`)
	if err != nil {
		return entites.Section{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Exec(
		sectionUp.SectionNumber,
		sectionUp.CurrentTemperature,
		sectionUp.MinimumTemperature,
		sectionUp.CurrentCapacity,
		sectionUp.MinimumCapacity,
		sectionUp.MaximumCapacity,
		sectionUp.WarehouseId,
		sectionUp.ProductTypeId,
		id)
	if err != nil {
		return entites.Section{}, err
	}

	totLines, err := rows.RowsAffected()
	if err != nil {
		return entites.Section{}, err
	}

	if totLines == 0 {
		return entites.Section{}, errors.New("erro ao altera section")
	}

	return sectionUp, nil
}

func omitFieldId(section entites.Section) entites.Section {
	section.Id = 0
	return section
}

func (r *repository) DeleteSection(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM sections WHERE id=?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return errors.New("sections not found")
	}
	return nil
}

func (r *repository) LastID() (int, error) {
	// var (
	// 	sectionsList  []entites.Section
	// 	erro          error
	// 	totalSections int
	// )
	// if erro = r.db.Read(&sectionsList); erro != nil {
	// 	return 0, erro
	// }
	// totalSections = len(sectionsList)
	// if totalSections > 0 {
	// 	return sectionsList[totalSections-1].Id, nil
	// }
	return 0, nil
}

// arquivo
func NewRepository(db *sql.DB) entites.Repository {
	return &repository{db: db}
}

//
// HELPERS
//
// func iterateAboutSectionList(rep *repository, sections []entites.Section, id int) error {
// 	for index := range sections {
// 		if sections[index].Id == id {
// 			if len(sections)-1 == index {
// 				sections = append([]entites.Section{}, sections[:index]...)
// 			} else {
// 				sections = append(sections[:index], sections[index+1:]...)
// 			}
// 			if err := rep.db.Write(sections); err != nil {
// 				return err
// 			}
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("section is not registered")
// }
