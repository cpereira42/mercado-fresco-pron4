package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	inboundOrders "github.com/cpereira42/mercado-fresco-pron4/internal/inbound_orders"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productsRecords"
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

	repoProd := products.NewRepositoryProductsDB(conn)
	serviceProd := products.NewService(repoProd)

	repoWarehouse := warehouse.NewRepository(conn)
	svcWarehouse := warehouse.NewService(repoWarehouse)
	w := handler.NewWarehouse(svcWarehouse)

	repoSeller := seller.NewRepositorySeller(conn)
	serviceSeller := seller.NewService(repoSeller)

	repoLocality := locality.NewRepositoryLocality(conn)
	serviceLocality := locality.NewService(repoLocality)

	repoPB := productbatch.NewRepositoryProductBatches(conn)
	servicePB := productbatch.NewServiceProductBatches(repoPB)
	productBatchesController := handler.NewProductBatChesController(servicePB)

	repSection := section.NewRepository(conn)
	serviceSection := section.NewService(repSection)
	sectionController := handler.NewSectionController(serviceSection)

	repositoryInboundOrders := inboundOrders.NewRepository(conn)
	serviceInboundOrders := inboundOrders.NewService(repositoryInboundOrders)

	repositoryEmployees := employee.NewRepository(conn)
	serviceEmployees := employee.NewService(repositoryEmployees)

	repoProdRecord := productsRecords.NewRepositoryProductsRecordsDB(conn)
	serviceProdRecord := productsRecords.NewService(repoProdRecord)

	r := gin.Default()
	handler.NewProduct(r, serviceProd)
	handler.NewProductRecords(r, serviceProdRecord)
	handler.NewInboundOrders(r, serviceInboundOrders)
	handler.NewEmployee(r, serviceEmployees)
	handler.NewSeller(r, serviceSeller)
	handler.NewLocality(r, serviceLocality)

	section := r.Group("/api/v1/sections")
	section.GET("/", sectionController.ListarSectionAll())
	section.GET("/:id", sectionController.ListarSectionOne())
	section.POST("/", sectionController.CreateSection())
	section.PATCH("/:id", sectionController.UpdateSection())
	section.DELETE("/:id", sectionController.DeleteSection())
	section.GET("/reportProducts", productBatchesController.ReadPB()) // new

	productBatches := r.Group("/api/v1/productBatches")          // new
	productBatches.POST("", productBatchesController.CreatePB()) // new

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
