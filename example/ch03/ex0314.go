package main

import "fmt"

func main() {
	x := [4]int{5, 6, 7, 8}  //liststart
	y := x[:2]
	z := x[2:]
	x[0] = 10
	fmt.Println("x:", x) // x: [10 6 7 8]
	fmt.Println("y:", y) // y: [10 6]  // @<ttb>{「10」に注意}
	fmt.Println("z:", z) // z: [7 8]  //listend
}
