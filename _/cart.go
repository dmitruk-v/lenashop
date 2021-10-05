package main

import (
	"log"
	"net/http"
	"strconv"

	"dmitrook.ru/lenashop/repository"
)

type cartPayload struct {
	AuthData
	Products []repository.CartProduct
}

func cartHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		isAuth, customer, err := checkAuth(r)
		if err != nil {
			log.Println(err)
			http.Error(w, "Something wrong happened, when checking authentication", 500)
			return
		}
		if !isAuth {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		products, err := repository.CartProducts(customer.CustomerId)
		if err != nil {
			log.Println(err)
			http.Error(w, "Something wrong happened, when fetching cart products", 500)
			return
		}
		render(w, "cart.page.html", cartPayload{AuthData{isAuth, customer}, products})
	}

	if r.Method == "POST" {
		isAuth, customer, err := checkAuth(r)
		if err != nil {
			log.Println(err)
			http.Error(w, "Something wrong happened, when checking authentication", 500)
			return
		}

		if !isAuth {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		if err := r.ParseForm(); err != nil {
			log.Println(err)
			http.Error(w, "Something wrong happened, when parsing form", 400)
			return
		}

		productId, err1 := strconv.ParseInt(r.PostForm.Get("product_id"), 10, 64)
		quantity, err2 := strconv.ParseInt(r.PostForm.Get("quantity"), 10, 64)
		if err1 != nil || err2 != nil {
			log.Println(err1, err2)
			http.Error(w, "Something wrong happened, when parsing form values", 400)
			return
		}

		if err := repository.CartAddProduct(customer.CustomerId, int(productId), int(quantity)); err != nil {
			log.Println(err)
			http.Error(w, "Something wrong happened, when adding product to cart", 500)
			return
		}

		http.Redirect(w, r, "/cart", http.StatusFound)
	}
}
