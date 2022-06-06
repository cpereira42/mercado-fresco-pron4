package main

import (
	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
)

 
  
func main() {
	dbSection := store.FileStore{FileName: "./internal/repositories/sections.json"}
	repSection := section.NewRepository(dbSection)
	serviceSection := section.NewService(repSection)
	sectionController := handler.NewSectionController(serviceSection)

	dbSeller := store.New(store.FileType, "../mercado-fresco-pron4/internal/repositories/sellers.json")
	repoSeller := seller.NewRepositorySeller(dbSeller)
	serviceSeller := seller.NewService(repoSeller)

	dbEmployees := store.New(store.FileType, "./internal/repositories/employees.json")
	repositoryEmployees := employee.NewRepository(dbEmployees)
	serviceEmployees := employee.NewService(repositoryEmployees)
	handlerEmployees := handler.NewEmployee(serviceEmployees)

	s := handler.NewSeller(serviceSeller)

	r := gin.Default()

	sellers := r.Group("/api/v1/sellers")
	sellers.GET("/", s.GetAll())
	sellers.GET("/:id", s.GetId())
	sellers.POST("/", s.Create())
	sellers.PATCH("/:id", s.Update())
	sellers.DELETE("/:id", s.Delete())

	routesEmployees := r.Group("/api/v1/employees")
	{
		routesEmployees.GET("/", handlerEmployees.GetAll())
		routesEmployees.GET("/:id", handlerEmployees.GetByID())
		routesEmployees.POST("/", handlerEmployees.Create())
		routesEmployees.PATCH("/:id", handlerEmployees.Update())
		routesEmployees.DELETE("/:id", handlerEmployees.Delete())
	}

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
