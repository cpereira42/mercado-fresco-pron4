package main

import (
	"database/sql"

	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/cpereira42/mercado-fresco-pron4/internal/carries"
	inboundOrders "github.com/cpereira42/mercado-fresco-pron4/internal/inbound_orders"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
	"github.com/cpereira42/mercado-fresco-pron4/internal/productsRecords"
	"github.com/cpereira42/mercado-fresco-pron4/internal/purchaseorders"
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"

	"github.com/cpereira42/mercado-fresco-pron4/cmd/server/handler"
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/cpereira42/mercado-fresco-pron4/internal/locality"

	"github.com/cpereira42/mercado-fresco-pron4/internal/products"

	"github.com/cpereira42/mercado-fresco-pron4/internal/seller"
	"github.com/cpereira42/mercado-fresco-pron4/internal/warehouse"
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

	repositoryBuyers := buyer.NewRepository(conn)
	serviceBuyers := buyer.NewService(repositoryBuyers)

	repositoryPurchase := purchaseorders.NewRepository(conn)
	servicePurchase := purchaseorders.NewService(repositoryPurchase)

	repoProd := products.NewRepositoryProductsDB(conn)
	serviceProd := products.NewService(repoProd)

	repoWarehouse := warehouse.NewRepository(conn)
	svcWarehouse := warehouse.NewService(repoWarehouse)

	repoCarries := carries.NewRepository(conn)
	svcCarries := carries.NewService(repoCarries)

	repoSeller := seller.NewRepositorySeller(conn)
	serviceSeller := seller.NewService(repoSeller)

	repoLocality := locality.NewRepositoryLocality(conn)
	serviceLocality := locality.NewService(repoLocality)

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
	handler.NewWarehouse(r, svcWarehouse)
	handler.NewRouteBuyer(r, serviceBuyers)
	handler.NewPurchase(servicePurchase)
	handler.NewCarry(r, svcCarries)

	repSection := section.NewRepository(conn)        // new
	serviceSection := section.NewService(repSection) // new
	handler.NewSectionController(r, serviceSection)  // new

	repoPB := productbatch.NewRepositoryProductBatches(conn)   // new
	servicePB := productbatch.NewServiceProductBatches(repoPB) // new
	handler.NewProductBatChesController(r, servicePB)          // new

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
