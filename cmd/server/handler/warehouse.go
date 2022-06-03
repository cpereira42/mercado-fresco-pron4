package handler

import (
	"net/http"

	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID                  int    `json:"id"`
	Address             string `json:"adress"`
	Telephone           string `json:"telephone"`
	Warehouse_code      string `json:"warehouse_code"`
	Minimum_capacity    int    `json:"minimum_capacity"`
	Minimum_temperature int    `json:"minimum_temperature"`
}

type Warehouse struct {
	service warehouse.Service
}

// func isTokenValid(ctx *gin.Context) bool {
// 	token := ctx.GetHeader("token")
// 	if token != os.Getenv("TOKEN") {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
// 		return true
// 	}
// 	return false

// }

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
