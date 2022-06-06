package main

import (
	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
	//"github.com/swaggo/swag/example/basic/docs"
)

func main() {

	dbEmployees := store.New(store.FileType, "./internal/repositories/employees.json")
	repositoryEmployees := employee.NewRepository(dbEmployees)
	serviceEmployees := employee.NewService(repositoryEmployees)
	handlerEmployees := handler.NewEmployee(serviceEmployees)

	r := gin.Default()

	routesEmployees := r.Group("/api/v1/employees")
	{
		routesEmployees.GET("/", handlerEmployees.GetAll())
		routesEmployees.GET("/:id", handlerEmployees.GetByID())
		routesEmployees.POST("/", handlerEmployees.Create())
		routesEmployees.PATCH("/:id", handlerEmployees.Update())
		routesEmployees.DELETE("/:id", handlerEmployees.Delete())
	}

	r.Run()
}
