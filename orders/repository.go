package orders

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"dmitrook.ru/lenashop/common"
	"dmitrook.ru/lenashop/products"
)

type Order struct {
	OrderId    int
	CustomerId int
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
}

type CartProduct struct {
	ProductId   int
	BuyQuantity int
}

func Orders(customerId int) ([]Order, error) {
	ctx := context.Background()
	rows, err := common.DbPool.Query(ctx, "SELECT * FROM customer_order WHERE customer_id = $1", customerId)
	if err != nil {
		return nil, fmt.Errorf("Orders(%v): %v", customerId, err)
	}
	defer rows.Close()
	var ords []Order
	for rows.Next() {
		var ord Order
		if err := rows.Scan(&ord.OrderId, &ord.CustomerId, &ord.CreatedAt, &ord.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Orders(%v): %v", customerId, err)
		}
		ords = append(ords, ord)
	}
	return ords, nil
}

func OrderById(orderId int) (Order, error) {
	ctx := context.Background()
	row := common.DbPool.QueryRow(ctx, "SELECT * FROM customer_order WHERE order_id = $1", orderId)
	var ord Order
	if err := row.Scan(&ord.OrderId, &ord.CustomerId, &ord.CreatedAt, &ord.UpdatedAt); err != nil {
		return ord, fmt.Errorf("OrderById(%v): %v", orderId, err)
	}
	return ord, nil
}

func CreateOrder(customerId int, cartId int) (int, error) {
	var orderId int
	ctx := context.Background()
	fail := func(err error) error {
		return fmt.Errorf("CreateOrder(%v, %v): %v", customerId, cartId, err)
	}
	tx, err := common.DbPool.Begin(ctx)
	if err != nil {
		return -1, fail(err)
	}
	defer tx.Rollback(ctx)
	// 1. Create customer_order
	// ----------------------------------------------------
	row := tx.QueryRow(ctx, "INSERT INTO customer_order (customer_id, created_at) VALUES ($1, $2) RETURNING order_id", customerId, time.Now())
	if err := row.Scan(&orderId); err != nil {
		return -1, fail(err)
	}
	// 2. Select products from cart
	// ----------------------------------------------------
	rows, err := tx.Query(ctx, "SELECT product_id, quantity FROM cart_has_product WHERE cart_id = $1", cartId)
	if err != nil {
		return -1, fail(err)
	}
	defer rows.Close()
	var cps []CartProduct
	for rows.Next() {
		var cp CartProduct
		if err := rows.Scan(&cp.ProductId, &cp.BuyQuantity); err != nil {
			return -1, fail(err)
		}
		cps = append(cps, cp)
	}
	// 3. Add each product to order
	// ----------------------------------------------------
	for _, cp := range cps {
		_, err := tx.Exec(ctx, "INSERT INTO customer_order_has_product (order_id, product_id, quantity) VALUES ($1, $2, $3)", orderId, cp.ProductId, cp.BuyQuantity)
		if err != nil {
			return -1, fail(err)
		}
	}
	// 4. Clear cart
	// ----------------------------------------------------
	_, err = tx.Exec(ctx, "DELETE FROM cart_has_product WHERE cart_id = $1", cartId)
	if err != nil {
		return -1, fail(err)
	}
	if err := tx.Commit(ctx); err != nil {
		return -1, fail(err)
	}
	return orderId, nil
}

func LastCreated(customerId int) (Order, error) {
	ctx := context.Background()
	var order Order
	row := common.DbPool.QueryRow(ctx, "SELECT * FROM customer_order WHERE customer_id = $1 ORDER BY order_id DESC LIMIT 1", customerId)
	if err := row.Scan(&order.OrderId, &order.CustomerId, &order.CreatedAt, &order.UpdatedAt); err != nil {
		return order, fmt.Errorf("LastCreatedOrder(%v): %v", customerId, err)
	}
	return order, nil
}

func Products(orderId int) ([]products.FullProduct, error) {
	ctx := context.Background()
	fail := func(err error) error {
		return fmt.Errorf("OrderProducts(%v): %v", orderId, err)
	}
	sqlQuery := `
	SELECT p.* FROM customer_order_has_product cohp
	INNER JOIN product p ON cohp.product_id = p.product_id
	WHERE cohp.order_id = $1
	`
	rows, err := common.DbPool.Query(ctx, sqlQuery, orderId)
	if err != nil {
		return nil, fail(err)
	}
	var fps []products.FullProduct
	for rows.Next() {
		var fp products.FullProduct
		if err := rows.Scan(&fp.ProductId, &fp.CategoryId, &fp.Title, &fp.Price, &fp.Quantity, &fp.Description, &fp.CreatedAt, &fp.UpdatedAt); err != nil {
			return nil, fail(err)
		}
		images, err := products.ImagesByProductId(fp.ProductId)
		if err != nil {
			return nil, fail(err)
		}
		fp.Images = images
		fps = append(fps, fp)
	}
	return fps, nil
}
