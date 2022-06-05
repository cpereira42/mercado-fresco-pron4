package section
 

type Section struct {
	Id 					int `json:"id"`
	SectionNumber 		int `json:"section_number" binding:"required"`
	CurrentTemperature 	int `json:"current_temperature" binding:"required"`
	MinimumTemperature 	int `json:"minimum_temperature" binding:"required"`
	CurrentCapacity 	int `json:"current_capacity" binding:"required"` 
	MinimumCapacity 	int `json:"minimum_capacity" binding:"required"` 
	MaximumCapacity 	int `json:"maximum_capacity" binding:"required"` 
	WareHouseId 		int `json:"warehouse_id" binding:"required"` 
	ProductTypeId 		int `json:"product_type_id" binding:"required"`
}
type ModifyParcial struct {
	SectionNumber 		int `json:"section_number" binding:"required"`
	WareHouseId 		int `json:"warehouse_id" binding:"required"` 
	ProductTypeId 		int `json:"product_type_id" binding:"required"`
	MinimumCapacity 	int `json:"minimum_capacity" binding:"required"` 
	MaximumCapacity 	int `json:"maximum_capacity" binding:"required"` 
}
