package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"

	//"github.com/cpereira42/mercado-fresco-pron4/internal/section"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type RequestError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewResponse(code int, data interface{}, err string) Response {
	if code < 300 {
		return Response{code, data, ""}
	}
	return Response{code, nil, err}
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "numeric":
		return "This field only accepts numbers"
	case "ascii":
		return "This field has an invalid date value"
	}
	return ""
}

func CheckIfErrorRequest(ctx *gin.Context, request any) bool {
	var (
		// type of errors
		out                 []RequestError
		unmarshalFieldError *json.UnmarshalFieldError // this erro is deprecated
		unmarshalTypeError  *json.UnmarshalTypeError
		validationErrors    validator.ValidationErrors
	)
	if err := ctx.ShouldBind(&request); err != nil {
		switch {
		case errors.As(err, &unmarshalFieldError):

			errString, sep := unmarshalFieldError.Error(), ":"
			strin := strings.Split(errString, sep)[1]
			requestError := RequestError{unmarshalFieldError.Field.Name, strings.TrimSpace(strin)}
			ctx.JSON(http.StatusUnprocessableEntity,
				Response{http.StatusUnprocessableEntity, requestError, ""})

		case errors.As(err, &validationErrors):

			out = make([]RequestError, len(validationErrors))
			typeAluno := reflect.TypeOf(request).Elem()
			for i, fe := range validationErrors {
				field, ok := typeAluno.FieldByName(fe.Field())
				if ok {
					out[i] = RequestError{field.Tag.Get("json"), msgForTag(fe.Tag())}
				}
			}
			ctx.JSON(http.StatusUnprocessableEntity,
				Response{http.StatusUnprocessableEntity, out, ""})

		case errors.As(err, &unmarshalTypeError):

			strin := strings.Split(unmarshalTypeError.Error(), ":")[1]
			requestError := RequestError{unmarshalTypeError.Field, strings.TrimSpace(strin)}
			ctx.JSON(http.StatusUnprocessableEntity,
				Response{http.StatusUnprocessableEntity, requestError, ""})

		default:

			ctx.JSON(http.StatusUnprocessableEntity,
				Response{http.StatusUnprocessableEntity, nil, err.Error()})

		}
		return true
	}
	return false
}
