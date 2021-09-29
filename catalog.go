package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"dmitrook.ru/lenashop/repository"
)

type catalogPayload struct {
	AuthData
	Products []repository.FullProduct
}

func catalogHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	isAuth, customer, err := checkAuth(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Sorry, something is going wrong on the server.", 500)
		return
	}

	if !isAuth {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	query := appendFromCookies(w, r, "catalog")

	products, err := repository.Products(query)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Sorry, something is going wrong on the server.", 500)
		return
	}

	render(w, "catalog.page.html", catalogPayload{AuthData{isAuth, customer}, products})
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
	// fmt.Println(query)
	return query
}
