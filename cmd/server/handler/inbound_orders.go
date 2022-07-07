package handler

import (
	"fmt"
	"net/http"
	"strconv"

	inboundOrders "github.com/cpereira42/mercado-fresco-pron4/internal/inbound_orders"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

func idChecker(ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "invalid ID"))
		return id, err
	}
	return id, nil
}

type InboundOrdersController struct {
	service inboundOrders.Service
}

func NewInboundOrders(ctx *gin.Engine, service inboundOrders.Service) {
	ioc := &InboundOrdersController{service: service}

	ior := ctx.Group("/api/v1")
	{
		ior.GET("/employees/reportInboundOrders", ioc.ReportInboundOrders())
		ior.POST("/inboundOrders", ioc.Create())
	}
}

func (c *InboundOrdersController) ReportInboundOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employeeID := ctx.Query("id")
		if employeeID != "" {
			id, err := idChecker(ctx)
			if err != nil {
				return
			}
			reportInboundOrder, err := c.service.GetByID(id)

			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, reportInboundOrder, ""))
		} else {
			reportInboundOrders, err := c.service.GetAll()
			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, reportInboundOrders, ""))
		}
	}
}

func (c *InboundOrdersController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request inboundOrders.InboundOrdersCreate
		if web.CheckIfErrorRequest(ctx, &request) {
			return
		}

		inboundOrders, err := c.service.Create(request)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, inboundOrders, ""))
	}
}
