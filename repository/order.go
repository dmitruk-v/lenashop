package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type OrderStatus int

const (
	OrderOpened OrderStatus = iota
	OrderClosed
	OrderFailed
)

type CustomerOrder struct {
	OrderId    int
	CustomerId int
	Status     OrderStatus
	CreateAt   time.Time
	UpdatedAt  sql.NullTime
}

func Orders() {}

func OrdersByCustomer(customer Customer) ([]CustomerOrder, error) {
	rows, err := dbPool.Query(context.Background(), "SELECT * FROM customer_order WHERE customer_id = $1", customer.CustomerId)
	if err != nil {
		return nil, fmt.Errorf("OrdersByCustomer(%v): %v", customer, err)
	}
	var orders []CustomerOrder
	for rows.Next() {
		var order CustomerOrder
		if err := rows.Scan(&order.OrderId, &order.CustomerId, &order.Status, &order.CreateAt, &order.UpdatedAt); err != nil {
			return nil, fmt.Errorf("OrdersByCustomer(%v): %v", customer, err)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func AddOrder(order CustomerOrder) error {
	_, err := dbPool.Exec(context.Background(), "INSERT INTO customer_order (customer_id, created_at) VALUES ($1, $2)", order.CustomerId, time.Now())
	if err != nil {
		return fmt.Errorf("AddOrder(%v): %v", order, err)
	}
	return nil
}

func HasOpenedOrder(customer Customer) {
	// row := db.QueryRow("SELECT * FROM customer_order WHERE status = $1", OrderOpened)
}
