package orders

import (
	"fmt"
	"net/http"

	"dmitrook.ru/lenashop/customers"
	"dmitrook.ru/lenashop/tools"
)

type orderCreatePayload struct {
	customers.AuthData
	OrderId int
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------
	// GET: Show cart products and forms
	// ----------------------------------------------------------------------
	if r.Method == "GET" {
		tools.Render(w, "checkout.page.html", nil)
	}
	// ----------------------------------------------------------------------
	// POST: Create new order
	// ----------------------------------------------------------------------
	if r.Method == "POST" {
		authData, err := customers.CheckAuth(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when checking authentication: %v", err), 500)
			return
		}
		if !authData.IsAuth {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		orderId, err := CreateOrder(authData.Customer.CustomerId, authData.Customer.CartId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when creating order: %v", err), 500)
			return
		}
		payload := orderCreatePayload{
			AuthData: authData,
			OrderId:  orderId,
		}
		tools.Render(w, "checkout-ok.page.html", payload)
	}
}
