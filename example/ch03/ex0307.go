package main

import "fmt"

func main() {
	x := make([]string, 0, 5) //liststart
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x) // x: [a b i j y]
	fmt.Println("y:", y) // y: [a b i j y]
	fmt.Println("z:", z) // z: [i j y]  //listend
}
