package purchaseorders

import (
	"database/sql"
	"fmt"

	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
)

type Repository interface {
	GetById(id int) (Purchase, error)
	Create(order_date, order_number, tracking_code string, buyer_id, product_record_id, order_status_id int) (Purchase, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetById(id int) (Purchase, error) {
	row := r.db.QueryRow(GET_PURCHASE_BY_ID)

	purchase := Purchase{}

	err := row.Scan(
		&purchase.ID,
		&purchase.Order_date,
		&purchase.Order_number,
		&purchase.Tracking_code,
		&purchase.Buyer_id,
		&purchase.Product_record_id,
		&purchase.Order_status_id,
	)

	if err != nil {
		return Purchase{}, fmt.Errorf("fail to get id")
	}

	return purchase, nil
}

func (r *repository) Create(
	order_date,
	order_number,
	tracking_code string,
	buyer_id,
	product_record_id,
	order_status_id int) (Purchase, error) {
	purchase := Purchase{
		Order_date:        order_date,
		Order_number:      order_number,
		Tracking_code:     tracking_code,
		Buyer_id:          buyer_id,
		Product_record_id: product_record_id,
		Order_status_id:   order_status_id,
	}

	stmt, err := r.db.Exec(CREATE_PURCHASE,
		order_date,
		order_number,
		tracking_code,
		buyer_id,
		product_record_id,
		order_status_id,
	)

	if err != nil {
		return Purchase{}, util.CheckError(err)
	}

	rowsAffected, _ := stmt.RowsAffected()
	if rowsAffected == 0 {
		return Purchase{}, fmt.Errorf("fail to create")
	}

	lastId, _ := stmt.LastInsertId()

	purchase.ID = int(lastId)

	return purchase, nil
}
