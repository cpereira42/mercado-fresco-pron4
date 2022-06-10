package employee

type Employee struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
}

type RequestEmployeeCreate struct {
	CardNumberID string `json:"card_number_id" binding:"required,numeric"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	WarehouseID  int    `json:"warehouse_id" binding:"required"`
}

type RequestEmployeeUpdate struct {
	CardNumberID string `json:"card_number_id" binding:"omitempty,required,numeric"`
	FirstName    string `json:"first_name" binding:"omitempty,required"`
	LastName     string `json:"last_name" binding:"omitempty,required"`
	WarehouseID  int    `json:"warehouse_id" binding:"omitempty,required"`
}
