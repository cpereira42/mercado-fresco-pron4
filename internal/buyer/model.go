package buyer

type Buyer struct {
	ID             int    `json:"id"`
	Card_number_ID string `json:"card_number_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
}
type RequestBuyerCreate struct {
	Card_number_ID string `json:"card_number_id" binding:"required,numeric"`
	First_name     string `json:"first_name" binding:"required"`
	Last_name      string `json:"last_name" binding:"required"`
}
type RequestBuyerUpdate struct {
	Card_number_ID string `json:"card_number_id" binding:"omitempty,numeric"`
	First_name     string `json:"first_name" binding:"omitempty"`
	Last_name      string `json:"last_name" binding:"omitempty"`
}
