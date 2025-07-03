package main //liststart

import (
	"fmt"
	"math/rand"
)


func main() {
	n := rand.Intn(10) // 0以上10未満の整数を戻す
	if n == 0 {
		fmt.Println("少し小さすぎます:", n)
	} else if n > 5 {
		fmt.Println("大きすぎます:", n)
	} else {
		fmt.Println("いい感じの数字です:", n)
	}
} //listend
