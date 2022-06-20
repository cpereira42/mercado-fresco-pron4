package section_test

import (  
	"fmt"
	"testing"
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	mocks "github.com/cpereira42/mercado-fresco-pron4/internal/section/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
 
func TestServiceListarSectionAll(t *testing.T) {

	t.Run("test de integração de repository e service, metodo ListarSectionAll, caso de sucesso", func(t *testing.T) {
		var sectionList []section.Section = []section.Section{
			{
				Id: 1,
				SectionNumber: 3,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},{
				Id: 2,
				SectionNumber: 313,
				CurrentTemperature: 745,
				MinimumTemperature: 344,
				CurrentCapacity: 1345,
				MinimumCapacity: 243,
				MaximumCapacity: 43456,
				WarehouseId: 784,
				ProductTypeId: 43456,
			},
		}

		mockRep := new(mocks.SectionRepository)

		mockRep.On("ListarSectionAll").
			Return(sectionList, nil). 
			Once()

		service := section.NewService(mockRep)

		obSectionList, err := service.ListarSectionAll()
		assert.Nil(t, err)
		assert.Equal(t, sectionList[0].SectionNumber, obSectionList[0].SectionNumber)
		assert.True(t, len(obSectionList) > 0)
 	})
	t.Run("test de integração de repository e service, metodo ListarSectionAll, caso de error", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)

		listaSectionNil := []section.Section{}
		
		errList := fmt.Errorf("não há sections registrado")
		
		mockRep.On("ListarSectionAll").
			Return(listaSectionNil, errList). 
			Once()

		service := section.NewService(mockRep)

		obSectionList, err := service.ListarSectionAll()
		assert.Error(t, err)
		assert.Equal(t, errList, err)
		assert.Equal(t, listaSectionNil, obSectionList)
		assert.True(t, len(obSectionList) == 0)
 	})

}


func TestServiceListarSectionOne(t *testing.T) {

	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de sucesso", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)
		
		var sList []section.Section = []section.Section{
			{
				Id: 1,
				SectionNumber: 1,
				CurrentTemperature: 1,
				MinimumTemperature: 1,
				CurrentCapacity: 1,
				MinimumCapacity: 1,
				MaximumCapacity: 1,
				WarehouseId: 1,
				ProductTypeId: 1,
			},
		}

		sectionOne := section.Section{
			Id: 1,
			SectionNumber: 1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity: 1,
			MinimumCapacity: 1,
			MaximumCapacity: 1,
			WarehouseId: 1,
			ProductTypeId: 1,
		}

		mockRep.On("ListarSectionAll").Return(sList, nil).Once()
		mockRep.On("ListarSectionOne",  mock.AnythingOfType("int")).
			Return(sectionOne, nil). 
			Once()

		service := section.NewService(mockRep)

		obSectionOne, err := service.ListarSectionOne(1)
		assert.Nil(t, err)
		assert.ObjectsAreEqual(sectionOne, obSectionOne)
 
	})
	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de error", func(t *testing.T) {
		
		var sectionListError []section.Section = []section.Section{
			{
				Id: 4,
				SectionNumber: 495,
				CurrentTemperature: 795,
				MinimumTemperature: 3,
				CurrentCapacity: 15,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},
		}

		mockRep := new(mocks.SectionRepository)

		expecErr := fmt.Errorf("Section is not registered")
		sectionNil := section.Section{}

		mockRep.On("ListarSectionAll").
			Return(sectionListError, nil).
			Once()
		mockRep.On("ListarSectionOne", mock.AnythingOfType("int")).
			Return(sectionNil, expecErr).
			Once()
 
		service := section.NewService(mockRep)
		obSectionOne, err := service.ListarSectionOne(10)
		
		assert.Error(t, err)
		assert.ObjectsAreEqual(sectionNil, obSectionOne)
		
	}) 

}


func TestServiceCreateSection(t *testing.T) {

	t.Run("metodo CreateSection, caso de sucesso", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)

		var sectionList []section.Section = []section.Section{
			{
				Id: 1,
				SectionNumber: 3,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},{
				Id: 2,
				SectionNumber: 313,
				CurrentTemperature: 745,
				MinimumTemperature: 344,
				CurrentCapacity: 1345,
				MinimumCapacity: 243,
				MaximumCapacity: 43456,
				WarehouseId: 784,
				ProductTypeId: 43456,
			},{
				Id: 3,
				SectionNumber: 490,
				CurrentTemperature: 795,
				MinimumTemperature: 3,
				CurrentCapacity: 15,
				MinimumCapacity: 23,
				MaximumCapacity: 3,
				WarehouseId: 78,
				ProductTypeId: 456,
			}, {
				Id: 4,
				SectionNumber: 495,
				CurrentTemperature: 795,
				MinimumTemperature: 3,
				CurrentCapacity: 15,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},
		}
	
		newSection := section.SectionRequestCreate{
			SectionNumber: 1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity: 1,
			MinimumCapacity: 1,
			MaximumCapacity: 1,
			WarehouseId: 1,
			ProductTypeId: 1,
		}
		expectSection := section.Section{		
			SectionNumber: 1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity: 1,
			MinimumCapacity: 1,
			MaximumCapacity: 1,
			WarehouseId: 1,
			ProductTypeId: 1,		
		}

		mockRep.On("ListarSectionAll").Return(sectionList, nil).Once()

		mockRep.On("CreateSection", mock.AnythingOfType("section.Section")).Return(expectSection, nil).Once()

		service := section.NewService(mockRep)

		expSection, _ := service.CreateSection(newSection)

		assert.ObjectsAreEqual(expectSection, expSection)

	})
	t.Run("metodo CreateSection, caso de caso de error ao listar sections dentro do metodo CriateSection", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)
	
		newSection := section.SectionRequestCreate{
			SectionNumber: 1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity: 1,
			MinimumCapacity: 1,
			MaximumCapacity: 1,
			WarehouseId: 1,
			ProductTypeId: 1,
		}
		expectSection := section.Section{}
		expectErrorList := fmt.Errorf("não há sections registrados")

		mockRep.On("ListarSectionAll").Return([]section.Section{}, expectErrorList).Once()
		mockRep.On("CreateSection", 
			mock.AnythingOfType("section.Section"),
		).
			Return(expectSection, nil).
			Once()

		service := section.NewService(mockRep)

		expSection, err := service.CreateSection(newSection)
		assert.Error(t, err)

		assert.ObjectsAreEqual(expectSection, expSection)
	})
	t.Run("metodo CreateSection, caso de caso de error ao criar um novo section", func(t *testing.T) {
		var sectionList []section.Section = []section.Section{
			{
				Id: 1,
				SectionNumber: 3,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},{
				Id: 2,
				SectionNumber: 313,
				CurrentTemperature: 745,
				MinimumTemperature: 344,
				CurrentCapacity: 1345,
				MinimumCapacity: 243,
				MaximumCapacity: 43456,
				WarehouseId: 784,
				ProductTypeId: 43456,
			},{
				Id: 3,
				SectionNumber: 490,
				CurrentTemperature: 795,
				MinimumTemperature: 3,
				CurrentCapacity: 15,
				MinimumCapacity: 23,
				MaximumCapacity: 3,
				WarehouseId: 78,
				ProductTypeId: 456,
			}, {
				Id: 4,
				SectionNumber: 495,
				CurrentTemperature: 795,
				MinimumTemperature: 3,
				CurrentCapacity: 15,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},
		}
		mockRep := new(mocks.SectionRepository)
	
		newSection := section.SectionRequestCreate{
			SectionNumber: 3,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity: 1,
			MinimumCapacity: 1,
			MaximumCapacity: 1,
			WarehouseId: 1,
			ProductTypeId: 1,
		}
		expectSection := section.Section{}
		expectErrSectionCreate := fmt.Errorf("section invalid, section_number field must be unique")

		mockRep.On("ListarSectionAll").Return(sectionList, nil).Once()
		mockRep.On("CreateSection", mock.AnythingOfType("section.Section")).
			Return(expectSection, expectErrSectionCreate).
			Once()

		service := section.NewService(mockRep)

		expSection, err := service.CreateSection(newSection)
		
		assert.Error(t, err)
		assert.Equal(t, expectErrSectionCreate, err)
		assert.ObjectsAreEqual(expectSection, expSection)
	})

}


func TestServiceUpdateSection(t *testing.T) {

	t.Run("test servoce no metodoSection, caso de sucesso", func(t *testing.T) {
		var sectionList []section.Section = []section.Section{
			{
				Id: 1,
				SectionNumber: 3,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},{
				Id: 2,
				SectionNumber: 313,
				CurrentTemperature: 745,
				MinimumTemperature: 344,
				CurrentCapacity: 1345,
				MinimumCapacity: 243,
				MaximumCapacity: 43456,
				WarehouseId: 784,
				ProductTypeId: 43456,
			},{
				Id: 3,
				SectionNumber: 490,
				CurrentTemperature: 795,
				MinimumTemperature: 3,
				CurrentCapacity: 15,
				MinimumCapacity: 23,
				MaximumCapacity: 3,
				WarehouseId: 78,
				ProductTypeId: 456,
			}, {
				Id: 4,
				SectionNumber: 495,
				CurrentTemperature: 795,
				MinimumTemperature: 3,
				CurrentCapacity: 15,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},
		}

		mockRep := new(mocks.SectionRepository)
	
		updateSection := section.SectionRequestUpdate{
			SectionNumber: 3,
			CurrentTemperature: 3,
			MinimumTemperature: 1,
			CurrentCapacity: 1,
			MinimumCapacity: 1,
			MaximumCapacity: 1,
			WarehouseId: 1,
			ProductTypeId: 2,
		}
	
		expectUpdateSection := section.Section{
			SectionNumber: 3,
			CurrentTemperature: 3,
			MinimumTemperature: 1,
			CurrentCapacity: 1,
			MinimumCapacity: 1,
			MaximumCapacity: 1,
			WarehouseId: 1,
			ProductTypeId: 2,
		}
		
		mockRep.On("ListarSectionAll"). 
			Return(sectionList, nil). 
			Once()
		mockRep.On("UpdateSection", 
			mock.AnythingOfType("int"),
			mock.Anything). 
			Return(expectUpdateSection, nil). 
			Once()
		
		service := section.NewService(mockRep)
		obUpdateSection, err := service.UpdateSection(1, updateSection)

		assert.Nil(t, err)
		assert.ObjectsAreEqual(expectUpdateSection, obUpdateSection) 
		 
	})
	t.Run("test servoce no metodoSection, caso de error section_number duplicado", func(t *testing.T) { 
		var sectionList []section.Section = []section.Section{
			{
				Id: 1,
				SectionNumber: 3,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},{
				Id: 2,
				SectionNumber: 313,
				CurrentTemperature: 745,
				MinimumTemperature: 344,
				CurrentCapacity: 1345,
				MinimumCapacity: 243,
				MaximumCapacity: 43456,
				WarehouseId: 784,
				ProductTypeId: 43456,
			}, 
		}
		
		mockRep := new(mocks.SectionRepository)
	
		updateSection := section.SectionRequestUpdate{
			SectionNumber: 313,
			CurrentTemperature: 79845,
			MinimumTemperature: 4,
			CurrentCapacity: 135,
			MinimumCapacity: 23,
			MaximumCapacity: 456,
			WarehouseId: 78,
			ProductTypeId: 456,
		}
	
		expectUpdateSection := section.Section{}

		expectedError := fmt.Errorf("this section_number %v is already registered", updateSection.SectionNumber)

		mockRep.On("ListarSectionAll"). 
			Return(sectionList, nil)
		mockRep.On("UpdateSection", 
			mock.AnythingOfType("int"),
			mock.AnythingOfType("section.Section")). 
			Return(expectUpdateSection, expectedError)
		
		service := section.NewService(mockRep)
		  		 
		obUpdateSectionConflict, errConflict := service.UpdateSection(1, updateSection)

		assert.Equal(t, expectedError, errConflict)
		assert.ObjectsAreEqual(expectUpdateSection, obUpdateSectionConflict)
	})
	t.Run("test servoce no metodoSection, caso de error, lista de section retorna vazia dentro do metodo update", func(t *testing.T) { 
		var sectionList []section.Section = []section.Section{}
		
		mockRep := new(mocks.SectionRepository)
	
		updateSection := section.SectionRequestUpdate{
			SectionNumber: 313,
			CurrentTemperature: 79845,
			MinimumTemperature: 4,
			CurrentCapacity: 135,
			MinimumCapacity: 23,
			MaximumCapacity: 456,
			WarehouseId: 78,
			ProductTypeId: 456,
		}
	
		expectUpdateSection := section.Section{}

		expectedError := fmt.Errorf("não há sections registrado")

		mockRep.On("ListarSectionAll"). 
			Return(sectionList, expectedError)
		mockRep.On("UpdateSection", 
			mock.AnythingOfType("int"),
			mock.AnythingOfType("section.Section")). 
			Return(expectUpdateSection, expectedError)
		
		service := section.NewService(mockRep)
		
		listNil, errListNil := service.ListarSectionAll()
		assert.Equal(t, expectedError, errListNil)
		assert.ObjectsAreEqual(sectionList, listNil)

		obUpdateSectionConflict, errConflict := service.UpdateSection(1, updateSection)
		assert.Equal(t, expectedError, errConflict)
		assert.ObjectsAreEqual(expectUpdateSection, obUpdateSectionConflict)
	})
	t.Run("test servoce no metodoSection, caso de error, section não encontrado", func(t *testing.T) { 
		var sectionListError []section.Section = []section.Section{
			{
				Id: 1,
				SectionNumber: 313,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},
		}		
		mockRep := new(mocks.SectionRepository)	
		updateSection := section.SectionRequestUpdate{
			SectionNumber: 313,
			CurrentTemperature: 79845,
			MinimumTemperature: 4,
			CurrentCapacity: 135,
			MinimumCapacity: 23,
			MaximumCapacity: 456,
			WarehouseId: 78,
			ProductTypeId: 456,
		}	
		expectUpdateSection := section.Section{}		
		expectError := fmt.Errorf("unable to update section")
		mockRep.On("ListarSectionAll"). 
			Return(sectionListError, nil)
		mockRep.On("UpdateSection", 
			mock.AnythingOfType("int"),
			mock.AnythingOfType("section.Section")). 
			Return(expectUpdateSection, expectError)		
		service := section.NewService(mockRep)		  		 
		
		obUpdateSectionNotFound, errNotFound := service.UpdateSection(5, updateSection)
		 
		assert.ObjectsAreEqual(expectUpdateSection, obUpdateSectionNotFound)
		assert.Error(t, errNotFound)
		
		 
	})

}


func TestServiceDeleteSection(t *testing.T) {

	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de sucesso", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)
		
		sectionListSucesso := []section.Section{
			{
				Id: 1,
				SectionNumber: 313,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},
			{
				Id: 2,
				SectionNumber: 313,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},{
				Id: 3,
				SectionNumber: 313,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},
		}

		sectionId := 3
		mockRep.On("ListarSectionAll").Return(sectionListSucesso, nil).Once()
		mockRep.On("DeleteSection",  mock.AnythingOfType("int")).
			Return(nil). 
			Once()

		service := section.NewService(mockRep)

		err := service.DeleteSection(sectionId)
		
		assert.Nil(t, err) 
	})
	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de error", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)
	
		sectionListError := []section.Section{
			{
				Id: 1,
				SectionNumber: 313,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},
			{
				Id: 2,
				SectionNumber: 313,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},{
				Id: 3,
				SectionNumber: 313,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WarehouseId: 78,
				ProductTypeId: 456,
			},
		}

		errSection := fmt.Errorf("section not Found")

		sectionId := 20
		mockRep.On("ListarSectionAll").Return(sectionListError, nil).Once()
		mockRep.On("DeleteSection",  mock.AnythingOfType("int")).
			Return(errSection). 
			Once()

		service := section.NewService(mockRep)

		err := service.DeleteSection(sectionId)
		
		assert.Error(t, err)
		assert.Equal(t, errSection, err)
	}) 

}
