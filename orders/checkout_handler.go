package orders

import (
	"fmt"
	"net/http"

	"dmitrook.ru/lenashop/cart"
	"dmitrook.ru/lenashop/customers"
	"dmitrook.ru/lenashop/tools"
)

type checkoutViewPayload struct {
	customers.AuthData
	Products   []cart.CartProduct
	TotalPrice float64
}

func (cvp *checkoutViewPayload) calcTotalPrice() {
	for _, p := range cvp.Products {
		cvp.TotalPrice += float64(p.BuyQuantity) * p.Price
	}
}

func newCheckoutViewPayload(authData customers.AuthData, products []cart.CartProduct) checkoutViewPayload {
	payload := checkoutViewPayload{
		AuthData: authData,
		Products: products,
	}
	payload.calcTotalPrice()
	return payload
}

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		authData, err := customers.CheckAuth(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when checking authentication: %v", err), 500)
			return
		}
		if !authData.IsAuth {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		products, err := cart.Products(authData.Customer.CartId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when fetching products: %v", err), 500)
			return
		}
		payload := newCheckoutViewPayload(authData, products)
		tools.Render(w, "checkout.page.html", payload)
	}
}
