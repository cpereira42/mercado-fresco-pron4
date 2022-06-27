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

func (s *Seller) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req seller.SellerRequestCreate

		if web.CheckIfErrorRequest(ctx, &req) {
			return
		}

		seller, err := s.service.Create(req.Cid, req.CompanyName, req.Adress, req.Telephone)

		if err != nil {
			ctx.JSON(
				http.StatusConflict,
				web.NewResponse(http.StatusConflict, nil, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusCreated,
			web.NewResponse(http.StatusCreated, seller, ""),
		)
	}
}

func (s *Seller) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sellers, err := s.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if len(sellers) == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, sellers, "Sellers is empty"))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, sellers, ""))
	}
}

func (s *Seller) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Invalid ID"))
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
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Invalid ID"))
			return
		}

		var req seller.SellerRequestUpdate
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Invalid body arguments"))
			return
		}

		seller, err := s.service.Update(int(id), req.Cid, req.CompanyName, req.Adress, req.Telephone)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, seller, ""))
	}
}

func (s *Seller) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid ID"))
			return
		}

		err = s.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, "Seller sucessfully removed"))
	}
}

type sellerRequest struct {
	Cid         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Adress      string `json:"address"`
	Telephone   string `json:"telephone"`
}
