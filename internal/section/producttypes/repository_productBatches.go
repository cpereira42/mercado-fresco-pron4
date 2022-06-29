package sectionproducttype

import (
	"database/sql"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
)

type repositoryProductBatches struct {
	conn *sql.DB
}

func NewRepositoryProductBatches(conn *sql.DB) section.RepositoryProductBatches {
	return &repositoryProductBatches{conn: conn}
}

func (repo *repositoryProductBatches) CreatePB(object section.ProductBatches) (section.ProductBatches, error) {
	return section.ProductBatches{}, nil
}

func (repo *repositoryProductBatches) ReadPB(id int64) (section.ProductBatchesResponse, error) {
	return section.ProductBatchesResponse{}, nil
}
