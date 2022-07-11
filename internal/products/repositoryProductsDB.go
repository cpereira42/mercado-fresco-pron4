package products

import (
	"database/sql"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"

	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	GetAll() ([]Product, error)
	GetId(id int) (Product, error)
	Delete(id int) error
	Create(p Product) (Product, error)
	Update(id int, prod Product) (Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepositoryProductsDB(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product

	rows, err := r.db.Query(QUERY_GETALL)
	if err != nil {
		return products, util.CheckError(err)
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
			return products, util.CheckError(err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) GetId(id int) (Product, error) {
	var product Product

	stmt, err := r.db.Prepare(QUERY_GETID)
	if err != nil {
		return product, fmt.Errorf("Fail to prepar query")
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
		return Product{}, util.CheckError(err)
	}
	return product, nil
}

func (r *repository) Delete(id int) error {

	stmt, err := r.db.Prepare(QUERY_DELETE)
	if err != nil {
		return util.CheckError(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return util.CheckError(err)
	}
	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}

func (r *repository) Create(p Product) (Product, error) {

	stmt, err := r.db.Prepare(QUERY_INSERT)

	if err != nil {
		return Product{}, util.CheckError(err)
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
		return Product{}, util.CheckError(err)
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return Product{}, fmt.Errorf("Fail to save")
	}
	lastId, _ := res.LastInsertId()
	p.Id = int(lastId)
	return p, nil
}

func (r *repository) Update(id int, p Product) (Product, error) {
	stmt, err := r.db.Prepare(QUERY_UPDATE)

	if err != nil {
		return Product{}, fmt.Errorf("Fail to prepar query")
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
		return Product{}, util.CheckError(err)
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return Product{}, fmt.Errorf("Fail to save")
	}
	return p, nil
}
