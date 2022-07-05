package section

import "errors"

var (
	// erro para section não encontrada
	ErrorNotFound error = errors.New("section is not found")
	// erro para section não alterada
	ErrorNotModify error = errors.New("section not modifycation")
	// erro para section que não foi alterada
	ErrorKeyTableSectionId error = errors.New("this section cannot be removed")
	// erro para listagem de section, quando não há registro
	ErrorFalhaInListAll error = errors.New("sections not this registered")
	// erro na execução do metodo exec do sql
	ErrorFalhaInserializerFields error = errors.New("falha ao serializar campos da section")
)
