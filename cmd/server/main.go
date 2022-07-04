package main

import (
	"database/sql"
<<<<<<< HEAD

	_ "github.com/go-sql-driver/mysql"

=======
>>>>>>> 39634193caf82ae505e42806c8384b5d8ed94405
	"fmt"
	"log"
	"os"

<<<<<<< HEAD
	//sectionRepository "github.com/cpereira42/mercado-fresco-pron4/internal/section/repository/file"
	"github.com/cpereira42/mercado-fresco-pron4/internal/locality"
	sectionRepository "github.com/cpereira42/mercado-fresco-pron4/internal/section/repository/mariadb"
=======
	_ "github.com/go-sql-driver/mysql"

	sectionRepository "github.com/cpereira42/mercado-fresco-pron4/internal/section/repository"
>>>>>>> 39634193caf82ae505e42806c8384b5d8ed94405
	sectionService "github.com/cpereira42/mercado-fresco-pron4/internal/section/service"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products"

	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var Conn *sql.DB

func main() {
<<<<<<< HEAD

=======
>>>>>>> 39634193caf82ae505e42806c8384b5d8ed94405
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".Env cant be load")
	}
	Conn, err = connection()
	if err != nil {
		log.Fatal(err)
	}

	dbBuyers := store.New(store.FileType, "./internal/repositories/buyer.json")
	repositoryBuyers := buyer.NewRepository(dbBuyers)
	serviceBuyers := buyer.NewService(repositoryBuyers)
	hdBuyers := handler.NewBuyer(serviceBuyers)

	dbProd := store.New(store.FileType, "./internal/repositories/products.json")
	repoProd := products.NewRepositoryProducts(dbProd)
	serviceProd := products.NewService(repoProd)

	dbWarehouse := store.New(store.FileType, "./internal/repositories/warehouse.json")
	repoWarehouse := warehouse.NewRepository(dbWarehouse)
	svcWarehouse := warehouse.NewService(repoWarehouse)
	w := handler.NewWarehouse(svcWarehouse)

<<<<<<< HEAD
	//dbSection := store.New(store.FileType, "./internal/repositories/sections.json")
	//repSection := sectionRepository.NewRepository(dbSection)

	repSection := sectionRepository.NewRepository(Conn)
=======
	dbSection := store.New(store.FileType, "./internal/repositories/sections.json")
	repSection := sectionRepository.NewRepository(dbSection)
>>>>>>> 39634193caf82ae505e42806c8384b5d8ed94405
	serviceSection := sectionService.NewService(repSection)
	sectionController := handler.NewSectionController(serviceSection)

	// dbSeller := store.New(store.FileType, "../mercado-fresco-pron4/internal/repositories/sellers.json")
	repoSeller := seller.NewRepositorySeller(Conn)
	serviceSeller := seller.NewService(repoSeller)

	repoLocality := locality.NewRepositoryLocality(Conn)
	serviceLocality := locality.NewService(repoLocality)
	l := handler.NewLocality(serviceLocality)

	dbEmployees := store.New(store.FileType, "./internal/repositories/employees.json")
	repositoryEmployees := employee.NewRepository(dbEmployees)
	serviceEmployees := employee.NewService(repositoryEmployees)
	handlerEmployees := handler.NewEmployee(serviceEmployees)

	s := handler.NewSeller(serviceSeller)
	p := handler.NewProduct(serviceProd)
	r := gin.Default()

	pr := r.Group("/api/v1/products")
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetId())
	pr.DELETE("/:id", p.Delete())
	pr.POST("/", p.Create())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.Update())

	sellers := r.Group("/api/v1/sellers")
	sellers.GET("/", s.GetAll())
	sellers.GET("/:id", s.GetId())
	sellers.POST("/", s.Create())
	sellers.PATCH("/:id", s.Update())
	sellers.DELETE("/:id", s.Delete())

	routesEmployees := r.Group("/api/v1/employees")
	routesEmployees.GET("/", handlerEmployees.GetAll())
	routesEmployees.GET("/:id", handlerEmployees.GetByID())
	routesEmployees.POST("/", handlerEmployees.Create())
	routesEmployees.PATCH("/:id", handlerEmployees.Update())
	routesEmployees.DELETE("/:id", handlerEmployees.Delete())

	section := r.Group("/api/v1/sections")
	section.GET("/", sectionController.ListarSectionAll())    // lista todos recursos
	section.GET("/:id", sectionController.ListarSectionOne()) // buscar recurso por id
	section.POST("/", sectionController.CreateSection())      // cria um novo recurso
	section.PATCH("/:id", sectionController.UpdateSection())  // modifica recursos
	section.DELETE("/:id", sectionController.DeleteSection()) // remove recursos

	wr := r.Group("api/v1/warehouse")
	wr.GET("/", w.GetAll)
	wr.POST("/", w.Create)
	wr.PATCH("/:id", w.Update)
	wr.GET("/:id", w.GetByID)
	wr.DELETE("/:id", w.Delete)

	buyers := r.Group("/api/v1/buyers")
	buyers.GET("/", hdBuyers.GetAll())
	buyers.GET("/:id", hdBuyers.GetID())
	buyers.POST("/", hdBuyers.Create())
	buyers.PATCH("/:id", hdBuyers.Update())
	buyers.DELETE("/:id", hdBuyers.Delete())

	localities := r.Group("/api/v1/localities")
	localities.POST("/", l.Create())
	localities.GET("/", l.GenerateReportAll())
	localities.GET("/:id", l.GenerateReportById())

	r.Run()
}

func connection() (*sql.DB, error) {
	user := os.Getenv("USER_DB")
	pass := os.Getenv("PASS_DB")
	port := os.Getenv("PORT_DB")
	host := os.Getenv("HOST_DB")
	database := os.Getenv("DATABASE")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, database)
	log.Println(dataSource)
	return sql.Open("mysql", dataSource)
}
