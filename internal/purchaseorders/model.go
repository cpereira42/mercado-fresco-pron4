package purchaseorders

type Purchase struct {
	ID                int    `json:"id"`
	Order_number      string `json:"order_number"`
	Order_date        string `json:"order_date"`
	Tracking_code     string `json:"tracking_code"`
	Buyer_id          int    `json:"buyer_id"`
	Product_record_id int    `json:"product_record_id"`
	Order_status_id   int    `json:"order_status_id"`
}

type RequestPurchaseCreate struct {
	Order_number      string `json:"order_number" binding:"required,numeric"`
	Order_date        string `json:"order_date" binding:"required"`
	Tracking_code     string `json:"tracking_code" binding:"required"`
	Buyer_id          int    `json:"buyer_id" binding:"required,numeric"`
	Product_record_id int    `json:"product_record_id" binding:"required,numeric"`
	Order_status_id   int    `json:"order_status_id" binding:"required,numeric"`
}
