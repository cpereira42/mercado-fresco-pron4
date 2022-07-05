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
	rows, err := repoSection.db.Query(SqlSelect)
	if err != nil {
		return []Section{}, ErrorFalhaInListAll
	}
	defer rows.Close()

	for rows.Next() {
		var sectionObj Section
		if err := rows.Scan(
			&sectionObj.Id,
			&sectionObj.SectionNumber, &sectionObj.CurrentCapacity, &sectionObj.CurrentTemperature,
			&sectionObj.MaximumCapacity, &sectionObj.MinimumCapacity, &sectionObj.MinimumTemperature,
			&sectionObj.ProductTypeId, &sectionObj.WarehouseId,
		); err != nil {
			return []Section{}, ErrorFalhaInListAll
		}
		sectionList = append(sectionList, sectionObj)
	}
	return sectionList, nil
}

func (repoSection *repository) ListarSectionOne(id int64) (Section, error) {
	rows := repoSection.db.QueryRow(SqlSelectByID, id)

	sectionObj := Section{}
	if err := rows.Scan(
		&sectionObj.Id,
		&sectionObj.SectionNumber, &sectionObj.CurrentCapacity, &sectionObj.CurrentTemperature,
		&sectionObj.MaximumCapacity, &sectionObj.MinimumCapacity, &sectionObj.MinimumTemperature,
		&sectionObj.ProductTypeId, &sectionObj.WarehouseId,
	); err != nil {
		return Section{}, ErrorNotFound
	}
	return sectionObj, nil
}

func (repoSection *repository) CreateSection(newSection Section) (Section, error) {
	sectionErro := Section{}
	result, err := repoSection.db.Exec(SqlCreateSection,
		&newSection.SectionNumber, &newSection.CurrentCapacity, &newSection.CurrentTemperature,
		&newSection.MaximumCapacity, &newSection.MinimumCapacity, &newSection.MinimumTemperature,
		&newSection.ProductTypeId, &newSection.WarehouseId)
	if err != nil {
		return sectionErro, checkError(err)
	}
	rowsAffexcted, err := result.RowsAffected()
	if rowsAffexcted == 0 {
		return sectionErro, checkError(err)
	}
	return newSection, nil
}

func (repoSection *repository) UpdateSection(sectionUp Section) (Section, error) {
	result, err := repoSection.db.Exec(SqlUpdateSection,
		&sectionUp.SectionNumber,
		&sectionUp.CurrentCapacity, &sectionUp.CurrentTemperature,
		&sectionUp.MaximumCapacity, &sectionUp.MinimumCapacity, &sectionUp.MinimumTemperature,
		&sectionUp.ProductTypeId, &sectionUp.WarehouseId, &sectionUp.Id,
	)
	if err != nil {
		return Section{}, checkError(err)
	}
	rowsAffexcted, err := result.RowsAffected()
	if err != nil {
		return Section{}, ErrorNotFound
	}
	if rowsAffexcted == 0 {
		return Section{}, ErrorNotModify
	}
	return sectionUp, nil
}

func (repoSection *repository) DeleteSection(id int64) error {
	result, err := repoSection.db.Exec(SqlDeleteSection, id)
	if err != nil {
		return ErrorKeyTableSectionId
	}
	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return ErrorNotFound
	}
	return nil
}
