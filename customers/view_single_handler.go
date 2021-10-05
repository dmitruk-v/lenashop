package customers

import "net/http"

type viewSinglePayload struct {
	AuthData
	Customer
}

func ViewSingleHandler(w http.ResponseWriter, r *http.Request) {}
