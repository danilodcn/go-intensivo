package database

import (
	"database/sql"

	"github.com/danilodcn/go-intensivo/internal/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	smtp, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES ($1, $2, $3, $4)")

	if err != nil {
		return err
	}

	_, err = smtp.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)

	return err
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("select count(*) from orders").
		Scan(total)

	if err != nil {
		return 0, err
	}
	return total, nil
}
