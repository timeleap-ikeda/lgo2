package main

import "fmt"

func main() {  //liststart1
	x := []int{10, 20, 30, 40, 50}
	for i, v := range x {
		fmt.Printf("%d, %d, %p\n", i, v, &v) // %p：ポインタのメモリアドレス
	}
}  //listend1
