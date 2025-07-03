package main

import "fmt"

func main() {
	x := make([]string, 0, 5)  //liststart
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]  // yが利用できるのは、xの2番目まで
	z := x[2:4:4]  // zが利用できるのは、xの4番目まで
	fmt.Println(cap(x), cap(y), cap(z)) // 5 2 2
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x) // x: [a b c d x]
	fmt.Println("y:", y) // y: [a b i j k]
	fmt.Println("z:", z) // z: [c d y]   //listend
}
