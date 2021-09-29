package main

import (
	"fmt"
	"net/http"
	"time"

	"dmitrook.ru/lenashop/repository"
	"golang.org/x/crypto/bcrypt"
)

// ----------------------------------------------------------------------
// Handle Registration
// ----------------------------------------------------------------------
func registerHandler(w http.ResponseWriter, r *http.Request) {

	type registerPayload struct {
		AuthData
	}

	type registerFailPayload struct {
		AuthData
		Messages []string
	}

	type registerOkPayload struct {
		AuthData
		NewCustomer repository.Customer
	}
	// ----------------------------------------------------------------------
	// GET
	// ----------------------------------------------------------------------
	if r.Method == "GET" {
		isAuth, customer, err := checkAuth(r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Sorry, something is going wrong on the server.", 500)
			return
		}
		render(w, "register.page.html", registerPayload{AuthData{isAuth, customer}})
	}

	// ----------------------------------------------------------------------
	// POST
	// ----------------------------------------------------------------------
	if r.Method == "POST" {
		isAuth, authCustomer, err := checkAuth(r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Sorry, something is going wrong on the server.", 500)
			return
		}

		vMessages, err := validateRegisterForm(r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Bad request", 400)
			return
		}
		if len(vMessages) > 0 {
			render(w, "register-fail.page.html", registerFailPayload{AuthData{isAuth, authCustomer}, vMessages})
			return
		}
		hasCust, err := repository.HasCustomer(r.PostForm.Get("email"))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Sorry, something is going wrong on the server.", 500)
			return
		}
		if hasCust {
			msg := fmt.Sprintf("Customer %q already registered.", r.PostForm.Get("email"))
			render(w, "register-fail.page.html", registerFailPayload{AuthData{isAuth, authCustomer}, []string{msg}})
			return
		}

		token, err := bcrypt.GenerateFromPassword([]byte(r.PostForm.Get("password")), 10)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Sorry, something is going wrong on the server.", 500)
			return
		}

		customer := repository.NewCustomer(
			r.PostForm.Get("email"),
			r.PostForm.Get("phone"),
			r.PostForm.Get("address"),
			string(token),
		)

		if err := repository.AddCustomer(customer); err != nil {
			fmt.Println(err)
			http.Error(w, "Sorry, something is going wrong on the server.", 500)
			return
		}

		exp := time.Now().UTC().Add(time.Minute * 10).Format(http.TimeFormat)
		w.Header().Add("Set-Cookie", fmt.Sprintf("customer=%s; Expires=%v; Secure; HttpOnly;", customer.Email, exp))
		w.Header().Add("Set-Cookie", fmt.Sprintf("token=%v; Expires=%v; Secure; HttpOnly;", customer.Token, exp))
		render(w, "register-ok.page.html", registerOkPayload{AuthData{isAuth, authCustomer}, customer})
	}
}
