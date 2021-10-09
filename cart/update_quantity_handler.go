package cart

import (
	"fmt"
	"net/http"
	"strconv"

	"dmitrook.ru/lenashop/customers"
)

type updateForm struct {
	ProductId   int
	BuyQuantity int
}

func parseUpdateForm(r *http.Request) (updateForm, error) {
	var uform updateForm
	pid, err := strconv.ParseInt(r.PostFormValue("product_id"), 10, 64)
	if err != nil {
		return uform, err
	}
	buyQuantity, err := strconv.ParseInt(r.PostFormValue("buy_quantity"), 10, 64)
	if err != nil {
		return uform, err
	}
	uform.ProductId = int(pid)
	uform.BuyQuantity = int(buyQuantity)
	return uform, nil
}

func UpdateQuantityHandler(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------
	// POST: Update product quantity in cart
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
		form, err := parseUpdateForm(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when parsing form: %v", err), 400)
		}
		if err := UpdateProductQuantity(authData.Customer.CartId, form.ProductId, form.BuyQuantity); err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when updating product quantity: %v", err), 500)
			return
		}
		http.Redirect(w, r, "/cart/products", http.StatusFound)
	}
}
