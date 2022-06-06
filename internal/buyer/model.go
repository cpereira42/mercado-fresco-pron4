package buyer

type Buyer struct {
	ID             int    `json:"id"`
	Card_number_ID string `json:"card_number_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
}
