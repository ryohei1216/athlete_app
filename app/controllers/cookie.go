package controllers

import (
	"fmt"
	"main/app/models"
	"net/http"
)


func setCookie (w http.ResponseWriter, r *http.Request, user models.User) {
	c := http.Cookie {
		Name: user.Mail,
		Value: "login cookie mail-address",
		HttpOnly: true,
	}
	fmt.Fprintln(w, c)
	http.SetCookie(w, &c)
}