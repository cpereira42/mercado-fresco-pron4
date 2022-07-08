package handler

import (
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/locality"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type Locality struct {
	service locality.Service
}

func NewLocality(l locality.Service) *Locality {
	return &Locality{service: l}
}

func (l *Locality) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req locality.LocalityRequestCreate

		if web.CheckIfErrorRequest(ctx, &req) {
			return
		}

		locality, err := l.service.Create(req.Id, req.LocalityName, req.ProvinceName, req.CountryName)

		if err != nil {
			ctx.JSON(
				http.StatusConflict,
				web.NewResponse(http.StatusConflict, nil, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusCreated,
			web.NewResponse(http.StatusCreated, locality, ""),
		)
	}
}

func (l *Locality) GenerateReportAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		report, err := l.service.GenerateReportAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, report, ""))
	}
}

func (l *Locality) GenerateReportById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Invalid ID"))
			return
		}
		report, err := l.service.GenerateReportById(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, report, ""))
	}
}
