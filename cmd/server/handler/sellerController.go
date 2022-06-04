package handler

import (
	"fmt"
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

func (s *Seller) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		if !s.service.CheckCid(req.Cid) {
			ctx.JSON(
				http.StatusConflict,
				web.NewResponse(http.StatusConflict, nil, "Cid já cadastrado"),
			)
			return
		}

		if req.Cid == 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "O Cid é obrigatório"),
			)
			return
		}

		if req.CompanyName == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Necessário informar nome da empresa"),
			)
			return
		}

		if req.Adress == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Necessário informar endereço"),
			)
			return
		}

		if req.Telephone == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Necessário informar telefone"),
			)
			return
		}

		seller, err := s.service.Create(req.Cid, req.CompanyName, req.Adress, req.Telephone)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(
			http.StatusOK,
			web.NewResponse(http.StatusOK, seller, ""),
		)
	}
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
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, sellers, ""))
	}
}

func (s *Seller) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "ID inválido"))
			return
		}

		seller, err := s.service.GetId(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, seller, ""))
	}
}

func (s *Seller) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Id inválido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if req.Cid == 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "O Cid é obrigatório"),
			)
			return
		}

		if req.CompanyName == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Necessário informar nome da empresa"),
			)
			return
		}

		if req.Adress == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Necessário informar endereço"),
			)
			return
		}

		if req.Telephone == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Necessário informar telefone"),
			)
			return
		}

		seller, err := s.service.Update(int(id), req.Cid, req.CompanyName, req.Adress, req.Telephone)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusBadRequest, seller, ""))
	}
}

func (s *Seller) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		err = s.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": fmt.Sprintf("O produto %d foi removido", id)})
	}
}

type request struct {
	Cid         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Adress      string `json:"address"`
	Telephone   string `json:"telephone"`
}
