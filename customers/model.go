package customers

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"dmitrook.ru/lenashop/common"
	"github.com/jackc/pgx/v4"
)

type Customer struct {
	CustomerId int
	Email      string
	Phone      sql.NullString
	Address    string
	Token      string
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
}

type CustomerWithCart struct {
	Customer
	CartId int
}

type FullCustomer struct {
	CustomerId int
	Email      string
	Phone      sql.NullString
	Address    string
	Token      string
	Cart       struct {
		CartId   int
		Products []struct {
			ProductId     int
			BuyQuantity   int
			StockQuantity int
		}
	}
}

func NewCustomer(email string, phone string, address string, token string) Customer {
	c := Customer{
		Email:   email,
		Phone:   sql.NullString{String: phone},
		Address: address,
		Token:   token,
	}
	return c
}

func Customers() ([]Customer, error) {
	ctx := context.Background()
	rows, err := common.DbPool.Query(ctx, "SELECT * FROM customer")
	if err != nil {
		return nil, fmt.Errorf("AllCustomers(): %v", err)
	}
	var cs []Customer
	var c Customer
	for rows.Next() {
		if err := rows.Scan(&c.CustomerId, &c.Email, &c.Phone, &c.Address, &c.Token, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, fmt.Errorf("AllCustomers(): %v", err)
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func CustomerById(id int) (Customer, error) {
	ctx := context.Background()
	row := common.DbPool.QueryRow(ctx, "SELECT * FROM customer WHERE customer_id = $1", id)
	var c Customer
	if err := row.Scan(&c.CustomerId, &c.Email, &c.Phone, &c.Address, &c.Token, &c.CreatedAt, &c.UpdatedAt); err != nil {
		return c, fmt.Errorf("CustomerById(%v): %v", id, err)
	}
	return c, nil
}

func CustomerByEmail(email string) (Customer, error) {
	ctx := context.Background()
	row := common.DbPool.QueryRow(ctx, "SELECT * FROM customer WHERE email = $1", email)
	var c Customer
	err := row.Scan(&c.CustomerId, &c.Email, &c.Phone, &c.Address, &c.Token, &c.CreatedAt, &c.UpdatedAt)
	if err != nil && err != pgx.ErrNoRows {
		return c, fmt.Errorf("CustomerByEmail(%v): %v", email, err)
	}
	return c, nil
}

func CustomerWithCartByEmail(email string) (CustomerWithCart, error) {
	ctx := context.Background()
	q := `
		SELECT customer.*, cart.cart_id
		FROM customer
		INNER JOIN cart ON customer.customer_id = cart.customer_id
		WHERE email = $1
	`
	row := common.DbPool.QueryRow(ctx, q, email)
	var c CustomerWithCart
	err := row.Scan(&c.CustomerId, &c.Email, &c.Phone, &c.Address, &c.Token, &c.CreatedAt, &c.UpdatedAt, &c.CartId)
	if err != nil && err != pgx.ErrNoRows {
		return c, fmt.Errorf("CustomerWithCartByEmail(%v): %v", email, err)
	}
	return c, nil
}

func CustomerExists(email string) (bool, error) {
	ctx := context.Background()
	row := common.DbPool.QueryRow(ctx, "SELECT customer_id FROM customer WHERE email = $1", email)
	var id int64
	if err := row.Scan(&id); err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("CustomerExists(%v): %v", email, err)
	}
	return true, nil
}

func CreateCustomer(c Customer) error {
	ctx := context.Background()
	// Error handler function.
	fail := func(err error) error {
		return fmt.Errorf("CreateCustomer(%v): %v", c, err)
	}
	tx, err := common.DbPool.Begin(ctx)
	if err != nil {
		return fail(err)
	}
	// If transaction is commited successfully, it will not rollback.
	// So it is safe to defer Rollback call.
	defer tx.Rollback(ctx)

	var createdCustomerId int
	row := tx.QueryRow(ctx, "INSERT INTO customer (email, phone, address, token, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING customer_id", c.Email, c.Phone, c.Address, c.Token, time.Now())
	if err != nil {
		return fail(err)
	}
	if err := row.Scan(&createdCustomerId); err != nil {
		return fail(err)
	}
	var createdCartId int
	row = tx.QueryRow(ctx, "INSERT INTO cart (customer_id, created_at) VALUES ($1, $2) RETURNING cart_id", createdCustomerId, time.Now())
	if err != nil {
		return fail(err)
	}
	if err := row.Scan(&createdCartId); err != nil {
		return fail(err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fail(err)
	}
	return nil
}

func DeleteCustomer(cid int) error {
	ctx := context.Background()
	// error helper function
	fail := func(err error) error {
		return fmt.Errorf("DeleteCustomer(%v): %v", cid, err)
	}
	tx, err := common.DbPool.Begin(ctx)
	if err != nil {
		return fail(err)
	}
	// If transaction is commited successfully, it will not rollback.
	// So it is safe to defer Rollback call.
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM cart WHERE customer_id = $1", cid)
	if err != nil {
		return fail(err)
	}
	_, err = tx.Exec(ctx, "DELETE FROM customer WHERE customer_id = $1", cid)
	if err != nil {
		return fail(err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fail(err)
	}
	return nil
}
