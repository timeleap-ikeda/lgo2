package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++ {
		total = total + i  // :=  を = に変える
		// :=ではスコープがfor文の中になるので、抜けると「var total int」で設定されたゼロ値に戻ってしまう
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Printf("total=%v\n", total)
}
