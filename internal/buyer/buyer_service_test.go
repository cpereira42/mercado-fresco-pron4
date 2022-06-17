package buyer_test

import (
 	 
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer/mock"
	"github.com/stretchr/testify/assert"
	
	tmock "github.com/stretchr/testify/mock"
)


func TestServiceCreate(t *testing.T) {
	
	var mockRep = new(mock.BuyerRepository)

	buyerObj := buyer.Buyer{
		ID:             1,
		Card_number_ID: "12",
		First_name:     "Jose",
		Last_name:      "Silva",
	} 
	bList := []buyer.Buyer{ 
		buyerObj,
	}

	t.Run(
		"Se contiveros campos necessários, será criado",
		func(t *testing.T) {
			mockRep.On("LastID").Return(1, nil).Once()
			mockRep.On("GetAll").Return(bList, nil).Once()
			mockRep.On("Create", 
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string") ).Return(buyerObj, nil).Once()
			
			service := buyer.NewService(mockRep)

			newBuyer, _ := service.Create(buyerObj.Card_number_ID,buyerObj.First_name, buyerObj.Last_name)

			assert.ObjectsAreEqual(buyerObj, newBuyer) 
			
			mockRep.AssertExpectations(t)

		},
	)
	/*
	t.Run(
		"Se o card_number_ID já está registrado, o buyer nao deverá ser criado",
		func(t *testing.T) {
			msgError := fmt.Errorf("a buyer with id 12, already exists")
			
			mockRep.On("Create", tmock.AnythingOfType("string"),tmock.AnythingOfType("string"),tmock.AnythingOfType("string")).Return(buyer.Buyer{}, msgError)
			service := buyer.NewService(mockRep)

			newBuyer, err := service.Create(buyerObj.Card_number_ID,
				buyerObj.First_name, buyerObj.Last_name)

			assert.Error(t, err)

			// assert.Equal(t, err, msgError)
			assert.ObjectsAreEqual(buyer.Buyer{}, newBuyer)

			mockRep.AssertExpectations(t)
		},
	)
	*/

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

		// assert.NoError(t, err)
		assert.Equal(t, bList, buyerList)
		assert.Equal(t, bList[0].Card_number_ID, buyerList[0].Card_number_ID)

		mockRep.AssertExpectations(t)
	})
}

func TestServiceGetId(t *testing.T) {
	t.Run("service metodo GetId", func(t *testing.T) {
		var buyerList []buyer.Buyer = []buyer.Buyer{
			{
				ID:             1,
				Card_number_ID: "123",
				First_name:     "S",
				Last_name:      "abc",
			},
			{
				ID:             2,
				Card_number_ID: "78",
				First_name:     "B",
				Last_name:      "dfc",
			},
		}
		
		var mockRep = new(mock.BuyerRepository)
		
		mockRep.On("GetId", tmock.Arguments{1}).Return(buyerList[0], nil).Once()

		service := buyer.NewService(mockRep)

		b, _ := service.GetId(1)
		assert.ObjectsAreEqual(b, buyerList[0])
		mockRep.AssertExpectations(t)
	})
}
