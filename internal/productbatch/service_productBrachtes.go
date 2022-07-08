package productbatch

type servicePB struct {
	repositoryPB RepositoryProductBatches
}

func NewServiceProductBatches(instance RepositoryProductBatches) ServicePB {
	return &servicePB{repositoryPB: instance}
}

func (s *servicePB) CreatePB(object ProductBatches) (ProductBatches, error) {
	obj, err := s.repositoryPB.CreatePB(object)
	if err != nil {
		return object, err
	}
	return obj, nil
}

func (s *servicePB) GetAll() ([]ProductBatchesResponse, error) {
	productBatchesResponseList, err := s.repositoryPB.GetAll()
	if err != nil {
		return []ProductBatchesResponse{}, err
	}
	return productBatchesResponseList, nil
}

func (s *servicePB) GetId(id int64) (ProductBatchesResponse, error) {
	productBatchesResponse, err := s.repositoryPB.GetId(id)
	if err != nil {
		return ProductBatchesResponse{}, err
	}
	return productBatchesResponse, nil
}
