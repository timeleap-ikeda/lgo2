package main

import "fmt"

func main() {
	xSlice := []int{1,2,3,4}  //liststart
	xArrayPointer := (*[4]int)(xSlice)
	fmt.Println("xArrayPointer:", xArrayPointer) // &[1 2 3 4]
	xSlice[0] = 10
	xArrayPointer[1] = 20
	fmt.Println("xSlice: ", xSlice) // [10 20 3 4]
	fmt.Println("xArrayPointer: ", xArrayPointer) // &[10 20 3 4]  //listend
}
