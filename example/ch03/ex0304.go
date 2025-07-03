package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"} //liststart
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x) // x: [a b c d]
	fmt.Println("y:", y) // y: [a b]    
	fmt.Println("z:", z) // z: [b c d]  
	fmt.Println("d:", d) // d: [b c]    
	fmt.Println("e:", e) // e: [a b c d]  //listend
}
