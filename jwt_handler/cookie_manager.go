package jwthandler

import (
	"fmt"
	"net/http"
)

func AssignCookie(w http.ResponseWriter, r *http.Request) {
	token := GenerateToken()
	cookie := &http.Cookie{
		Name:  "SID",
		Value: token,

		MaxAge:   604800,
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}

func CheckCookies(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("SID")

	if err == http.ErrNoCookie {
		return false
	}

	fmt.Println(cookie.Value)
	return true
}
