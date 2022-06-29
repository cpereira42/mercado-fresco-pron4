package section

/*
 * Modelo de estrutura de entradas de requests post/patch e de db
 * Entidade: Section{}, [usado para respostas do db no repository]
 * Request POST: SectionRequestCreate{}
 * Request PATCH: SectionRequestUpdate{}
 * Request POST: ProductBatches{}
 * Request GET: ProductBatchesResponse{}
 */
type Section struct {
	Id                 int64 `json:"id,omitempty" binding:"numeric,omitempty"`
	SectionNumber      int   `json:"section_number" binding:"alfa,omitempty"`
	CurrentTemperature int   `json:"current_temperature" binding:"numeric,omitempty"`
	MinimumTemperature int   `json:"minimum_temperature" binding:"numeric,omitempty"`
	CurrentCapacity    int   `json:"current_capacity" binding:"numeric,omitempty"`
	MinimumCapacity    int   `json:"minimum_capacity" binding:"numeric,omitempty"`
	MaximumCapacity    int   `json:"maximum_capacity" binding:"numeric,omitempty"`
	WarehouseId        int64 `json:"warehouse_id" binding:"numeric,omitempty"`
	ProductTypeId      int64 `json:"product_type_id" binding:"numeric,omitempty"`
}
type SectionRequestCreate struct {
	SectionNumber      int   `json:"section_number" binding:"required,numeric"`
	CurrentTemperature int   `json:"current_temperature" binding:"required,numeric"`
	MinimumTemperature int   `json:"minimum_temperature" binding:"required,numeric"`
	CurrentCapacity    int   `json:"current_capacity" binding:"required,numeric"`
	MinimumCapacity    int   `json:"minimum_capacity" binding:"required,numeric"`
	MaximumCapacity    int   `json:"maximum_capacity" binding:"required,numeric"`
	WarehouseId        int64 `json:"warehouse_id" binding:"required,numeric"`
	ProductTypeId      int64 `json:"product_type_id" binding:"required,numeric"`
}
type SectionRequestUpdate struct {
	SectionNumber      int   `json:"section_number" binding:"required,numeric,omitempty"`
	CurrentTemperature int   `json:"current_temperature" binding:"numeric,omitempty"`
	MinimumTemperature int   `json:"minimum_temperature" binding:"numeric,omitempty"`
	CurrentCapacity    int   `json:"current_capacity" binding:"numeric,omitempty"`
	MinimumCapacity    int   `json:"minimum_capacity" binding:"numeric,omitempty"`
	MaximumCapacity    int   `json:"maximum_capacity" binding:"numeric,omitempty"`
	WarehouseId        int64 `json:"warehouse_id" binding:"numeric,omitempty"`
	ProductTypeId      int64 `json:"product_type_id" binding:"numeric,omitempty"`
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
	CreateSection(section SectionRequestCreate) (Section, error)
	ListarSectionOne(id int64) (Section, error)
	UpdateSection(id int64, sectionUp SectionRequestUpdate) (Section, error)
	DeleteSection(id int64) error
}

/*
 * -- DATABASE QUERIES
 */
const (
	sqlSelect = `
		SELECT id,section_number,current_capacity,current_temperature,maximum_capacity,minimum_capacity,minimum_temperature,product_type_id,warehouse_id
		FROM mercadofresco.sections`
	sqlSelectByID = `
		SELECT id,section_number,current_capacity,current_temperature,maximum_capacity,minimum_capacity,minimum_temperature,product_type_id,warehouse_id
		FROM mercadofresco.sections WHERE id=?`
	sqlCreateSection = `
		INSERT INTO mercadofresco.sections (section_number,current_capacity,current_temperature,maximum_capacity,minimum_capacity,minimum_temperature,product_type_id,warehouse_id)
		VALUES (?,?,?,?,?,?,?,?)`
	sqlUpdateSection = `
		UPDATE mercadofresco.sections 
		SET section_number=?,current_capacity=?,current_temperature=?,maximum_capacity=?,minimum_capacity=?,
		minimum_temperature=?,product_type_id=?, warehouse_id=? WHERE id=?`
	sqlDeleteSection = `DELETE FROM mercadofresco.sections WHERE id=?`
)
