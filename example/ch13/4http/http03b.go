package main

import(
	"fmt"
	"net/http"
	"log"
)


func main() {

	personMux := http.NewServeMux()
	personMux.HandleFunc("get /greet/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			name := r.PathValue("name")
			w.Write([]byte(fmt.Sprintf("Hello, %s!\n", name)))
		})
	
	dogMux := http.NewServeMux()
	dogMux.HandleFunc("get /bye/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(fmt.Sprintf("Good bye, %s!\n", name)))
		})

	mux := http.NewServeMux()
	mux.Handle("/person/", http.StripPrefix("/person", personMux))
	// "/person/" のパターンは personMux が処理する
	mux.Handle("/dog/", http.StripPrefix("/dog", dogMux))
	// "/dog/" のパターンは dogMux が処理する 
	fmt.Printf("%s", "次でテスト:\ncurl http://localhost:8080/dog/greet\ncurl http://localhost:8080/person/greet\n")

	log.Fatal(http.ListenAndServe(":8080", mux))		

}
