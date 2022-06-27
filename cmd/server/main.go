package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/cpereira42/mercado-fresco-pron4/internal/products"
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".Env cant be load")
	}

	conn, err2 := ConnectDB()
	if err2 != nil {
		log.Fatal("could not open the conection: ", err2)
	}

	dbBuyers := store.New(store.FileType, "./internal/repositories/buyer.json")
	repositoryBuyers := buyer.NewRepository(dbBuyers)
	serviceBuyers := buyer.NewService(repositoryBuyers)
	hdBuyers := handler.NewBuyer(serviceBuyers)

	repoProd := products.NewRepositoryProductsDB(conn)

	//dbProd := store.New(store.FileType, "./internal/repositories/products.json")
	//repoProd := products.NewRepositoryProductsDB(conn)
	serviceProd := products.NewService(repoProd)

	dbWarehouse := store.New(store.FileType, "./internal/repositories/warehouse.json")
	repoWarehouse := warehouse.NewRepository(dbWarehouse)
	svcWarehouse := warehouse.NewService(repoWarehouse)
	w := handler.NewWarehouse(svcWarehouse)

	dbSection := store.New(store.FileType, "./internal/repositories/sections.json")
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

	r.Run()
}

func ConnectDB() (*sql.DB, error) {
	user := os.Getenv("USER_DB")
	pass := os.Getenv("PASS_DB")
	host := os.Getenv("HOST_DB")
	port := os.Getenv("PORT_DB")
	table := os.Getenv("DATABASE")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, table)
	log.Println(dataSource)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("could not ping the database: ", err)
		return nil, err
	}

	log.Println("connected")
	return db, nil
}
