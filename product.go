package main

import (
	"fmt"
	"net/http"

	"dmitrook.ru/lenashop/repository"
)

type productPayload struct {
	AuthData
	Product repository.Product
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	isAuth, authCustomer, err := checkAuth(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Sorry, something is going wrong on the server.", 500)
		return
	}
	render(w, "product.page.html", productPayload{AuthData{isAuth, authCustomer}, repository.Product{}})
}
