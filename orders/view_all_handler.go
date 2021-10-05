package orders

import (
	"fmt"
	"net/http"

	"dmitrook.ru/lenashop/customers"
	"dmitrook.ru/lenashop/tools"
)

type viewAllPayload struct {
	customers.AuthData
	Orders []Order
}

func ViewAllHandler(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------
	// GET: Show products list and forms
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
		ords, err := Orders(authData.Customer.CustomerId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when fetching order products: %v", err), 500)
			return
		}
		payload := viewAllPayload{
			AuthData: authData,
			Orders:   ords,
		}
		tools.Render(w, "orders.page.html", payload)
	}
}
