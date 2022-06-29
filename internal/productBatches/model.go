package productBatches 


type ProductBatches struct {
	BatchNumber        int    `json:"batch_number" binding:"required,numeric"`
	CurrentQuantity    int    `json:"current_quantity" binding:"required,numeric"`
	CurrentTemperature int    `json:"current_temperature" binding:"required"`
	DueDate            string `json:"due_date" binding:"require,alpha"`
	InitialQuantity    int    `json:"initial_quantity" binding:"required,numeric"`
	ManufacturingDate  string `json:"manufacturing_date" binding:"required,numeric"`
	ManufacturingHour  int    `json:"manufacturing_hour" binding:"required,numeric"`
	MinimumTemperature int    `json:"minumum_temperature" binding:"required,numeric"`
	ProductId          int    `json:"product_id" binding:"required,numeric"`
	SectionId          int    `json:"section_id" binding:"required,numeric"`
}

type ProductBatchesResponse struct {
	SectionId     int `json:"section_id" binding:"required"`
	SectionNumber int `json:"section_number" binding:"required"`
	ProductsCount int `json:"products_count" binding:"required"`
}


type RepositoryProductBatches interface {
	CreatePB(productBatches ProductBatches) (ProductBatches, error)
	ReadPB(sectionId int64) (ProductBatchesResponse, error)
}

type ServicePB interface {
	CreatePB(productBatches ProductBatches) (ProductBatches, error)
	ReadPB(sectionId int64) (ProductBatchesResponse, error)
}