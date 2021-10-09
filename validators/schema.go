package validators

import (
	"net/http"
	"net/url"
)

type RequestValidatorFunc func(r *http.Request) (errorMessages []string)

type FormValidatorFunc func(form url.Values) (errorMessages []string)
type QueryValidatorFunc func(query url.Values) (errorMessages []string)
type CookieValidatorFunc func(cookies []http.Cookie) (errorMessages []string)
