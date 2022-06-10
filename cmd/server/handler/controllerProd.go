package handler

import (
	"fmt"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{service: p}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *Product) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "Invalid ID"))
			return
		}
		p, err := c.service.GetId(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "Invalid ID"))
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(204, web.NewResponse(204, fmt.Sprintf("product %d was deleted", id), ""))
	}
}

func (c *Product) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request products.RequestProductsCreate
		if web.CheckIfErrorRequest(ctx, &request) {
			return
		}
		p, err := c.service.Create(request)
		if err != nil {
			ctx.JSON(422, web.NewResponse(422, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(201, p, ""))
	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "Invali ID"))
			return
		}
		var request products.RequestProductsUpdate
		if web.CheckIfErrorRequest(ctx, &request) {
			return
		}
		p, err := c.service.Update(int(id), request)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}
