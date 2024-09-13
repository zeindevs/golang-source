package order

import (
	"database/sql"

	"github.com/zeindevs/go-ecommerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateOrder(o *types.Order) (int, error) {
	var id int
	if err := s.db.QueryRow("INSERT INTO orders (user_id, total, status, address) VALUES ($1, $2, $3, $4) RETURNING id", o.UserID, o.Total, o.Status, o.Address).
		Scan(&id); err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Store) CreateOrderItem(oi *types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO orders_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", oi.OrderID, oi.ProductID, oi.Quantity, oi.Price)
	return err
}
