package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/gin-gonic/gin"
)

type BuyerController struct {
	service buyer.Service
}

func NewBuyer(buyer buyer.Service) *BuyerController {
	return &BuyerController{
		service: buyer,
	}
}

func (c *BuyerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buyer, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": buyer})
	}
}

func (c *BuyerController) GetID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}

		buyer, err := c.service.GetId(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": buyer})
	}
}

func (c *BuyerController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request buyerRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		buyer, err := c.service.Create(request.Card_number_ID, request.First_name, request.Last_name)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, buyer)

	}
}

func (c *BuyerController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}

		var request buyerRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if request.Card_number_ID == "" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Buyer's ID is mandatory"})
			return
		}
		if request.First_name == "" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Buyer's name is mandatory"})
			return
		}
		if request.Last_name == "" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Buyer's lastname is mandatory"})
			return
		}

		buyer, err := c.service.Update(int(id), request.Card_number_ID, request.First_name, request.Last_name)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, buyer)
	}
}

func (c *BuyerController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{"data": fmt.Sprintf("buyer %d removed", id)})
	}
}

type buyerRequest struct {
	ID             int    `json:"id"`
	Card_number_ID string `json:"card_number_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
}
