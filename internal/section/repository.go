package section

import (
	"database/sql"
	"errors"
	"fmt"
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
		return sectionList, fmt.Errorf("sections not this registered")
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
			return []Section{}, fmt.Errorf("section field invalid")
		}
		sectionList = append(sectionList, sectionObj)
	}
	return sectionList, nil
}

func (repoSection *repository) ListarSectionOne(id int64) (Section, error) {
	rows := repoSection.db.QueryRow(SqlSelectByID, id)

	sectionObj := Section{}
	if err := rows.Scan(&sectionObj.Id, &sectionObj.SectionNumber, &sectionObj.CurrentCapacity, &sectionObj.CurrentTemperature,
		&sectionObj.MaximumCapacity, &sectionObj.MinimumCapacity, &sectionObj.MinimumTemperature,
		&sectionObj.ProductTypeId, &sectionObj.WarehouseId); err != nil {
		return Section{}, fmt.Errorf("section %v not found", id)
	}

	return sectionObj, nil
}

func (repoSection *repository) CreateSection(newSection Section) (Section, error) {
	sectionErro := Section{}

	result, err := repoSection.db.Exec(SqlCreateSection,
		&newSection.SectionNumber,
		&newSection.CurrentCapacity,
		&newSection.CurrentTemperature,
		&newSection.MaximumCapacity,
		&newSection.MinimumCapacity,
		&newSection.MinimumTemperature,
		&newSection.ProductTypeId,
		&newSection.WarehouseId,
	)
	if err != nil {
		return sectionErro, errors.New("inserção de section falho, campos invalidos")
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return sectionErro, err
	}

	newSection.Id = lastID

	return newSection, nil
}

func (repoSection *repository) UpdateSection(sectionUp Section) (Section, error) {
	stmt, err := repoSection.db.Prepare(SqlUpdateSection)
	if err != nil {
		return sectionUp, errors.New("falha ao executar query sql")
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		&sectionUp.SectionNumber,
		&sectionUp.CurrentCapacity,
		&sectionUp.CurrentTemperature,
		&sectionUp.MaximumCapacity,
		&sectionUp.MinimumCapacity,
		&sectionUp.MinimumTemperature,
		&sectionUp.ProductTypeId,
		&sectionUp.WarehouseId,
		&sectionUp.Id,
	)
	if err != nil {
		return Section{}, errors.New("falha ao atualizar o section")
	}

	rowsAffexcted, _ := result.RowsAffected()
	if rowsAffexcted == 0 {
		return Section{}, errors.New("section não atualizada")
	}

	sectionUp.Id = 0
	return sectionUp, nil
}

func (repoSection *repository) DeleteSection(id int64) error {
	result, err := repoSection.db.Query(SqlDeleteSection, id)
	if err != nil {
		return err
	}
	defer result.Close()

	if result.Err() != nil {
		return err
	}
	return nil
}

func (repo *repository) getProductTypes(id int64) (ProductTypes, error) {
	result := repo.db.QueryRow("select id from products_types where id=?", id)
	productTypes := ProductTypes{}
	err := result.Scan(
		&productTypes.ID,
	)
	if err != nil {
		return productTypes, errors.New("product_type_id id not found")
	}
	return productTypes, nil
}

func (repo *repository) getWarehouse(id int64) (int, error) {
	result := repo.db.QueryRow("select count(*) total from products_types where id=?", id)

	var total int

	err := result.Scan(
		&total,
	)
	if err != nil {
		return total, errors.New("warehouse_id not found")
	}
	if total == 0 {
		return total, errors.New("warehouse_id not found")
	}
	return total, nil
}
