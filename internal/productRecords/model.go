package productsRecords

type ProductRecords struct {
	Id             int    `json:"id"`
	LastUpdateDate string `json:"last_update_date"`
	PurchasePrice  int    `json:"purchase_price"`
	SalePrice      int    `json:"sale_price"`
	productTypeId  int    `json:"product_type_id"`
	//Description    string `json:"description" binding:"required"`
}

type RequestProductRecordsCreate struct {
	Id             int    `json:"id" binding:"required"`
	LastUpdateDate string `json:"last_update_date" binding:"required"`
	PurchasePrice  int    `json:"purchase_price" binding:"required"`
	SalePrice      int    `json:"sale_price" binding:"required"`
	productTypeId  int    `json:"product_type_id" binding:"required"`
}

type Repository interface {
	GetId(id int) (ProductRecords, error)
	Create(p ProductRecords) (ProductRecords, error)
	//CheckCode(code string) error
}
