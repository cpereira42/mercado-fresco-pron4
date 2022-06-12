package section

import "github.com/cpereira42/mercado-fresco-pron4/pkg/store"

type Section struct {
	Id                 int `json:"id,omitempty" binding:"numeric"`
	SectionNumber      int `json:"section_number,omitempty" binding:"numeric"`
	CurrentTemperature int `json:"current_temperature,omitempty" binding:"numeric"`
	MinimumTemperature int `json:"minimum_temperature,omitempty" binding:"numeric"`
	CurrentCapacity    int `json:"current_capacity,omitempty" binding:"numeric"`
	MinimumCapacity    int `json:"minimum_capacity,omitempty" binding:"numeric"`
	MaximumCapacity    int `json:"maximum_capacity,omitempty" binding:"numeric"`
	WareHouseId        int `json:"warehouse_id,omitempty" binding:"numeric"`
	ProductTypeId      int `json:"product_type_id,omitempty" binding:"numeric"`
}
type SectionRequest struct {
	SectionNumber      int `json:"section_number" binding:"required,numeric"`
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
