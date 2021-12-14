package cart

import (
	"context"
	"fmt"
	"time"

	"dmitrook.ru/lenashop/common"
	"dmitrook.ru/lenashop/products"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Model struct {
	db pgxpool.Pool
}

type Cart struct {
	CartId     int
	CustomerId int
	CreateAt   time.Time
}

type FullCart struct {
	CartId     int
	CustomerId int
	Products   []CartProduct
	CreateAt   time.Time
}

type CartProduct struct {
	products.FullProduct
	BuyQuantity int
}

type CartBL struct {
	Cart
	Products []CartProduct
}

// --------------------------------------------------------------------

func Create(customerId int) (cartId int, err error) {
	ctx := context.Background()
	const q = `
		INSERT INTO cart (customer_id, created_at) VALUES ($1, $2)
		RETURNING cart_id
	`
	row := common.DbPool.QueryRow(ctx, q, customerId, time.Now())
	if err := row.Scan(&cartId); err != nil {
		return cartId, fmt.Errorf("Create(%v): %v", customerId, err)
	}
	return cartId, nil
}

func Update(cart FullCart) error {
	ctx := context.Background()
	tx, err := common.DbPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("Update(%v): %v", cart, err)
	}
	defer tx.Rollback(ctx)
	q := `
		UPDATE cart_has_product
		SET quantity = $1
		WHERE cart_id = $2 AND product_id = $3
	`
	for _, p := range cart.Products {
		// --- bussiness logic! check buy and stock quantity
		_, err := tx.Exec(ctx, q, p.BuyQuantity, cart.CartId, p.ProductId)
		if err != nil {
			return fmt.Errorf("Update(%v): %v", cart, err)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("Update(%v): %v", cart, err)
	}
	return nil
}

// --------------------------------------------------------------------

func ById(cartId int) (Cart, error) {
	ctx := context.Background()
	const q = `
		SELECT * 
		FROM cart
		WHERE cart_id = $1
	`
	row := common.DbPool.QueryRow(ctx, q, cartId)
	var cart Cart
	if err := row.Scan(&cart.CartId, &cart.CustomerId, &cart.CreateAt); err != nil {
		return cart, fmt.Errorf("CartById(%v): %v", cartId, err)
	}
	return cart, nil
}

func ByCustomerId(customerId int) (Cart, error) {
	ctx := context.Background()
	const q = `
		SELECT * 
		FROM cart
		WHERE customer_id = $1
	`
	row := common.DbPool.QueryRow(ctx, q, customerId)
	var cart Cart
	if err := row.Scan(&cart.CartId, &cart.CustomerId, &cart.CreateAt); err != nil {
		return cart, fmt.Errorf("CartByCustomerId(%v): %v", customerId, err)
	}
	return cart, nil
}

func AddProduct(cartId int, productId int, buyQuantity int) error {
	ctx := context.Background()
	const q = `
		INSERT INTO cart_has_product (cart_id, product_id, quantity)
		VALUES ($1, $2, $3)
	`
	_, err := common.DbPool.Exec(ctx, q, cartId, productId, buyQuantity)
	if err != nil {
		// -------------------------------------------------
		// TODO Something with same product already in cart
		// -------------------------------------------------
		if err.(*pgconn.PgError).Code == "23505" {
			return nil
		}
		return fmt.Errorf("AddProduct(%v, %v, %v): %v", cartId, productId, buyQuantity, err)
	}
	return nil
}

func RemoveProduct(cartId int, productId int) error {
	ctx := context.Background()
	const q = `
		DELETE FROM cart_has_product
		WHERE cart_id = $1 AND product_id = $2
	`
	_, err := common.DbPool.Exec(ctx, q, cartId, productId)
	if err != nil {
		return fmt.Errorf("CartRemoveProduct(%v, %v): %v", cartId, productId, err)
	}
	return nil
}

func UpdateProductQuantity(cartId int, productId int, buyQuantity int) error {
	ctx := context.Background()
	const q1 = `
		SELECT quantity
		FROM product
		WHERE product_id = $1
	`
	row := common.DbPool.QueryRow(ctx, q1, productId)
	var stockQuantity int
	if err := row.Scan(&stockQuantity); err != nil {
		return fmt.Errorf("UpdateProductQuantity(%v, %v, %v): %v", cartId, productId, buyQuantity, err)
	}
	// --- bussiness logic! prevent negative quantity
	if buyQuantity < 1 {
		buyQuantity = 1
	}
	// --- bussiness logic! prevent buying more than stock
	if buyQuantity > stockQuantity {
		buyQuantity = stockQuantity
	}
	// --- bussiness logic! if stock is empty
	if stockQuantity == 0 {
		return fmt.Errorf("UpdateProductQuantity(%v, %v, %v): ", cartId, productId, buyQuantity)
	}

	const q2 = `
		UPDATE cart_has_product
		SET quantity = $3
		WHERE cart_id = $1 AND product_id = $2
	`
	_, err := common.DbPool.Exec(ctx, q2, cartId, productId, buyQuantity)
	if err != nil {
		return fmt.Errorf("UpdateProductQuantity(%v, %v, %v): %v", cartId, productId, buyQuantity, err)
	}
	return nil
}

func Products(cartId int) ([]CartProduct, error) {
	ctx := context.Background()
	fail := func(err error) error {
		return fmt.Errorf("CartProducts(%v): %v", cartId, err)
	}
	// ---------------------------------------------------
	// 1. Get all products in cart
	// ---------------------------------------------------
	const q = `
		SELECT chp.product_id, chp.quantity, pr.category_id, pr.title, pr.price, pr.quantity, pr.description, pr.created_at
		FROM cart_has_product chp
		INNER JOIN product pr ON chp.product_id = pr.product_id
		WHERE cart_id = $1
		ORDER BY pr.product_id
	`
	rows, err := common.DbPool.Query(ctx, q, cartId)
	if err != nil {
		return nil, fail(err)
	}
	var cartProducts []CartProduct
	for rows.Next() {
		var cpr CartProduct
		if err := rows.Scan(&cpr.ProductId, &cpr.BuyQuantity, &cpr.CategoryId, &cpr.Title, &cpr.Price, &cpr.Quantity, &cpr.Description, &cpr.CreatedAt); err != nil {
			return nil, fail(err)
		}
		// ---------------------------------------------------
		// 2. Get all images for each product in cart
		// ---------------------------------------------------
		images, err := products.ImagesByProductId(cpr.ProductId)
		if err != nil {
			return nil, fail(err)
		}
		cpr.Images = images
		cartProducts = append(cartProducts, cpr)
	}
	if err := rows.Err(); err != nil {
		return nil, fail(err)
	}
	return cartProducts, nil
}
