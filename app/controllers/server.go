package controllers

import (
	"fmt"
	"io/ioutil"
	"main/app/models"
	"main/app/models/twitter_model"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"text/template"

	"github.com/dghubble/go-twitter/twitter"
)

//HTMLを生成する
func generateHTML(w http.ResponseWriter, r *http.Request, data interface{}, filenames ...string) {
	// Cookieがあれば"header"を"private_header"に変更
	if _, err := models.GetCookie(w, r); err == nil {
		for index, filename := range filenames {
			if filename == "header" {
				filenames[index] = "private_header"
			}
		}
	}
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("./app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}


//トップ画面
func top(w http.ResponseWriter, r *http.Request) {
	images := models.GetAllImg()
	sort.Slice(images, func(i, j int) bool {
		return images[i].Good > images[j].Good
	})
	generateHTML(w, r, images, "layout", "top", "header")
}

//asian表示
func asian(w http.ResponseWriter, r *http.Request) {
	images := models.GetImgByRace("asian")
	sort.Slice(images, func(i, j int) bool {
		return images[i].Good > images[j].Good
	})
	generateHTML(w, r, images, "layout", "top", "header")
}

//white表示
func white(w http.ResponseWriter, r *http.Request) {
	images := models.GetImgByRace("white")
	sort.Slice(images, func(i, j int) bool {
		return images[i].Good > images[j].Good
	})
	generateHTML(w, r, images, "layout", "top", "header")
}

//black表示
func black(w http.ResponseWriter, r *http.Request) {
	images := models.GetImgByRace("black")
	sort.Slice(images, func(i, j int) bool {
		return images[i].Good > images[j].Good
	})
	generateHTML(w, r, images, "layout", "top", "header")
}

//画像アップロードトップ画面
func upload(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, nil, "layout", "upload", "header")
}


//画像をアップロードするPOSTメソッド
func uploading(w http.ResponseWriter, r *http.Request){
	file, header, err := r.FormFile("uploading")
	if err != nil {
		fmt.Println("読み込み失敗")
	}
	defer file.Close()
	
	data, _ := ioutil.ReadAll(file)
	newFile, err := os.Create(filepath.Join("./app/views/images/", header.Filename))
	if err != nil {
		fmt.Println("ファイル作成失敗")
		fmt.Println(err)
	}
	defer newFile.Close()
	
	fileStr := string(data)
	_, err = newFile.Write([]byte(fileStr))
	if err != nil {
		fmt.Println(err)
	}
	race := r.FormValue("race")
	name := r.FormValue("name")
	img := models.Image{
		Id: RandomString(10),
		Race: race,
		Filename: header.Filename,
		Good: 0,
		Nope: 0,
		Name: name,
	}
	// 同じ名前のファイルがある場合処理しない
	if models.GetImgByFilename(img.Filename).Filename != ""  {
		fmt.Println("同じ名前のファイルが存在します")
	} else {
		img.InsertImg()
	}
	//topにリダイレクト
	http.Redirect(w, r, "/", http.StatusFound)
}

func RandomString(n int) string {
    var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    b := make([]rune, n)
    for i := range b {
        b[i] = letter[rand.Intn(len(letter))]
    }
    return string(b)
}


// 画像を削除するDELETEメソッド
func delete(w http.ResponseWriter, r *http.Request) {

	img  := models.Image {
		Id: r.FormValue("id"),
		Filename: r.FormValue("filename"),
	}
	
	
	img.DeleteImg()
	
	//個人Dbから削除
	img.DeleteImgByUser()
	
	
	deleteFilePath := "./app/views/images/" + img.Filename
	err := os.Remove(deleteFilePath)
	if err != nil {
		panic(err)
	}
	//topにリダイレクト
	// race := "/" +r.FormValue("race")
	http.Redirect(w, r, "/", http.StatusFound)
}


//画像情報の更新PUT or DELETEメソッド（Good, Nope）
//cookieがあればfavoriteリストに画像を登録
func evaluate(w http.ResponseWriter, r *http.Request) {
	getImg := models.GetImgById(r.FormValue("id"))

	switch r.FormValue("preference") {
		case "good":
			getImg.Good += 1
			// Cookieの取得
			h, err := models.GetCookie(w, r)
			if err == nil {
				//個人Dbに登録
				models.RegisterGoodImg(w, r, h)
			}

		case "nope":
			if getImg.Nope +1 >= 10 {
				uri := "http://127.0.0.1:8080/delete"
				req, err := http.NewRequest("DELETE", uri, nil)
				if err != nil {
					panic(err)
				}
				//クエリパラメータを追加してエンコードする
				q := req.URL.Query()
				q.Add("id", getImg.Id)
				q.Add("filename", getImg.Filename)
				req.URL.RawQuery = q.Encode()

				client := &http.Client{}
				client.Do(req)
			}

			//個人Dbから削除
			getImg.DeleteImgByUser()
			
		getImg.Nope += 1
	}
	getImg.UpdateImg()

	// race := "/" + getImg.Race

	//topページにリダイレクト
	http.Redirect(w, r, "/", http.StatusFound)
}

//SignUp
func signUp (w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, nil, "layout", "signup", "header")
}

//ユーザ登録するAPI
func register (w http.ResponseWriter, r *http.Request) {
	var user = models.User {
		Name: r.FormValue("name"),
		Mail: r.FormValue("mail"),
		Password: r.FormValue("password"),
	}
	fmt.Fprintln(w, user)
	user.SignUp(w)
}

func login (w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, nil, "layout", "login", "header")
}

//ログインするAPI
func logging (w http.ResponseWriter, r *http.Request) {
	
	var user = models.User {
		Mail: r.FormValue("mail"),
		Password: r.FormValue("password"),
	}
	isLogin := user.Login()
	if isLogin {
		models.SetCookie(w, r, user)
	}
	//topにリダイレクト
	http.Redirect(w, r, "/", http.StatusFound)
}

func myPage(w http.ResponseWriter, r *http.Request) {
	h, err := models.GetCookie(w, r)
	if err != nil {
		fmt.Println("Cookie取得失敗")
		fmt.Println(err)
	}
	imagesId:= models.GetImgByUser(h)

	var images []string
	for _, id := range imagesId {
		image := models.GetImgById(id)
		images = append(images, image.Filename)
	}
	generateHTML(w, r, images, "layout", "mypage", "header")
}

func athletepage(w http.ResponseWriter, r *http.Request) {
	athleteName := r.FormValue("name")
	images := models.GetImgByName(athleteName)
	tweets := twitter_model.SearchTweets(athleteName)
	wiki := models.GetWiki(athleteName)
	data := struct {
		Images    []models.Image
		Tweets 		[]twitter.Tweet
		Wiki      map[string]interface{}
	}{
		Images: images,
		Tweets: tweets,
		Wiki: wiki,
	}
	// for _, tweet := range tweets {
	// 	fmt.Println("***************************")
	// 	fmt.Println(tweet.Entities)
	// 	fmt.Println("***************************")
	// }
	models.GetWiki(athleteName)
	generateHTML(w, r, data, "layout", "athletepage", "header")
}




func InitServer() {
	
	files := http.FileServer(http.Dir("app/views"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/asian", asian)
	http.HandleFunc("/white", white)
	http.HandleFunc("/black", black)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/upload/uploading", uploading)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/evaluate", evaluate)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logging", logging)
	http.HandleFunc("/deleteCookie", models.DeleteCookie)
	http.HandleFunc("/mypage", myPage)
	http.HandleFunc("/athletepage", athletepage)

	//React API用
	http.HandleFunc("/reactGetImg", models.ReactGetAllImg)

  http.ListenAndServe("127.0.0.1:8080", nil)
}