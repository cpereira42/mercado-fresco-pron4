package section

import "github.com/cpereira42/mercado-fresco-pron4/pkg/store"

type Section struct {
	Id                 int `json:"id"`
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WareHouseId        int `json:"warehouse_id"`
	ProductTypeId      int `json:"product_type_id"`
}
type SectionRequest struct {
	SectionNumber      int `json:"section_number" binding:"numeric"`
	CurrentTemperature int `json:"current_temperature" binding:"required,numeric"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required,numeric"`
	CurrentCapacity    int `json:"current_capacity" binding:"required,numeric"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required,numeric"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"required,numeric"`
	WareHouseId        int `json:"warehouse_id" binding:"required,numeric"`
	ProductTypeId      int `json:"product_type_id" binding:"required,numeric"`
}

type Repository interface {
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int) (Section, error)
	CreateSection(newSection Section) (Section, error)
	UpdateSection(id int, sectionUp Section) (Section, error)
	DeleteSection(id int) error
	lastID() (int, error)
}

type repository struct {
	db store.Store
}
