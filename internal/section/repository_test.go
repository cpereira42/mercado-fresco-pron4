package section

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryListarSectionAll(t *testing.T) {
	t.Run("should return a valid product list", func(t *testing.T) {
		//fileStore := store.New(store.FileType, "")

		sectionList := []Section{
			{
				SectionNumber: 3,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WareHouseId: 78,
				ProductTypeId: 456,
			}, {
				SectionNumber: 313,
				CurrentTemperature: 745,
				MinimumTemperature: 344,
				CurrentCapacity: 1345,
				MinimumCapacity: 243,
				MaximumCapacity: 43456,
				WareHouseId: 784,
				ProductTypeId: 43456,
			},
		}

		sectionListInJson, _ := json.Marshal(sectionList)

		fileStore := store.MockStore{
			ReadMock: func (data interface{}) error {
				return json.Unmarshal(sectionListInJson, data)
			},
			WriteMock: func(data interface{}) error {
				return nil
			},
		} 
		repository := NewRepository(&fileStore)

		result, _ := repository.ListarSectionAll()

		assert.Equal(t, result, sectionList, "sould be equal")
	})

	t.Run("should return err when Store returns an error", func(t *testing.T) {
		
		expectedErr := errors.New("error on connect store/ database")
		 
		fileStoreMock := store.MockStore{
			WriteMock: func(data interface{}) error {
				return nil
			},
			ReadMock: func (data interface{}) error  {
				return expectedErr
			},
		}

		repository := NewRepository(&fileStoreMock)

		_, err := repository.ListarSectionAll()

		assert.Equal(t, err, expectedErr, "sould be equal")
	})
}

func TestRepositoryListarSectionOne(t *testing.T){

	t.Run("buscar por um section no db", func (t *testing.T)  {
		sectionList := []Section{
			{
				Id: 1,
				SectionNumber: 3,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WareHouseId: 78,
				ProductTypeId: 456,
			}, {
				Id: 2,
				SectionNumber: 313,
				CurrentTemperature: 745,
				MinimumTemperature: 344,
				CurrentCapacity: 1345,
				MinimumCapacity: 243,
				MaximumCapacity: 43456,
				WareHouseId: 784,
				ProductTypeId: 43456,
			},
		}	 
		sectionByte, _ := json.Marshal(sectionList)


		mockStore := store.MockStore{
			ReadMock: func (data interface{}) error  {
				return json.Unmarshal(sectionByte, data)
			},
			WriteMock: func(data interface{}) error {
				return nil
			},
		}
		repository := NewRepository(&mockStore)

		result, _ := repository.ListarSectionOne(1)

		assert.Equal(t, result, sectionList[0])
	})
	
	t.Run("buscar por um section passando um id invalido", func (t *testing.T)  {
		sectionList := []Section{
			{
				Id: 1,
				SectionNumber: 3,
				CurrentTemperature: 79845,
				MinimumTemperature: 4,
				CurrentCapacity: 135,
				MinimumCapacity: 23,
				MaximumCapacity: 456,
				WareHouseId: 78,
				ProductTypeId: 456,
			}, {
				Id: 2,
				SectionNumber: 313,
				CurrentTemperature: 745,
				MinimumTemperature: 344,
				CurrentCapacity: 1345,
				MinimumCapacity: 243,
				MaximumCapacity: 43456,
				WareHouseId: 784,
				ProductTypeId: 43456,
			},
		}	 
		sectionByte, _ := json.Marshal(sectionList)

		errorExpect := fmt.Errorf("Section is not registered")

		mockStore := store.MockStore{
			ReadMock: func (data interface{}) error  {
				return json.Unmarshal(sectionByte, data)
			},
			WriteMock: func(data interface{}) error {
				return nil
			},
		}
		repository := NewRepository(&mockStore)

		_, err := repository.ListarSectionOne(10)

		assert.Equal(t, err.Error(), errorExpect.Error())
	})

}

func TestRepositoryCreateSection(t *testing.T) {
	reqSection := Section{
		Id: 1,
		SectionNumber: 5,
		CurrentTemperature: 7985,
		MinimumTemperature: 4,
		CurrentCapacity: 135,
		MinimumCapacity: 23,
		MaximumCapacity: 456,
		WareHouseId: 78,
		ProductTypeId: 456,
	}
	var sectionList []Section  

	t.Run("store success", func(t *testing.T) {
		dataInByte, _ := json.Marshal(sectionList)

		mockSection := store.MockStore{
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(dataInByte, &data)
			},
			WriteMock: func(data interface{}) error {
				_, err := json.Marshal(data)
				return err
			},
		}

		repo := NewRepository(&mockSection)
		
		novoSection, _ := repo.CreateSection(reqSection) 		  
		assert.Equal(t, reqSection, novoSection) 
	})
}


func TestUpdateSection(t *testing.T) {
	 
	reqSection := Section{
		Id: 1,
		SectionNumber: 5444,
		CurrentTemperature: 564,
		MinimumTemperature: 22,
		CurrentCapacity: 111,
		MaximumCapacity: 99,
		WareHouseId: 6756,
		ProductTypeId: 4444,
	}
	reqSection2 := Section{
		Id: 1,
		SectionNumber: 5444,
		CurrentTemperature: 564,
		MinimumTemperature: 22,
		CurrentCapacity: 111,
		MinimumCapacity: 888,
		MaximumCapacity: 99,
		WareHouseId: 6756,
		ProductTypeId: 4444,
	}

	t.Run("Update Section", func(t *testing.T) {
		// cria o array de sections
		var sectionsList []Section
		sectionsList = append(sectionsList, Section{
				Id: 1,
				SectionNumber: 5444,
				CurrentTemperature: 564,
				MinimumTemperature: 22,
				CurrentCapacity: 111,
				MinimumCapacity: 888,
				MaximumCapacity: 99,
				WareHouseId: 123,
				ProductTypeId: 4444,
			},		
		)
		// criar o array de byte
		dataInByte, _ := json.Marshal(sectionsList)
		// cria o mock com as funções mockada, Read(), Write()
		mockStore := store.MockStore{
			WriteMock: func(data interface{}) error {
				return json.Unmarshal(dataInByte, &data)
			},
			ReadMock: func(data interface{}) error {
				err := json.Unmarshal(dataInByte, data)
				return err
			},
		}
		// criar o repository
		repository := NewRepository(&mockStore)
		// chama o metodo a ser testado
		resultado, _ := repository.UpdateSection(1, reqSection)
		//verificar resultado retornado com o resultado esperado
		assert.Equal(t, reqSection2, resultado)
	})

}

func TestRepositoryDelete(t *testing.T){

	id := 1

	sectionsList := []Section{
		{
			Id: 1,
			SectionNumber: 5444,
			CurrentTemperature: 564,
			MinimumTemperature: 22,
			CurrentCapacity: 111,
			MinimumCapacity: 888,
			MaximumCapacity: 99,
			WareHouseId: 6756,
			ProductTypeId: 4444,
		},
	}
	dataInByte, _ := json.Marshal(sectionsList)
	mockStore := store.MockStore{
		WriteMock: func(data interface{}) error {
			return json.Unmarshal(dataInByte, &data)
		},
		ReadMock: func(data interface{}) error {
			return json.Unmarshal(dataInByte, data)
		},
	}

	repository := NewRepository(&mockStore)
	result := repository.DeleteSection(id)
	assert.NoError(t, result)
}
