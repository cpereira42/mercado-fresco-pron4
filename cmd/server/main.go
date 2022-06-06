package main

import (
	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {

	dbBuyers := store.New(store.FileType, "./internal/repositories/buyer.json")
	repositoryBuyers := buyer.NewRepository(dbBuyers)
	serviceBuyers := buyer.NewService(repositoryBuyers)
	hdBuyers := handler.NewBuyer(serviceBuyers)

	r := gin.Default()

	buyers := r.Group("/api/v1/buyers")

	buyers.GET("/", hdBuyers.GetAll())
	buyers.GET("/:id", hdBuyers.GetID())
	buyers.POST("/", hdBuyers.Create())
	buyers.PATCH("/:id", hdBuyers.Update())
	buyers.DELETE("/:id", hdBuyers.Delete())

	r.Run()
}
