package tools

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type ContextKey string

var keyPatReg = regexp.MustCompile(`{([a-zA-Z0-9_-]+):([^}]+)}`)

type Route struct {
	pNames  []string
	reg     *regexp.Regexp
	handler http.HandlerFunc
}

func (r Route) String() string {
	return fmt.Sprintf("Route{reg=%v,handler=%v}", r.reg, r.handler)
}

type MyMux struct {
	mux    http.ServeMux
	routes []Route
}

func (mm *MyMux) String() string {
	return fmt.Sprintf("MyMux{mux=%v,routes=%v}", mm.mux, mm.routes)
}

func NewMyMux() *MyMux {
	return &MyMux{
		mux:    *http.NewServeMux(),
		routes: []Route{},
	}
}

func (myMux *MyMux) Handle(pattern string, handler http.Handler) {
	myMux.mux.Handle(pattern, handler)
}

// ---------------------------------------------------------------------
// TOKNOW
// I do not understand how it can be done with registering regexp pattern in mux.HandleFunc.
// Because mux.HandleFunc requires static string literal.
// ---------------------------------------------------------------------
func (myMux *MyMux) HandleFunc(pattern string, handler http.HandlerFunc) {
	pNames := []string{}
	replPattern := keyPatReg.ReplaceAllStringFunc(pattern, func(s string) string {
		pair := strings.Split(s[1:len(s)-1], ":")
		pNames = append(pNames, pair[0])
		return fmt.Sprintf("(%v)", pair[1])
	})
	route := Route{
		pNames:  pNames,
		reg:     regexp.MustCompile(replPattern),
		handler: handler,
	}
	myMux.routes = append(myMux.routes, route)
	// myMux.mux.HandleFunc(replPattern, handler)
}

func (myMux *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if we have request to the static content, let standard server mux serve it
	if strings.HasPrefix(r.URL.Path, "/assets/") {
		myMux.mux.ServeHTTP(w, r)
		return
	}
	// otherwise do our routing work
	for _, route := range myMux.routes {
		if route.reg.MatchString(r.URL.Path) {
			fmt.Println("--- MATCHED ---", route.reg, r.URL)
			matches := route.reg.FindStringSubmatch(r.URL.Path)
			pvals := matches[1:]
			if len(pvals) != len(route.pNames) {
				log.Fatal("Error. The lengths of params keys and params values are not equal.")
			}
			// There are two ways to add query params:
			// 1. --------------------------------------
			// q := r.URL.Query()
			// for idx, pname := range route.pNames {
			// 	q.Add(pname, pvals[idx])
			// 	r.URL.RawQuery = q.Encode()
			// }
			// 2. --------------------------------------
			for idx, pname := range route.pNames {
				r.URL.RawQuery += fmt.Sprintf("&%v=%v", pname, pvals[idx])
			}
			route.handler(w, r)
			// myMux.mux.ServeHTTP(w, r)
			return
		}
	}
	// if url does not match any route
	log.Printf("Url: %q does not match any existing route.", r.URL.Path)
}
