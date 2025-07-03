// http02c.go に「パスワイルドカード変数」を2つ使った例
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// "GET /greet/{name}" にマッチするハンドラ
	mux.HandleFunc("GET /greet/{name}", func(w http.ResponseWriter, r *http.Request) {  //liststart1
		name := r.PathValue("name") // URL の {name} を取得
		w.Write([]byte(fmt.Sprintf("Hello, %s!\n", name)))
	})  //listend1

	// サーバを起動
	port := 8080
	log.Printf("「http://localhost:%d/greet/」に続けて名前を入力", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
