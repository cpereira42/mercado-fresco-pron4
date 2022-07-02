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
	BatchNumber        string  `json:"batch_number" binding:"required,numeric"`
	CurrentQuantity    int     `json:"current_quantity" binding:"required,numeric"`
	CurrentTemperature float64 `json:"current_temperature" binding:"required"`
	DueDate            string  `json:"due_date" binding:"required"`
	InitialQuantity    int     `json:"initial_quantity" binding:"required,numeric"`
	ManufacturingDate  string  `json:"manufacturing_date" binding:"required"`
	ManufacturingHour  string  `json:"manufacturing_hour" binding:"required"`
	MinimumTemperature float64 `json:"minimum_temperature" binding:"required,numeric"`
	ProductId          int     `json:"product_id" binding:"required,numeric"`
	SectionId          int     `json:"section_id" binding:"required,numeric"`
}

type ProductBatchesResponse struct {
	SectionId     int `json:"section_id" binding:"required"`
	SectionNumber int `json:"section_number" binding:"required"`
	ProductsCount int `json:"products_count" binding:"required"`
}

type RepositoryProductBatches interface {
	CreatePB(productBatches ProductBatches) (ProductBatches, error)
	ReadPBSectionTodo() ([]ProductBatchesResponse, error)
	ReadPBSectionId(id int64) (ProductBatchesResponse, error)
	GetByBatcheNumber(batch_number string) (bool, error)
}

type ServicePB interface {
	CreatePB(productBatches ProductBatches) (ProductBatches, error)
	ReadPBSectionTodo() ([]ProductBatchesResponse, error)
	ReadPBSectionId(id int64) (ProductBatchesResponse, error)
}
