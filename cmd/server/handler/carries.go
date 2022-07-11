package handler

import (
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/carries"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type Carries struct {
	service carries.Service
}

func NewCarry(carry carries.Service) *Carries {
	return &Carries{
		service: carry,
	}
}
func (c *Carries) GetAllReport(ctx *gin.Context) {
	carries, err := c.service.GetAllReport()
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, carries, ""))
}
func (c *Carries) GetByIDReport(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invalid ID"))
		return
	}

	locality, err := c.service.GetByIDReport(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, locality, ""))
}

func (c *Carries) Create(ctx *gin.Context) {
	var request carries.RequestCarriesCreate
	if web.CheckIfErrorRequest(ctx, &request) {
		return
	}

	carry, err := c.service.Create(request.Cid, request.CompanyName, request.Address, request.Telephone, request.LocalityID)
	if err != nil {
		if err.Error() == "Carry already exists" {
			ctx.JSON(409, web.NewResponse(409, nil, err.Error()))
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, carry, ""))
}
