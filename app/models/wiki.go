package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)



func GetWiki(name string) map[string]interface{} {
	uri := "http://ja.wikipedia.org/w/rest.php/v1/search/page"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println(err)
	}
	//クエリパラメータを追加してエンコードする
	q := req.URL.Query()
	q.Add("q", name)
	q.Add("limit", "1")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	//JSONはkey string, value []interface{}で与えられる
	var mapDate map[string][]interface{}

	json.Unmarshal([]byte(body), &mapDate)
	pages := mapDate["pages"][0]

	wiki := pages.(map[string]interface{})

	return wiki
}

