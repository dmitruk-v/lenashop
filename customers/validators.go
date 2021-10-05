package customers

import (
	"fmt"
	"net/http"
	"regexp"
)

var emailReg = regexp.MustCompile(`^[a-zA-Z0-9_.]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`)
var phoneReg = regexp.MustCompile(`^\d{3}-\d{2}-\d{2}$`)
var passReg = regexp.MustCompile(`^\S{6,}$`)

func ValidateRegisterForm(r *http.Request) ([]string, error) {
	var err error
	var eMsgs []string
	if err = r.ParseForm(); err != nil {
		return eMsgs, fmt.Errorf("validateRegisterForm(): %v", err)
	}
	if !emailReg.MatchString(r.PostForm.Get("email")) {
		eMsgs = append(eMsgs, "Email is not valid. Example: user@site.com")
	}
	if !phoneReg.MatchString(r.PostForm.Get("phone")) {
		eMsgs = append(eMsgs, "Phone is not valid. Example: 333-22-11")
	}
	if r.PostForm.Get("address") == "" {
		eMsgs = append(eMsgs, "Address is not valid")
	}
	password := r.PostForm.Get("password")
	if !passReg.MatchString(password) {
		eMsgs = append(eMsgs, "Password is not valid. Must be at least 6 characters long.")
	}
	passwordRep := r.PostForm.Get("password-rep")
	if password != passwordRep {
		eMsgs = append(eMsgs, "Passwords are not qual")
	}
	return eMsgs, nil
}

func ValidateLoginForm(r *http.Request) ([]string, error) {
	var err error
	var eMsgs []string
	if err = r.ParseForm(); err != nil {
		return eMsgs, fmt.Errorf("validateLoginForm(): %v", err)
	}
	if !emailReg.MatchString(r.PostForm.Get("email")) {
		eMsgs = append(eMsgs, "Email is not valid. Example: user@site.com")
	}
	password := r.PostForm.Get("password")
	if !passReg.MatchString(password) {
		eMsgs = append(eMsgs, "Password is not valid.")
	}
	return eMsgs, nil
}
