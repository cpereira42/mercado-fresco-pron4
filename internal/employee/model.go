package employee

type Employee struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
}

type Request struct {
	CardNumberID string `json:"card_number_id" validate:"nonzero" `
	FirstName    string `json:"first_name" validate:"nonzero" `
	LastName     string `json:"last_name" validate:"nonzero"`
	WarehouseID  int    `json:"warehouse_id" validate:"nonzero"`
}
