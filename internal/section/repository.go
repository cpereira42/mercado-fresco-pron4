package section

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
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
			return []Section{}, err
		}
		sectionList = append(sectionList, sectionObj)
	}
	return sectionList, nil
}

func (repoSection *repository) ListarSectionOne(id int64) (Section, error) {
	rows := repoSection.db.QueryRow(sqlSelectByID, id)

	sectionObj := Section{}
	if err := rows.Scan(&sectionObj.Id, &sectionObj.SectionNumber, &sectionObj.CurrentCapacity, &sectionObj.CurrentTemperature,
		&sectionObj.MaximumCapacity, &sectionObj.MinimumCapacity, &sectionObj.MinimumTemperature,
		&sectionObj.ProductTypeId, &sectionObj.WarehouseId); err != nil {
		return Section{}, fmt.Errorf("section %v not found", id)
	}

	return sectionObj, nil
}

func (repoSection *repository) CreateSection(newSection Section) (Section, error) {
	if _, err := repoSection.getWarhouseById(repoSection.db, newSection.WarehouseId); err != nil {
		return Section{}, err
	}

	if _, err := repoSection.getProductBatchtes(repoSection.db, newSection.ProductTypeId); err != nil {
		return Section{}, err
	}

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

	newSection.Id = lastID

	return newSection, nil
}

func (repoSection *repository) UpdateSection(sectionUp Section) (Section, error) {
	if _, err := repoSection.getWarhouseById(repoSection.db, sectionUp.WarehouseId); err != nil {
		return Section{}, err
	}

	if _, err := repoSection.getProductBatchtes(repoSection.db, sectionUp.ProductTypeId); err != nil {
		return Section{}, err
	}
	_, err := repoSection.db.Exec(sqlUpdateSection,
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
		return Section{}, err
	}
	sectionUp.Id = 0
	return sectionUp, nil
}

func (repoSection *repository) DeleteSection(id int64) error {
	result, err := repoSection.db.Query(sqlDeleteSection, id)
	if err != nil {
		return errors.New("falha ao remove o sections")
	}
	defer result.Close()

	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (repo *repository) getProductBatchtes(db *sql.DB, id int64) (ProductTypes, error) {
	result := db.QueryRow("select id from products_types where id=?", id)
	productTypes := ProductTypes{}
	err := result.Scan(
		&productTypes.ID,
	)
	if err != nil {
		return productTypes, errors.New("product_type_id id not found")
	}
	return productTypes, nil
}

func (repo *repository) getWarhouseById(db *sql.DB, id int64) (warehouse.Warehouse, error) {
	result := db.QueryRow("select id from warehouse where id=?", id)
	warehouse := warehouse.Warehouse{}
	err := result.Scan(
		&warehouse.ID,
	)
	if err != nil {
		return warehouse, errors.New("warehouse_id id not found")
	}
	return warehouse, nil
}
