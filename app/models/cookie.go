package models

import (
	"fmt"
	"net/http"
)

func SetCookie (w http.ResponseWriter, r *http.Request, user User) {
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

func GetCookie (w http.ResponseWriter, r *http.Request)  ([]*http.Cookie, error)  {
	//Cookieの確認するため最初のcookieだけでerror確認
	_, err := r.Cookie("mail")
	if err != nil {
		return nil, fmt.Errorf("Cookieはセットされていない")
		} else {
		h := r.Cookies()
		return  h, nil
	}
}

func DeleteCookie (w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("mail")
	if err != nil {
		fmt.Println(err)
		fmt.Println("delete_cookie取得失敗")
	}
	c2, err := r.Cookie("password")
	if err != nil {
		fmt.Println(err)
		fmt.Println("delete_cookie取得失敗")
	}

	fmt.Println(c1, c2)

	c1.MaxAge = -1
	c2.MaxAge = -1

	http.SetCookie(w, c1)
	http.SetCookie(w, c2)

	//topにリダイレクト
	http.Redirect(w, r, "/", http.StatusFound)
}