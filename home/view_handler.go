package home

import (
	"fmt"
	"net/http"

	"dmitrook.ru/lenashop/customers"
	"dmitrook.ru/lenashop/tools"
)

type viewPayload struct {
	customers.AuthData
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	authData, err := customers.CheckAuth(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Something wrong happened, when checking authentication: %v", err), 500)
		return
	}
	tools.Render(w, "home.page.html", viewPayload{AuthData: authData})
}
