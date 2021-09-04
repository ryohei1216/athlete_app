package controllers

import (
	"fmt"
	"main/app/models"
	"net/http"
)

func setCookie (w http.ResponseWriter, r *http.Request, user models.User) {
	c1 := http.Cookie{
		Name:     "mail",
		Value:    user.Mail,
		HttpOnly: true,
	}

	c2 := http.Cookie {
		Name: "password",
		Value: user.Password,
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}

func getCookie (w http.ResponseWriter, r *http.Request) bool {
	h := r.Header["Cookie"]
	if h != nil {
		fmt.Println("Cookie取得完了")
		return true
	} else {
		return false
	}
}

func deleteCookie (w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("mail")
	if err != nil {
		fmt.Println("cookie取得失敗")
	}
	c2, err := r.Cookie("password")
	if err != nil {
		fmt.Println("cookie取得失敗")
	}

	fmt.Println(c1, c2)

	c1.MaxAge = -1
	c2.MaxAge = -1

	http.SetCookie(w, c1)
	http.SetCookie(w, c2)

	//topにリダイレクト
	http.Redirect(w, r, "/", http.StatusFound)
}