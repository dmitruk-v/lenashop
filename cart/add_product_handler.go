package cart

import (
	"fmt"
	"net/http"
	"strconv"

	"dmitrook.ru/lenashop/customers"
)

type addForm struct {
	ProductId   int
	BuyQuantity int
}

func parseAddForm(r *http.Request) (addForm, error) {
	var aform addForm
	pid, err := strconv.ParseInt(r.PostFormValue("product_id"), 10, 64)
	if err != nil {
		return aform, err
	}
	bq, err := strconv.ParseInt(r.PostFormValue("buy_quantity"), 10, 64)
	if err != nil {
		return aform, err
	}
	aform.ProductId = int(pid)
	aform.BuyQuantity = int(bq)
	return aform, nil
}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------
	// POST: Add product to cart
	// ----------------------------------------------------------------------
	if r.Method == "POST" {
		authData, err := customers.CheckAuth(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when checking authentication: %v", err), 500)
			return
		}
		if !authData.IsAuth {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		form, err := parseAddForm(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when parsing form: %v", err), 400)
			return
		}
		if err := AddProduct(authData.Customer.CartId, form.ProductId, form.BuyQuantity); err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when adding product to cart: %v", err), 500)
			return
		}
		http.Redirect(w, r, "/cart/products", http.StatusFound)
	}
}
