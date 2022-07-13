package handler

import (
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/purchaseorders"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type PurchaseController struct {
	service purchaseorders.Service
}

func NewPurchase(purchase purchaseorders.Service) *PurchaseController {
	return &PurchaseController{
		service: purchase,
	}
}

func NewRoutePurchase(r *gin.Engine, servicePurchase purchaseorders.Service) {
	controllerPurchase := NewPurchase(servicePurchase)
	purchase := r.Group("/api/v1/purchase")
	purchase.GET("/:id", controllerPurchase.GetById())
	purchase.POST("/", controllerPurchase.Create())
}

func (c *PurchaseController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invalid ID"))
			return
		}

		purchase, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, purchase, ""))
	}
}

func (c *PurchaseController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request purchaseorders.RequestPurchaseCreate
		if web.CheckIfErrorRequest(ctx, &request) {
			return
		}

		purchase, err := c.service.Create(
			request.Order_date,
			request.Order_number,
			request.Tracking_code,
			request.Buyer_id,
			request.Product_record_id,
			request.Order_status_id,
		)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, purchase, ""))

	}
}
