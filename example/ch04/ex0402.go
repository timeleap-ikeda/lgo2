package main

import "fmt"

func main() { //liststart
	x := 10
	if x > 5 {
		a, x := 5, 20
		fmt.Println(a, x) // 5 20
	}
	fmt.Println(x) // 10
} //listend
