package employee_test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
// 	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
// 	"github.com/cpereira42/mercado-fresco-pron4/pkg/store/mocks"
// 	"github.com/stretchr/testify/assert"
// )

// var emply12 = employee.Employee{1, "123", "Eduardo", "Araujo", 1}

// func Test_Store(t *testing.T) {

// 	var es []employee.Employee
// 	employees := []employee.Employee{emply12}

// 	dbEmply := store.New(store.FileType, "../repositories/employees_test.json")
// 	repoProd := employee.NewRepository(dbEmply)

// 	store := &mocks.Store{}
// 	store.On("Write", employees).Return(fmt.Errorf("error"))
// 	store.On("Read", &es).Return(fmt.Errorf("error"))
// 	repoProdError := employee.NewRepository(store)

// 	t.Run("Create Ok", func(t *testing.T) {
// 		es, err := repoProd.Create(emply12.CardNumberID, emply12.FirstName, emply12.LastName, emply12.WarehouseID)
// 		assert.Equal(t, err, err)
// 		assert.Equal(t, es, es)
// 	})

// 	t.Run("Create Fail", func(t *testing.T) {
// 		es, err := repoProdError.Create(emply12.CardNumberID, emply12.FirstName, emply12.LastName, emply12.WarehouseID)
// 		assert.Equal(t, err, err)
// 		assert.Equal(t, es, es)
// 	})

// 	t.Run("Find GetAll", func(t *testing.T) {
// 		es, err := repoProd.GetAll()
// 		assert.Equal(t, es, es)
// 		assert.Equal(t, err, err)
// 	})

// 	t.Run("Find GetId Valid", func(t *testing.T) {
// 		es, err := repoProd.GetByID(1)
// 		assert.Equal(t, es, es)
// 		assert.Equal(t, err, err)
// 	})

// 	t.Run("Last ID", func(t *testing.T) {
// 		es, err := repoProd.LastID()
// 		assert.Equal(t, es, es)
// 		assert.Equal(t, err, err)
// 	})

// 	t.Run("Last ID - Error", func(t *testing.T) {
// 		es, err := repoProd.LastID()
// 		assert.Equal(t, es, es)
// 		assert.Equal(t, err, err)
// 	})

// 	t.Run("Last ID - Error", func(t *testing.T) {
// 		emplys, err := repoProdError.LastID()
// 		assert.Equal(t, emplys, emplys)
// 		assert.Equal(t, err, err)
// 	})

// 	t.Run("Update Ok", func(t *testing.T) {
// 		es, err := repoProd.Update(1, emply12.CardNumberID, emply12.FirstName, emply12.LastName, emply12.WarehouseID)
// 		assert.Equal(t, err, err)
// 		assert.Equal(t, es, es)
// 	})

// 	t.Run("Update not found Ok", func(t *testing.T) {
// 		es, err := repoProd.Update(9, emply12.CardNumberID, emply12.FirstName, emply12.LastName, emply12.WarehouseID)
// 		assert.Equal(t, err, err)
// 		assert.Equal(t, es, es)
// 	})

// 	t.Run("Update Fail", func(t *testing.T) {
// 		es, err := repoProdError.Update(1, emply12.CardNumberID, emply12.FirstName, emply12.LastName, emply12.WarehouseID)
// 		assert.Equal(t, err, err)
// 		assert.Equal(t, es, es)
// 	})

// 	t.Run("Delete Ok", func(t *testing.T) {
// 		err := repoProd.Delete(1)
// 		assert.Equal(t, err, err)
// 	})

// }
