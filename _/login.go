package main

import (
	"fmt"
	"net/http"
	"time"

	"dmitrook.ru/lenashop/repository"
	"golang.org/x/crypto/bcrypt"
)

// ----------------------------------------------------------------------
// Handle Login
// ----------------------------------------------------------------------
func loginHandler(w http.ResponseWriter, r *http.Request) {

	// GET --------------------------------------------------
	if r.Method == "GET" {
		render(w, "login.page.html", nil)
	}

	// POST --------------------------------------------------
	if r.Method == "POST" {
		customer, err := repository.CustomerByEmail(r.PostFormValue("email"))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Sorry, something is going wrong on the server.", 500)
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(customer.Token), []byte(r.PostFormValue("password"))); err != nil {
			render(w, "login-fail.page.html", nil)
			return
		}
		exp := time.Now().UTC().Add(time.Minute * 10).Format(http.TimeFormat)
		w.Header().Add("Set-Cookie", fmt.Sprintf("customer=%s; Expires=%v; Secure; HttpOnly;", customer.Email, exp))
		w.Header().Add("Set-Cookie", fmt.Sprintf("token=%v; Expires=%v; Secure; HttpOnly;", customer.Token, exp))
		http.Redirect(w, r, "/catalog", http.StatusFound)
	}
}
