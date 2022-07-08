package productbatch

// • id: Identificación unica
// • batch_number: número do lote
// • current_quantity: quantidade atual por lote
// • current_temperature: Temperatura atual.
// • due_date: data de validade do produto
// • initial_quantity: quantidade inicial
// • manufacturing_date: data de fabricação
// • manufacturing_hour: hora de fabricação
// • minimum_temperature: Temperatura mínima
// • product_id: ID do produto
// • section_id: ID da seção
type ProductBatches struct {
	Id                 int     `json:"id,omitempty"`
	BatchNumber        string  `json:"batch_number" binding:"numeric,required"`
	CurrentQuantity    int     `json:"current_quantity" binding:"numeric,required"`
	CurrentTemperature float64 `json:"current_temperature" binding:"required"`
	DueDate            string  `json:"due_date" validate:"datetime,required"`
	InitialQuantity    int     `json:"initial_quantity" binding:"numeric,required"`
	ManufacturingDate  string  `json:"manufacturing_date" validate:"datetime,required"`
	ManufacturingHour  string  `json:"manufacturing_hour" validate:"datetime,required"`
	MinimumTemperature float64 `json:"minimum_temperature" binding:"numeric,required"`
	ProductId          int     `json:"product_id" binding:"numeric,required"`
	SectionId          int     `json:"section_id" binding:"numeric,required"`
}

type ProductBatchesResponse struct {
	SectionId     int `json:"section_id" binding:"required"`
	SectionNumber int `json:"section_number" binding:"required"`
	ProductsCount int `json:"products_count" binding:"required"`
}

type RepositoryProductBatches interface {
	CreatePB(productBatches ProductBatches) (ProductBatches, error)
	GetAll() ([]ProductBatchesResponse, error)
	GetId(id int64) (ProductBatchesResponse, error)
}

type ServicePB interface {
	CreatePB(productBatches ProductBatches) (ProductBatches, error)
	GetAll() ([]ProductBatchesResponse, error)
	GetId(id int64) (ProductBatchesResponse, error)
}
