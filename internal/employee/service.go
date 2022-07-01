package employee

import "fmt"

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
		return Employee{}, fmt.Errorf("employee with id %d not found", id)
	}
	return employee, nil
}

func (s service) Create(cardNumberID, firstName, lastName string, warehouseID int) (Employee, error) {
	employees, err := s.repository.GetAll()
	if err != nil {
		return Employee{}, err
	}

	exists := false
	for i := range employees {
		if employees[i].CardNumberID == cardNumberID {
			exists = true
		}
	}
	if exists {
		return Employee{}, fmt.Errorf("employee with this card number id %s exists", cardNumberID)
	}

	if _, err := s.repository.GetByIDWarehouse(warehouseID); err != nil {
		return Employee{}, fmt.Errorf("warehouse id does not exists")
	}

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

	employees, err := s.repository.GetAll()
	if err != nil {
		return Employee{}, err
	}
	exists := false

	for i := range employees {
		if employees[i].CardNumberID == cardNumberID && id != employees[i].ID {
			exists = true
		}
	}
	if exists {
		return Employee{}, fmt.Errorf("employee with this card number id %s exists", cardNumberID)
	}

	employee = Employee{CardNumberID: cardNumberID, FirstName: firstName, LastName: lastName, WarehouseID: warehouseID}
	for i := range employees {
		if employees[i].ID == id {
			employee.ID = id
			if cardNumberID == "" {
				employee.CardNumberID = employees[i].CardNumberID
			}
			if firstName == "" {
				employee.FirstName = employees[i].FirstName
			}
			if lastName == "" {
				employee.LastName = employees[i].LastName
			}
			if warehouseID == 0 {
				employee.WarehouseID = employees[i].WarehouseID
			}
			employees[i] = employee
		}
	}
	if _, err := s.repository.GetByIDWarehouse(employee.WarehouseID); err != nil {
		return Employee{}, fmt.Errorf("warehouse id does not exists")
	}

	employee, err = s.repository.Update(employee.ID, employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID)
	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (s service) Delete(id int) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}

	err := s.repository.Delete(id)

	if err != nil {
		return err
	}
	return err
}
