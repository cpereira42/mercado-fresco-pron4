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
		return []Employee{}, err
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
	employee, err := s.repository.Create(cardNumberID, firstName, lastName, warehouseID)

	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (s service) Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	employee, err := s.GetByID(id)
	if err != nil {
		return Employee{}, err
	}

	if cardNumberID != "" {
		employee.CardNumberID = cardNumberID
	}
	if firstName != "" {
		employee.FirstName = firstName
	}
	if lastName != "" {
		employee.LastName = lastName
	}
	if warehouseID > 0 {
		employee.WarehouseID = warehouseID
	}

	employee, err = s.repository.Update(employee.ID, employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID)
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
