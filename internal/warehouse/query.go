package warehouse

const (
	GetAll = `SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id FROM warehouse`
	//LastID          = "SELECT MAX(id) FROM warehouse"
	CreateWarehouse = `INSERT INTO warehouse (address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id) VALUES (?, ?, ?, ?, ?, ?)`
	UpdateWarehouse = `UPDATE warehouse SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=?, locality_id=? WHERE id=?`
	GetId           = `SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id FROM warehouse WHERE id=?`
	DeleteWarehouse = `DELETE FROM warehouse WHERE id=?`
)
