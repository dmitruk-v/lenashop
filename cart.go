package main

import "net/http"

func cartHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
