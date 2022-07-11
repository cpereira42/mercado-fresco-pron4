package handler

import (
	"fmt"
	"net/http"

	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	service products.Service
}

func NewProduct(ctx *gin.Engine, p products.Service) {
	ep := &Product{service: p}
	pr := ctx.Group("/api/v1/products")
	pr.GET("/", ep.GetAll())
	pr.GET("/:id", ep.GetId())
	pr.DELETE("/:id", ep.Delete())
	pr.POST("/", ep.Create())
	pr.PUT("/:id", ep.Update())
	pr.PATCH("/:id", ep.Update())

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
		id, err := util.IDChecker(ctx)
		if err != nil {
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
		id, err := util.IDChecker(ctx)
		if err != nil {
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
		id, err := util.IDChecker(ctx)
		if err != nil {
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
