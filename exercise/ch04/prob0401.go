package main

import (
	"fmt"
	"math/rand"  // rand.Intnのため
)

func main() {
	mySlice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		mySlice = append(mySlice, rand.Intn(100))
		// rand.Intn(100)は0以上100未満の整数を戻す
	}
	fmt.Println(mySlice)
}

