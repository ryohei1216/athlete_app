package models

import (
	"encoding/json"
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
)

type Post struct {
	Title string `json:"title"`
  Tag   string `json:"tag"`
  URL   string `json:"url"`
}


func GetTOkenHandler(w http.ResponseWriter, r *http.Request) {
	//headerのセット
	token := jwt.New(jwt.SigningMethodES256)

	//claimのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = "taro"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	

	secret := "2FMd5FNSqS/nW2wWJy5S3ppjSHhUnLt8HuwBkTD6HqfPfBBDlykwLA=="
	//電子署名
	tokenString, _ := token.SignedString([]byte(secret))

	//JWTを返却
	w.Write([]byte(tokenString))
}

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
    ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
        return []byte("38ryohei38"), nil
    },
    SigningMethod: jwt.SigningMethodHS256,
})

//authハンドラー
func Public (w http.ResponseWriter, r *http.Request) {
	post := &Post {
		Title: "VueCLIからVue.js入門①【VueCLIで出てくるファイルを概要図で理解】",
    Tag:   "Vue.js",
    URL:   "https://qiita.com/po3rin/items/3968f825f3c86f9c4e21",
	}
	json.NewEncoder(w).Encode(post)
}


//authハンドラー
var Private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    post := &Post{
        Title: "VGolangとGoogle Cloud Vision APIで画像から文字認識するCLIを速攻でつくる",
        Tag:   "Go",
        URL:   "https://qiita.com/po3rin/items/bf439424e38757c1e69b",
    }
    json.NewEncoder(w).Encode(post)
})



