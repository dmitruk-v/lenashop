package customers

import (
	"fmt"
	"net/http"

	"dmitrook.ru/lenashop/tools"
)

type registerPayload struct {
	AuthData
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------
	// GET: Show registration form
	// ----------------------------------------------------------------------
	if r.Method == "GET" {
		authData, err := CheckAuth(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when checking auth: %v", err), 500)
			return
		}
		payload := registerPayload{
			AuthData: authData,
		}
		tools.Render(w, "register.page.html", payload)
	}
}
