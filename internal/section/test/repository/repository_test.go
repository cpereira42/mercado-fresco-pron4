package section

import (
	"encoding/json"
	"errors"
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
	}, {
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


func TestRepositoryListarSectionAll(t *testing.T) {
	t.Run("should return a valid product list", func(t *testing.T) {
		//fileStore := store.New(store.FileType, "")
 
		sectionListInJson, _ := json.Marshal(sectionList)

		fileStore := store.MockStore{
			ReadMock: func (data interface{}) error {
				return json.Unmarshal(sectionListInJson, data)
			},
			WriteMock: func(data interface{}) error {
				return nil
			},
		} 
		repository := section.NewRepository(&fileStore)

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

		repository := section.NewRepository(&fileStoreMock)

		_, err := repository.ListarSectionAll()

		assert.Equal(t, err, expectedErr, "sould be equal")
	})
}

func TestRepositoryListarSectionOne(t *testing.T){

	t.Run("buscar por um section no db", func (t *testing.T)  {
		  
		sectionByte, _ := json.Marshal(sectionList)


		mockStore := store.MockStore{
			ReadMock: func (data interface{}) error  {
				return json.Unmarshal(sectionByte, data)
			},
			WriteMock: func(data interface{}) error {
				return nil
			},
		}
		repository := section.NewRepository(&mockStore)

		result, _ := repository.ListarSectionOne(1)

		assert.Equal(t, result, sectionList[0])
	})
	
	t.Run("buscar por um section passando um id invalido", func (t *testing.T)  {
		  
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
		repository := section.NewRepository(&mockStore)

		_, err := repository.ListarSectionOne(10)

		assert.Equal(t, err.Error(), errorExpect.Error())
	})

}

func TestRepositoryCreateSection(t *testing.T) {
	reqSection := section.Section{
		SectionNumber: 5,
		CurrentTemperature: 7985,
		MinimumTemperature: 4,
		CurrentCapacity: 135,
		MinimumCapacity: 23,
		MaximumCapacity: 456,
		WarehouseId: 78,
		ProductTypeId: 456,
	}
	var sectionList []section.Section  

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

		repo := section.NewRepository(&mockSection)
		
		novoSection, _ := repo.CreateSection(reqSection) 		  
		assert.Equal(t, reqSection, novoSection) 
	})
}


func TestUpdateSection(t *testing.T) {
	 
	reqSection := section.Section{ 
		SectionNumber: 3,
		CurrentTemperature: 79845,
		MinimumTemperature: 4,
		CurrentCapacity: 135,
		MinimumCapacity: 23,
		MaximumCapacity: 456,
		WarehouseId: 78,
		ProductTypeId: 12,
	} 

	t.Run("Update section.Section", func(t *testing.T) {
		 
		// criar o array de byte
		dataInByte, _ := json.Marshal(sectionList)
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
		repository := section.NewRepository(&mockStore)
		// chama o metodo a ser testado
		resultado, _ := repository.UpdateSection(1, reqSection)
		//verificar resultado retornado com o resultado esperado
		assert.Equal(t, reqSection, resultado)
	})

}

func TestRepositoryDelete(t *testing.T){

	id := 1

	sectionsList := []section.Section{
		{
			Id: 1,
			SectionNumber: 5444,
			CurrentTemperature: 564,
			MinimumTemperature: 22,
			CurrentCapacity: 111,
			MinimumCapacity: 888,
			MaximumCapacity: 99,
			WarehouseId: 6756,
			ProductTypeId: 4444,
		},
	}
	t.Run("delete one section", func(t *testing.T) {
		dataInByte, _ := json.Marshal(sectionsList)
		mockStore := store.MockStore{
			WriteMock: func(data interface{}) error {
				return json.Unmarshal(dataInByte, &data)
			},
			ReadMock: func(data interface{}) error {
				return json.Unmarshal(dataInByte, data)
			},
		}
	
		repository := section.NewRepository(&mockStore)
		result := repository.DeleteSection(id)
		assert.NoError(t, result)
	})
}
