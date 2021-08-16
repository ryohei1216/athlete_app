package models

import (
	"fmt"
	"log"
)

type User struct {
	Name     string
	Mail     string
	Password string
}


//SignUp
func (user *User) SignUp() {
	cmd := "SELECT * FROM users where mail = $1"
	//userを更新するかも？しないほうがいい？
	//データがあるか確認
	checkMail := user.Mail
	_ = db.QueryRow(cmd, user.Mail).Scan(user.Name, user.Mail, user.Password)
	if checkMail == user.Mail {
		fmt.Println(user.Mail)
		fmt.Println("同じメールアドレスで登録されています")
		return
	}

	cmd = "INSERT INTO users VALUES ($1, $2, $3)"
	cmd2 := "CREATE TABLE IF NOT EXISTS $1 (favorite varchar(50))"
	_, err := db.Exec(cmd, user.Name, user.Mail, user.Password)
	if err != nil {
		fmt.Println("データ挿入失敗")
	}
	_, err = db.Exec(cmd2, user.Mail)
	if err != nil {
		log.Println(err)
	}
}

//Login
func (user *User) Login() {	
	cmd := "SELECT * FROM users WHERE mail = $1 AND password = $2"
	err := db.QueryRow(cmd, user.Mail, user.Password)
	if err != nil {
		fmt.Println("mail または password が間違っています。")
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