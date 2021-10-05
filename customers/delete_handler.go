package customers

import (
	"fmt"
	"net/http"
	"strconv"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------
	// POST: Delete customer
	// ----------------------------------------------------------------------
	if r.Method == "POST" {
		authData, err := CheckAuth(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when checking auth: %v", err), 500)
			return
		}
		if !authData.IsAuth {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		cid, err := strconv.ParseInt(r.PostFormValue("customer_id"), 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when parsing form: %v", err), 400)
			return
		}
		if err := DeleteCustomer(int(cid)); err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when deleting customer: %v", err), 500)
			return
		}
		http.Redirect(w, r, "/customers", http.StatusFound)
	}
}
