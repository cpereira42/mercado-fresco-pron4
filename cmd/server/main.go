package main

import (
	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
)


func main() {
	store := store.FileStore{FileName: "sections.json"}

	rep := section.NewRepository(store)

	service := section.NewService(rep)

	sectionController := handler.NewSectionController(service)

	router := gin.Default()
	group := router.Group("/api/v1/sections")
	{
		group.GET("/", sectionController.ListarSectionAll()) 	// lista todos recursos
		group.GET("/:id", sectionController.ListarSectionOne()) // buscar recurso por id
		group.POST("/", sectionController.CreateSection()) 		// cria um novo recurso
		group.PUT("/:id", sectionController.UpdateSection()) 	// modifica recursos
		group.PATCH("/:id", sectionController.ModifyParcial()) 	// modifica recursos
		group.DELETE("/:id", sectionController.DeleteSection()) // remove recursos
	}

	router.Run()
}