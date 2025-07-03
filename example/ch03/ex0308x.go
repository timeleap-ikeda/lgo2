package main

import "fmt"

func main() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println("x:", x) // x: [a b c d]
	fmt.Println("y:", y) // y: [a b]
	fmt.Println("z:", z) // z: [c d]
	fmt.Println(len(x), len(y), len(z)) // 4 2 2
	fmt.Println(cap(x), cap(y), cap(z)) // 5 2 2

	y = append(y, "i", "j", "k")
	fmt.Println("x:", x) // x: [a b c d]
	fmt.Println("y:", y) // y: [a b i j k]
	fmt.Println("z:", z) // z: [c d]
	fmt.Println(len(x), len(y), len(z)) // 4 5 2
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 2

	x = append(x, "x")
	fmt.Println("x:", x) // x: [a b c d x]
	fmt.Println("y:", y) // y: [a b i j k]
	fmt.Println("z:", z) // z: [c d]
	fmt.Println(len(x), len(y), len(z)) // 5 5 2
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 2

	z = append(z, "y")
	fmt.Println("x:", x) // x: [a b c d x]
	fmt.Println("y:", y) // y: [a b i j k]
	fmt.Println("z:", z) // z: [c d y]
	fmt.Println(len(x), len(y), len(z)) // 5 5 3
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 4
}
