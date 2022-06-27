package products

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type RepositoryDB interface {
	GetAll() ([]Product, error)
	GetId(id int) (Product, error)
	Delete(id int) error
	LastID() (int, error)
	Create(p Product) (Product, error)
	Update(id int, prod Product) (Product, error)
	CheckCode(code string) error
}

type repositoryDB struct {
	db *sql.DB
}

func NewRepositoryProductsDB(db *sql.DB) RepositoryDB {
	return &repositoryDB{
		db: db,
	}
}

func (r *repositoryDB) GetAll() ([]Product, error) {
	var products []Product

	rows, err := r.db.Query("SELECT * FROM products")
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product

		err := rows.Scan(&product.Id,
			&product.ProductCode,
			&product.Description,
			&product.Width,
			&product.Length,
			&product.Height,
			&product.NetWeight,
			&product.ExpirationRate,
			&product.RecommendedFreezingTemperature,
			&product.FreezingRate,
			&product.ProductTypeId,
			&product.SellerId,
		)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *repositoryDB) GetId(id int) (Product, error) {
	var product Product

	stmt, err := r.db.Prepare("SELECT * FROM products Where id = ?")
	if err != nil {
		return product, err
	}
	err = stmt.QueryRow(id).Scan(&product.Id,
		&product.ProductCode,
		&product.Description,
		&product.Width,
		&product.Length,
		&product.Height,
		&product.NetWeight,
		&product.ExpirationRate,
		&product.RecommendedFreezingTemperature,
		&product.FreezingRate,
		&product.ProductTypeId,
		&product.SellerId)
	defer stmt.Close()

	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (r *repositoryDB) CheckCode(code string) error {

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
}

func (r *repositoryDB) Delete(id int) error {

	stmt, err := r.db.Prepare("DELETE FROM products WHERE id=?")
	if err != nil {
		return err
	}

	defer stmt.Close() // Impedir vazamento de memória

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return fmt.Errorf("product %d not found", id)
	}

	return nil
}

func (r *repositoryDB) Create(p Product) (Product, error) {

	stmt, err := r.db.Prepare(`INSERT INTO products (
		product_code, 
		description, 
		width, 
		length,	
		height,	
		net_weight,	
		expiration_rate, 
		recommended_freezing_temperature, 
		freezing_rate, 
		product_type_id,
		seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		p.ProductCode,
		p.Description,
		p.Width,
		p.Length,
		p.Height,
		p.NetWeight,
		p.ExpirationRate,
		p.RecommendedFreezingTemperature,
		p.FreezingRate,
		p.ProductTypeId,
		p.SellerId)
	if err != nil {
		return Product{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return Product{}, fmt.Errorf("Fail to save")
	}
	return p, nil
}

func (r *repositoryDB) LastID() (int, error) {
	/*var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}*/
	return 0, nil
}

func (r *repositoryDB) Update(id int, p Product) (Product, error) {
	stmt, err := r.db.Prepare(`UPDATE products SET 
		product_code=?,
		description=?, 
		width=? ,
		length=?,	
		height=?,	
		net_weight=?,	
		expiration_rate=?, 
		recommended_freezing_temperature=? ,
		freezing_rate=? ,
		product_type_id=?,
		seller_id=? WHERE id=?`)

	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		p.ProductCode,
		p.Description,
		p.Width,
		p.Length,
		p.Height,
		p.NetWeight,
		p.ExpirationRate,
		p.RecommendedFreezingTemperature,
		p.FreezingRate,
		p.ProductTypeId,
		p.SellerId,
		id)
	if err != nil {
		return Product{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return Product{}, fmt.Errorf("Fail to save")
	}
	return p, nil
}
