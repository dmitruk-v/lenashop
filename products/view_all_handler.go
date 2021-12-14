package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"dmitrook.ru/lenashop/customers"
	"dmitrook.ru/lenashop/tools"
)

type viewAllPayload struct {
	customers.AuthData
	Products []FullProduct
}

func ViewAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var err error
		authData, err := customers.CheckAuth(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something wrong happened, when checking authentication: %v", err), 500)
			return
		}
		if !authData.IsAuth {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		query := appendFromCookies(w, r, "catalog")
		products, err := ProductsWithImages(query)
		if err != nil {
			http.Error(w, fmt.Sprintf("Sorry, something wrong happened when fetching products: %v", err), 500)
			return
		}

		if _, ok := r.Header["X-Requested-With"]; ok {
			encoder := json.NewEncoder(w)
			if err := encoder.Encode(products); err != nil {
				http.Error(w, fmt.Sprintf("Sorry, something wrong happened when encoding json: %v", err), 500)
				return
			}
			return
		}

		viewAllPayload := viewAllPayload{
			AuthData: authData,
			Products: products,
		}
		tools.Render(w, "products.page.html", viewAllPayload)
	}
}

// If value is present in query, it will add it to cookie, otherwise it will try to load
// value from cookie. If success, value will be added to query. If failed, cookie will be set with default value.
func appendFromCookies(w http.ResponseWriter, r *http.Request, prefix string) url.Values {
	var catalogDefaultValues = map[string][]string{
		"sort":  {"price desc"},
		"limit": {"10"},
		"bla":   {"xz"},
		"yep":   {"nope"},
	}
	query := r.URL.Query()
	for key := range catalogDefaultValues {
		cookieName := fmt.Sprintf("%s-%s", prefix, key)
		if len(query[key]) == 0 {
			cookie, err := r.Cookie(cookieName)
			if err == nil {
				query[key] = strings.Split(cookie.Value, ",")
			}
		} else {
			w.Header().Add("Set-Cookie", fmt.Sprintf("%s=%s", cookieName, strings.Join(query[key], ",")))
		}
	}
	return query
}
