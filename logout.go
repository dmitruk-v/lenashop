package main

import (
	"fmt"
	"net/http"
	"time"
)

// ----------------------------------------------------------------------
// Handle Logout
// ----------------------------------------------------------------------
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// GET --------------------------------------------------
	if r.Method == "GET" {
		exp := time.Now().UTC().Format(http.TimeFormat)
		w.Header().Add("Set-Cookie", fmt.Sprintf("customer=%s; Expires=%v; Secure; HttpOnly;", "", exp))
		w.Header().Add("Set-Cookie", fmt.Sprintf("token=%v; Expires=%v; Secure; HttpOnly;", "", exp))
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
