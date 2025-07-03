package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("少し小さすぎます:", n)
	} else if n > 5 {
		fmt.Println("大きすぎます:", n)
	} else {
		fmt.Println("いい感じの数字です:", n)
	}
	fmt.Println(n) // エラー。nは未定義になる
}
