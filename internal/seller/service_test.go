package seller_test

import (
	"fmt"
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
		seller1,
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
				Return(seller1, nil)
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
			mockRespository.On("LastID").Return(1, nil).Maybe()
			mockRespository.On("GetAll").Return(sListError, nil).Once()
			mockRespository.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(seller.Seller{}, msgError)

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
		},
	)
	t.Run(
		"If the application could not connect to the DB to get all sellers - CheckCid",
		func(t *testing.T) {
			msgError := fmt.Errorf("Could not read file")
			mockRespository.On("LastID").Return(1, nil)
			mockRespository.On("GetAll").Return(nil, msgError).Once()
			mockRespository.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(seller.Seller{}, nil)

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
		},
	)
	// t.Run(
	// 	"If the application could not connect to the DB - LastID",
	// 	func(t *testing.T) {
	// 		msgError := fmt.Errorf("Could not read file")
	// 		mockRespository.On("GetAll").Return(sListSuccess, nil)
	// 		mockRespository.On("LastID").Return(0, msgError)
	// 		mockRespository.On("Create",
	// 			tmock.AnythingOfType("int"),
	// 			tmock.AnythingOfType("int"),
	// 			tmock.AnythingOfType("string"),
	// 			tmock.AnythingOfType("string"),
	// 			tmock.AnythingOfType("string")).
	// 			Return(seller.Seller{}, nil).Maybe()
	// 		service := seller.NewService(mockRespository)

	// 		newSeller, err := service.Create(
	// 			210,
	// 			seller1.CompanyName,
	// 			seller1.Adress,
	// 			seller1.Telephone,
	// 		)

	// 		assert.Error(t, err)

	// 		assert.EqualError(t, err, msgError.Error())

	// 		assert.ObjectsAreEqual(seller.Seller{}, newSeller)

	// 		mockRespository.AssertExpectations(t)
	// 	},
	// )
	// t.Run(
	// 	"If the application could not connect to the DB - Create",
	// 	func(t *testing.T) {
	// 		msgError := fmt.Errorf("Could not read file")
	// 		mockRespository.On("GetAll").Return(sListSuccess, nil).Maybe()
	// 		mockRespository.On("LastID").Return(1, nil)
	// 		mockRespository.On("Create",
	// 			tmock.AnythingOfType("int"),
	// 			tmock.AnythingOfType("int"),
	// 			tmock.AnythingOfType("string"),
	// 			tmock.AnythingOfType("string"),
	// 			tmock.AnythingOfType("string")).
	// 			Return(seller.Seller{}, msgError).Maybe()
	// 		service := seller.NewService(mockRespository)

	// 		newSeller, err := service.Create(
	// 			210,
	// 			seller1.CompanyName,
	// 			seller1.Adress,
	// 			seller1.Telephone,
	// 		)

	// 		assert.Error(t, err)

	// 		assert.EqualError(t, err, msgError.Error())

	// 		assert.ObjectsAreEqual(seller.Seller{}, newSeller)

	// 		mockRespository.AssertExpectations(t)
	// 	},
	// )
}

func TestServiceGetAll(t *testing.T) {
	mockRespository := new(mocks.RepositorySeller)
	seller1 := seller.Seller{Id: 1, Cid: 200, CompanyName: "MELI", Adress: "Rua B", Telephone: "9999-8888"}
	seller2 := seller.Seller{8, 201, "Digital House", "Avenida Brasil", "7777-5555"}

	sListSuccess := []seller.Seller{
		seller1, seller2,
	}
	t.Run(
		"Receives data from Repository GetAll",
		func(t *testing.T) {
			mockRespository.On("GetAll").Return(sListSuccess, nil).Once()
			service := seller.NewService(mockRespository)
			sellerList, err := service.GetAll()
			assert.NoError(t, err)
			assert.Equal(t, sListSuccess, sellerList)
			mockRespository.AssertExpectations(t)
		},
	)
	t.Run(
		"Fails to receive data from Repository GetAll",
		func(t *testing.T) {
			msgError := fmt.Errorf("Could not read file")
			mockRespository.On("GetAll").Return(nil, msgError).Once()
			service := seller.NewService(mockRespository)
			_, err := service.GetAll()
			assert.Error(t, err)

			assert.EqualError(t, err, msgError.Error())

			mockRespository.AssertExpectations(t)
		},
	)
}

func TestServiceGetId(t *testing.T) {
	mockRespository := new(mocks.RepositorySeller)
	seller1 := seller.Seller{Id: 1, Cid: 200, CompanyName: "MELI", Adress: "Rua B", Telephone: "9999-8888"}
	// seller2 := seller.Seller{8, 201, "Digital House", "Avenida Brasil", "7777-5555"}

	t.Run(
		"Receives data with found Id from Repository GetId",
		func(t *testing.T) {
			mockRespository.On("GetId", tmock.AnythingOfType("int")).Return(seller1, nil).Once()
			service := seller.NewService(mockRespository)
			sellerFound, err := service.GetId(1)
			assert.NoError(t, err)
			assert.Equal(t, seller1, sellerFound)
			mockRespository.AssertExpectations(t)
		},
	)
	t.Run(
		"Tests ID not found error",
		func(t *testing.T) {
			msgError := fmt.Errorf("Seller 2 not found")
			mockRespository.On("GetId", tmock.AnythingOfType("int")).Return(seller.Seller{}, msgError)
			service := seller.NewService(mockRespository)
			_, err := service.GetId(2)
			assert.Error(t, err)

			assert.EqualError(t, err, msgError.Error())

			mockRespository.AssertExpectations(t)
		},
	)
}

func TestServiceUpdate(t *testing.T) {
	mockRespository := new(mocks.RepositorySeller)
	// seller1 := seller.Seller{Id: 1, Cid: 200, CompanyName: "MELI", Adress: "Rua B", Telephone: "9999-8888"}
	seller2 := seller.Seller{8, 201, "Digital House", "Avenida Brasil", "7777-5555"}

	t.Run(
		"Success updating seller",
		func(t *testing.T) {
			mockRespository.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(seller2, nil).Once()
			service := seller.NewService(mockRespository)
			updatedSeller, err := service.Update(8, 201, "Digital House", "Avenida Brasil", "7777-5555")
			assert.NoError(t, err)
			assert.Equal(t, seller2, updatedSeller)
			mockRespository.AssertExpectations(t)
		},
	)
	t.Run(
		"Tests ID not found error",
		func(t *testing.T) {
			msgError := fmt.Errorf("Seller 8 not found")
			mockRespository.On("Update",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(seller.Seller{}, msgError).Once()
			service := seller.NewService(mockRespository)
			_, err := service.Update(8, 201, "Digital House", "Avenida Brasil", "7777-5555")
			assert.Error(t, err)

			assert.EqualError(t, err, msgError.Error())

			mockRespository.AssertExpectations(t)
		},
	)
}

func TestServiceDelete(t *testing.T) {
	mockRespository := new(mocks.RepositorySeller)
	t.Run(
		"Success deleting seller",
		func(t *testing.T) {
			mockRespository.On("Delete",
				tmock.AnythingOfType("int")).
				Return(nil).Once()
			service := seller.NewService(mockRespository)
			err := service.Delete(8)
			assert.NoError(t, err)
			assert.Equal(t, nil, err)
			mockRespository.AssertExpectations(t)
		},
	)
	t.Run(
		"Tests ID not found error",
		func(t *testing.T) {
			msgError := fmt.Errorf("Seller 8 not found")
			mockRespository.On("Delete",
				tmock.AnythingOfType("int")).
				Return(msgError).Once()
			service := seller.NewService(mockRespository)
			err := service.Delete(8)
			assert.Error(t, err)

			assert.EqualError(t, err, msgError.Error())

			mockRespository.AssertExpectations(t)
		},
	)
}
