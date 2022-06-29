package productBatches

type servicePB struct {
	repositoryPB RepositoryProductBatches
}

func NewServiceProductBatches(instance RepositoryProductBatches) RepositoryProductBatches {
	return &servicePB{repositoryPB: instance}
}

func (s *servicePB) CreatePB(object ProductBatches) (ProductBatches, error) {
	return s.repositoryPB.CreatePB(object)
}

func (s *servicePB) ReadPB(id int64) (ProductBatchesResponse, error) {
	return s.repositoryPB.ReadPB(id)
}
