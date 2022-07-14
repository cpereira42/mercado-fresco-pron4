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
	PurchasePrice float64 `json:"purchase_price" binding:"required"`
	SalePrice     float64 `json:"sale_price" binding:"required"`
	ProductId     int     `json:"product_id" binding:"required"`
}

type ReturnProductRecords struct {
	ProductId    int    `json:"product_id"`
	RecordsCount int    `json:"records_count"`
	Description  string `json:"description"`
}

type Repository interface {
	GetIdRecords(id int) (ReturnProductRecords, error)
	GetAllRecords() ([]ReturnProductRecords, error)
	Create(p ProductRecords) (ProductRecords, error)
}

const (
	QUERY_GETALL = `SELECT p.id, p.description, COUNT(pr.product_id) AS records_count
	FROM product_records pr 
	JOIN products p ON p.id = pr.product_id
	GROUP BY p.id`

	QUERY_GETID = `SELECT pr.product_id, p.description, COUNT(pr.product_id) AS records_count
	FROM product_records pr 
	JOIN products p ON p.id = pr.product_id
	WHERE pr.product_id = ?
	GROUP BY pr.product_id`

	QUERY_INSERT = `INSERT INTO product_records (last_update_date, purchase_price, sale_price,product_id) VALUES (?, ?, ?, ?)`
)
