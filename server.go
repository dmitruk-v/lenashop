package main

import (
	"fmt"
	"log"
	"net/http"
)

func serve() {
	// static files handler
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// auth
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	// commerce
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/catalog", catalogHandler)
	http.HandleFunc("/product", productHandler)
	http.HandleFunc("/order", orderHandler)
	http.HandleFunc("/cart", cartHandler)

	fmt.Println("--- Server listen at 127.0.0.1:4000 ---")
	if err := http.ListenAndServe("127.0.0.1:4000", nil); err != nil {
		log.Fatal(err)
	}
}
