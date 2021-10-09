package customers

import (
	"fmt"
	"net/http"
	"time"

	"dmitrook.ru/lenashop/tools"
	"golang.org/x/crypto/bcrypt"
)

type registerFailPayload struct {
	AuthData
	Messages []string
}

type registerOkPayload struct {
	AuthData
	NewCustomer Customer
}

// Makes a handler to handle creation of customer
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------
	// POST: Register customer
	// ----------------------------------------------------------------------
	if r.Method == "POST" {
		authData, err := CheckAuth(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when checking auth: %v", err), 500)
			return
		}
		vMessages, err := ValidateRegisterForm(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when validating form: %v", err), 400)
			return
		}

		var failPayload = registerFailPayload{
			AuthData: authData,
		}

		if len(vMessages) > 0 {
			failPayload.Messages = vMessages
			tools.Render(w, "register-fail.page.html", failPayload)
			return
		}
		custExists, err := CustomerExists(r.PostForm.Get("email"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when checking customer existence: %v", err), 500)
			return
		}
		if custExists {
			msg := fmt.Sprintf("Customer %q already registered.", r.PostForm.Get("email"))
			failPayload.Messages = []string{msg}
			tools.Render(w, "register-fail.page.html", failPayload)
			return
		}

		token, err := bcrypt.GenerateFromPassword([]byte(r.PostForm.Get("password")), 10)
		if err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when checking token: %v", err), 500)
			return
		}

		customer := NewCustomer(
			r.PostForm.Get("email"),
			r.PostForm.Get("phone"),
			r.PostForm.Get("address"),
			string(token),
		)

		if err := CreateCustomer(customer); err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when creating customer: %v", err), 500)
			return
		}

		var okPayload = registerOkPayload{
			AuthData:    authData,
			NewCustomer: customer,
		}

		exp := time.Now().UTC().Add(time.Minute * 10).Format(http.TimeFormat)
		w.Header().Add("Set-Cookie", fmt.Sprintf("customer=%s; Expires=%v; Secure; HttpOnly;", customer.Email, exp))
		w.Header().Add("Set-Cookie", fmt.Sprintf("token=%v; Expires=%v; Secure; HttpOnly;", customer.Token, exp))
		tools.Render(w, "register-ok.page.html", okPayload)
	}
}
