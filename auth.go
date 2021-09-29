package main

import (
	"fmt"
	"net/http"

	"dmitrook.ru/lenashop/repository"
)

type AuthData struct {
	IsAuth   bool
	Customer repository.Customer
}

func checkAuth(r *http.Request) (bool, repository.Customer, error) {
	var customer repository.Customer
	var err error

	email, custErr := r.Cookie("customer")
	token, tokenErr := r.Cookie("token")
	if custErr != nil || tokenErr != nil {
		return false, customer, nil
	}
	customer, err = repository.CustomerByEmail(email.Value)
	if err != nil {
		return false, customer, fmt.Errorf("checkAuth(): %v", err)
	}
	if customer.Token == "" || customer.Token != token.Value {
		return false, customer, nil
	}
	return true, customer, nil
}
