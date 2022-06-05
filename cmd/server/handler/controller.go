package handler

import (
	"fmt"
	"os"
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
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
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
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "ID inválido"))
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
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "ID inválido"))
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(204, web.NewResponse(204, fmt.Sprintf("O produto %d foi removido", id), ""))
	}
}

func checkFields(p products.Product) error {

	if p.Description == "" {
		return fmt.Errorf("o campo descrição é obrigatório")
	}

	if p.Width <= 0.0 {
		return fmt.Errorf("o campo width é obrigatório e não pode ser menor que 0")
	}

	if p.Length <= 0.0 {
		return fmt.Errorf("o campo Length é obrigatório e não pode ser menor que 0")
	}

	if p.Height <= 0.0 {
		return fmt.Errorf("o campo Height é obrigatório e não pode ser menor que 0")
	}
	if p.NetWeight <= 0.0 {
		return fmt.Errorf("o campo NetWeight é obrigatório e não pode ser menor que 0")
	}

	if p.NetWeight <= 0.0 {
		return fmt.Errorf("o campo NetWeight é obrigatório e não pode ser menor que 0")
	}
	if p.RecommendedFreezingTemperature <= 0.0 {
		return fmt.Errorf("o campo RecommendedFreezingTemperature é obrigatório e não pode ser menor que 0")
	}

	if p.FreezingRate <= 0.0 {
		return fmt.Errorf("o campo FreezingRate é obrigatório e não pode ser menor que 0")
	}

	if p.ExpirationRate <= "" {
		return fmt.Errorf("o campo ExpirationRate é obrigatório")
	}

	if p.ProductType_Id <= 0 {
		return fmt.Errorf("o campo ProductType_Id é obrigatório e não pode ser menor que 0")
	}

	if p.SellerId <= 0 {
		return fmt.Errorf("o campo SellerId é obrigatório e não pode ser menor que 0")
	}

	if p.Product_code <= "" {
		return fmt.Errorf("o campo Product_code é obrigatório")
	}
	return nil
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
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
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "ID inválido"))
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
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "ID inválido"))
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
