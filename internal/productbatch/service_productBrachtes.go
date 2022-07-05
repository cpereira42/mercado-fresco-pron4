package productbatch

import "fmt"

type servicePB struct {
	repositoryPB RepositoryProductBatches
}

func NewServiceProductBatches(instance RepositoryProductBatches) ServicePB {
	return &servicePB{repositoryPB: instance}
}

func (s *servicePB) CreatePB(object ProductBatches) (ProductBatches, error) {
	if ok, _ := s.repositoryPB.GetByBatcheNumber(object.BatchNumber); ok {
		return object, fmt.Errorf("this batch_number %v is already registered", object.BatchNumber)
	}

	if err := s.repositoryPB.SearchProductById(object.ProductId); err != nil {
		return object, err
	}
	if err := s.repositoryPB.SearchSectionId(int64(object.SectionId)); err != nil {
		return object, err
	}

	obj, err := s.repositoryPB.CreatePB(object)
	if err != nil {
		return object, err
	}
	return obj, nil
}

func (s *servicePB) ReadPBSectionTodo() ([]ProductBatchesResponse, error) {
	return s.repositoryPB.ReadPBSectionTodo()
}

func (s *servicePB) ReadPBSectionId(id int64) (ProductBatchesResponse, error) {
	return s.repositoryPB.ReadPBSectionId(id)
}
