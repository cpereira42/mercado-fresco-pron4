package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error any         `json:"error,omitempty"`
}

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewResponse(code int, data interface{}, err any) Response {
	if code < 300 {
		return (Response{code, data, ""})
	}
	return (Response{code, nil, err})
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
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), msgForTag(fe.Tag())}
			}
			ctx.JSON(http.StatusUnprocessableEntity, NewResponse(http.StatusUnprocessableEntity, nil, out))
		}
		return true
	}
	return false
}
