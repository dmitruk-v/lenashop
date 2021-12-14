package clean

import "net/http"

type Validator interface {
	Validate(r *http.Request) bool
}

type CartAddProduct struct {
	ProductId   int
	BuyQuantity int
	validator   Validator
}

func (cap CartAddProduct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// some logic
}

func NewCartAddProduct(productId int, buyQuantity int, v Validator) CartAddProduct {
	return CartAddProduct{
		ProductId:   productId,
		BuyQuantity: buyQuantity,
		validator:   v,
	}
}

// func withValidator(pattern string, handler http.Handler, vr Validator) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 	})
// }

// --------------------------------------------------------------

type CartAddProductValidator struct{}

func (c CartAddProductValidator) Validate(r *http.Request) bool {
	return false
}

func run2() {

	cap := NewCartAddProduct(1, 5, CartAddProductValidator{})

	http.Handle("/", cap)
	// handler := func(w http.ResponseWriter, r *http.Request) {}
	// withValidator("/", http.HandlerFunc(handler), CartAddProductValidator{})
}
