package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{dict}/{word}", func(w http.ResponseWriter, r *http.Request) {
		dict := r.PathValue("dict") // URL の {dict} を取得
		word := r.PathValue("word") // URL の {word} を取得
		w.Write([]byte(fmt.Sprintf("単語「%s」を辞書「%s」で検索します\n", word, dict)))
	})

	// サーバを起動
	port := 8080
	log.Printf("「http://localhost:%d/辞書名/単語」の形式で入力してください", port)
	
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
