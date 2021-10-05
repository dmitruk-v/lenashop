package customers

import (
	"fmt"
	"net/http"
)

type AuthData struct {
	IsAuth   bool
	Customer CustomerWithCart
}

func CheckAuth(r *http.Request) (AuthData, error) {
	var customer CustomerWithCart
	var err error

	email, custErr := r.Cookie("customer")
	token, tokenErr := r.Cookie("token")
	if custErr != nil || tokenErr != nil {
		return AuthData{false, customer}, nil
	}
	customer, err = CustomerWithCartByEmail(email.Value)
	if err != nil {
		return AuthData{false, customer}, fmt.Errorf("CheckAuth(): %v", err)
	}
	if customer.Token == "" || customer.Token != token.Value {
		return AuthData{false, customer}, nil
	}
	return AuthData{true, customer}, nil
}
