package products

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"dmitrook.ru/lenashop/customers"
	"dmitrook.ru/lenashop/tools"
)

type viewSinglePayload struct {
	customers.AuthData
	Product FullProduct
}

// var productIdReg = regexp.MustCompile(`^/product/(\d+)$`)

func extractId(r *http.Request) (int, error) {
	pidStr := strings.TrimPrefix(r.URL.Path, "/product/")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return -1, fmt.Errorf("can not parse %q to int", pidStr)
	}
	return pid, nil
}

func ViewSingleHandler(w http.ResponseWriter, r *http.Request) {
	authData, err := customers.CheckAuth(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Something wrong happened, when checking authentication: %v", err), 500)
		return
	}
	if !authData.IsAuth {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	id, err := extractId(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Something wrong happened, when extracting product id: %v", err), 400)
		return
	}
	product, err := ProductById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Something wrong happened, when fetching product: %v", err), 500)
		return
	}
	payload := viewSinglePayload{
		AuthData: authData,
		Product:  product,
	}
	tools.Render(w, "product.page.html", payload)
}
