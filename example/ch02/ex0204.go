package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	x = x + 1 // @<ttb>{コンパイルできない！}
	y = "bye" // @<ttb>{コンパイルできない！}

	fmt.Println(x)
	fmt.Println(y)
}
