package main

import (
	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
	//"github.com/swaggo/swag/example/basic/docs"
)

func main() {

	dbProd := store.New(store.FileType, "../../internal/repositories/products.json")
	repo := products.NewRepositoryProducts(dbProd)
	service := products.NewService(repo)

	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/api/v1/products")
	// pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetId())
	/*pr.POST("/", p.Store())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())*/

	//pr.PUT("/:id", p.Update())

	r.Run()
}
