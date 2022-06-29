package productsRecords

type ProductRecords struct {
	Id             int     `json:"id"`
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      int     `json:"product_id"`
	//Description    string `json:"description" binding:"required"`
}

type RequestProductRecordsCreate struct {
	LastUpdateDate string  `json:"last_update_date" binding:"required"`
	PurchasePrice  float64 `json:"purchase_price" binding:"required"`
	SalePrice      float64 `json:"sale_price" binding:"required"`
	ProductId      int     `json:"product_id" binding:"required"`
}

type ReturnProductRecords struct {
	ProductId    int    `json:"product_id"`
	RecordsCount int    `json:"records_count"`
	Description  string `json:"description"`
}

type Repository interface {
	GetIdRecords(id int) (ReturnProductRecords, error)
	Create(p ProductRecords) (ProductRecords, error)
	//CheckCode(code string) error
}
