package main

import "fmt"

func main() {
	xSlice := []int{1, 2, 3, 4}  //liststart
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice) // [10 2 3 4]
	fmt.Println(xArray) // [1 2 3 4]
	fmt.Println(smallArray) // [1 2]   //listend
}
