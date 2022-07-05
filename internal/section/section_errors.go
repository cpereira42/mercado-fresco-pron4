package section

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// erro para section não encontrada
	ErrorNotFound error = errors.New("section is not found")
	// erro para section não alterada
	ErrorNotModify error = errors.New("section not modifycation")
	// erro para section que não foi alterada
	ErrorKeyTableSectionId error = errors.New("this section cannot be removed")
	// erro para listagem de section, quando não há registro
	ErrorFalhaInListAll error = fmt.Errorf("sections not this registered")
	// erro na execução do metodo exec do sql
	ErrorFalhaInserializerFields error = fmt.Errorf("falha ao serializar campos da section")
)

func checkError(sqlError error) error {
	switch {
	case strings.Contains(sqlError.Error(), "no rows in result set"):
		return fmt.Errorf("data not found")
	case strings.Contains(sqlError.Error(), "Duplicate entry"):
		err := strings.Split(sqlError.Error(), "'")
		msg := fmt.Sprint(err[3], " is Unique, and ", err[1], " already registred")
		return fmt.Errorf(msg)
	case strings.Contains(sqlError.Error(), "Cannot add"):
		err := strings.Split(sqlError.Error(), "`")
		msg := fmt.Sprint(err[7], " is not registred ON ", err[9])
		return fmt.Errorf(msg)
	case strings.Contains(sqlError.Error(), "arguments do not match: expected 7, but got 8 arguments"):
		return fmt.Errorf("fields invalid")
	}
	return sqlError
}
