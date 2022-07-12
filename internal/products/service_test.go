package products_test

import (
	"fmt"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products/mocks"
	"github.com/stretchr/testify/assert"
)

var lastId = 3

var prod1 = products.Product{
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

var prod2 = products.Product{
	Id:                             2,
	Description:                    "prod2",
	ExpirationRate:                 1,
	FreezingRate:                   2,
	Height:                         3.3,
	Length:                         4.3,
	NetWeight:                      5.5,
	ProductCode:                    "prod2",
	RecommendedFreezingTemperature: 6.6,
	Width:                          7.7,
	ProductTypeId:                  8,
	SellerId:                       9}

var prod3 = products.Product{
	Id:                             3,
	Description:                    "prod3",
	ExpirationRate:                 1,
	FreezingRate:                   2,
	Height:                         3.3,
	Length:                         4.3,
	NetWeight:                      5.5,
	ProductCode:                    "prod3",
	RecommendedFreezingTemperature: 6.6,
	Width:                          7.7,
	ProductTypeId:                  8,
	SellerId:                       9}
var prod4 = products.Product{
	Id:                             4,
	Description:                    "prod4",
	ExpirationRate:                 1,
	FreezingRate:                   2,
	Height:                         3.3,
	Length:                         4.3,
	NetWeight:                      5.5,
	ProductCode:                    "prod4",
	RecommendedFreezingTemperature: 6.6,
	Width:                          7.7,
	ProductTypeId:                  8,
	SellerId:                       1}

var prodNew = products.Product{
	Id:                             0,
	Description:                    "prod4",
	ExpirationRate:                 1,
	FreezingRate:                   2,
	Height:                         3.3,
	Length:                         4.3,
	NetWeight:                      5.5,
	ProductCode:                    "prod4",
	RecommendedFreezingTemperature: 6.6,
	Width:                          7.7,
	ProductTypeId:                  8,
	SellerId:                       1}

var prodCreate = products.RequestProductsCreate{
	Description:                    "prod4",
	ExpirationRate:                 1,
	FreezingRate:                   2,
	Height:                         3.3,
	Length:                         4.3,
	NetWeight:                      5.5,
	ProductCode:                    "prod4",
	RecommendedFreezingTemperature: 6.6,
	Width:                          7.7,
	ProductTypeId:                  8,
	SellerId:                       1}

var prodUp = products.RequestProductsUpdate{
	Description:   "prod3",
	NetWeight:     9.9,
	ProductCode:   "prod3",
	ProductTypeId: 8,
	SellerId:      9}

func Test_RepositoryFindAll(t *testing.T) {

	produtos := []products.Product{prod1, prod2}
	t.Run("Find All Ok", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetAll").Return(produtos, nil)

		service := products.NewService(repo)
		ps, err := service.GetAll()

		assert.Nil(t, err)
		assert.True(t, len(ps) == 2)
		assert.Equal(t, produtos, ps)
		repo.AssertExpectations(t)
	})

	t.Run("Find All Fail", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetAll").Return([]products.Product{}, fmt.Errorf("Fail to get all"))

		service := products.NewService(repo)
		ps, err := service.GetAll()
		assert.True(t, len(ps) == 0)
		assert.Equal(t, "Fail to get all", err.Error())
		repo.AssertExpectations(t)
	})
}

func Test_RepositoryFindId(t *testing.T) {

	produtos := []products.Product{prod1, prod2}
	t.Run("GetId Ok", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetId", 1).Return(produtos[0], nil)

		service := products.NewService(repo)
		ps, err := service.GetId(1)

		assert.Nil(t, err)
		assert.Equal(t, produtos[0], ps)
		repo.AssertExpectations(t)
	})

	t.Run("GetId Fail", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetId", 3).Return(products.Product{}, fmt.Errorf("Product not Found"))
		service := products.NewService(repo)
		_, err := service.GetId(3)
		assert.NotNil(t, err)
		assert.Equal(t, "Product not Found", err.Error())
		repo.AssertExpectations(t)
	})
}

func Test_RepositoryDelete(t *testing.T) {

	t.Run("Delete Ok", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("Delete", 1).Return(nil)
		service := products.NewService(repo)
		err := service.Delete(1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("TestDeleteFail", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("Delete", 9).Return(fmt.Errorf("produto não encontrado"))
		service := products.NewService(repo)
		err := service.Delete(9)
		assert.Equal(t, "produto não encontrado", err.Error())
		repo.AssertExpectations(t)
	})
}

func Test_RepositoryUpdate(t *testing.T) {
	produtos := []products.Product{prod1, prod2, prod3}
	t.Run("Update Ok", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetId", 3).Return(produtos[2], nil)
		prod3.NetWeight = 9.9
		repo.On("Update", 3, prod3).Return(prod3, nil)
		service := products.NewService(repo)
		ps, err := service.Update(3, prodUp)

		assert.Nil(t, err)
		assert.Equal(t, prod3, ps)
		repo.AssertExpectations(t)
	})

	t.Run("Update Fail", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetId", 5).Return(products.Product{}, fmt.Errorf("Product not found"))
		service := products.NewService(repo)
		_, err := service.Update(5, prodUp)

		assert.Equal(t, "Product not found", err.Error())
		repo.AssertExpectations(t)

	})
}

func Test_RepositoryCreate(t *testing.T) {

	t.Run("Create Ok", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("Create", prodNew).Return(prod4, nil)
		service := products.NewService(repo)
		ps, err := service.Create(prodCreate)
		lastId++
		ps.Id = 4
		assert.Nil(t, err)
		assert.Equal(t, prod4, ps)
		repo.AssertExpectations(t)
	})

	t.Run("Create Fail to Save ", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("Create", prodNew).Return(products.Product{}, fmt.Errorf("Fail to save"))
		service := products.NewService(repo)
		_, err := service.Create(prodCreate)

		assert.NotNil(t, err)
		assert.Equal(t, "Fail to save", err.Error())
		repo.AssertExpectations(t)

	})
}
