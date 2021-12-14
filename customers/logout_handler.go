package customers

import (
	"fmt"
	"net/http"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// ----------------------------------------------------------------------
	// POST: Logout user
	// ----------------------------------------------------------------------
	if r.Method == "GET" {
		exp := time.Now().UTC().Format(http.TimeFormat)
		w.Header().Add("Set-Cookie", fmt.Sprintf("customer=%s; Expires=%v; HttpOnly;", "", exp))
		w.Header().Add("Set-Cookie", fmt.Sprintf("token=%v; Expires=%v; HttpOnly;", "", exp))
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
