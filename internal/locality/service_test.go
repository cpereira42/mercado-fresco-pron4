package locality_test

import (
	"fmt"
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/locality"
	"github.com/cpereira42/mercado-fresco-pron4/internal/locality/mocks"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

func TestServiceCreate(t *testing.T) {
	mockRespository := new(mocks.RepositoryLocality)
	locality1 := locality.Locality{1000, "Itabaiana", "Sergipe", "Brasil"}

	t.Run(
		"If all necessary fields are complete, a new locality is created",
		func(t *testing.T) {
			mockRespository.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(locality1, nil).Once()
			service := locality.NewService(mockRespository)
			newLocality, err := service.Create(1000, "Itabaiana", "Sergipe", "Brasil")
			assert.NoError(t, err)
			assert.Equal(t, "Itabaiana", newLocality.LocalityName)
			assert.Equal(t, "Sergipe", newLocality.ProvinceName)
			assert.Equal(t, "Brasil", newLocality.CountryName)
			mockRespository.AssertExpectations(t)
		},
	)
	t.Run(
		"If the informed Locality Name is already registered, the locality should not be created",
		func(t *testing.T) {
			msgError := fmt.Errorf("Locality already registered")
			mockRespository.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(locality.Locality{}, msgError).Once()

			service := locality.NewService(mockRespository)

			newLocality, err := service.Create(1000, "Itabaiana", "Sergipe", "Brasil")

			assert.Error(t, err)

			assert.EqualError(t, err, msgError.Error())

			assert.ObjectsAreEqual(locality.Locality{}, newLocality)

			mockRespository.AssertExpectations(t)
		},
	)
	t.Run(
		"If the application could not connect to the DB - Create",
		func(t *testing.T) {
			mockRespository := new(mocks.RepositoryLocality)
			msgError := fmt.Errorf("Could not connect to database")
			mockRespository.On("Create",
				tmock.AnythingOfType("int"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string"),
				tmock.AnythingOfType("string")).
				Return(locality.Locality{}, msgError).Once()
			service := locality.NewService(mockRespository)

			newLocality, err := service.Create(1000, "Itabaiana", "Sergipe", "Brasil")

			assert.Error(t, err)

			assert.EqualError(t, err, msgError.Error())

			assert.ObjectsAreEqual(locality.Locality{}, newLocality)

		},
	)
}

func TestServiceGenerateReportAll(t *testing.T) {
	mockRespository := new(mocks.RepositoryLocality)
	report1 := locality.LocalityReport{1, "Itabaiana", 2}
	report2 := locality.LocalityReport{2, "Aracaju", 4}
	reportList := []locality.LocalityReport{report1, report2}
	t.Run("Generate a report with all localitis",
		func(t *testing.T) {
			mockRespository.On("GenerateReportAll").Return(reportList, nil).Once()
			service := locality.NewService(mockRespository)
			generatedReport, err := service.GenerateReportAll()
			assert.NoError(t, err)
			assert.Equal(t, reportList, generatedReport)
			mockRespository.AssertExpectations(t)
		},
	)
	t.Run(
		"If the application could not connect to the DB - GenerateReportAll",
		func(t *testing.T) {
			mockRespository := new(mocks.RepositoryLocality)
			msgError := fmt.Errorf("Could not connect to database")
			mockRespository.On("GenerateReportAll").Return([]locality.LocalityReport{}, msgError).Once()
			service := locality.NewService(mockRespository)

			generatedReport, err := service.GenerateReportAll()

			assert.Error(t, err)

			assert.EqualError(t, err, msgError.Error())

			assert.ObjectsAreEqual([]locality.LocalityReport{}, generatedReport)

		},
	)
}

func TestServiceGenerateReportById(t *testing.T) {
	mockRespository := new(mocks.RepositoryLocality)
	report1 := locality.LocalityReport{1, "Itabaiana", 2}
	t.Run("Generate a report with a informed locality",
		func(t *testing.T) {
			mockRespository.On("GenerateReportById", 1).Return(report1, nil).Once()
			service := locality.NewService(mockRespository)
			generatedReport, err := service.GenerateReportById(1)
			assert.NoError(t, err)
			assert.Equal(t, report1, generatedReport)
			mockRespository.AssertExpectations(t)
		},
	)
	t.Run(
		"Id not found error - GenerateReportById",
		func(t *testing.T) {
			mockRespository := new(mocks.RepositoryLocality)
			msgError := fmt.Errorf("Locality 2 not found")
			mockRespository.On("GenerateReportById", tmock.AnythingOfType("int")).Return(locality.LocalityReport{}, msgError).Once()
			service := locality.NewService(mockRespository)

			generatedReport, err := service.GenerateReportById(1)

			assert.Error(t, err)

			assert.EqualError(t, err, msgError.Error())

			assert.ObjectsAreEqual(locality.LocalityReport{}, generatedReport)

		},
	)
	t.Run(
		"If the application could not connect to the DB - GenerateReportById",
		func(t *testing.T) {
			mockRespository := new(mocks.RepositoryLocality)
			msgError := fmt.Errorf("Could not connect to database")
			mockRespository.On("GenerateReportById", 1).Return(locality.LocalityReport{}, msgError).Once()
			service := locality.NewService(mockRespository)

			generatedReport, err := service.GenerateReportById(1)

			assert.Error(t, err)

			assert.EqualError(t, err, msgError.Error())

			assert.ObjectsAreEqual(locality.LocalityReport{}, generatedReport)

		},
	)
}
