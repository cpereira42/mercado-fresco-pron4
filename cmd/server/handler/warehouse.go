package handler

import (
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	service warehouse.Service
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		service: w,
	}
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

	if web.CheckIfErrorInRequest(ctx, &r) {
		return
	}

	w, err := c.service.Create(r.Address, r.Telephone, r.Warehouse_code, r.Minimum_capacity, r.Minimum_temperature)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
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
	if web.CheckIfErrorInRequest(ctx, &r) {
		return
	}
	// if err := ctx.ShouldBindJSON(&r); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Error reading request body"))
	// }

	w, err := c.service.Update(id, r.Address, r.Telephone, r.Warehouse_code, r.Minimum_capacity, r.Minimum_temperature)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, w, ""))
}

func (c *Warehouse) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invalid ID"))
		return
	}

	w, err := c.service.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, w, ""))
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
