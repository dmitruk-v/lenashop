package main

import (
	"fmt"
	"net/http"
)

type homePayload struct {
	AuthData
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	isAuth, customer, err := checkAuth(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Sorry, something is going wrong on the server.", 500)
		return
	}

	if !isAuth {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	render(w, "home.page.html", homePayload{AuthData{isAuth, customer}})
}
