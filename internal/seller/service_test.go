package seller_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"
	"github.com/cpereira42/mercado-fresco-pron4/internal/seller/mocks"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

func TestServiceCreate(t *testing.T) {
	mockRespository := new(mocks.RepositorySeller)
	seller1 := seller.Seller{Id: 1, Cid: 200, CompanyName: "MELI", Adress: "Rua B", Telephone: "9999-8888"}
	seller2 := seller.Seller{8, 201, "Digital House", "Avenida Brasil", "7777-5555"}

	sListSuccess := []seller.Seller{
		seller2,
	}
	sListError := []seller.Seller{
		seller1, seller2,
	}
	t.Run(
		"If all necessary fields are complete, a new seller is created",
		func(t *testing.T) {
			mockRespository.On("LastID").Return(1, nil).Once()
			mockRespository.On("GetAll").Return(sListSuccess, nil).Once()
			mockRespository.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(seller1, nil).Once()
			service := seller.NewService(mockRespository)
			newSeller, err := service.Create(seller1.Cid, seller1.Adress, seller1.CompanyName, seller1.Telephone)
			assert.NoError(t, err)
			assert.Equal(t, 200, newSeller.Cid)
			assert.Equal(t, "MELI", newSeller.CompanyName)
			assert.Equal(t, "Rua B", newSeller.Adress)
			assert.Equal(t, "9999-8888", newSeller.Telephone)
			mockRespository.AssertExpectations(t)
		},
	)
	t.Run(
		"If the informed CID is already registered, the seller should not be created",
		func(t *testing.T) {
			msgError := fmt.Errorf("Cid already registered")
			mockRespository.On("LastID").Return(1, nil)
			mockRespository.On("GetAll").Return(sListError, nil)
			mockRespository.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(seller.Seller{}, msgError).Maybe()

			service := seller.NewService(mockRespository)

			newSeller, err := service.Create(
				200,
				seller1.CompanyName,
				seller1.Adress,
				seller1.Telephone,
			)

			assert.Error(t, err)

			assert.EqualError(t, err, msgError.Error())

			assert.ObjectsAreEqual(seller.Seller{}, newSeller)

			mockRespository.AssertExpectations(t)
			log.Println(newSeller, err)
		})
}
