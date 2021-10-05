package main

import (
	"fmt"
	"log"
	"net/http"

	"dmitrook.ru/lenashop/cart"
	"dmitrook.ru/lenashop/customers"
	"dmitrook.ru/lenashop/home"
	"dmitrook.ru/lenashop/orders"
	"dmitrook.ru/lenashop/products"
)

func serve() {

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", home.ViewHandler)
	http.HandleFunc("/products", products.ViewAllHandler)
	http.HandleFunc("/product/", products.ViewSingleHandler)

	http.HandleFunc("/register", customers.RegisterHandler)
	http.HandleFunc("/login", customers.LoginHandler)
	http.HandleFunc("/logout", customers.LogoutHandler)

	http.HandleFunc("/customers", customers.ViewAllHandler)
	http.HandleFunc("/customers/create", customers.CreateHandler)
	http.HandleFunc("/customers/update", customers.UpdateHandler)
	http.HandleFunc("/customers/delete", customers.DeleteHandler)
	http.HandleFunc("/customer/", customers.ViewSingleHandler)

	http.HandleFunc("/cart/products", cart.ViewHandler)
	http.HandleFunc("/cart/products/add", cart.AddProductHandler)
	http.HandleFunc("/cart/products/remove", cart.RemoveProductHandler)
	http.HandleFunc("/cart/products/update", cart.UpdateQuantityHandler)

	http.HandleFunc("/orders", orders.ViewAllHandler)
	http.HandleFunc("/checkout", orders.CheckoutHandler)
	http.HandleFunc("/orders/create", orders.CreateHandler)
	http.HandleFunc("/orders/delete", orders.DeleteHandler)

	// http.HandleFunc("/products/create", adminCreateProductHandler)
	// http.HandleFunc("/products/update", adminUpdateProductHandler)
	// http.HandleFunc("/products/delete", adminDeleteProductHandler)

	fmt.Println("--- Server listen at 127.0.0.1:4000 ---")
	if err := http.ListenAndServe("127.0.0.1:4000", nil); err != nil {
		log.Fatal(err)
	}
}
