package section

/*
 * Modelo de estrutura de entradas de requests post/patch e de db
 * Entidade: Section{}, [usado para respostas do db no repository]
 * Request POST: SectionRequestCreate{}
 * Request PATCH: SectionRequestUpdate{}
 */
type Section struct {
	Id                 int     `json:"id" binding:"numeric,omitempty"`
	SectionNumber      string  `json:"section_number" binding:"alfa,omitempty"`
	CurrentTemperature float64 `json:"current_temperature" binding:"numeric,omitempty"`
	MinimumTemperature float64 `json:"minimum_temperature" binding:"numeric,omitempty"`
	CurrentCapacity    int     `json:"current_capacity" binding:"numeric,omitempty"`
	MinimumCapacity    int     `json:"minimum_capacity" binding:"numeric,omitempty"`
	MaximumCapacity    int     `json:"maximum_capacity" binding:"numeric,omitempty"`
	WarehouseId        int     `json:"warehouse_id" binding:"numeric,omitempty"`
	ProductTypeId      int     `json:"product_type_id" binding:"numeric,omitempty"`
}
type SectionRequestCreate struct {
	SectionNumber      string  `json:"section_number" binding:"required,alphanum"`
	CurrentTemperature float64 `json:"current_temperature" binding:"required,numeric"`
	MinimumTemperature float64 `json:"minimum_temperature" binding:"required,numeric"`
	CurrentCapacity    int     `json:"current_capacity" binding:"required,numeric"`
	MinimumCapacity    int     `json:"minimum_capacity" binding:"required,numeric"`
	MaximumCapacity    int     `json:"maximum_capacity" binding:"required,numeric"`
	WarehouseId        int     `json:"warehouse_id" binding:"required,numeric"`
	ProductTypeId      int     `json:"product_type_id" binding:"required,numeric"`
}
type SectionRequestUpdate struct {
	SectionNumber      string  `json:"section_number" binding:"required,numeric,omitempty"`
	CurrentTemperature float64 `json:"current_temperature" binding:"numeric,omitempty"`
	MinimumTemperature float64 `json:"minimum_temperature" binding:"numeric,omitempty"`
	CurrentCapacity    int     `json:"current_capacity" binding:"numeric,omitempty"`
	MinimumCapacity    int     `json:"minimum_capacity" binding:"numeric,omitempty"`
	MaximumCapacity    int     `json:"maximum_capacity" binding:"numeric,omitempty"`
	WarehouseId        int     `json:"warehouse_id" binding:"numeric,omitempty"`
	ProductTypeId      int     `json:"product_type_id" binding:"numeric,omitempty"`
}

/*
	novos requisitos

	É necessário que os Product_Batches sejam compostos por • id: Identificación unica
	• batch_number: número do lote
	• current_quantity: quantidade atual por lote
	• current_temperature: Temperatura atual.
	• due_date: data de validade do produto
	• initial_quantity: quantidade inicial
	• manufacturing_date: data de fabricação
	• manufacturing_hour: hora de fabricação
	• minumum_temperature: Temperatura mínima • product_id: ID do produto
	• section_id: ID da seção

	restrições de criação
	* o product id deve existir
	* o section id deve existir
	batch_number é um campo único

	etapa de criação
	POST |	/api/v1/productBatches
	* Quando a entrada de dados for bem-sucedida, um código 201 será retornado junto com o objeto inserido.
	* Se o batch_number já existir, ele retornará um erro 409 Conflict.
	* Se o objeto JSON não contiver os campos necessários, um código 422 será retornado.

	GET  |	/api/v1/sections/reportProduc ts?id=1
	* Quando a solicitação for bem-sucedida, o backend retornará um relatório com o número de Produtos em
		cada Seção ou o número de produtos de uma determinada seção.


	Request :
	{
		"batch_number": 111,
		"current_quantity": 200,
		"current_temperature": 20,
		"due_date": "2022-04-04",
		"initial_quantity": 10,
		"manufacturing_date": "2020-04-04",
		"manufacturing_hour": 10,
		"minumum_temperature":5,
		"product_id": 1,
		"section_id": 1
	}

	Response:
	{
		"section_id": 1,
		"section_number": 23,
		"products_count": 200
	}

*/

type ProductBatches struct {
	BatchNumber        int    `json:"batch_number" binding:"required,numeric"`
	CurrentQuantity    int    `json:"current_quantity" binding:"required,numeric"`
	DueDate            string `json:"due_date" binding:"require,alpha"`
	InitialQuantity    int    `json:"initial_quantity" binding:"required,numeric"`
	ManufacturingDate  string `json:"manufacturing_date" binding:"required,numeric"`
	ManufacturingHour  int    `json:"manufacturing_hour" binding:"required,numeric"`
	MinimumTemperature int    `json:"minumum_temperature" binding:"required,numeric"`
	ProductId          int    `json:"product_id" binding:"required,numeric"`
	SectionId          int    `json:"section_id" binding:"required,numeric"`
}

type ProdctuBatchesResponse struct {
	SectionId     int `json:"section_id" binding:"required"`
	SectionNumber int `json:"section_number" binding:"required"`
	ProductsCount int `json:"products_count" binding:"required"`
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
		minimum_temperature=?,product_type_id=?, warehouse_id=? WHERE id = ?;`
	sqlDeleteSection = `DELETE FROM mercadofresco.sections WHERE id = ?;`
)
