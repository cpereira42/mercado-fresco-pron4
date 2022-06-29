package productsRecords

import (
	"database/sql"
	"fmt"

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

func (r *repository) GetAllRecords() ([]ReturnProductRecords, error) {
	var ProductRecords []ReturnProductRecords

	rows, err := r.db.Query(` SELECT p.id, p.description, COUNT(pr.product_id) AS records_count
	FROM product_records pr 
	JOIN products p ON p.id = pr.product_id
	GROUP BY p.id`)

	if err != nil {
		return ProductRecords, fmt.Errorf("ProductsRecords not Found")
	}
	defer rows.Close()

	for rows.Next() {
		var ProductRecord ReturnProductRecords

		err := rows.Scan(
			&ProductRecord.ProductId,
			&ProductRecord.Description,
			&ProductRecord.RecordsCount,
		)
		if err != nil {
			return ProductRecords, err
		}
		ProductRecords = append(ProductRecords, ProductRecord)
	}

	return ProductRecords, nil
}

func (r *repository) GetIdRecords(id int) (ReturnProductRecords, error) {

	var ProductRecords ReturnProductRecords

	stmt, err := r.db.Prepare(`SELECT pr.product_id, p.description, COUNT(pr.product_id) AS records_count
	FROM product_records pr 
	JOIN products p ON p.id = pr.product_id
	WHERE pr.product_id = ?
	GROUP BY pr.product_id`)

	if err != nil {
		return ProductRecords, fmt.Errorf("ProductsRecords not Found")
	}

	err = stmt.QueryRow(id).Scan(
		&ProductRecords.ProductId,
		&ProductRecords.Description,
		&ProductRecords.RecordsCount,
	)

	if err != nil {
		return ProductRecords, fmt.Errorf("ProductsRecords not Found")
	}
	defer stmt.Close()
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
