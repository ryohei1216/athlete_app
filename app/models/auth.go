package models

import (
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
)

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
	
	//環境変数設定
	// os.Setenv("SIGNINGKEY", "athlete_app")

	//電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	//JWTを返却
	w.Write([]byte(tokenString))
}

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
    ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("SIGNINGKEY")), nil
    },
    SigningMethod: jwt.SigningMethodHS256,
})

