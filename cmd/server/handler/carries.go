package handler

import (
	"net/http"

	"github.com/cpereira42/mercado-fresco-pron4/internal/carries"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type Carries struct {
	service carries.Service
}

func NewCarry(ctx *gin.Engine, carry carries.Service) {
	c := &Carries{service: carry}
	carries := ctx.Group("/api/v1/carries")
	carries.POST("/", c.Create)
	localities := ctx.Group("/api/v1/localities")
	localities.GET("/", c.GetReport)
}

func (c *Carries) GetReport(ctx *gin.Context) {
	carryID := ctx.Query("id")

	if carryID != "" {
		id, err := util.IDChecker(ctx)
		if err != nil {
			return
		}
		carryReport, err := c.service.GetByIDReport(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invalid ID"))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, carryReport, ""))
	} else {
		locality, err := c.service.GetAllReport()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, locality, ""))
	}
}

func (c *Carries) Create(ctx *gin.Context) {
	var request carries.RequestCarriesCreate
	if web.CheckIfErrorRequest(ctx, &request) {
		return
	}

	carry, err := c.service.Create(request.Cid, request.CompanyName, request.Address, request.Telephone, request.LocalityID)
	if err != nil {
		ctx.JSON(http.StatusConflict, web.NewResponse(409, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, carry, ""))
}
