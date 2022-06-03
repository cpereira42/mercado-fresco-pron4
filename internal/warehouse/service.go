package warehouse

type Service interface {
	GetAll() ([]Warehouse, error)
	Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error)
	GetByID(id int) (Warehouse, error)                                                                                      // GET
	Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) // PATCH
	Delete(id int) error                                                                                                    // DELETE
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
	return s.repository.GetAll()
}

func (s *service) Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Warehouse{}, err
	}

	lastID++

	warehouse, err := s.repository.Create(lastID, address, telephone, warehouse_code, minimum_capacity, minimum_temperature)
	if err != nil {
		return Warehouse{}, err
	}

	return warehouse, nil

}

func (s *service) Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {
	warehouse, err := s.repository.Update(id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature)
	if err != nil {
		return Warehouse{}, err
	}
	return warehouse, nil
}

func (s *service) GetByID(id int) (Warehouse, error) {
	warehouse, err := s.repository.GetByID(id)
	if err != nil {
		return Warehouse{}, err
	}
	return warehouse, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
