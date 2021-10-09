package cart

import (
	"fmt"
	"net/http"
	"strconv"

	"dmitrook.ru/lenashop/customers"
)

type removeForm struct {
	ProductId int
}

func parseRemoveForm(r *http.Request) (removeForm, error) {
	var rform removeForm
	pid, err := strconv.ParseInt(r.PostFormValue("product_id"), 10, 64)
	if err != nil {
		return rform, err
	}
	rform.ProductId = int(pid)
	return rform, nil
}

func RemoveProductHandler(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------
	// POST: Remove product from cart
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
		form, err := parseRemoveForm(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when parsing form: %v", err), 400)
			return
		}
		if err := RemoveProduct(authData.Customer.CartId, form.ProductId); err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when removing product: %v", err), 500)
			return
		}
		http.Redirect(w, r, "/cart/products", http.StatusFound)
	}
}
