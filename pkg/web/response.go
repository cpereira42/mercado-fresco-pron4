package web

import (
	"errors"
	"net/http"
	"strconv"

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
