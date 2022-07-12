package warehouse

type Service interface {
	GetAll() ([]Warehouse, error)
	Create(address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_id int) (RequestWarehouseCreate, error)
	GetByID(id int) (Warehouse, error)
	Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_Id int) (Warehouse, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Warehouse, error) {
	warehouse, err := s.repository.GetAll()
	if err != nil {
		return []Warehouse{}, err
	}
	return warehouse, nil
}

func (s *service) Create(address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_Id int) (RequestWarehouseCreate, error) {

	warehouse, err := s.repository.Create(address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_Id)
	if err != nil {
		return RequestWarehouseCreate{}, err
	}

	return warehouse, nil

}

func (s *service) Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature, locality_id int) (Warehouse, error) {
	warehouse, err := s.repository.GetByID(id)
	if err != nil {
		return Warehouse{}, err
	}
	if address != "" {
		warehouse.Address = address
	}
	if telephone != "" {
		warehouse.Telephone = telephone
	}
	if warehouse_code != "" {
		warehouse.Warehouse_code = warehouse_code
	}
	if minimum_capacity != 0 {
		warehouse.Minimum_capacity = minimum_capacity
	}
	if minimum_temperature != 0 {
		warehouse.Minimum_temperature = minimum_temperature
	}
	if locality_id != 0 {
		warehouse.Locality_id = locality_id
	}

	w, err = s.repository.Update(warehouse.ID, warehouse.Address, warehouse.Telephone, warehouse.Warehouse_code, warehouse.Minimum_capacity, warehouse.Minimum_temperature, warehouse.Locality_id)
	if err != nil {
		return Warehouse{}, err
	}
	return w, nil
}

func (s *service) GetByID(id int) (Warehouse, error) {
	w, err := s.repository.GetByID(id)
	if err != nil {
		return Warehouse{}, err
	}
	return w, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
