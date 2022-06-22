package products_test

import (
	"fmt"
	"testing"

	//"github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/mocks"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store/mocks"
	"github.com/stretchr/testify/assert"
	//products "github.com/meliBootcamp/go-web/aula03/ex01a/internal/products/repository"
)

var prod12 = products.Product{
	Id:                             1,
	Description:                    "prod1",
	ExpirationRate:                 1,
	FreezingRate:                   2,
	Height:                         3.3,
	Length:                         4.3,
	NetWeight:                      5.5,
	ProductCode:                    "prod1",
	RecommendedFreezingTemperature: 6.6,
	Width:                          7.7,
	ProductTypeId:                  8,
	SellerId:                       9}

func Test_Store(t *testing.T) {

	var ps []products.Product
	produtos := []products.Product{prod12}

	dbProd := store.New(store.FileType, "../repositories/products_test.json")
	repoProd := products.NewRepositoryProducts(dbProd)

	store := &mocks.Store{}
	store.On("Write", produtos).Return(fmt.Errorf("error"))
	store.On("Read", &ps).Return(fmt.Errorf("error"))
	repoProdError := products.NewRepositoryProducts(store)

	t.Run("Last ID == 0", func(t *testing.T) {
		ps, err := repoProd.LastID()
		assert.Equal(t, ps, ps)
		assert.Equal(t, err, err)
	})

	t.Run("Create Ok", func(t *testing.T) {
		ps, err := repoProd.Create(prod12)
		assert.Equal(t, err, err)
		assert.Equal(t, ps, ps)
	})

	t.Run("Create Fail", func(t *testing.T) {
		ps, err := repoProdError.Create(prod12)
		assert.Equal(t, err, err)
		assert.Equal(t, ps, ps)
	})

	/*t.Run("Delete Fail file", func(t *testing.T) {
		err := repoProdError.Delete(1)
		assert.Equal(t, err, err)
		assert.Equal(t, ps, ps)
	})*/

	t.Run("Find GetAll", func(t *testing.T) {
		ps, err := repoProd.GetAll()
		assert.Equal(t, ps, ps)
		assert.Equal(t, err, err)
	})

	t.Run("Find GetId Valid", func(t *testing.T) {
		ps, err := repoProd.GetId(1)
		assert.Equal(t, ps, ps)
		assert.Equal(t, err, err)
	})

	t.Run("Find GetId Invalid", func(t *testing.T) {
		ps, err := repoProd.GetId(9)
		assert.Equal(t, ps, ps)
		assert.Equal(t, err, err)
	})

	t.Run("Find CheckCode duplicated", func(t *testing.T) {
		err := repoProd.CheckCode("prod1")
		assert.Equal(t, err, err)
	})

	t.Run("Find CheckCode Ok", func(t *testing.T) {
		err := repoProd.CheckCode("prod10")
		assert.Equal(t, err, err)
	})

	t.Run("Last ID", func(t *testing.T) {
		ps, err := repoProd.LastID()
		assert.Equal(t, ps, ps)
		assert.Equal(t, err, err)
	})

	t.Run("Last ID - Error", func(t *testing.T) {
		ps, err := repoProd.LastID()
		assert.Equal(t, ps, ps)
		assert.Equal(t, err, err)
	})

	t.Run("Last ID - Error", func(t *testing.T) {
		ps, err := repoProdError.LastID()
		assert.Equal(t, ps, ps)
		assert.Equal(t, err, err)
	})

	t.Run("Update Ok", func(t *testing.T) {
		ps, err := repoProd.Update(1, prod12)
		assert.Equal(t, err, err)
		assert.Equal(t, ps, ps)
	})

	t.Run("Update not found Ok", func(t *testing.T) {
		ps, err := repoProd.Update(9, prod12)
		assert.Equal(t, err, err)
		assert.Equal(t, ps, ps)
	})

	t.Run("Update Fail", func(t *testing.T) {
		ps, err := repoProdError.Update(1, prod12)
		assert.Equal(t, err, err)
		assert.Equal(t, ps, ps)
	})

	t.Run("Delete Ok", func(t *testing.T) {
		err := repoProd.Delete(1)
		assert.Equal(t, err, err)
	})
	t.Run("Delete Not Found", func(t *testing.T) {
		err := repoProd.Delete(10)
		assert.Equal(t, err, err)
	})

}
