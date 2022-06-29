package section_test
/*
import (
	"errors"
	"fmt"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	mocks "github.com/cpereira42/mercado-fresco-pron4/internal/section/mock"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	mocksWareshouse "github.com/cpereira42/mercado-fresco-pron4/internal/warehouse/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var sectionObject = section.Section{
	Id:                 3,
	SectionNumber:      "123",
	CurrentTemperature: 1,
	MinimumTemperature: 1,
	CurrentCapacity:    1,
	MinimumCapacity:    1,
	MaximumCapacity:    1,
	WarehouseId:        14,
	ProductTypeId:      1,
}
var newSectionRequest = section.SectionRequestCreate{
	SectionNumber:      "1",
	CurrentTemperature: 1,
	MinimumTemperature: 1,
	CurrentCapacity:    1,
	MinimumCapacity:    1,
	MaximumCapacity:    1,
	WarehouseId:        13,
	ProductTypeId:      1,
}
var sectionList []section.Section = []section.Section{
	{
		Id:                 1,
		SectionNumber:      "3",
		CurrentTemperature: 79845,
		MinimumTemperature: 4,
		CurrentCapacity:    135,
		MinimumCapacity:    23,
		MaximumCapacity:    456,
		WarehouseId:        12,
		ProductTypeId:      456,
	}, {
		Id:                 2,
		SectionNumber:      "313",
		CurrentTemperature: 745,
		MinimumTemperature: 344,
		CurrentCapacity:    1345,
		MinimumCapacity:    243,
		MaximumCapacity:    43456,
		WarehouseId:        13,
		ProductTypeId:      43456,
	},
	sectionObject,
}

func TestServiceListarSectionAll(t *testing.T) {
	t.Run("test de integração de repository e service, metodo ListarSectionAll, caso de sucesso", func(t *testing.T) {
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
		 
		mockRep.On("ListarSectionAll").Return(sectionList, nil).Once()
		mockRep.On("ListarSectionOne", mock.AnythingOfType("int")).
			Return(sectionObject, nil).
			Once()
		service := section.NewService(mockRep)
		obSectionOne, err := service.ListarSectionOne(1)
		assert.Nil(t, err)
		assert.ObjectsAreEqual(sectionObject, obSectionOne)
	})
	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de error", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)
		expecErr := fmt.Errorf("section.Section is not registered")
		sectionNil := section.Section{}
		mockRep.On("ListarSectionAll").
			Return(sectionList, nil).
			Once()
		mockRep.On("ListarSectionOne", mock.AnythingOfType("int")).
			Return(sectionNil, expecErr).
			Once()
		service := section.NewService(mockRep)
		obSectionOne, err := service.ListarSectionOne(3)
		assert.Error(t, err)
		assert.ObjectsAreEqual(sectionNil, obSectionOne)
	})
}

func TestServiceCreateSection(t *testing.T) {
	t.Run("metodo CreateSection, caso de sucesso", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)
		mockRep.On("ListarSectionAll").Return(sectionList, nil).Once()
		mockRep.On("CreateSection", mock.AnythingOfType("section.Section")).
			Return(sectionObject, nil).
			Once()
		service := section.NewService(mockRep)
		expSection, _ := service.CreateSection(newSectionRequest)
		assert.Equal(t, sectionObject.SectionNumber, expSection.SectionNumber)
	})
	t.Run("metodo CreateSection, caso de caso de error ao listar sections dentro do metodo CriateSection", func(t *testing.T) {
		mockSecRep := new(mocks.SectionRepository)
		mockWareRepo := new(mocksWareshouse.Repository)
		expectErrorList := fmt.Errorf("não há sections registrados")
	 
		mockWareRepo.On("GetByID").Return(warehouse.Warehouse{}, errors.New("Warehouse not found")).Once()
		mockSecRep.On("ListarSectionAll").Return([]section.Section{}, expectErrorList).Once()
		mockSecRep.On("CreateSection",
			mock.AnythingOfType("section.Section")).
			Return(section.Section{}, expectErrorList).
			Once()
		service := section.NewService(mockSecRep)
		expSection, err := service.CreateSection(newSectionRequest)
		assert.Error(t, err)
		assert.Equal(t, expectErrorList, err)
		assert.Equal(t, section.Section{}.SectionNumber, expSection.SectionNumber)
	})
	t.Run("metodo CreateSection, caso de caso de error ao criar um novo section", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)
		expectErrSectionCreate := fmt.Errorf("section invalid, section_number field must be unique")
		mockRep.On("ListarSectionAll").Return(sectionList, nil).Once()
		mockRep.On("CreateSection", mock.AnythingOfType("section.Section")).
			Return(section.Section{}, expectErrSectionCreate).
			Once()
		service := section.NewService(mockRep)
		
		newSectionRequest.SectionNumber = "123" // injeta campo duplicado

		expSection, err := service.CreateSection(newSectionRequest)
		assert.Error(t, err)
		assert.Equal(t, expectErrSectionCreate, err)
		assert.Equal(t, section.Section{}.SectionNumber, expSection.SectionNumber)
	})
}

func TestServiceUpdateSection(t *testing.T) {
	t.Run("test servoce no metodo UpdateSection, caso de sucesso", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)
		updateSection := section.SectionRequestUpdate{
			SectionNumber:      "30",
			CurrentTemperature: 3,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseId:        1,
			ProductTypeId:      2,
		} 
		mockRep.On("ListarSectionOne", mock.AnythingOfType("int")). 
			Return(sectionObject, nil).Once()
		mockRep.On("ListarSectionAll").
			Return(sectionList, nil).
			Once()
		mockRep.On("UpdateSection",
			mock.AnythingOfType("int"),
			mock.Anything).
			Return(sectionObject, nil).
			Once()

		service := section.NewService(mockRep)

		obUpdateSection, err := service.UpdateSection(3, updateSection)

		assert.Nil(t, err)

		assert.ObjectsAreEqual(sectionObject, obUpdateSection)
	})
	t.Run("test servoce no metodo UpdateSection, caso de error section_number duplicado", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)
		updateSection := section.SectionRequestUpdate{
			SectionNumber:      "313",
			CurrentTemperature: 79845,
			MinimumTemperature: 4,
			CurrentCapacity:    135,
			MinimumCapacity:    23,
			MaximumCapacity:    456,
			WarehouseId:        78,
			ProductTypeId:      456,
		}
		expectUpdateSection := section.Section{}
		expectedError := fmt.Errorf("this section_number %v is already registered", updateSection.SectionNumber)
		notFoundErr := fmt.Errorf("sections not found")

		mockRep.On("ListarSectionOne", mock.AnythingOfType("int")). 
			Return(expectUpdateSection, notFoundErr).Once()

		mockRep.On("ListarSectionAll").
			Return(sectionList, nil).Once()
		mockRep.On("UpdateSection",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("section.Section")).
			Return(expectUpdateSection, expectedError).Once()
		service := section.NewService(mockRep)
		obUpdateSectionConflict, errConflict := service.UpdateSection(1, updateSection)
		assert.Equal(t, notFoundErr, errConflict)
		assert.ObjectsAreEqual(expectUpdateSection, obUpdateSectionConflict)
	})
	t.Run("test service no metodo UpdateSection, caso de error, lista de section retorna vazia dentro do metodo update", func(t *testing.T) {
		var sectionList []section.Section = []section.Section{}
		mockRep := new(mocks.SectionRepository)
		updateSection := section.SectionRequestUpdate{
			SectionNumber:      "313",
			CurrentTemperature: 79845,
			MinimumTemperature: 4,
			CurrentCapacity:    135,
			MinimumCapacity:    23,
			MaximumCapacity:    456,
			WarehouseId:        78,
			ProductTypeId:      456,
		}
		expectUpdateSection := section.Section{}
		expectedError := fmt.Errorf("não há sections registrado")
		updateErr := fmt.Errorf("sections not found")
		mockRep.On("ListarSectionOne", mock.AnythingOfType("int")). 
			Return(expectUpdateSection, fmt.Errorf("sections not found")).Once()
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

		assert.Equal(t, updateErr, errConflict)

		assert.ObjectsAreEqual(expectUpdateSection, obUpdateSectionConflict)
	})
	t.Run("test service no metodo UpdateSection, caso de error, section não encontrado", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)

		updateSection := section.SectionRequestUpdate{
			SectionNumber:      "313",
			CurrentTemperature: 79845,
			MinimumTemperature: 4,
			CurrentCapacity:    135,
			MinimumCapacity:    23,
			MaximumCapacity:    456,
			WarehouseId:        78,
			ProductTypeId:      456,
		}
		expectUpdateSection := section.Section{}
		
		mockRep.On("ListarSectionOne", mock.AnythingOfType("int")). 
			Return(expectUpdateSection, fmt.Errorf("não há sections registrado")).Once()
		
			expectError := fmt.Errorf("unable to update section")
		
		mockRep.On("ListarSectionAll").
			Return(sectionList, nil)

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
		mockRep.On("ListarSectionAll").Return(sectionList, nil).Once()
		mockRep.On("DeleteSection", mock.AnythingOfType("int")).
			Return(nil).
			Once()
		service := section.NewService(mockRep)
		err := service.DeleteSection(3)
		assert.Nil(t, err)
	})
	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de error", func(t *testing.T) {
		mockRep := new(mocks.SectionRepository)
		errSection := fmt.Errorf("section not Found") 
		mockRep.On("ListarSectionAll").Return(sectionList, nil).Once()
		mockRep.On("DeleteSection", mock.AnythingOfType("int")).
			Return(errSection).
			Once()
		service := section.NewService(mockRep)
		err := service.DeleteSection(20)
		assert.Error(t, err)
		assert.Equal(t, errSection, err)
	})
}
*/