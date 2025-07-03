package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(body string) string {
		return prefix + " " + body
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob")) // 「Hello Bob」と出力
	fmt.Println(helloPrefix("Maria")) // 「Hello Maria」と出力
}
