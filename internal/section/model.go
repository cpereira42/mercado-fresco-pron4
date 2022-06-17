package section

import "github.com/cpereira42/mercado-fresco-pron4/pkg/store"


/* 
 * Modelo de estrutura de entradas de requests post/patch e de db
 * Entidade: Section{}, [usado para respostas do db no repository]
 * Request POST: SectionRequestCreate{}
 * Request PATCH: SectionRequestUpdate{}
 */
type Section struct {
	Id                 int `json:"id"`
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseId        int `json:"warehouse_id"`
	ProductTypeId      int `json:"product_type_id"`
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
	SectionNumber      int `json:"section_number" binding:"omitempty,required,numeric"`
	CurrentTemperature int `json:"current_temperature" binding:"omitempty,required,numeric"`
	MinimumTemperature int `json:"minimum_temperature" binding:"omitempty,required,numeric"`
	CurrentCapacity    int `json:"current_capacity"	binding:"omitempty,required,numeric"`
	MinimumCapacity    int `json:"minimum_capacity"	binding:"omitempty,required,numeric"`
	MaximumCapacity    int `json:"maximum_capacity"	binding:"omitempty,required,numeric"`
	WarehouseId        int `json:"warehouse_id" binding:"omitempty,required,numeric"`
	ProductTypeId      int `json:"product_type_id" binding:"omitempty,required,numeric"`
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
	LastID() (int, error)
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
