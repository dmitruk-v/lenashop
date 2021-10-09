package models

// import "fmt"

// // High level module
// // --------------------------------------------------
// type CustomerHandler struct{}

// func (ch *CustomerHandler) handle(form map[string]string, fvr CustomerFormValidator) {
// 	if fvr.Validate(form) {
// 		fmt.Println("valid!")
// 	} else {
// 		fmt.Println("invalid!")
// 	}
// }

// type ProductHandler struct{}

// func (ph *ProductHandler) handle(form map[string]string, fvr ProductFormValidator) {
// 	if fvr.Validate(form) {
// 		fmt.Println("valid!")
// 	} else {
// 		fmt.Println("invalid!")
// 	}
// }

// // Low level module
// // --------------------------------------------------
// type CustomerFormValidator struct{}

// func (v *CustomerFormValidator) Validate(form map[string]string) bool {
// 	fmt.Println("validate customer form:", form)
// 	return true
// }

// type ProductFormValidator struct{}

// func (v *ProductFormValidator) Validate(form map[string]string) bool {
// 	fmt.Println("validate product form:", form)
// 	return true
// }

// // run code
// // --------------------------------------------------
// func run() {
// 	cfValidator := CustomerFormValidator{}

// 	ch := CustomerHandler{}
// 	ch.handle(map[string]string{
// 		"email":    "valera.dmitruk@gmail.com",
// 		"password": "strangepass",
// 	}, cfValidator)

// 	pfValidator := ProductFormValidator{}

// 	ph := ProductHandler{}
// 	ph.handle(map[string]string{
// 		"title": "tennis ball",
// 		"price": "125",
// 	}, pfValidator)
// }
