package customers

import (
	"fmt"
	"net/http"
	"time"

	"dmitrook.ru/lenashop/tools"
	"golang.org/x/crypto/bcrypt"
)

type loginForm struct {
	Email    string
	Password string
}

func parseLoginForm(r *http.Request) (loginForm, error) {
	var lform loginForm
	lform.Email = r.PostFormValue("email")
	lform.Password = r.PostFormValue("password")
	return lform, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------
	// GET: Show login page
	// ----------------------------------------------------------------------
	if r.Method == "GET" {
		tools.Render(w, "login.page.html", nil)
	}
	// ----------------------------------------------------------------------
	// POST: Login user
	// ----------------------------------------------------------------------
	if r.Method == "POST" {
		form, err := parseLoginForm(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when parsing form: %v", err), 400)
			return
		}
		customer, err := CustomerByEmail(form.Email)
		if err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when fetching customer: %v", err), 500)
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(customer.Token), []byte(form.Password)); err != nil {
			tools.Render(w, "login-fail.page.html", nil)
			return
		}
		exp := time.Now().UTC().Add(time.Minute * 10).Format(http.TimeFormat)
		w.Header().Add("Set-Cookie", fmt.Sprintf("customer=%s; Expires=%v; HttpOnly;", customer.Email, exp))
		w.Header().Add("Set-Cookie", fmt.Sprintf("token=%v; Expires=%v; HttpOnly;", customer.Token, exp))
		http.Redirect(w, r, "/products", http.StatusFound)
	}
}
