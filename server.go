package main

import (
	"fmt"
	"log"
	"net/http"

	"dmitrook.ru/lenashop/cart"
	"dmitrook.ru/lenashop/common"
	"dmitrook.ru/lenashop/customers"
	"dmitrook.ru/lenashop/home"
	"dmitrook.ru/lenashop/orders"
	"dmitrook.ru/lenashop/products"
)

type Server struct {
	mux http.ServeMux
}

func (s *Server) WithValidator(v common.Validator, handler http.Handler) http.Handler {
	// return func(w http.ResponseWriter, r *http.Request) {}
	return nil
}

// type AddProductValidator struct{}

// func (apv AddProductValidator) ValidateQuery(r *http.Request)   {}
// func (apv AddProductValidator) ValidateForm(r *http.Request)    {}
// func (apv AddProductValidator) ValidateCookies(r *http.Request) {}

func serve() {

	// ------------------------------------------------------------------------
	// s := Server{}
	// apv := AddProductValidator{}

	// handler := s.WithValidator(apv, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	apv.ValidateCookies(r)
	// }))
	// s.mux.Handle("/cart/products/add", handler)
	// ------------------------------------------------------------------------

	mux := http.DefaultServeMux

	// mux.Handle("/bla", authMiddleware(http.HandlerFunc(blaHandler), "Valera"))

	fs := http.FileServer(http.Dir("./static/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", home.ViewHandler)
	mux.HandleFunc("/products", products.ViewAllHandler)
	mux.HandleFunc("/product/", products.ViewSingleHandler)

	mux.HandleFunc("/register", customers.RegisterHandler)
	mux.HandleFunc("/login", customers.LoginHandler)
	mux.HandleFunc("/logout", customers.LogoutHandler)

	mux.HandleFunc("/customers", customers.ViewAllHandler)
	mux.HandleFunc("/customers/create", customers.CreateHandler)
	mux.HandleFunc("/customers/update", customers.UpdateHandler)
	mux.HandleFunc("/customers/delete", customers.DeleteHandler)
	mux.HandleFunc("/customer/", customers.ViewSingleHandler)

	mux.HandleFunc("/cart/products", cart.ViewHandler)
	mux.HandleFunc("/cart/products/add", cart.AddProductHandler)
	mux.HandleFunc("/cart/products/remove", cart.RemoveProductHandler)
	mux.HandleFunc("/cart/products/update", cart.UpdateQuantityHandler)

	mux.HandleFunc("/orders", orders.ViewAllHandler)
	mux.HandleFunc("/checkout", orders.CheckoutHandler)
	mux.HandleFunc("/orders/create", orders.CreateHandler)
	mux.HandleFunc("/orders/delete", orders.DeleteHandler)

	// http.HandleFunc("/products/create", adminCreateProductHandler)
	// http.HandleFunc("/products/update", adminUpdateProductHandler)
	// http.HandleFunc("/products/delete", adminDeleteProductHandler)

	fmt.Println("--- Server listen at 127.0.0.1:4000 ---")
	if err := http.ListenAndServe("0.0.0.0:4000", nil); err != nil {
		log.Fatal(err)
	}
}
