package productbatch

import (
	"database/sql"
	"errors"

	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
)

type repositoryProductBatches struct {
	db *sql.DB
}

func NewRepositoryProductBatches(conn *sql.DB) RepositoryProductBatches {
	return &repositoryProductBatches{db: conn}
}

func (repo *repositoryProductBatches) CreatePB(object ProductBatches) (ProductBatches, error) {
	res, err := repo.db.Exec(sqlCreatePB,
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
		return object, err
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

func (repo *repositoryProductBatches) ReadPBSectionTodo() ([]ProductBatchesResponse, error) {
	// consulta sem o section_id
	productBatList := []ProductBatchesResponse{}
	rows, err := repo.db.Query(sqlrelatorioTodo)
	if err != nil {
		return productBatList, err
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
			return []ProductBatchesResponse{}, err
		}
		productBatList = append(productBatList, productBatchesRes)
	}
	return productBatList, nil
}

func (repo *repositoryProductBatches) ReadPBSectionId(id int64) (ProductBatchesResponse, error) {
	messageErr := errors.New("section_id not found")
	result, err := repo.db.Query(sqlrelatorioSectioId, id)
	if err != nil {
		return ProductBatchesResponse{}, messageErr
	}
	if result.Next() {
		productBatchesResponse := ProductBatchesResponse{}
		if err := result.Scan(
			&productBatchesResponse.SectionId,
			&productBatchesResponse.SectionNumber,
			&productBatchesResponse.ProductsCount); err != nil {
			return ProductBatchesResponse{}, messageErr
		}
		return productBatchesResponse, nil
	}
	return ProductBatchesResponse{}, messageErr
}

func (repo *repositoryProductBatches) SearchSectionId(id int64) error {
	repoSection := section.NewRepository(repo.db)
	_, err := repoSection.ListarSectionOne(id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repositoryProductBatches) SearchProductById(id int) error {
	dbProducts := store.New(store.FileType, "./internal/repositories/products.json")
	repoProduct := products.NewRepositoryProducts(dbProducts)

	if _, err := repoProduct.GetId(id); err != nil {
		return err
	}
	return nil
}

func (repo *repositoryProductBatches) GetByBatcheNumber(batch_number string) (bool, error) {
	rows, err := repo.db.Query(sqlBatcheNumber, batch_number)
	if err != nil {
		return false, err
	}
	var objPB ProductBatches
	if err := rows.Scan(&objPB.BatchNumber); err != nil {
		return false, err
	}
	return true, nil
}
