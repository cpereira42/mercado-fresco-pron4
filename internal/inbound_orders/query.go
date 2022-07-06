package inbound_orders

const (
	GET_ALL_REPORT_INBOUND_ORDERS = `SELECT e.id, e.card_number_id, e.first_name, e.last_name, w.id AS warehouse_id, COUNT(ib.id) AS inbound_orders_count
	FROM inbound_orders AS ib
	INNER JOIN employees AS e ON e.id = ib.employee_id
	INNER JOIN warehouse AS w ON w.id = ib.warehouse_id
	INNER JOIN products_batches AS pb ON pb.id = ib.product_batch_id
	GROUP BY e.id`
	GET_REPORT_INBOUND_ORDER_BY_ID = `SELECT e.id, e.card_number_id, e.first_name, e.last_name, w.id AS warehouse_id, COUNT(ib.id) AS inbound_orders_count
	FROM inbound_orders AS ib
	INNER JOIN employees AS e ON e.id = ib.employee_id
	INNER JOIN warehouse AS w ON w.id = ib.warehouse_id
	INNER JOIN products_batches AS pb ON pb.id = ib.product_batch_id
	WHERE e.id = ?
	GROUP BY e.id`
	CREATE_INBOUND_ORDERS = "INSERT INTO employees (card_number_id, first_name,last_name,warehouse_id) VALUES(?,?,?,?)"
)
