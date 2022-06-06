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

func checkFields(p products.Product) error {

	if p.Description == "" {
		return fmt.Errorf("description is mandatory")
	}

	if p.Width <= 0.0 {
		return fmt.Errorf("width is mandatory and cannot be less than 0")
	}

	if p.Length <= 0.0 {
		return fmt.Errorf("Length is mandatory and cannot be less than 0")
	}

	if p.Height <= 0.0 {
		return fmt.Errorf("Height is mandatory and cannot be less than 0")
	}
	if p.NetWeight <= 0.0 {
		return fmt.Errorf("NetWeight is mandatory and cannot be less than 0")
	}

	if p.NetWeight <= 0.0 {
		return fmt.Errorf("NetWeight is mandatory and cannot be less than 0")
	}
	if p.RecommendedFreezingTemperature <= 0.0 {
		return fmt.Errorf("RecommendedFreezingTemperature is mandatory and cannot be less than 0")
	}

	if p.FreezingRate <= 0.0 {
		return fmt.Errorf("FreezingRate is mandatory and cannot be less than 0")
	}

	if p.ExpirationRate <= 0.0 {
		return fmt.Errorf("ExpirationRate é obrigatório")
	}

	if p.ProductType_Id <= 0 {
		return fmt.Errorf("ProductType_Id is mandatory and cannot be less than 0")
	}

	if p.SellerId <= 0 {
		return fmt.Errorf("SellerId is mandatory and cannot be less than 0")
	}

	if p.Product_code <= "" {
		return fmt.Errorf("Product_code é obrigatório")
	}
	return nil
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var prod products.Product
		if err := ctx.Bind(&prod); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}

		if err := c.service.CheckCode(prod.Product_code); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}

		if err := checkFields(prod); err != nil {
			ctx.JSON(422, web.NewResponse(422, nil, err.Error()))
			return
		}
		p, err := c.service.Store(prod)
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
		var prod products.Product
		if err := ctx.ShouldBindJSON(&prod); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if err := checkFields(prod); err != nil {
			ctx.JSON(422, web.NewResponse(422, nil, err.Error()))
			return
		}
		if err := c.service.CheckCode(prod.Product_code); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}

		p, err := c.service.Update(int(id), prod)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *Product) UpdatePatch() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "Invalid ID"))
			return
		}

		var prod products.Product
		if err := ctx.ShouldBindJSON(&prod); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if err := c.service.CheckCode(prod.Product_code); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}

		p, err := c.service.UpdatePatch(int(id), prod)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}
