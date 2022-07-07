package handler

import (
	"fmt"
	"net/http"
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
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

func (c *Product) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invalid ID"))
			return
		}
		p, err := c.service.GetId(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invalid ID"))
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, fmt.Sprintf("product %d was deleted", id), ""))
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
			fmt.Println(err)
			if err.Error() == "product_code is unique, and "+request.ProductCode+" already registered" {
				ctx.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, err.Error()))
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			}
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, p, ""))
	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Invali ID"))
			return
		}

		var request products.RequestProductsUpdate
		if web.CheckIfErrorRequest(ctx, &request) {
			return
		}

		p, err := c.service.Update(int(id), request)
		if err != nil {
			if err.Error() == "data not found" {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			}
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}
