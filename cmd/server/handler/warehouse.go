package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	service warehouse.Service
}

func NewWarehouse(ctx *gin.Engine, service warehouse.Service) {
	w := &Warehouse{service: service}

	wr := ctx.Group("api/v1/warehouse")
	wr.GET("/", w.GetAll)
	wr.POST("/", w.Create)
	wr.PATCH("/:id", w.Update)
	wr.GET("/:id", w.GetByID)
	wr.DELETE("/:id", w.Delete)
}

func (c *Warehouse) GetAll(ctx *gin.Context) {
	w, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, w, ""))

}

func (c *Warehouse) Create(ctx *gin.Context) {
	var r warehouse.RequestWarehouseCreate

	if web.CheckIfErrorRequest(ctx, &r) {
		return
	}

	w, err := c.service.Create(r.Address, r.Telephone, r.Warehouse_code, r.Minimum_capacity, r.Minimum_temperature, r.Locality_id)
	if err != nil {
		if err.Error() == "Warehouse already exists" {
			ctx.JSON(409, web.NewResponse(409, nil, err.Error()))
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, w, ""))
}

func (c *Warehouse) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Warehouse Not Found"))
		return
	}

	var r warehouse.RequestWarehouseUpdate
	if web.CheckIfErrorRequest(ctx, &r) {
		return
	}

	w, err := c.service.Update(id, r.Address, r.Telephone, r.Warehouse_code, r.Minimum_capacity, r.Minimum_temperature, r.Locality_id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, w, ""))
}

func (c *Warehouse) GetByID(ctx *gin.Context) {
	id, err := util.IDChecker(ctx)
	if err != nil {
		return
	}

	warehouse, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, warehouse, ""))
}

func (c *Warehouse) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invalid ID"))
		return
	}
	err = c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
}
