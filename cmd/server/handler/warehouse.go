package handler

import (
	"io"
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	service warehouse.Service
}
type request struct {
	ID                  int    `json:"id"`
	Address             string `json:"adress" binding:"required"`
	Telephone           string `json:"telephone" binding:"required"`
	Warehouse_code      string `json:"warehouse_code" binding:"required"`
	Minimum_capacity    int    `json:"minimum_capacity" binding:"required"`
	Minimum_temperature int    `json:"minimum_temperature" binding:"required"`
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
	var r request
	if err := ctx.ShouldBindJSON(&r); err != nil {
		if err == io.EOF {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest, nil, "Error reading request body"),
			)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
		return
	}
	w, err := c.service.Create(r.ID, r.Address, r.Telephone, r.Warehouse_code, r.Minimum_capacity, r.Minimum_temperature)
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

	var r warehouse.Warehouse
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Error reading request body"))
	}

	// if r.Address == "" {
	// 	ctx.JSON(422, gin.H{"error": "address is required"})
	// }
	// if r.Telephone == "" {
	// 	ctx.JSON(422, gin.H{"error": "telephone is required"})
	// }
	// if r.Warehouse_code == "" {
	// 	ctx.JSON(422, gin.H{"error": "warehouse_code is required"})
	// }
	// if r.Minimum_capacity == 0 {
	// 	ctx.JSON(422, gin.H{"error": "minimum_capacity is required"})
	// }
	// if r.Minimum_temperature == 0 {
	// 	ctx.JSON(422, gin.H{"error": "minimum_temperature is required"})
	// }

	w, err := c.service.Update(id, r.Address, r.Telephone, r.Warehouse_code, r.Minimum_capacity, r.Minimum_temperature)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Warehouse Code already exists"))
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
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
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
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
}
