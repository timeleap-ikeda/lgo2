package main

import "fmt"


func main() {  //liststart2

	add := func(i int, j int) int { return i + j }
	
	x := add(2, 3)
	fmt.Println(x)
	add = changeAdd()
	y := add(2, 3)
	fmt.Println(y)

}

func changeAdd() func(i int, j int) int {
	return func(i int, j int) int { return i + j + j }
}  //listend2
