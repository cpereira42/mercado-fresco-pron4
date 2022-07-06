package util

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

func CheckError(sqlError error) error {
	switch {
	case strings.Contains(sqlError.Error(), "no rows in result set"):
		return fmt.Errorf("data not found")
	case strings.Contains(sqlError.Error(), "Duplicate entry"):
		err := strings.Split(sqlError.Error(), "'")
		msg := fmt.Sprint(err[3], " is unique, and ", err[1], " already registered")
		return fmt.Errorf(msg)
	case strings.Contains(sqlError.Error(), "Cannot add"):
		err := strings.Split(sqlError.Error(), "`")
		msg := fmt.Sprint(err[7], " is not registered on ", err[9])
		return fmt.Errorf(msg)
	}
	return sqlError
}

func CreateRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}
