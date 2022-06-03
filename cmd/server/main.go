package main

import (
	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	dbSeller := store.New(store.FileType, "/Users/josmoura/Study/Projects/mercado-fresco-pron4/internal/repositories/sellers.json")
	//"/Users/josmoura/Study/Projects/mercado-fresco-pron4/internal/repositories/sellers.json"
	repo := seller.NewRepositorySeller(dbSeller)
	service := seller.NewService(repo)

	s := handler.NewSeller(service)

	r := gin.Default()

	sellers := r.Group("/api/v1/sellers")
	sellers.GET("/", s.GetAll())
	sellers.GET("/:id", s.GetId())
	sellers.POST("/", s.Create())
	/*pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())*/

	//pr.PUT("/:id", p.Update())

	r.Run()
}
