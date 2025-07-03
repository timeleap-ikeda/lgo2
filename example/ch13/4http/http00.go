// http *クライアント*
package main

import(
	"fmt"
	"net/http" // https://pkg.go.dev/netおよび https://pkg.go.dev/net/http 参照
	"io"
)

/*
 本文に記載されているとおり、タイムアウトがないため、本番環境での使用は避けてください。
*/

func main() {
	url := "https://www.example.com/"
	res, err := http.DefaultClient.Get(url)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
	fmt.Printf("%s", body)
}
