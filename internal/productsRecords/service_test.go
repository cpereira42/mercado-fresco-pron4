package productsRecords_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cpereira42/mercado-fresco-pron4/internal/productsRecords"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productsRecords/mocks"
	"github.com/stretchr/testify/assert"
)

var prod1 = productsRecords.ReturnProductRecords{

	ProductId:    2,
	RecordsCount: 1,
	Description:  "notebook"}

var prod2 = productsRecords.ReturnProductRecords{

	ProductId:    6,
	RecordsCount: 8,
	Description:  "celular"}

var prodNew2 = productsRecords.RequestProductRecordsCreate{

	PurchasePrice: 155,
	SalePrice:     105,
	ProductId:     5}

var prodRet = productsRecords.ProductRecords{
	PurchasePrice: 155,
	SalePrice:     105,
	ProductId:     5}

func Test_RepositoryFindAll(t *testing.T) {

	produtos := []productsRecords.ReturnProductRecords{prod1, prod2}
	t.Run("Find All Ok", func(t *testing.T) {
		repo := &mocks.Repository{}

		repo.On("GetAllRecords").Return(produtos, nil)

		service := productsRecords.NewService(repo)
		ps, err := service.GetAllRecords()

		assert.Nil(t, err)
		assert.Equal(t, produtos, ps)
		repo.AssertExpectations(t)
	})

	t.Run("Find All Fail", func(t *testing.T) {
		repo := &mocks.Repository{}
		repo.On("GetAllRecords").Return([]productsRecords.ReturnProductRecords{}, fmt.Errorf("Fail to get all"))

		service := productsRecords.NewService(repo)
		ps, err := service.GetAllRecords()
		assert.True(t, len(ps) == 0)
		assert.Equal(t, fmt.Errorf("Fail to get all"), err)

		repo.AssertExpectations(t)
	})
}

func Test_GetId(t *testing.T) {

	t.Run("Find Id Ok", func(t *testing.T) {
		repo := &mocks.Repository{}

		repo.On("GetIdRecords", 2).Return(prod1, nil)

		service := productsRecords.NewService(repo)
		ps, err := service.GetIdRecords(2)

		assert.Nil(t, err)
		assert.Equal(t, prod1, ps)
		repo.AssertExpectations(t)
	})

	t.Run("Find Id Not found", func(t *testing.T) {
		repo := &mocks.Repository{}

		repo.On("GetIdRecords", 3).Return(productsRecords.ReturnProductRecords{}, fmt.Errorf("data not found"))

		service := productsRecords.NewService(repo)
		ps, err := service.GetIdRecords(3)

		assert.NotNil(t, err)
		assert.Equal(t, productsRecords.ReturnProductRecords{}, ps)
		assert.Equal(t, "data not found", err.Error())
		repo.AssertExpectations(t)
	})
}

func Test_Create(t *testing.T) {

	t.Run("Create Ok", func(t *testing.T) {
		repo := &mocks.Repository{}
		currentTime := time.Now()
		theTime := time.Date(currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
			currentTime.Hour(),
			currentTime.Minute(),
			currentTime.Second(),
			100,
			time.Local).Format("2006-01-02 15:04:05")

		prodRet.LastUpdateDate = theTime
		repo.On("Create", prodRet).Return(prodRet, nil)

		service := productsRecords.NewService(repo)
		ps, err := service.Create(prodNew2)

		assert.Nil(t, err)
		assert.Equal(t, prodRet, ps)
		repo.AssertExpectations(t)
	})

	t.Run("Create Fail", func(t *testing.T) {
		repo := &mocks.Repository{}
		currentTime := time.Now()
		theTime := time.Date(currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
			currentTime.Hour(),
			currentTime.Minute(),
			currentTime.Second(),
			100,
			time.Local).Format("2006-01-02 15:04:05")

		prodRet.LastUpdateDate = theTime
		repo.On("Create", prodRet).Return(productsRecords.ProductRecords{}, fmt.Errorf("Error to save"))

		service := productsRecords.NewService(repo)
		ps, err := service.Create(prodNew2)

		assert.NotNil(t, err)
		assert.Equal(t, productsRecords.ProductRecords{}, ps)
		assert.Equal(t, "Error to save", err.Error())
		repo.AssertExpectations(t)
	})

}
