package productbatch

import (
	"database/sql"
	"errors"
	"regexp"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
)

type repositoryProductBatches struct {
	db *sql.DB
}

func NewRepositoryProductBatches(conn *sql.DB) RepositoryProductBatches {
	return &repositoryProductBatches{db: conn}
}

func (repo *repositoryProductBatches) validateFieldsDateAndTime(object ProductBatches) (bool, error) {
	dataValidRe := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	validateFullRe := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)
	if !validateFullRe.MatchString(object.ManufacturingHour) {
		return false, errors.New("manufacturing_hour is invalid must use ':' to separate the time, ex: '00:00:00'")
	}
	if !dataValidRe.MatchString(object.ManufacturingDate) {
		return false, errors.New("manufacturing_date is invalid must use '-' to separate date ex: '000-00-00'")
	}
	if !dataValidRe.MatchString(object.DueDate) {
		return false, errors.New("due_date is invalid must use '-' to separate date ex: '000-00-00'")
	}
	return true, nil
}

func (repo *repositoryProductBatches) CreatePB(object ProductBatches) (ProductBatches, error) {
	ok, err := repo.validateFieldsDateAndTime(object)
	if !ok {
		return ProductBatches{}, err
	}

	res, err := repo.db.Exec(SqlCreatePB,
		&object.BatchNumber,
		&object.CurrentQuantity,
		&object.CurrentTemperature,
		&object.DueDate,
		&object.InitialQuantity,
		&object.ManufacturingDate,
		&object.ManufacturingHour,
		&object.MinimumTemperature,
		&object.ProductId,
		&object.SectionId,
	)
	if err != nil {
		return object, util.CheckError(err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return object, err
	}

	if rows == 0 {
		return object, errors.New("falha ao registrar um novo products_batches")
	}

	return object, nil
}

func (repo *repositoryProductBatches) GetAll() ([]ProductBatchesResponse, error) {
	// consulta sem o section_id
	productBatList := []ProductBatchesResponse{}
	rows, err := repo.db.Query(SqlrelatorioTodo)
	if err != nil {
		return productBatList, errors.New("query sql invalid")
	}
	defer rows.Close()

	for rows.Next() {
		var productBatchesRes ProductBatchesResponse

		err := rows.Scan(
			&productBatchesRes.SectionId,
			&productBatchesRes.SectionNumber,
			&productBatchesRes.ProductsCount,
		)
		if err != nil {
			return []ProductBatchesResponse{}, errors.New("failed to serialize product_batches_response fields")
		}
		productBatList = append(productBatList, productBatchesRes)
	}
	return productBatList, nil
}

func (repo *repositoryProductBatches) GetId(id int64) (ProductBatchesResponse, error) {
	messageErr := errors.New("section_id not found")
	result := repo.db.QueryRow(SqlrelatorioSectioId, id)

	productBatchesResponse := ProductBatchesResponse{}
	if err := result.Scan(
		&productBatchesResponse.SectionId,
		&productBatchesResponse.SectionNumber,
		&productBatchesResponse.ProductsCount); err != nil {
		return ProductBatchesResponse{}, messageErr
	}
	return productBatchesResponse, nil
}
