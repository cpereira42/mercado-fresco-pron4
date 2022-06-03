package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id                             int     `json:"id"`
	Product_code                   int     `json:"ui"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	NetHeight                      float64 `json:"net_height"`
	ExpirationRate                 string  `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	ProductType_Id                 float64 `json:"product_type_id"`
	SellerId                       float64 `json:"seller_id"`
}

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
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
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
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("O produto %d foi removido", id), ""))
	}
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
		p, err := c.service.Store(prod)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
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

		/*if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "O nome do produto é obrigatório"))
			return
		}

		if req.Tipo == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "O tipo do produto é obrigatório"))
			return
		}

		if req.Count == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "A Quantidade do produto é obrigatória"))
			return
		}

		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "O preço do produto é obrigatório"))
			return
		}*/
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

		if prod.Description == "" {
			ctx.JSON(401, web.NewResponse(400, nil, "O nome do produto é obrigatório"))
			return
		}
		p, err := c.service.UpdatePatch(int(id), prod)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}
