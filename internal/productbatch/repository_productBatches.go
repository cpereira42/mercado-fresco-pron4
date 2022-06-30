package productbatch

import (
	"database/sql"
	"errors"

	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
)

const (
	sqlrelatorioTodo = `SELECT s.id AS 'section_id', s.section_number AS 'section_number', count(*) AS 'products_count'
		FROM mercadofresco.products_batches as pbs
		INNER JOIN mercadofresco.sections AS s ON s.id = pbs.section_id 
		INNER JOIN mercadofresco.products AS pcts ON pbs.product_id = pcts.id
		GROUP BY s.id`

	sqlrelatorioSectioId = `SELECT s.id as 'section_id', s.section_number, count(*) as 'products_count'
		from mercadofresco.products_batches as pbs
		INNER JOIN mercadofresco.sections as s on s.id = pbs.section_id 
		INNER JOIN mercadofresco.products as pcts on pbs.product_id = pcts.id
		WHERE s.id=?`
)

type repositoryProductBatches struct {
	db *sql.DB
}

func NewRepositoryProductBatches(conn *sql.DB) RepositoryProductBatches {
	return &repositoryProductBatches{db: conn}
}

func (repo *repositoryProductBatches) CreatePB(object ProductBatches) (ProductBatches, error) {
	if err := repo.SearchProductById(object.ProductId); err != nil {
		return object, err
	}
	if err := repo.SearchSectionId(int64(object.SectionId)); err != nil {
		return object, err
	}

	query := `insert into mercadofresco.products_batches (batch_number,current_quantity
		,current_temperature,due_date,initial_quantity,manufacturing_date,manufacturing_hour,minimum_temperature,product_id,section_id)
		values(?,?,?,?,?,?,?,?,?,?)`
	res, err := repo.db.Exec(query,
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
	query := `SELECT batch_number FROM mercadofresco.products_batches WHERE batch_number=?`
	rows, err := repo.db.Query(query, batch_number)
	if err != nil {
		return false, err
	}
	var objPB ProductBatches
	if err := rows.Scan(&objPB.BatchNumber); err != nil {
		return false, err
	}
	return true, nil
}
