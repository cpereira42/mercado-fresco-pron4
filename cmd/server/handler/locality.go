package handler

import (
	"fmt"
	"net/http"

	"github.com/cpereira42/mercado-fresco-pron4/internal/locality"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
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

func (c *Locality) ReportLocality() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		localityID := ctx.Query("id")
		if localityID != "" {
			id, err := util.IDChecker(ctx)
			if err != nil {
				return
			}
			reportLocality, err := c.service.GenerateReportById(id)

			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, reportLocality, ""))
		} else {
			reportLocalities, err := c.service.GenerateReportAll()
			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, reportLocalities, ""))
		}
	}
}
