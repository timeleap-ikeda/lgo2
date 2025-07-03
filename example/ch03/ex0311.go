package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}  //liststart
	y := make([]int, 2)  // 長さ2のスライスを作る
	copy(y, x[2:])  // x[2]からコピー
	fmt.Println(y) // [3 4]  //listend
}
