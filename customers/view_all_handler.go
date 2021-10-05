package customers

import (
	"fmt"
	"net/http"

	"dmitrook.ru/lenashop/tools"
)

type viewAllPayload struct {
	AuthData
	Customers []Customer
}

func ViewAllHandler(w http.ResponseWriter, r *http.Request) {
	authData, err := CheckAuth(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Sorry, something wrong happened when checking auth: %v", err), 500)
		return
	}
	if !authData.IsAuth {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	customers, err := Customers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Sorry, something wrong happened when fetching customers: %v", err), 500)
		return
	}

	tools.Render(w, "customers.page.html", viewAllPayload{
		AuthData:  authData,
		Customers: customers,
	})
}
