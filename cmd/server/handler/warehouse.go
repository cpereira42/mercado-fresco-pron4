package handler

import (
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	service warehouse.Service
}
type request struct {
	ID                  int    `json:"id"`
	Address             string `json:"adress"`
	Telephone           string `json:"telephone"`
	Warehouse_code      string `json:"warehouse_code"`
	Minimum_capacity    int    `json:"minimum_capacity"`
	Minimum_temperature int    `json:"minimum_temperature"`
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		service: w,
	}
}

func (c *Warehouse) GetAll(ctx *gin.Context) {
	w, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, w)

}

func (c *Warehouse) Create(ctx *gin.Context) {
	var r request
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	w, err := c.service.Create(r.ID, r.Address, r.Telephone, r.Warehouse_code, r.Minimum_capacity, r.Minimum_temperature)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, w)
}

func (c *Warehouse) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var r request
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if r.Address == "" {
		ctx.JSON(400, gin.H{"error": "address is required"})
	}
	if r.Telephone == "" {
		ctx.JSON(400, gin.H{"error": "telephone is required"})
	}
	if r.Warehouse_code == "" {
		ctx.JSON(400, gin.H{"error": "warehouse_code is required"})
	}
	if r.Minimum_capacity == 0 {
		ctx.JSON(400, gin.H{"error": "minimum_capacity is required"})
	}
	if r.Minimum_temperature == 0 {
		ctx.JSON(400, gin.H{"error": "minimum_temperature is required"})
	}

	w, err := c.service.Update(id, r.Address, r.Telephone, r.Warehouse_code, r.Minimum_capacity, r.Minimum_temperature)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, w)
}

func (c *Warehouse) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	w, err := c.service.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, w)
}
