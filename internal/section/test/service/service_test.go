package service_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/stretchr/testify/assert"
)

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

func TestServiceListarSectionAll(t *testing.T) {

	sectionByte, _ := json.Marshal(sectionList)

	mockStore := store.MockStore{
		ReadMock: func(data interface{}) error {
			return json.Unmarshal(sectionByte, data)
		},
		WriteMock: func(data interface{}) error {
			return nil
		},
	}

	repository := section.NewRepository(&mockStore)

	service := section.NewService(repository)

	t.Run("get all sections success", func(t *testing.T) {

		resultado, _ := service.ListarSectionAll()

		assert.Equal(t, resultado, sectionList)
	})
}

func TestServiceListarSectionOne(t *testing.T) {

	sectionByte, _ := json.Marshal(sectionList)

	mockStore := store.MockStore{
		ReadMock: func(data interface{}) error {
			return json.Unmarshal(sectionByte, data)
		},
		WriteMock: func(data interface{}) error {
			return nil
		},
	}

	sectionOne := sectionList[0]

	repository := section.NewRepository(&mockStore)
	service := section.NewService(repository)
	
	t.Run("test method of ListarSectionOne(id) with return of success", func(t *testing.T) {
		resultado, _ := service.ListarSectionOne(sectionOne.Id)
	
		assert.Equal(t, resultado, sectionOne)
	})
	
	t.Run("test method of ListarSectionOne(id) service with return of error", func(t *testing.T) {
	
		errorExpected := fmt.Errorf("Section is not registered")
	
		_, errRetornado := repository.ListarSectionOne(50)
	
		assert.Equal(t, errRetornado.Error(), errorExpected.Error())
	})
}

func TestServiceCreate(t *testing.T) {

	newSection := section.SectionRequestCreate{SectionNumber: 1,CurrentTemperature: 1,
		MinimumTemperature: 1, CurrentCapacity: 1, MinimumCapacity: 1,
		MaximumCapacity: 1, WarehouseId: 1,ProductTypeId: 1}

	expected := section.Section{SectionNumber: 1, CurrentTemperature: 1,
		MinimumTemperature: 1, CurrentCapacity: 1, MinimumCapacity: 1, MaximumCapacity: 1,
		WarehouseId: 1, ProductTypeId: 1}

	sectionbyte, _ := json.Marshal(sectionList)

	mockStore := store.MockStore{
		WriteMock: func (data interface{}) error  {
			_, err :=json.Marshal(&data)
			return err
		},
		ReadMock: func(data interface{}) error {
			return json.Unmarshal(sectionbyte, data)
		},
	}
	t.Run("testing method of service CreateSection(newSection)", func(t *testing.T) {
		repository := section.NewRepository(&mockStore)
		service := section.NewService(repository)

		resultado, _ := service.CreateSection(newSection)

		assert.Equal(t, resultado, expected)
	})
}


func TestServiceUpdated(t *testing.T) {
	sectionUp := section.SectionRequestUpdate{
		SectionNumber: 1,
		CurrentTemperature: 2,
		MinimumTemperature: 3, 
		WarehouseId: 7,
		ProductTypeId: 8,
	}
	expectedUp := section.Section{
		SectionNumber: 1,
		CurrentTemperature: 2,
		MinimumTemperature: 3,
		CurrentCapacity: 135,
		MinimumCapacity: 23,
		MaximumCapacity: 456,
		WarehouseId: 7,
		ProductTypeId: 8,
	}
	
	sectionByte, _ := json.Marshal(sectionList)

	mockStore := store.MockStore{
		ReadMock: func(data interface{}) error {
			return json.Unmarshal(sectionByte, data)
		},
		WriteMock: func(data interface{}) error {
			_, err := json.Marshal(data)
			return err
		},
	}

	service := section.NewService(section.NewRepository(&mockStore))

	resultado, _ := service.UpdateSection(1, sectionUp)

	t.Run("testing method of service UpdateSection(id, sectionUp), success", func(t *testing.T) {
		assert.Equal(t, resultado, expectedUp)
	})
}

func TestServiceDelete(t *testing.T) {
	id := 1

	sectionByte, _ := json.Marshal(sectionList)

	mockStore := store.MockStore{
		ReadMock: func(data interface{}) error {
			return json.Unmarshal(sectionByte, data)
		},
		WriteMock: func(data interface{}) error {
			_, err := json.Marshal(data)
			return err
		},
	}

	service := section.NewService(section.NewRepository(&mockStore))

	resultado := service.DeleteSection(id)

	t.Run("testing method of service DeleteSection(id), success", func(t *testing.T) {
		assert.Equal(t, resultado, nil)
		assert.NoError(t, resultado)
	})

	t.Run("testing method of service DeleteSection(1*5), error", func(t *testing.T) {
		
		resultado = service.DeleteSection(id*5)
		
		expectedErr := fmt.Errorf("section is not registered")

		assert.Equal(t, resultado.Error(), expectedErr.Error())
		assert.Error(t, resultado)
	})
}

