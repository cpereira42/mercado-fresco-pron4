package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"

	"strconv"
	//"strings"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response {
	if code < 300 {
		return (Response{strconv.FormatInt(int64(code), 10), data, ""})
	}
	return (Response{strconv.FormatInt(int64(code), 10), nil, err})
}

type RequestError struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "numeric":
		return "This field only accepts numbers"
	}
	return ""
}

func CheckIfErrorRequest(ctx *gin.Context, req any) bool {
	if err := ctx.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors 
		if errors.As(err, &ve) {
			out := make([]RequestError, len(ve))
			for i, fe := range ve {
				out[i] = RequestError{ fe.Field(), msgForTag(fe.Tag())}
			}
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": http.StatusUnprocessableEntity,
				"error": out})
		}
		return true
	}
	return false
}

/* 
	Implementação de validação no bind das request em rotas post/patch
	esse método contém melhorias seguindo a lógica do metodo acima no código
*/
func getBindRequest(context *gin.Context, request *section.SectionRequest) error {
	return context.ShouldBind(&request)
}

func CheckIfErrorInRequest(ctx *gin.Context, request *section.SectionRequest) bool {
	var (
		// type of errors
		out []RequestError
		unmarshalFieldError *json.UnmarshalFieldError // this erro is deprecated
		unmarshalTypeError *json.UnmarshalTypeError
		validationErrors validator.ValidationErrors
	)
	if err := getBindRequest(ctx, request); err != nil {
		switch {
		case errors.As(err, &unmarshalFieldError):
			
			errString, sep  := unmarshalFieldError.Error(), ":"
			strin := strings.Split(errString, sep)[1]
			requestError := RequestError{ unmarshalFieldError.Field.Name, strings.TrimSpace(strin) } 
			ctx.JSON(http.StatusUnprocessableEntity, 
				Response{strconv.FormatInt(int64(http.StatusUnprocessableEntity),10), requestError, ""})

		case errors.As(err, &validationErrors):
			
			out = make([]RequestError, len(validationErrors))
			typeAluno := reflect.TypeOf(*request)			
			for i, fe := range validationErrors {
				field, ok :=typeAluno.FieldByName(fe.Field())
				if ok {
					out[i] = RequestError{ field.Tag.Get("json"), msgForTag(fe.Tag())}
				}
			}
			ctx.JSON(http.StatusUnprocessableEntity, 
				Response{strconv.FormatInt(int64(http.StatusUnprocessableEntity),10), out, ""})

		case  errors.As(err, &unmarshalTypeError) :
			
			strin := strings.Split(unmarshalTypeError.Error(), ":")[1]
			requestError := RequestError{ unmarshalTypeError.Field, strings.TrimSpace(strin) }
			ctx.JSON(http.StatusUnprocessableEntity,
				Response{strconv.FormatInt(int64(http.StatusUnprocessableEntity),10), requestError, ""})

		default:
			
			ctx.JSON(http.StatusUnprocessableEntity, 
				Response{strconv.FormatInt(int64(http.StatusUnprocessableEntity),10), nil, err.Error()})

		}
		return true
	}
	return false
}
