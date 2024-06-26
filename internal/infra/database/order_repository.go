package database

import (
	"database/sql"

	"github.com/allanmaral/go-expert-clean-arch-challenge/internal/entity"
)

type OrderRepository struct {
	db *sql.DB
}

// Make sure OrderRepository implements the interface
var _ entity.OrderRepositoryInterface = (*OrderRepository)(nil)

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetAll() ([]entity.Order, error) {
	rows, err := r.db.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(
			&order.ID,
			&order.Price,
			&order.Tax,
			&order.FinalPrice,
		); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
