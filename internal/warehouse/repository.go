package warehouse

import (
	"errors"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
)

type Repository interface {
	GetAll() ([]Warehouse, error)                                                                                           // GET
	Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) // POST
	LastID() (int, error)                                                                                                   // CONTADOR
	Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) // PATCH
	GetByID(id int) (Warehouse, error)                                                                                      // GET
	Delete(id int) error                                                                                                    // DELETE
}

var wr []Warehouse

//var lastID int

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Warehouse, error) {
	if err := r.db.Read(&wr); err != nil {
		return []Warehouse{}, err

	}
	return wr, nil
}

func (r *repository) LastID() (int, error) {
	if err := r.db.Read(&wr); err != nil {
		return 0, err
	}
	if len(wr) == 0 {
		return 0, nil
	}
	return wr[len(wr)-1].ID, nil
}

func (r *repository) Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {
	if err := r.db.Read(&wr); err != nil {
		return Warehouse{}, err
	}
	w := Warehouse{id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature}
	wr = append(wr, w)
	if err := r.db.Write(&wr); err != nil {
		return Warehouse{}, err
	}
	//lastID = w.ID
	return w, nil

}

func (r *repository) Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {
	if err := r.db.Read(&wr); err != nil {
		return Warehouse{}, err
	}
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
	if err := r.db.Write(&wr); err != nil {
		return Warehouse{}, err
	}
	return w, nil
}

func (r *repository) GetByID(id int) (Warehouse, error) {
	if err := r.db.Read(&wr); err != nil {
		return Warehouse{}, err
	}
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

func (r *repository) Delete(id int) error {
	if err := r.db.Read(&wr); err != nil {
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
	wr = append(wr[:index], wr[index+1:]...)
	return nil
}
