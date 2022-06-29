package productsRecords

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	db *sql.DB
}

func NewRepositoryProductsRecordsDB(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetIdRecords(id int) (ReturnProductRecords, error) {

	var ProductRecords ReturnProductRecords

	stmt, err := r.db.Prepare("SELECT product_id FROM product_records Where id = ?")
	if err != nil {
		return ProductRecords, err
	}
	err = stmt.QueryRow(id).Scan(&ProductRecords.ProductId)
	if err != nil {
		return ProductRecords, fmt.Errorf("ProductRecord not Found")
	}

	stmt, err = r.db.Prepare("SELECT description FROM products Where id = ?")
	if err != nil {
		return ProductRecords, err
	}
	err = stmt.QueryRow(ProductRecords.ProductId).Scan(&ProductRecords.Description)
	if err != nil {
		return ProductRecords, fmt.Errorf("Products not Found")
	}

	defer stmt.Close()

	/*stmt, err = r.db.Prepare(`SELECT products.id, products.description, COUNT(product_records.product_id)
	  FROM mercadofresco.products
	  INNER JOIN product_records ON products.id = product_records.product_id
	  WHERE product_records.id = ?
	  GROUP BY product_id`)*/

	stmt, err = r.db.Prepare(`SELECT pr.product_id, p.description, COUNT(pr.product_id) AS records_count
	FROM product_records pr 
	JOIN products p ON p.id = pr.product_id
	WHERE pr.product_id = ?
	GROUP BY pr.product_id`)

	if err != nil {
		log.Fatal("err", err)
		return ProductRecords, fmt.Errorf("Products not sssFound")
	}

	err = stmt.QueryRow(id).Scan(
		&ProductRecords.ProductId,
		&ProductRecords.Description,
		&ProductRecords.RecordsCount,
	)

	if err != nil {
		return ProductRecords, fmt.Errorf("Products not wwwFound")
	}
	return ProductRecords, nil
}

func (r *repository) Create(p ProductRecords) (ProductRecords, error) {

	stmt, err := r.db.Prepare(`INSERT INTO product_records (
		last_update_date, 
		purchase_price, 
		sale_price,
		product_id
		) VALUES (?, ?, ?, ?)`)

	if err != nil {
		return ProductRecords{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		p.LastUpdateDate,
		p.PurchasePrice,
		p.SalePrice,
		p.ProductId)
	if err != nil {
		return ProductRecords{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return ProductRecords{}, fmt.Errorf("Fail to save")
	}

	lastId, _ := res.LastInsertId()
	p.Id = int(lastId)
	return p, nil
}
