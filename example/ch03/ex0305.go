package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"} //liststart
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x) // x: [x y z d]
	fmt.Println("y:", y) // y: [x y]
	fmt.Println("z:", z) // z: [y z d] //listend
}
