package employee

const (
	GET_ALL_EMPLOYEES  = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees"
	CREATE_EMPLOYEE    = "INSERT INTO employees (card_number_id, first_name,last_name,warehouse_id) VALUES(?,?,?,?)"
	UPDATE_EMPLOYEE    = "UPDATE employees SET card_number_id=?, first_name=?,	last_name=?, warehouse_id=? WHERE id=?"
	GET_EMPLOYEE_BY_ID = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id=?"
	DELETE_EMPLOYEE    = "DELETE FROM employees WHERE id=?"
)
