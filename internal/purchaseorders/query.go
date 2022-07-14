package purchaseorders

const (
	GET_ALL_PURCHASE   = "SELECT id, order_date, order_number, tracking_code, buyer_id, product_record_id, order_status_id FROM purchase_orders"
	GET_PURCHASE_BY_ID = "SELECT id, order_date, order_number, tracking_code, buyer_id, product_record_id, order_status_id FROM purchase_orders WHERE id=?"
	CREATE_PURCHASE    = "INSERT INTO purchase_orders(order_date, order_number, tracking_code, buyer_id, product_record_id, order_status_id) VALUES(?,?,?,?,?,?,?)"
)
