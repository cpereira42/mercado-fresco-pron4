package warehouse

import "errors"

type Repository interface {
	GetAll() ([]Warehouse, error)                                                                                           // GET
	Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) // POST
	LastID() (int, error)                                                                                                   // CONTADOR
	Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) // PATCH
	GetByID(id int) (Warehouse, error)                                                                                      // GET
	// DeleteWarehouse(id int)                                                                                                          // DELETE
}

var wr []Warehouse
var lastID int

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Warehouse, error) {
	return wr, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {
	w := Warehouse{id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature}
	wr = append(wr, w)
	lastID = w.ID
	return w, nil

}

func (r *repository) Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {
	w := Warehouse{id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature}
	update := false
	for i := range wr {
		if wr[i].ID == id {
			w.ID = id
			wr[i] = w
			update = true
		}
	}
	if !update {
		return Warehouse{}, errors.New("Warehouse not found")
	}
	return w, nil
}

func (r *repository) GetByID(id int) (Warehouse, error) {
	var w Warehouse
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
	return w, nil

}

// func (r *repository) Delete(id int) error {
// 	delete := false
// 	var index int
// 	for i := range wr {
// 		if wr[i].ID == id {
// 			delete = true
// 			index = i
// 		}
// 	}
// 	if !delete {
// 		return errors.New("Warehouse not found")
// 	}
// 	wr = append(wr[:index], wr[index+1:]...)
// 	return nil
// }
