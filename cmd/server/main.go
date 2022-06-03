package main

import (
	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := warehouse.NewRepository()
	svc := warehouse.NewService(repo)
	w := handler.NewWarehouse(svc)

	router := gin.Default()

	wr := router.Group("api/v1/warehouse")
	wr.GET("/", w.GetAll)
	wr.POST("/", w.Create)
	wr.PUT("/:id", w.Update)
	router.Run()
}
