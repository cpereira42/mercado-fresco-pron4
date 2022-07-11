package productbatch_test

import (
	"errors"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreatePB(t *testing.T) {
	t.Run("sucesso", func(t *testing.T) {
		mockRepo := new(mocks.RepositoryProductBatches)
		mockRepo.On("CreatePB", mock.Anything).Return(productBatches, nil).Once()
		service := productbatch.NewServiceProductBatches(mockRepo)
		result, err := service.CreatePB(productBatches)
		assert.NoError(t, err)
		assert.Equal(t, productBatches, result)
	})
	t.Run("error", func(t *testing.T) {
		mockRepo := new(mocks.RepositoryProductBatches)

		messageError := errors.New("batch_number_UNIQUE is unique, and 111 already registered")
		mockRepo.On("CreatePB", mock.Anything).Return(productBatches, messageError).Once()
		service := productbatch.NewServiceProductBatches(mockRepo)
		result, err := service.CreatePB(productBatches)
		assert.Error(t, err)
		assert.Equal(t, productBatches, result)
	})
}

func TestServiceGetAll(t *testing.T) {
	t.Run("sucesso", func(t *testing.T) {
		mockRepo := new(mocks.RepositoryProductBatches)

		mockRepo.On("GetAll", mock.Anything).Return(productBatchesList, nil).Once()
		service := productbatch.NewServiceProductBatches(mockRepo)
		result, err := service.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, productBatchesList, result)
	})
	t.Run("error", func(t *testing.T) {
		mockRepo := new(mocks.RepositoryProductBatches)
		messageError := errors.New("query sql invalid")
		mockRepo.On("GetAll", mock.Anything).Return([]productbatch.ProductBatchesResponse{}, messageError).Once()
		service := productbatch.NewServiceProductBatches(mockRepo)
		result, err := service.GetAll()
		assert.Error(t, err)
		assert.Equal(t, []productbatch.ProductBatchesResponse{}, result)
	})
}

func TestServiceGetId(t *testing.T) {
	t.Run("sucesso", func(t *testing.T) {
		mockRepo := new(mocks.RepositoryProductBatches)
		mockRepo.On("GetId", mock.Anything).Return(productBatchesRes, nil).Once()
		service := productbatch.NewServiceProductBatches(mockRepo)
		paramId := 1
		result, err := service.GetId(int64(paramId))
		assert.NoError(t, err)
		assert.Equal(t, productBatchesRes, result)
	})
	t.Run("error", func(t *testing.T) {
		mockRepo := new(mocks.RepositoryProductBatches)
		messageError := errors.New("section_id not found")
		mockRepo.On("GetId", mock.Anything).Return(productbatch.ProductBatchesResponse{}, messageError).Once()
		service := productbatch.NewServiceProductBatches(mockRepo)
		paramId := 1
		result, err := service.GetId(int64(paramId))
		assert.Error(t, err)
		assert.Equal(t, productbatch.ProductBatchesResponse{}, result)
	})
}
