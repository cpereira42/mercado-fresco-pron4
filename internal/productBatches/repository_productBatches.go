package productBatches

import (
	"database/sql"
)

type repositoryProductBatches struct {
	conn *sql.DB
}

func NewRepositoryProductBatches(conn *sql.DB) RepositoryProductBatches {
	return &repositoryProductBatches{conn: conn}
}

func (repo *repositoryProductBatches) CreatePB(object ProductBatches) (ProductBatches, error) {
	return ProductBatches{}, nil
}

func (repo *repositoryProductBatches) ReadPB(id int64) (ProductBatchesResponse, error) {
	return ProductBatchesResponse{}, nil
}
