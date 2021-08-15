package controllers

import (
	"fmt"
	"io/ioutil"
	"main/app/models"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"text/template"
)

//HTMLを生成する
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
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
	generateHTML(w, images, "layout", "top", "header")
}

//asian表示
func asian(w http.ResponseWriter, r *http.Request) {
	images := models.GetImgByRace("asian")
	generateHTML(w, images, "layout", "top", "header")
}

//white表示
func white(w http.ResponseWriter, r *http.Request) {
	images := models.GetImgByRace("white")
	generateHTML(w, images, "layout", "top", "header")
}

//black表示
func black(w http.ResponseWriter, r *http.Request) {
	images := models.GetImgByRace("black")
	generateHTML(w, images, "layout", "top", "header")
}

//画像アップロードトップ画面
func upload(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "upload", "header")
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
		panic(err)
	}
	defer newFile.Close()
	
	fileStr := string(data)
	_, err = newFile.Write([]byte(fileStr))
	if err != nil {
		panic(err)
	}
	race := r.FormValue("race")
	img := models.Image{
		Id: RandomString(10),
		Race: race,
		Filename: header.Filename,
		Good: 0,
		Nope: 0,
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
	
	race := "/" +r.FormValue("race")

	img.DeleteImg()

	deleteFilePath := "./app/views/images/" + img.Filename
	err := os.Remove(deleteFilePath)
	if err != nil {
		panic(err)
	}
	//topにリダイレクト
	http.Redirect(w, r, race, http.StatusFound)
}


//画像情報の更新PUT or DELETEメソッド（Good, Nope）
func evaluate(w http.ResponseWriter, r *http.Request) {
	getImg := models.GetImgById(r.FormValue("id"))

	switch r.FormValue("preference"){
		case "good":
			getImg.Good += 1

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
		getImg.Nope += 1
	}
	getImg.UpdateImg()

	race := "/" + getImg.Race

	//元のページにリダイレクト
	http.Redirect(w, r, race, http.StatusFound)
}

//SignUpのAPI
func signUp (w http.ResponseWriter, r *http.Request) {
	
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
	// http.HandleFunc("/public", models.Public)
	// http.Handle("/private", models.JwtMiddleware.Handler(models.Private))
	// http.HandleFunc("/auth", models.GetTOkenHandler)

  http.ListenAndServe("127.0.0.1:8080", nil)
}