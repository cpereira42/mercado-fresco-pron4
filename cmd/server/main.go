package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	inboundOrders "github.com/cpereira42/mercado-fresco-pron4/internal/inbound_orders"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/cpereira42/mercado-fresco-pron4/internal/locality"

	"github.com/cpereira42/mercado-fresco-pron4/internal/products"

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
	conn, err := connection()
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

	repoWarehouse := warehouse.NewRepository(conn)
	svcWarehouse := warehouse.NewService(repoWarehouse)
	w := handler.NewWarehouse(svcWarehouse)

	repoSeller := seller.NewRepositorySeller(conn)
	serviceSeller := seller.NewService(repoSeller)

	repoLocality := locality.NewRepositoryLocality(conn)
	serviceLocality := locality.NewService(repoLocality)
	l := handler.NewLocality(serviceLocality)

	repositoryInboundOrders := inboundOrders.NewRepository(conn)
	serviceInboundOrders := inboundOrders.NewService(repositoryInboundOrders)

	repositoryEmployees := employee.NewRepository(conn)
	serviceEmployees := employee.NewService(repositoryEmployees)

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

	handler.NewInboundOrders(r, serviceInboundOrders)
	handler.NewEmployee(r, serviceEmployees)

	repSection := section.NewRepository(conn)        // new
	serviceSection := section.NewService(repSection) // new
	handler.NewSectionController(r, serviceSection)  // new

	repoPB := productbatch.NewRepositoryProductBatches(conn)   // new
	servicePB := productbatch.NewServiceProductBatches(repoPB) // new
	handler.NewProductBatChesController(r, servicePB)          // new

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
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, database)
	log.Println("connection successful")
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	return conn, nil
}
