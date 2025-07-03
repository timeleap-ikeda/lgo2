package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mySlice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		mySlice = append(mySlice, rand.Intn(100))
	}
	for i, v := range mySlice {
		switch {
		case v%6 == 0:
			fmt.Printf("%d: Six! (%d)\n", i, v)
		case v%2 == 0:
			fmt.Printf("%d: Two! (%d)\n", i, v)
		case v%3 == 0:
			fmt.Printf("%d: Three! (%d)\n", i, v)
		default:
			fmt.Printf("%d: Never mind (%d)\n", i, v)
		}
	}
}
