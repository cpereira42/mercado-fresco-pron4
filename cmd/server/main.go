package main

import (
	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	dbSection := store.FileStore{FileName: "./internal/repositories/sections.json"}

	repSection := section.NewRepository(dbSection)

	serviceSection := section.NewService(repSection)

	sectionController := handler.NewSectionController(serviceSection)

	r := gin.Default()
	section := r.Group("/api/v1/sections")
	{
		section.GET("/", sectionController.ListarSectionAll()) 		// lista todos recursos
		section.GET("/:id", sectionController.ListarSectionOne()) 	// buscar recurso por id
		section.POST("/", sectionController.CreateSection()) 		// cria um novo recurso
		section.PATCH("/:id", sectionController.UpdateSection()) 	// modifica recursos
		section.DELETE("/:id", sectionController.DeleteSection()) 	// remove recursos
	}

	r.Run()
}