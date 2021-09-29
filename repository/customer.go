package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type Customer struct {
	CustomerId int64
	Email      string
	Phone      sql.NullString
	Address    string
	Token      string
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
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

func AllCustomers() ([]Customer, error) {
	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		return nil, fmt.Errorf("AllCustomers(): %v", err)
	}
	var cs []Customer
	var c Customer
	for rows.Next() {
		if err := rows.Scan(&c.CustomerId, &c.Email, &c.Phone, &c.Address, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, fmt.Errorf("AllCustomers(): %v", err)
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func CustomerById(id int64) (Customer, error) {
	row := db.QueryRow("SELECT * FROM customer WHERE customer_id = $1", id)
	var c Customer
	if err := row.Scan(&c.CustomerId, &c.Email, &c.Phone, &c.Address, &c.CreatedAt, &c.UpdatedAt); err != nil {
		return c, fmt.Errorf("CustomerById(%v): %v", id, err)
	}
	return c, nil
}

func CustomerByEmail(email string) (Customer, error) {
	row := db.QueryRow("SELECT * FROM customer WHERE email = $1", email)
	var c Customer
	err := row.Scan(&c.CustomerId, &c.Email, &c.Phone, &c.Address, &c.Token, &c.CreatedAt, &c.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return c, fmt.Errorf("CustomerByEmail(%v): %v", email, err)
	}
	return c, nil
}

func HasCustomer(email string) (bool, error) {
	row := db.QueryRow("SELECT customer_id FROM customer WHERE email = $1", email)
	var id int64
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("HasCustomer(%v): %v", email, err)
	}
	return true, nil
}

func AddCustomer(c Customer) error {
	var processError = func(tx *sql.Tx, err error) error {
		if txErr := tx.Rollback(); txErr != nil {
			return fmt.Errorf("AddCustomer(%v): %v", c, txErr)
		}
		return fmt.Errorf("AddCustomer(%v): %v", c, err)
	}
	tx, err := db.Begin()
	if err != nil {
		return processError(tx, err)
	}
	/*customerResult*/ res, err := tx.Exec("INSERT INTO customer (email, phone, address, token, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING customer_id", c.Email, c.Phone, c.Address, c.Token, time.Now())
	if err != nil {
		return processError(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

	/*cartResult*/
	_, err = tx.Exec("INSERT INTO cart (customer_id, created_at) VALUES ($1, $2)", c.CustomerId, time.Now())
	if err != nil {
		return processError(tx, err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("AddCustomer(%v): %v", c, err)
	}
	return nil
}
