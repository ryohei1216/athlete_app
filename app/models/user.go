package models

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name     string
	Mail     string
	Password string
}


//SignUp
func (user User) SignUp(w http.ResponseWriter) {
	cmd := "SELECT * FROM users where mail = $1"
	//userを更新するかも？しないほうがいい？
	//データがあるか確認
	err := db.QueryRow(cmd, user.Mail).Scan(&user.Name, &user.Mail, &user.Password)
	if err == nil {
		fmt.Fprintln(w,"同じメールアドレスで登録されています")
		return
	}

	cmd = "INSERT INTO users VALUES ($1, $2, $3)"
	cmd2 := "CREATE TABLE IF NOT EXISTS " + user.Mail + "(good_list varchar(50))"
	_, err = db.Exec(cmd, user.Name, user.Mail, user.Password)
	if err != nil {
		fmt.Println("データ挿入失敗")
	}
	_, err = db.Exec(cmd2)
	if err != nil {
		log.Println(err)
	}
}

//Login
func (user User) Login() bool {	
	cmd := "SELECT * FROM users WHERE mail = $1 AND password = $2"
	err := db.QueryRow(cmd, user.Mail, user.Password).Scan(&user.Name, &user.Mail, &user.Password)
	if err != nil {
		fmt.Println("mail または password が間違っています。")
		return false
	} else {
		fmt.Println("ログイン成功")
		return true
	}
}

//特定のユーザーの取得(Id)
func GetUser(mail string, password string) User {
	cmd := "SELECT * FROM images WHERE mail = $1 AND password = $2"

	var user User
	err := db.QueryRow(cmd, mail, password).Scan(&user.Name, &user.Mail, &user.Password)
	if err != nil {
		fmt.Println(err)
	}
	return user
}