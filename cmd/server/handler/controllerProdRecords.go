package handler

import (
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/productsRecords"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductRecords struct {
	service productsRecords.Service
}

func NewProductRecords(ctx *gin.Engine, p productsRecords.Service) /* *ProductRecords */ {
	prc := &ProductRecords{service: p}
	ctx.GET("/api/v1/products/reportRecords/", prc.GetId())
	ctx.POST("/api/v1/productsRecords/", prc.Create())
}

func (c *ProductRecords) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if ctx.Query("id") == "" {
			p, err := c.service.GetAllRecords()
			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invalid ID"))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
		} else {
			id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invalid ID"))
				return
			}
			p, err := c.service.GetIdRecords(int(id))
			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
		}
	}
}

func (c *ProductRecords) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request productsRecords.RequestProductRecordsCreate
		if web.CheckIfErrorRequest(ctx, &request) {
			return
		}
		p, err := c.service.Create(request)
		if err != nil {
			if err.Error() == "product_id is not registered on products" {
				ctx.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, err.Error()))
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			}
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, p, ""))
	}
}
