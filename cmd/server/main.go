package main

import (
	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	dbSeller := store.New(store.FileType, "../mercado-fresco-pron4/internal/repositories/sellers.json")
	repo := seller.NewRepositorySeller(dbSeller)
	service := seller.NewService(repo)

	s := handler.NewSeller(service)

	r := gin.Default()

	sellers := r.Group("/api/v1/sellers")
	sellers.GET("/", s.GetAll())
	sellers.GET("/:id", s.GetId())
	sellers.POST("/", s.Create())
	sellers.PATCH("/:id", s.Update())
	sellers.DELETE("/:id", s.Delete())

	r.Run()
}
