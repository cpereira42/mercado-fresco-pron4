package productsRecords

import (
	"database/sql"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
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

	rows, err := r.db.Query(QUERY_GETALL)

	if err != nil {
		return ProductRecords, util.CheckError(err)
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
			return ProductRecords, fmt.Errorf("Fail to scan")
		}
		ProductRecords = append(ProductRecords, ProductRecord)
	}

	return ProductRecords, nil
}

func (r *repository) GetIdRecords(id int) (ReturnProductRecords, error) {

	var ProductRecords ReturnProductRecords

	stmt, err := r.db.Prepare(QUERY_GETID)

	if err != nil {
		return ProductRecords, util.CheckError(err)
	}

	err = stmt.QueryRow(id).Scan(
		&ProductRecords.ProductId,
		&ProductRecords.Description,
		&ProductRecords.RecordsCount,
	)

	if err != nil {
		return ProductRecords, util.CheckError(err)
	}
	defer stmt.Close()
	return ProductRecords, nil
}

func (r *repository) Create(p ProductRecords) (ProductRecords, error) {

	stmt, err := r.db.Prepare(QUERY_INSERT)

	if err != nil {
		return ProductRecords{}, util.CheckError(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		p.LastUpdateDate,
		p.PurchasePrice,
		p.SalePrice,
		p.ProductId)
	if err != nil {
		return ProductRecords{}, util.CheckError(err)
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return ProductRecords{}, fmt.Errorf("Fail to save")
	}

	lastId, _ := res.LastInsertId()
	p.Id = int(lastId)
	return p, nil
}
