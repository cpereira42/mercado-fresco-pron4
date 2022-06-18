package buyer_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer/mock"
	"github.com/stretchr/testify/assert"

	tmock "github.com/stretchr/testify/mock"
)


func TestServiceCreate(t *testing.T) {
	var mockRep = new(mock.BuyerRepository)	

	buyerObj := buyer.Buyer{
		Card_number_ID: "12",
		First_name:     "Jose",
		Last_name:      "Silva",
	}
	buyerObj2 := buyer.Buyer{
		Card_number_ID: "13",
		First_name:     "Maria",
		Last_name:      "Jose",
	}
	bListSuccess := []buyer.Buyer{ 
		buyerObj2,
	}
	bListError := []buyer.Buyer{ 
		buyerObj,
	}

	t.Run(
		"Se contiveros campos necessários, será criado",
		func(t *testing.T) {
			mockRep.On("LastID").Return(1, nil).Once()
			mockRep.On("GetAll").Return(bListSuccess, nil).Once()
			mockRep.On("Create", 
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(buyerObj, nil).
				Once()
			
			service := buyer.NewService(mockRep)

			newBuyer, err := service.Create(buyerObj.Card_number_ID,buyerObj.First_name, buyerObj.Last_name)

			assert.NoError(t, err)

			assert.Equal(t, "Jose", newBuyer.First_name) 
			
			mockRep.AssertExpectations(t)
		},
	)

	t.Run(
		"Se o card_number_ID já está registrado, o buyer nao deverá ser criado",
		func(t *testing.T) {
			msgError := fmt.Errorf("a buyer with id 12, already exists")
			mockRep.On("LastID").Return(1,nil).Once()
			mockRep.On("GetAll").Return(bListError, nil).Once()
			mockRep.On("Create", 
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(buyer.Buyer{}, msgError).Once()
			
			service := buyer.NewService(mockRep)

			buyerExpected, err := service.Create(
				buyerObj2.Card_number_ID,
				buyerObj2.First_name, 
				buyerObj2.Last_name)

			assert.Error(t, err)

			assert.ObjectsAreEqual(buyer.Buyer{}, buyerExpected)

			mockRep.AssertExpectations(t)
		},
	)
}

func TestServiceGetAll(t *testing.T) {
	b := buyer.Buyer{
		ID:             1,
		Card_number_ID: "123",
		First_name:     "S",
		Last_name:      "abc",
	}
	var mockRep = new(mock.BuyerRepository)

	var buyerList []buyer.Buyer = make([]buyer.Buyer, 0)
	buyerList = append(buyerList, b)

	t.Run("listar todos buyers, caso de error", func(t *testing.T) {
		mockRep.On("GetAll").Return(buyerList, nil).Once()

		service := buyer.NewService(mockRep)

		bList, _ := service.GetAll()

		assert.Equal(t, bList, buyerList)

		assert.Equal(t, bList[0].Card_number_ID, buyerList[0].Card_number_ID)

		mockRep.AssertExpectations(t)
	})
}

func TestServiceGetId(t *testing.T) {
	buyerObj := buyer.Buyer{
		ID:             1,
		Card_number_ID: "123",
		First_name:    "S",
		Last_name:     "abc",
	}
	t.Run("service metodo GetId", func(t *testing.T) {
		var mockRep = new(mock.BuyerRepository)
		
		mockRep.On("GetId", tmock.AnythingOfType("int")).
			Return(buyerObj, nil).
			Once()

		service := buyer.NewService(mockRep)

		b, err := service.GetId(1)

		assert.NoError(t, err)
		assert.ObjectsAreEqual(b, buyerObj)

		mockRep.AssertExpectations(t)
	})

	t.Run("service metodo GetId, useCase error", func(t *testing.T) {
		
		var mockRep = new(mock.BuyerRepository)
		
		mockRep.On("GetId", 
			tmock.AnythingOfType("int")).
			Return(buyer.Buyer{}, errors.New("falha ao buscar um buyer")).
			Once()

		service := buyer.NewService(mockRep)

		bList, err := service.GetId(1)

		assert.NotNil(t, bList)

		assert.Error(t, err)

		mockRep.AssertExpectations(t)
	})
}


func TestServiceUpdate(t *testing.T) {
	buyerUpdate := buyer.Buyer{
		ID: 1,
		Card_number_ID: "45",
		First_name: "Marta",
		Last_name: "Lima",
	}
	t.Run("useCase Metodo Updade do service", func(t *testing.T) {
		var mockRep = new(mock.BuyerRepository)

		mockRep.On("Update", 
			tmock.AnythingOfType("int"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			).
			Return(buyerUpdate, nil).
			Once()

		service := buyer.NewService(mockRep)

		result, err := service.Update(1,buyerUpdate.Card_number_ID, buyerUpdate.First_name, buyerUpdate.Last_name)

		assert.NoError(t, err)
		assert.ObjectsAreEqual(result, buyerUpdate)

		mockRep.AssertExpectations(t)
	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("metodo do service Delete", func(t *testing.T) {
		id := 1

		var mockRep = new(mock.BuyerRepository)

		mockRep.On("Delete", tmock.AnythingOfType("int")).
			Return(nil).
			Once()

			service := buyer.NewService(mockRep)

			result := service.Delete(id)

			assert.NoError(t, result)

			mockRep.AssertExpectations(t)
	})

	t.Run("tente delete um recurso buyer com id invalido", func(t *testing.T) {
		id := 10

		var mockRep = new(mock.BuyerRepository)
		
		mockRep.On("Delete", tmock.AnythingOfType("int")).
			Return(errors.New("buyer 10 não esta registrado")).
			Once()

			service := buyer.NewService(mockRep)

			result := service.Delete(id)

			assert.Error(t, result)

			mockRep.AssertExpectations(t)
	})
}

