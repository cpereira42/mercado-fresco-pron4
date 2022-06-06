package warehouse

type Warehouse struct {
	ID                  int    `json:"id"`
	Address             string `json:"adress"`
	Telephone           string `json:"telephone"`
	Warehouse_code      string `json:"warehouse_code"`
	Minimum_capacity    int    `json:"minimum_capacity"`
	Minimum_temperature int    `json:"minimum_temperature"`
}
