package warehouse

type Warehouse struct {
	ID                  int    `json:"id"`
	Address             string `json:"address"`
	Telephone           string `json:"telephone"`
	Warehouse_code      string `json:"warehouse_code"`
	Minimum_capacity    int    `json:"minimum_capacity"`
	Minimum_temperature int    `json:"minimum_temperature"`
	Locality_id         int    `json:"locality_id"`
}
type RequestWarehouseCreate struct {
	Address             string `json:"address" binding:"required"`
	Telephone           string `json:"telephone" binding:"required,numeric"`
	Warehouse_code      string `json:"warehouse_code" binding:"required"`
	Minimum_capacity    int    `json:"minimum_capacity" binding:"required"`
	Minimum_temperature int    `json:"minimum_temperature" binding:"required"`
	Locality_id         int    `json:"locality_id" binding:"required"`
}
type RequestWarehouseUpdate struct {
	Address             string `json:"address" binding:"omitempty,required"`
	Telephone           string `json:"telephone" binding:"omitempty,required,numeric"`
	Warehouse_code      string `json:"warehouse_code" binding:"omitempty,required"`
	Minimum_capacity    int    `json:"minimum_capacity" binding:"omitempty,required"`
	Minimum_temperature int    `json:"minimum_temperature" binding:"omitempty,required"`
	Locality_id         int    `json:"locality_id" binding:"omitempty,required"`
}
