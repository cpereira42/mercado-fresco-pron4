package carries_test

import (
	"errors"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/carries"
	"github.com/cpereira42/mercado-fresco-pron4/internal/carries/mocks"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var carryOne = carries.Carries{"cid1", "companyname1", "rua 1", "11112222", 1}

var localityOne = carries.Localities{1, "São Paulo", 1}
var localityTwo = carries.Localities{2, "Nova York", 2}

var carriesGetReports = []carries.Localities{localityOne, localityTwo}

func TestServiceCreate(t *testing.T) {

	t.Run("If all fields are valid, should return a new carry", func(t *testing.T) {
		mockRepository := new(mocks.Repository)
		mockRepository.On("Create",
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int")).Return(carryOne, nil)

		service := carries.NewService(mockRepository)

		newCarry, err := service.Create("cid1", "companyname1", "rua 1", "11112222", 1)

		assert.Nil(t, err)
		assert.ObjectsAreEqualValues(carryOne, newCarry)
	})

	t.Run("If Create fail, should return a an error", func(t *testing.T) {
		mockRepository := new(mocks.Repository)
		errorMsg := "fail to create"

		mockRepository.On("Create",
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("int")).Return(carries.Carries{}, errors.New(errorMsg))

		service := carries.NewService(mockRepository)

		_, err := service.Create("cid1", "companyname1", "rua 1", "11112222", 1)

		assert.Error(t, err)

	})
}

func TestServiceGetAllReport(t *testing.T) {
	t.Run("Genearate a report of all localities - OK", func(t *testing.T) {
		mockRepository := new(mocks.Repository)
		mockRepository.On("GetAllReport").Return(carriesGetReports, nil)

		service := carries.NewService(mockRepository)

		carries, err := service.GetAllReport()

		assert.Nil(t, err)
		assert.ObjectsAreEqualValues(carriesReports, carries)
	})
	t.Run("Try to Genereate a report of all localities - Fail", func(t *testing.T) {
		mockRepository := new(mocks.Repository)
		errorMsg := "fail to get all localities"

		mockRepository.On("GetAllReport").Return([]carries.Localities{}, errors.New(errorMsg))

		service := carries.NewService(mockRepository)

		_, err := service.GetAllReport()

		assert.Error(t, err)

	})
}

func TestServiceGetByIDReport(t *testing.T) {
	t.Run("Get a report of a locality by ID - OK", func(t *testing.T) {
		mockRepository := new(mocks.Repository)
		mockRepository.On("GetByIDReport", tmock.AnythingOfType("int")).Return(localityOne, nil)

		service := carries.NewService(mockRepository)

		locality, err := service.GetByIDReport(1)

		assert.Nil(t, err)
		assert.ObjectsAreEqualValues(localityOne, locality)
	})
	t.Run("Get a report of a locality by ID - Fail", func(t *testing.T) {
		mockRepository := new(mocks.Repository)
		errorMsg := "fail to get locality"

		mockRepository.On("GetByIDReport", tmock.AnythingOfType("int")).Return(carries.Localities{}, errors.New(errorMsg))

		service := carries.NewService(mockRepository)

		_, err := service.GetByIDReport(1)

		assert.Error(t, err)

	})

}
