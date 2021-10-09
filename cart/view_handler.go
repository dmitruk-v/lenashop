package cart

import (
	"fmt"
	"net/http"

	"dmitrook.ru/lenashop/customers"
	"dmitrook.ru/lenashop/tools"
)

type cartViewPayload struct {
	customers.AuthData
	Products   []CartProduct
	TotalPrice float64
}

func (cp *cartViewPayload) calcTotalPrice() float64 {
	total := 0.0
	for _, p := range cp.Products {
		total += float64(p.BuyQuantity) * p.Price
	}
	return total
}

func newCartViewPayload(authData customers.AuthData, products []CartProduct) cartViewPayload {
	payload := cartViewPayload{
		AuthData: authData,
		Products: products,
	}
	tp := payload.calcTotalPrice()
	payload.TotalPrice = tp
	return payload
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------
	// GET: Show products list in cart
	// ----------------------------------------------------------------------
	if r.Method == "GET" {
		authData, err := customers.CheckAuth(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when checking authentication: %v", err), 500)
			return
		}
		if !authData.IsAuth {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		products, err := Products(authData.Customer.CartId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when fetching cart products: %v", err), 500)
			return
		}
		payload := newCartViewPayload(authData, products)
		tools.Render(w, "cart.page.html", payload)
	}
}
