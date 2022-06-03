package employee

type Service interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (Employee, error)
	Create(cardNumberID, firstName, lastName string, warehouseID int) (Employee, error)
	Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error)
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

func (s service) GetAll() ([]Employee, error) {
	employees, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (s service) GetByID(id int) (Employee, error) {
	employee, err := s.repository.GetByID(id)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func (s service) Create(cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Employee{}, nil
	}
	lastID++

	employee, err := s.repository.Create(lastID, cardNumberID, firstName, lastName, warehouseID)
	if err != nil {
		return Employee{}, nil
	}

	return employee, nil
}

func (s service) Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	employee, err := s.repository.Update(id, cardNumberID, firstName, lastName, warehouseID)
	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
