package section

import "github.com/cpereira42/mercado-fresco-pron4/pkg/store"


/* 
 * Modelo de estrutura de entradas de requests post/patch e de db
 * Entidade: Section{}, [usado para respostas do db no repository]
 * Request POST: SectionRequestCreate{}
 * Request PATCH: SectionRequestUpdate{}
 */
type Section struct {
	Id                 int `json:"id,omitempty" binding:"numeric"`
	SectionNumber      int `json:"section_number,omitempty" binding:"numeric"`
	CurrentTemperature int `json:"current_temperature,omitempty" binding:"numeric"`
	MinimumTemperature int `json:"minimum_temperature,omitempty" binding:"numeric"`
	CurrentCapacity    int `json:"current_capacity,omitempty" binding:"numeric"`
	MinimumCapacity    int `json:"minimum_capacity,omitempty" binding:"numeric"`
	MaximumCapacity    int `json:"maximum_capacity,omitempty" binding:"numeric"`
	WarehouseId        int `json:"warehouse_id,omitempty" binding:"numeric"`
	ProductTypeId      int `json:"product_type_id,omitempty" binding:"numeric"`
} 
type SectionRequestCreate struct {
	SectionNumber      int `json:"section_number" binding:"required,numeric"`
	CurrentTemperature int `json:"current_temperature" binding:"required,numeric"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required,numeric"`
	CurrentCapacity    int `json:"current_capacity" binding:"required,numeric"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required,numeric"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"required,numeric"`
	WarehouseId        int `json:"warehouse_id" binding:"required,numeric"`
	ProductTypeId      int `json:"product_type_id" binding:"required,numeric"`
}
type SectionRequestUpdate struct {
	SectionNumber      int `json:"section_number" binding:"required,numeric,omitempty"`
	CurrentTemperature int `json:"current_temperature" binding:"numeric,omitempty"`
	MinimumTemperature int `json:"minimum_temperature" binding:"numeric,omitempty"`
	CurrentCapacity    int `json:"current_capacity" binding:"numeric,omitempty"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"numeric,omitempty"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"numeric,omitempty"`
	WarehouseId        int `json:"warehouse_id" binding:"numeric,omitempty"`
	ProductTypeId      int `json:"product_type_id" binding:"numeric,omitempty"`
}

/* 
 * Estrutura do repository.go
 * Repository interface{}
 * repository struct{}
 */

type Repository interface {
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int) (Section, error)
	CreateSection(section Section) (Section, error)
	UpdateSection(id int, section Section) (Section, error)
	DeleteSection(id int) error
	lastID() (int, error)
}

type repository struct {
	db store.Store
}


/*
 * estrutura do service.go
 *	Service interface{}
 * service struct{}
 */

type Service interface {
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int) (Section, error)
	CreateSection(section SectionRequestCreate) (Section, error)
	UpdateSection(id int, sectionUp SectionRequestUpdate) (Section, error)
	DeleteSection(id int) error
}

type service struct {
	repository Repository
}
