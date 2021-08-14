package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


type Image struct {
	Id 	  		string
	Race  		string
	Filename  string
	Good 			int
	Nope    	int
}

//全画像情報の取得
func GetAllImg() ([]Image) {
	rows, err := db.Query("SELECT * FROM images")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var img Image
	var images []Image

	for rows.Next() {
		if err := rows.Scan(&img.Id, &img.Race, &img.Filename, &img.Good, &img.Nope); err != nil {
			log.Fatal(err)
		}
		images = append(images, img)
	}
	return images
}

//特定の画像情報の取得(Id)
func  GetImgById(id string) Image{
	cmd := "SELECT * FROM images WHERE id = $1"
	
	var img Image
	err := db.QueryRow(cmd, id).Scan(&img.Id, &img.Race, &img.Filename, &img.Good, &img.Nope)
	if err != nil {
		fmt.Println(err)
	}
	return img
}

//特定の画像情報の取得(Id)
func  GetImgByFilename(filename string) Image {
	cmd := "SELECT * FROM images WHERE filename = $1"
	
	var img Image
	err := db.QueryRow(cmd, filename).Scan(&img.Id, &img.Race, &img.Filename, &img.Good, &img.Nope)
	if err != nil {
		fmt.Println(err)
	}
	return img
}

//全画像情報の取得(race別)
func GetImgByRace(race string) ([]Image) {
	rows, err := db.Query("SELECT * FROM images WHERE race = $1", race)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var images []Image
	
	for rows.Next() {
		var img Image
		if err := rows.Scan(&img.Id, &img.Race, &img.Filename, &img.Good, &img.Nope); err != nil {
			log.Fatal(err)
		}
		images = append(images, img)
	}
	return images
}


//画像の保存とDBへの格納
func (img Image) InsertImg() {
	cmd := "INSERT INTO images (id, race, filename, Good, Nope) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Exec(cmd, img.Id, img.Race, img.Filename, img.Good, img.Nope)
	if err != nil {
		fmt.Println("InsertImg")
		log.Fatal(err)
	}
}


//画像の削除
func (img Image) DeleteImg() {
	cmd := "DELETE FROM images WHERE id = $1"
	_, err := db.Exec(cmd, img.Id)
	if err != nil {
		log.Fatal(err)
	}
}

//画像情報の更新
func (img Image) UpdateImg() {
	cmd := "UPDATE images SET good = $1, nope = $2 WHERE id = $3"
	_, err := db.Exec(cmd, img.Good, img.Nope, img.Id)
	if err != nil {
		panic(err)
	}
}





