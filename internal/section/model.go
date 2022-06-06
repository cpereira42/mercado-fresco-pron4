package section

import "github.com/cpereira42/mercado-fresco-pron4/pkg/store"
 

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
	CurrentTemperature 	int `json:"current_temperature" binding:"required"`
	WareHouseId 		int `json:"warehouse_id" binding:"required"` 
	CurrentCapacity		int `json:"current_capacity" binding:"required"`
	ProductTypeId 		int `json:"product_type_id" binding:"required"` 
}

type Repository interface {
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int) (Section, error)
	CreateSection(newSection Section) (Section, error) 
	UpdateSection(id int, sectionUp Section) (Section, error)
	DeleteSection(id int) error
	ModifyParcial(id int, section *ModifyParcial) (*ModifyParcial, error)
	lastID() (int, error)
}

type repository struct {
	db store.Store
} 