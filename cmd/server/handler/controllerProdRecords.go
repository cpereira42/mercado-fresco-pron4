package handler

import (
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/productsRecords"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductRecords struct {
	service productsRecords.Service
}

func NewProductRecords(p productsRecords.Service) *ProductRecords {
	return &ProductRecords{service: p}
}

func (c *ProductRecords) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if ctx.Query("id") == "" {
			p, err := c.service.GetAllRecords()
			if err != nil {
				ctx.JSON(404, web.NewResponse(404, nil, "Invalid ID"))
				return
			}
			ctx.JSON(200, web.NewResponse(200, p, ""))
		} else {
			id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
			if err != nil {
				ctx.JSON(404, web.NewResponse(404, nil, "Invalid ID"))
				return
			}
			p, err := c.service.GetIdRecords(int(id))
			if err != nil {
				ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
				return
			}
			ctx.JSON(200, web.NewResponse(200, p, ""))
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
			if err.Error() == "Product not found" {
				ctx.JSON(409, web.NewResponse(409, nil, err.Error()))
			} else {
				ctx.JSON(422, web.NewResponse(422, nil, err.Error()))
			}
			return
		}
		ctx.JSON(201, web.NewResponse(201, p, ""))
	}
}
