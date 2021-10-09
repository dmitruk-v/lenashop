package models

import "fmt"

// type FormValidator interface {
// 	Validate(form map[string]string) (errorMessages []string)
// }

type FormValidatorFunc func(form map[string]string) (errorMessages []string)

// High level module
// --------------------------------------------------
type CustomerHandler struct{}

func (ch *CustomerHandler) handle(form map[string]string, formValidator FormValidatorFunc) {
	errMsgs := formValidator(form)
	if len(errMsgs) == 0 {
		fmt.Println("valid!")
	} else {
		fmt.Println("invalid!")
	}
}

type ProductHandler struct{}

func (ph *ProductHandler) handle(form map[string]string, formValidator FormValidatorFunc) {
	errMsgs := formValidator(form)
	if len(errMsgs) == 0 {
		fmt.Println("valid!")
	} else {
		fmt.Println("invalid!")
	}
}

// Low level module
// --------------------------------------------------
func validateCustomerForm(form map[string]string) (errorMessages []string) {
	email, ok := form["email"]
	if !ok || email == "" {
		errorMessages = append(errorMessages, "email is required")
	}
	password, ok := form["password"]
	if !ok || password == "" {
		errorMessages = append(errorMessages, "password is required")
	}
	return errorMessages
}

func validateProductForm(form map[string]string) (errorMessages []string) {
	title, ok := form["title"]
	if !ok || title == "" {
		errorMessages = append(errorMessages, "title is required")
	}
	price, ok := form["price"]
	if !ok || price == "" {
		errorMessages = append(errorMessages, "price is required")
	}
	return errorMessages
}

// run code
// --------------------------------------------------
func run() {
	ch := CustomerHandler{}
	ch.handle(map[string]string{
		"email":    "valera.dmitruk@gmail.com",
		"password": "strangepass",
	}, validateCustomerForm)

	ph := ProductHandler{}
	ph.handle(map[string]string{
		"title": "tennis ball",
		"price": "125",
	}, validateProductForm)
}
