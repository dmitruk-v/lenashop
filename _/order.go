package main

import (
	"fmt"
	"log"
	"net/http"
)

func orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}
		id := r.PostForm.Get("product_id")
		fmt.Println(id)
		// TODO Check if customer has not completed order.
		// If he has, then add product to that order, otherwise create new order and add product to it
	}
}
