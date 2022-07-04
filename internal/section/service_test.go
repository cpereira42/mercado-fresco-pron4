package section_test

import (
	"testing"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
)

var sectionList []section.Section = []section.Section{
	{
		Id:                 1,
		SectionNumber:      3,
		CurrentTemperature: 79845,
		MinimumTemperature: 4,
		CurrentCapacity:    135,
		MinimumCapacity:    23,
		MaximumCapacity:    456,
		WarehouseId:        12,
		ProductTypeId:      456,
	}, {
		Id:                 2,
		SectionNumber:      313,
		CurrentTemperature: 745,
		MinimumTemperature: 344,
		CurrentCapacity:    1345,
		MinimumCapacity:    243,
		MaximumCapacity:    43456,
		WarehouseId:        13,
		ProductTypeId:      43456,
	},
}

func TestServiceListarSectionAll(t *testing.T) {
	t.Run("test de integração de repository e service, metodo ListarSectionAll, caso de sucesso", func(t *testing.T) {
	})
	t.Run("test de integração de repository e service, metodo ListarSectionAll, caso de error", func(t *testing.T) {
	})
}

func TestServiceListarSectionOne(t *testing.T) {
	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de sucesso", func(t *testing.T) {
	})
	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de error", func(t *testing.T) {
	})
}

func TestServiceCreateSection(t *testing.T) {

	t.Run("metodo CreateSection, caso de sucesso", func(t *testing.T) {

	})
	t.Run("metodo CreateSection, caso de caso de error ao listar sections dentro do metodo CriateSection", func(t *testing.T) {

	})
	t.Run("metodo CreateSection, caso de caso de error ao criar um novo section", func(t *testing.T) {

	})
}

func TestServiceUpdateSection(t *testing.T) {
	t.Run("test servoce no metodo UpdateSection, caso de sucesso", func(t *testing.T) {
	 
	})
	t.Run("test servoce no metodo UpdateSection, caso de error section_number duplicado", func(t *testing.T) {

	})
	t.Run("test service no metodo UpdateSection, caso de error, lista de section retorna vazia dentro do metodo update", func(t *testing.T) {

	})
	t.Run("test service no metodo UpdateSection, caso de error, section não encontrado", func(t *testing.T) {

	})
}

func TestServiceDeleteSection(t *testing.T) {
	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de sucesso", func(t *testing.T) {

	})
	t.Run("test de integração de repository e service, metodo ListarSectionOne, caso de error", func(t *testing.T) {

	})
}
