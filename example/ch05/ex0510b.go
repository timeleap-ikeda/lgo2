package main

import "fmt"

var (  //liststart1
	add = func(i int, j int) int { return i + j }
	sub = func(i int, j int) int { return i - j }
	mul = func(i int, j int) int { return i * j }
	div = func(i int, j int) int { return i / j }
)  //listend1

func main() {  //liststart2
	x := add(2, 3)
	fmt.Println(x) // 5
	changeAdd()
	y := add(2, 3) // 8
	fmt.Println(y)
}

func changeAdd() {
	add = func(i int, j int) int { return i + j + j }
}  //listend2
