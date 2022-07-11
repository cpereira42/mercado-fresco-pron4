package section_test

import (
	"errors"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	mocksSection "github.com/cpereira42/mercado-fresco-pron4/internal/section/mocks"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

var sectionList []section.Section = []section.Section{
	{
		Id:                 1,
		SectionNumber:      3,
		CurrentTemperature: 79845,
		MinimumTemperature: 4,
		CurrentCapacity:    135,
		MinimumCapacity:    23,
		MaximumCapacity:    456,
		WarehouseId:        12,
		ProductTypeId:      456,
	}, {
		Id:                 2,
		SectionNumber:      313,
		CurrentTemperature: 745,
		MinimumTemperature: 344,
		CurrentCapacity:    1345,
		MinimumCapacity:    243,
		MaximumCapacity:    43456,
		WarehouseId:        13,
		ProductTypeId:      43456,
	},
}

var sectionIntance section.Section = section.Section{
	Id:                 3,
	SectionNumber:      45,
	CurrentTemperature: 56,
	MinimumTemperature: 3464,
	CurrentCapacity:    22433,
	MinimumCapacity:    346,
	MaximumCapacity:    566,
	WarehouseId:        1,
	ProductTypeId:      1,
}

var sectionIntanceCreate section.SectionRequestCreate = section.SectionRequestCreate{
	SectionNumber:      45,
	CurrentTemperature: 56,
	MinimumTemperature: 3464,
	CurrentCapacity:    22433,
	MinimumCapacity:    346,
	MaximumCapacity:    566,
	WarehouseId:        1,
	ProductTypeId:      1,
}

var sectionIntanceUpdate section.SectionRequestUpdate = section.SectionRequestUpdate{
	SectionNumber:      45,
	CurrentTemperature: 56,
	MinimumTemperature: 3464,
	CurrentCapacity:    22433,
	MinimumCapacity:    346,
	MaximumCapacity:    566,
	WarehouseId:        1,
	ProductTypeId:      1,
}
var sectionIntanceError section.Section = section.Section{}

func TestServiceListarSectionAll(t *testing.T) {
	t.Run("test de integração de repository e service, metodo ListarSectionAll, caso de sucesso", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("ListarSectionAll").Return(sectionList, nil).Once()
		service := section.NewService(mockRepository)
		objSectionList, err := service.ListarSectionAll()
		assert.Nil(t, err)
		assert.True(t, len(objSectionList) > 0)
		sectionNumberField := 313
		assert.Equal(t, sectionNumberField, sectionList[1].SectionNumber)
	})
	t.Run("test de integração de repository e service, metodo ListarSectionAll, caso de error", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		expectErro := errors.New("sections not this registered")
		mockRepository.On("ListarSectionAll").Return([]section.Section{}, expectErro).Once()
		service := section.NewService(mockRepository)
		objSectionList, err := service.ListarSectionAll()
		assert.Error(t, err)
		assert.Equal(t, expectErro, err)
		assert.ObjectsAreEqual([]section.Section{}, objSectionList)
	})
}

func TestServiceListarSectionOne(t *testing.T) {
	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de sucesso", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("ListarSectionOne", tmock.Anything).
			Return(sectionIntance, nil).
			Once()

		paramId := 1

		service := section.NewService(mockRepository)
		objSection, err := service.ListarSectionOne(int64(paramId))
		assert.NoError(t, err)
		assert.Equal(t, objSection.SectionNumber, sectionIntance.SectionNumber)
	})
	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de error", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("ListarSectionOne", tmock.Anything).
			Return(sectionIntanceError, section.ErrorNotFound).
			Once()

		paramId := 2

		service := section.NewService(mockRepository)
		objSection, err := service.ListarSectionOne(int64(paramId))
		assert.Error(t, err)
		assert.Equal(t, section.ErrorNotFound, err)
		assert.Equal(t, sectionIntanceError.SectionNumber, objSection.SectionNumber)
	})
}

func TestServiceCreateSection(t *testing.T) {
	t.Run("metodo CreateSection, caso de sucesso", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("CreateSection", tmock.Anything).Return(sectionIntance, nil).Once()
		service := section.NewService(mockRepository)
		newSection, err := service.CreateSection(sectionIntanceCreate)
		assert.NoError(t, err)
		assert.ObjectsAreEqual(sectionIntance, newSection)
	})
	t.Run("metodo CreateSection, warehouse_id invalid", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("CreateSection", tmock.Anything).
			Return(sectionIntanceError, errors.New("warehouse_id is not registred ON warehouse")).
			Once()
		service := section.NewService(mockRepository)
		sectionIntanceCreate.WarehouseId = 90
		newSection, err := service.CreateSection(sectionIntanceCreate)
		assert.Error(t, err)
		assert.ObjectsAreEqual(sectionIntanceError, newSection)
	})
	t.Run("metodo CreateSection, product_type_id ivalid", func(t *testing.T) {
		expectError := errors.New("product_type_id is not registred ON products_types")
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("CreateSection", tmock.Anything).
			Return(sectionIntanceError, expectError).
			Once()
		service := section.NewService(mockRepository)
		sectionIntanceCreate.ProductTypeId = 91
		newSection, err := service.CreateSection(sectionIntanceCreate)
		assert.Error(t, err)
		assert.ObjectsAreEqual(sectionIntanceError, newSection)
	})
	t.Run("metodo CreateSection, section_number duplicate", func(t *testing.T) {
		expectError := errors.New("section_number_UNIQUE is Unique, and 151 already registred")
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("CreateSection", tmock.Anything).
			Return(sectionIntanceError, expectError).
			Once()
		service := section.NewService(mockRepository)
		sectionIntanceCreate.SectionNumber = 151
		newSection, err := service.CreateSection(sectionIntanceCreate)
		assert.Error(t, err)
		assert.Equal(t, sectionIntanceError.SectionNumber, newSection.SectionNumber)
	})
}

func TestServiceUpdateSection(t *testing.T) {
	t.Run("test servoce no metodo UpdateSection, caso de sucesso", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("ListarSectionOne", tmock.AnythingOfType("int64")).
			Return(sectionList[1], nil).
			Once()
		mockRepository.On("UpdateSection",
			tmock.AnythingOfType("section.Section"),
		).
			Return(sectionList[1], nil).
			Once()
		service := section.NewService(mockRepository)

		sectionId := int64(1)
		sectionResult, err := service.UpdateSection(sectionId, sectionIntanceUpdate)
		assert.Nil(t, err)
		assert.Equal(t, sectionIntanceUpdate.SectionNumber, sectionResult.SectionNumber)
	})
	t.Run("test servoce no metodo UpdateSection, caso de error section_number duplicado", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		expectError := errors.New("section_number_UNIQUE is Unique, and 3 already registred")
		mockRepository.On("ListarSectionOne", tmock.AnythingOfType("int64")).
			Return(sectionIntance, nil).
			Once()
		mockRepository.On("UpdateSection", tmock.AnythingOfType("section.Section")).
			Return(sectionIntanceError, expectError).
			Once()
		service := section.NewService(mockRepository)

		sectionId := int64(2)
		sectionResult, err := service.UpdateSection(sectionId, sectionIntanceUpdate)
		assert.Error(t, err)
		assert.Equal(t, expectError, err)
		assert.Equal(t, sectionIntanceError.SectionNumber, sectionResult.SectionNumber)
	})
	t.Run("test service no metodo UpdateSection, caso de error, warehouse_id is not registred", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		expectError := errors.New("warehouse_id is not registred ON warehouse")
		mockRepository.On("ListarSectionOne", tmock.AnythingOfType("int64")).
			Return(sectionIntance, nil).
			Once()
		mockRepository.On("UpdateSection", tmock.AnythingOfType("section.Section")).
			Return(sectionIntanceError, expectError).
			Once()
		service := section.NewService(mockRepository)

		sectionId := int64(2)
		sectionResult, err := service.UpdateSection(sectionId, sectionIntanceUpdate)
		assert.Error(t, err)
		assert.Equal(t, expectError, err)
		assert.Equal(t, sectionIntanceError.SectionNumber, sectionResult.SectionNumber)
	})
	t.Run("test service no metodo UpdateSection, caso de error, product_type_id is not registred", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		expectError := errors.New("product_type_id is not registred ON products_types")
		mockRepository.On("ListarSectionOne", tmock.AnythingOfType("int64")).
			Return(sectionIntance, nil).
			Once()
		mockRepository.On("UpdateSection", tmock.AnythingOfType("section.Section")).
			Return(sectionIntanceError, expectError).
			Once()
		service := section.NewService(mockRepository)

		sectionId := int64(2)
		sectionResult, err := service.UpdateSection(sectionId, sectionIntanceUpdate)
		assert.Error(t, err)
		assert.Equal(t, expectError, err)
		assert.Equal(t, sectionIntanceError.SectionNumber, sectionResult.SectionNumber)
	})
}

func TestServiceDeleteSection(t *testing.T) {
	t.Run("test de integração de repository e service, metodo DeleteSection, caso de sucesso", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("ListarSectionOne", tmock.Anything).
			Return(sectionIntance, nil).
			Once()
		mockRepository.On("DeleteSection", tmock.Anything).
			Return(nil).
			Once()
		service := section.NewService(mockRepository)
		sectionDeleteId := 1
		err := service.DeleteSection(int64(sectionDeleteId))
		assert.NoError(t, err)
	})
	t.Run("test de integração de repository e service, metodo DeleteSection, caso de error no db", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("ListarSectionOne", tmock.Anything).
			Return(sectionIntance, nil).
			Once()
		mockRepository.On("DeleteSection", tmock.Anything).
			Return(section.ErrorKeyTableSectionId).
			Once()
		service := section.NewService(mockRepository)
		sectionDeleteId := 1
		err := service.DeleteSection(int64(sectionDeleteId))
		assert.Error(t, err)
		assert.Equal(t, section.ErrorKeyTableSectionId, err)
	})
	t.Run("test de integração de repository e service, metodo DeleteSection, caso de error not found", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("ListarSectionOne", tmock.Anything).
			Return(sectionIntance, nil).
			Once()
		mockRepository.On("DeleteSection", tmock.Anything).
			Return(section.ErrorNotFound).
			Once()
		service := section.NewService(mockRepository)
		sectionDeleteId := 10
		err := service.DeleteSection(int64(sectionDeleteId))
		assert.Error(t, err)
		assert.Equal(t, section.ErrorNotFound, err)
	})
	t.Run("test de integração de repository e service, metodo DeleteSection, erro in list one section", func(t *testing.T) {
		mockRepository := new(mocksSection.Repository)
		mockRepository.On("ListarSectionOne", tmock.Anything).
			Return(sectionIntanceError, section.ErrorNotFound).
			Once()
		mockRepository.On("DeleteSection", tmock.Anything).
			Return(section.ErrorNotFound).
			Once()
		service := section.NewService(mockRepository)
		sectionDeleteId := 10
		err := service.DeleteSection(int64(sectionDeleteId))
		assert.Error(t, err)
		assert.Equal(t, section.ErrorNotFound, err)
	})
}
