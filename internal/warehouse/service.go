package warehouse

import (
	"errors"
)

type Service interface {
	GetAll() ([]Warehouse, error)
	Create(address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error)
	GetByID(id int) (Warehouse, error)
	Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error)
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
	return s.repository.GetAll()
}

func (s *service) Create(address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Warehouse{}, err
	}

	lastID++

	wr, err := s.repository.GetAll()
	if err != nil {
		return Warehouse{}, err
	}

	exists := false
	for i := range wr {
		if wr[i].Warehouse_code == warehouse_code {
			exists = true
		}
	}
	if exists {
		return Warehouse{}, errors.New("Warehouse already exists")
	}

	warehouse, err := s.repository.Create(lastID, address, telephone, warehouse_code, minimum_capacity, minimum_temperature)
	if err != nil {
		return Warehouse{}, err
	}

	return warehouse, nil

}

func (s *service) Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {
	wr, err := s.repository.GetAll()
	if err != nil {
		return Warehouse{}, err
	}

	exists := false
	idExists := false
	for i := range wr {
		if wr[i].ID == id {
			w = wr[i]
			idExists = true
		}

	}
	if !idExists {
		return Warehouse{}, errors.New("invalid id")
	}
	for i := range wr {
		if wr[i].Warehouse_code == warehouse_code && id != wr[i].ID {
			exists = true
		}
	}
	if exists {
		return Warehouse{}, errors.New("Warehouse code already exists")
	}
	w := Warehouse{id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature}
	for i := range wr {
		if wr[i].ID == id {
			w.ID = id
			if address == "" {
				w.Address = wr[i].Address
			}
			if telephone == "" {
				w.Telephone = wr[i].Telephone
			}
			if warehouse_code == "" {
				w.Warehouse_code = wr[i].Warehouse_code
			}
			if minimum_capacity == 0 {
				w.Minimum_capacity = wr[i].Minimum_capacity
			}
			if minimum_temperature == 0 {
				w.Minimum_temperature = wr[i].Minimum_temperature
			}
			wr[i] = w
			exists = true
		}
	}
	// if !update {
	// 	return Warehouse{}, errors.New("invalid id")
	// }
	w, err = s.repository.Update(w.ID, w.Address, w.Telephone, w.Warehouse_code, w.Minimum_capacity, w.Minimum_temperature)
	if err != nil {
		return Warehouse{}, err
	}
	return w, nil
}

func (s *service) GetByID(id int) (Warehouse, error) {
	wr, err := s.repository.GetAll()
	if err != nil {
		return Warehouse{}, err
	}
	exists := false
	for i := range wr {
		if wr[i].ID == id {
			w = wr[i]
			exists = true
		}

	}
	if !exists {
		return Warehouse{}, errors.New("Warehouse not found")
	}
	w, err = s.repository.GetByID(id)
	if err != nil {
		return Warehouse{}, err
	}
	return w, nil
}

func (s *service) Delete(id int) error {
	wr, err := s.repository.GetAll()
	if err != nil {
		return err
	}
	delete := false
	var index int
	for i := range wr {
		if wr[i].ID == id {
			delete = true
			index = i
		}
	}
	if !delete {
		return errors.New("Warehouse not found")
	}
	err = s.repository.Delete(index)
	if err != nil {
		return err
	}
	return nil
}
