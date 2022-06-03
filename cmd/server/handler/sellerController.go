package handler

import (
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type Seller struct {
	service seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{service: s}
}

func (s *Seller) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sellers, err := s.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(402, nil, err.Error()))
			return
		}
		if len(sellers) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Sellers nao retornou nada",
				"data":    sellers,
			})
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(200, sellers, ""))
	}
}

func (s *Seller) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(402, nil, "ID inválido"))
			return
		}

		seller, err := s.service.GetId(int(id))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(402, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(http.StatusOK, seller, ""))
	}
}
