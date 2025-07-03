package main

import "fmt"

func main() {
	f := func(j int) {
		fmt.Println("無名関数の中で", j, "を出力")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}
}
