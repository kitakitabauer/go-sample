package json_html_server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == "POST" {
		// リクエストボディをJSONに変換
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		// ファイルを {id}.txt とする
		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename) // ファイルを生成
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// ファイルにNameを書き込む
		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)
	http.ListenAndService(":3000", nil)
}
