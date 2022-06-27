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

func (r *repository) GetId(id int) (ProductRecords, error) {
	var ProductRecords ProductRecords

	stmt, err := r.db.Prepare("SELECT * FROM products_records Where id = ?")
	if err != nil {
		return ProductRecords, err
	}
	err = stmt.QueryRow(id).Scan(&ProductRecords.Id,
		&ProductRecords.LastUpdateDate,
		&ProductRecords.PurchasePrice,
		&ProductRecords.SalePrice,
		&ProductRecords.productTypeId)
	defer stmt.Close()

	if err != nil {
		return ProductRecords, err
	}
	return ProductRecords, nil
}

/*func (r *repositoryDB) CheckCode(code string) error {

	stmt, err := r.db.Prepare("SELECT product_code FROM products Where product_code = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(code).Scan(&code)
	if err != nil {
		return fmt.Errorf("Product already registred")
	}
	return nil
}*/

func (r *repository) Create(p ProductRecords) (ProductRecords, error) {

	stmt, err := r.db.Prepare(`INSERT INTO products_records (
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
		p.productTypeId)
	if err != nil {
		return ProductRecords{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return ProductRecords{}, fmt.Errorf("Fail to save")
	}
	return p, nil
}
