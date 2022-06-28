package section

import (
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
 

func (repoSection *repository) ListarSectionAll() ([]Section, error) {
	var sectionList []Section = []Section{}
	rows, err := repoSection.db.Query(sqlSelect)
	if err != nil {
		return sectionList, err
	}
	defer rows.Close()

	for rows.Next() {
		var sectionObj Section
		err := rows.Scan(
			&sectionObj.Id,
			&sectionObj.SectionNumber,
			&sectionObj.CurrentCapacity,
			&sectionObj.CurrentTemperature,
			&sectionObj.MaximumCapacity,
			&sectionObj.MinimumCapacity,
			&sectionObj.MinimumTemperature,
			&sectionObj.ProductTypeId,
			&sectionObj.WarehouseId,
		)
		if err != nil {
			return []Section{}, err
		}
		sectionList = append(sectionList, sectionObj)
	}
	return sectionList, nil
}

func (repoSection *repository) ListarSectionOne(id int) (Section, error) {
	rows := repoSection.db.QueryRow(sqlSelectByID, id)

	sectionObj := Section{}
	if err := rows.Scan(&sectionObj.Id, &sectionObj.SectionNumber, &sectionObj.CurrentCapacity, &sectionObj.CurrentTemperature,
		&sectionObj.MaximumCapacity, &sectionObj.MinimumCapacity, &sectionObj.MinimumTemperature,
		&sectionObj.ProductTypeId, &sectionObj.WarehouseId); err != nil {
		return Section{}, err
	}

	return sectionObj, nil
}

func (repoSection *repository) CreateSection(newSection Section) (Section, error) {
	stmt, err := repoSection.db.Prepare(sqlCreateSection)
	if err != nil {
		return Section{}, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		newSection.SectionNumber, newSection.CurrentCapacity, newSection.CurrentTemperature, newSection.MaximumCapacity,
		newSection.MinimumCapacity, newSection.MinimumTemperature, newSection.ProductTypeId, newSection.WarehouseId)
	if err != nil {
		return Section{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return Section{}, err
	}

	newSection.Id = int(lastID)

	return newSection, nil
}

func (repoSection *repository) UpdateSection(id int, sectionUp Section) (Section, error) {
	stmt, err := repoSection.db.Prepare(sqlUpdateSection)
	if err != nil {
		return Section{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		&sectionUp.SectionNumber,
		&sectionUp.CurrentCapacity,
		&sectionUp.CurrentTemperature,
		&sectionUp.MaximumCapacity,
		&sectionUp.MinimumCapacity,
		&sectionUp.MinimumTemperature,
		&sectionUp.ProductTypeId,
		&sectionUp.WarehouseId,
		&id,
	)
	if err != nil {
		return Section{}, err
	}
	return sectionUp, nil
}

func (repoSection *repository) DeleteSection(id int) error {
	result, err := repoSection.db.Query(sqlDeleteSection, id)
	if err != nil {
		return err
	} 
	defer result.Close()

	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
