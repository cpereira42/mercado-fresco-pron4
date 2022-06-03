package main

import (
	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {

	dbWarehouse := store.New(store.FileType, "./internal/repositories/warehouse.json")
	repoWarehouse := warehouse.NewRepository(dbWarehouse)
	svcWarehouse := warehouse.NewService(repoWarehouse)
	w := handler.NewWarehouse(svcWarehouse)

	router := gin.Default()

	wr := router.Group("api/v1/warehouse")
	{
		wr.GET("/", w.GetAll)
		wr.POST("/", w.Create)
		wr.PUT("/:id", w.Update)
		wr.GET("/:id", w.GetByID)
		wr.DELETE("/:id", w.Delete)
		router.Run()
	}
}
