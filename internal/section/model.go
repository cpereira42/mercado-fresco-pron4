package section

/*
 * Modelo de estrutura de entradas de requests post/patch e de db
 * Entidade: Section{}, [usado para respostas do db no repository]
 * Request POST: SectionRequestCreate{}
 * Request PATCH: SectionRequestUpdate{}
 */
type Section struct {
	Id                 int64 `json:"id" binding:"numeric"`
	SectionNumber      int   `json:"section_number" binding:"numeric"`
	CurrentCapacity    int   `json:"current_capacity" binding:"numeric"`
	CurrentTemperature int   `json:"current_temperature" binding:"numeric"`
	MaximumCapacity    int   `json:"maximum_capacity" binding:"numeric"`
	MinimumCapacity    int   `json:"minimum_capacity" binding:"numeric"`
	MinimumTemperature int   `json:"minimum_temperature" binding:"numeric"`
	WarehouseId        int64 `json:"warehouse_id" binding:"numeric"`
	ProductTypeId      int64 `json:"product_type_id" binding:"numeric"`
}
type SectionRequestCreate struct {
	SectionNumber      int   `json:"section_number" binding:"required,numeric"`
	CurrentCapacity    int   `json:"current_capacity" binding:"required,numeric"`
	CurrentTemperature int   `json:"current_temperature" binding:"required,numeric"`
	MaximumCapacity    int   `json:"maximum_capacity" binding:"required,numeric"`
	MinimumCapacity    int   `json:"minimum_capacity" binding:"required,numeric"`
	MinimumTemperature int   `json:"minimum_temperature" binding:"required,numeric"`
	WarehouseId        int64 `json:"warehouse_id" binding:"required,numeric"`
	ProductTypeId      int64 `json:"product_type_id" binding:"required,numeric"`
}
type SectionRequestUpdate struct {
	SectionNumber      int   `json:"section_number" binding:"omitempty,numeric"`
	CurrentCapacity    int   `json:"current_capacity" binding:"omitempty,numeric"`
	CurrentTemperature int   `json:"current_temperature" binding:"omitempty,numeric"`
	MaximumCapacity    int   `json:"maximum_capacity" binding:"omitempty,numeric"`
	MinimumCapacity    int   `json:"minimum_capacity" binding:"omitempty,numeric"`
	MinimumTemperature int   `json:"minimum_temperature" binding:"omitempty,numeric"`
	WarehouseId        int64 `json:"warehouse_id" binding:"omitempty,numeric"`
	ProductTypeId      int64 `json:"product_type_id" binding:"omitempty,numeric"`
}

/*
 * Estrutura do repository.go
 * Repository interface{}
 */
type Repository interface {
	ListarSectionAll() ([]Section, error)
	CreateSection(section Section) (Section, error)
	UpdateSection(section Section) (Section, error)
	ListarSectionOne(id int64) (Section, error)
	DeleteSection(id int64) error
}
type ProductTypes struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

/*
 * estrutura do service.go
 *	Service interface{}
 */
type Service interface {
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int64) (Section, error)
	CreateSection(section SectionRequestCreate) (SectionRequestCreate, error)
	UpdateSection(id int64, sectionUp SectionRequestUpdate) (Section, error)
	DeleteSection(id int64) error
}

/*
 * -- DATABASE QUERIES
 */
const (
	SqlSelect = `SELECT id,section_number,current_capacity,current_temperature,maximum_capacity,minimum_capacity,
	minimum_temperature,product_type_id,warehouse_id FROM mercadofresco.sections`
	SqlSelectByID = `SELECT id,section_number,current_capacity,current_temperature,maximum_capacity,minimum_capacity,
	minimum_temperature,product_type_id,warehouse_id FROM mercadofresco.sections WHERE id=?`
	SqlCreateSection = `INSERT INTO mercadofresco.sections (section_number,current_capacity,current_temperature,
	maximum_capacity,minimum_capacity,minimum_temperature,product_type_id,warehouse_id) VALUES (?,?,?,?,?,?,?,?)`
	SqlUpdateSection = `UPDATE mercadofresco.sections SET section_number=?,current_capacity=?,current_temperature=?,
	maximum_capacity=?,minimum_capacity=?,minimum_temperature=?,product_type_id=?, warehouse_id=? WHERE id=?`
	SqlDeleteSection = `DELETE FROM mercadofresco.sections WHERE id=?`
)
