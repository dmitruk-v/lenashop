package orders

import (
	"fmt"
	"net/http"
	"strconv"

	"dmitrook.ru/lenashop/customers"
	"dmitrook.ru/lenashop/tools"
)

type viewSinglePayload struct {
	customers.AuthData
	Order
}

type viewSingleQuery struct {
	OrderId int
}

func parseViewSingleQuery(r *http.Request) (viewSingleQuery, error) {
	var vsq viewSingleQuery
	query := r.URL.Query()
	orderId, err := strconv.ParseInt(query.Get("order_id"), 10, 64)
	if err != nil {
		return vsq, fmt.Errorf("parseViewSingleQuery(request): %v", err)
	}
	vsq.OrderId = int(orderId)
	return vsq, nil
}

func ViewSingleHandler(w http.ResponseWriter, r *http.Request) {
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
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		query, err := parseViewSingleQuery(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when parsing query params: %v", err), 400)
			return
		}

		ord, err := OrderById(query.OrderId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when fetching order products: %v", err), 500)
			return
		}

		payload := viewSinglePayload{
			AuthData: authData,
			Order:    ord,
		}
		tools.Render(w, "order.page.html", payload)
	}
}
