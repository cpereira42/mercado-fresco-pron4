package sectionproducttype

import "github.com/cpereira42/mercado-fresco-pron4/internal/section"

type servicePB struct {
	repositoryPB section.RepositoryProductBatches
}

func NewServiceProductBatches(instance section.RepositoryProductBatches) section.RepositoryProductBatches {
	return &servicePB{repositoryPB: instance}
}

func (s *servicePB) CreatePB(object section.ProductBatches) (section.ProductBatches, error) {
	return s.repositoryPB.CreatePB(object)
}

func (s *servicePB) ReadPB(id int64) (section.ProductBatchesResponse, error) {
	return s.repositoryPB.ReadPB(id)
}
