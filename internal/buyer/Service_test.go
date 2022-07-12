package buyer_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer/mocks"
	"github.com/stretchr/testify/assert"

	tmock "github.com/stretchr/testify/mock"
)

func TestServiceCreateBuyer(t *testing.T) {
	var mockRep = new(mocks.Repository)

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
	/*bListSuccess := []buyer.Buyer{
		buyerObj2,
	}*/

	t.Run(
		"create_ok",
		func(t *testing.T) {
			/*mockRep.On("LastID").Return(1, nil).Once()
			mockRep.On("GetAll").Return(bListSuccess, nil).Once()*/
			mockRep.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(buyerObj, nil).
				Once()

			service := buyer.NewService(mockRep)

			newBuyer, err := service.Create(buyerObj.Card_number_ID, buyerObj.First_name, buyerObj.Last_name)

			assert.NoError(t, err)

			assert.Equal(t, "Jose", newBuyer.First_name)

			mockRep.AssertExpectations(t)
		},
	)

	t.Run(
		"create_fail_lastID",
		func(t *testing.T) {
			msgError := fmt.Errorf("buyer with id %s, not found", buyerObj.Card_number_ID)
			/*mockRep.On("LastID").Return(0, errors.New("error to get last id")).Once()
			mockRep.On("GetAll").Return(bListSuccess, nil).Maybe()*/
			mockRep.On("Create",
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(buyer.Buyer{}, msgError).
				Maybe()
			service := buyer.NewService(mockRep)

			_, err := service.Create("500", "nome", "sobrenome")

			assert.Error(t, err)

			mockRep.AssertExpectations(t)
		},
	)

	t.Run(
		"create_conflict",
		func(t *testing.T) {
			msgError := fmt.Errorf("a buyer with id %s, already exists", buyerObj.Card_number_ID)
			/*mockRep.On("LastID").Return(1, nil).Once()
			mockRep.On("GetAll").Return(bListSuccess, nil).Maybe()*/
			mockRep.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(buyer.Buyer{}, msgError).Maybe()

			service := buyer.NewService(mockRep)

			buyerExpected, err := service.Create(
				buyerObj2.Card_number_ID,
				buyerObj2.First_name,
				buyerObj2.Last_name)

			assert.Error(t, err)
			assert.ObjectsAreEqual(buyer.Buyer{}, buyerExpected)
			//assert.Equal(t, msgError, err)
			mockRep.AssertExpectations(t)
		},
	)
}

func TestServiceGetAllBuyer(t *testing.T) {
	b := buyer.Buyer{
		ID:             1,
		Card_number_ID: "123",
		First_name:     "S",
		Last_name:      "abc",
	}
	var mockRep = new(mocks.Repository)

	var buyerList []buyer.Buyer = make([]buyer.Buyer, 0)
	buyerList = append(buyerList, b)

	t.Run("find_all", func(t *testing.T) {
		mockRep.On("GetAll").Return(buyerList, nil).Once()

		service := buyer.NewService(mockRep)

		bList, _ := service.GetAll()

		assert.Equal(t, bList, buyerList)

		assert.Equal(t, bList[0].Card_number_ID, buyerList[0].Card_number_ID)

		mockRep.AssertExpectations(t)
	})

	t.Run("find_all_fail", func(t *testing.T) {
		mockRep.On("GetAll").Return(nil, errors.New("Fail to find all")).Once()

		service := buyer.NewService(mockRep)

		_, err := service.GetAll()

		assert.Error(t, err)
		mockRep.AssertExpectations(t)
	})
}

func TestServiceGetIdBuyer(t *testing.T) {
	buyerObj := buyer.Buyer{
		ID:             1,
		Card_number_ID: "123",
		First_name:     "S",
		Last_name:      "abc",
	}
	t.Run("find_by_id_existent", func(t *testing.T) {
		var mockRep = new(mocks.Repository)

		mockRep.On("GetId", tmock.AnythingOfType("int")).
			Return(buyerObj, nil).
			Once()

		service := buyer.NewService(mockRep)

		b, err := service.GetId(1)

		assert.NoError(t, err)
		assert.ObjectsAreEqual(b, buyerObj)

		mockRep.AssertExpectations(t)
	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {

		var mockRep = new(mocks.Repository)

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

func TestServiceUpdateBuyer(t *testing.T) {
	buyerUpdate := buyer.Buyer{
		ID:             1,
		Card_number_ID: "45",
		First_name:     "Marta",
		Last_name:      "Lima",
	}
	t.Run("update_existent", func(t *testing.T) {
		var mockRep = new(mocks.Repository)

		mockRep.On("Update",
			tmock.AnythingOfType("int"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
		).
			Return(buyerUpdate, nil).
			Once()

		service := buyer.NewService(mockRep)

		result, err := service.Update(1, buyerUpdate.Card_number_ID, buyerUpdate.First_name, buyerUpdate.Last_name)

		assert.NoError(t, err)
		assert.ObjectsAreEqual(result, buyerUpdate)

		mockRep.AssertExpectations(t)
	})

	t.Run("update_non_existent", func(t *testing.T) {
		var mockRep = new(mocks.Repository)

		mockRep.On("Update",
			tmock.AnythingOfType("int"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
			tmock.AnythingOfType("string"),
		).
			Return(buyerUpdate, nil).
			Once()

		service := buyer.NewService(mockRep)

		result, err := service.Update(9, buyerUpdate.Card_number_ID, buyerUpdate.First_name, buyerUpdate.Last_name)

		assert.NoError(t, err)
		assert.ObjectsAreEqual(result, buyerUpdate)

		mockRep.AssertExpectations(t)
	})
}

func TestServiceDeleteBuyer(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		id := 1

		var mockRep = new(mocks.Repository)

		mockRep.On("Delete", tmock.AnythingOfType("int")).
			Return(nil).
			Once()

		service := buyer.NewService(mockRep)

		result := service.Delete(id)

		assert.NoError(t, result)

		mockRep.AssertExpectations(t)
	})

	t.Run("delete_non_existent", func(t *testing.T) {
		id := 10

		var mockRep = new(mocks.Repository)

		mockRep.On("Delete", tmock.AnythingOfType("int")).
			Return(errors.New("buyer 10 não está registrado")).
			Once()

		service := buyer.NewService(mockRep)

		result := service.Delete(id)

		assert.Error(t, result)

		mockRep.AssertExpectations(t)
	})
}
