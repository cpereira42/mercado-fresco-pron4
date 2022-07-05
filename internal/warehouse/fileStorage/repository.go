// package warehouse

// import (
// 	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
// )

// type Repository interface {
// 	GetAll() ([]Warehouse, error)
// 	Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error)
// 	LastID() (int, error)
// 	Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error)
// 	GetByID(id int) (Warehouse, error)
// 	Delete(id int) error
// }

// var wr []Warehouse
// var w Warehouse

// type repository struct {
// 	db store.Store
// }

// func NewRepository(db store.Store) Repository {
// 	return &repository{
// 		db: db,
// 	}
// }

// func (r *repository) GetAll() ([]Warehouse, error) {
// 	if err := r.db.Read(&wr); err != nil {
// 		return []Warehouse{}, err

// 	}
// 	return wr, nil
// }

// func (r *repository) LastID() (int, error) {
// 	if err := r.db.Read(&wr); err != nil {
// 		return 0, err
// 	}
// 	if len(wr) == 0 {
// 		return 0, nil
// 	}
// 	return wr[len(wr)-1].ID, nil
// }

// func (r *repository) Create(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {
// 	if err := r.db.Read(&wr); err != nil {
// 		return Warehouse{}, err
// 	}
// 	w := Warehouse{id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature}

// 	wr = append(wr, w)
// 	if err := r.db.Write(&wr); err != nil {
// 		return Warehouse{}, err
// 	}
// 	return w, nil

// }

// func (r *repository) Update(id int, address, telephone, warehouse_code string, minimum_capacity, minimum_temperature int) (Warehouse, error) {
// 	w := Warehouse{id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature}

// 	if err := r.db.Write(&wr); err != nil {
// 		return Warehouse{}, err
// 	}
// 	return w, nil
// }

// func (r *repository) GetByID(id int) (Warehouse, error) {

// 	if err := r.db.Read(&wr); err != nil {
// 		return Warehouse{}, err
// 	}

// 	// exists := false
// 	// for i := range wr {
// 	// 	if wr[i].ID == id {
// 	// 		w = wr[i]
// 	// 		exists = true
// 	// 	}

// 	// }
// 	// if !exists {
// 	// 	return Warehouse{}, errors.New("Warehouse not found")
// 	// }
// 	return w, nil

// }

// func (r *repository) Delete(id int) error {
// 	if err := r.db.Read(&wr); err != nil {
// 		return err
// 	}

// 	wr = append(wr[:id], wr[id+1:]...)
// 	if err := r.db.Write(&wr); err != nil {
// 		return err
// 	}
// 	return nil
// }
package qualquercoisa
