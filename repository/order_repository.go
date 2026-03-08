package repository

import (
	"context"
	"database/sql"
	"golang-api/domain"
	"golang-api/domain/web"
)

type OrderRepository interface {
	Save(ctx context.Context, tx *sql.Tx, order web.Order) web.Order
	SaveItem(ctx context.Context, tx *sql.Tx, item web.OrderItem)
	FindAll(ctx context.Context, db *sql.DB) []domain.Order
}

type orderRepository struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) Save(ctx context.Context, tx *sql.Tx, order web.Order) web.Order {

	query := "INSERT INTO orderss(customer_name, total, payment, status, created_at) VALUES (?, ?, ?, ?, NOW())"

	result, err := tx.ExecContext(ctx, query,
		order.CustomerName,
		order.Total,
		order.Payment,
		order.Status,
	)

	if err != nil {
		panic(err)
	}

	id, _ := result.LastInsertId()
	order.ID = int(id)

	return order
}

func (r *orderRepository) SaveItem(ctx context.Context, tx *sql.Tx, item web.OrderItem) {

	query := "INSERT INTO order_items(product_id, quantity) VALUES (?, ?)"

	_, err := tx.ExecContext(ctx, query,
		item.ProductID,
		item.Quantity,
	)

	if err != nil {
		panic(err)
	}
}

func (r *orderRepository) FindAll(ctx context.Context, db *sql.DB) []domain.Order {

	query := "SELECT id, customer_name, total, payment, status, created_at FROM orderss"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var orders []domain.Order

	for rows.Next() {
		order := domain.Order{}

		err := rows.Scan(
			&order.ID,
			&order.CustomerName,
			&order.Total,
			&order.Payment,
			&order.Status,
			&order.Created_at,
		)

		if err != nil {
			panic(err)
		}

		orders = append(orders, order)
	}

	return orders
}