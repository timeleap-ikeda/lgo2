package main

import (
	"fmt"
)


func main() {
	for i := 1; i <= 100; i++ { //liststart
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println(i, "3でも5でも割り切れる")
		case i%3 == 0:
			fmt.Println(i, "3で割り切れる")
		case i%5 == 0:
			fmt.Println(i, "5で割り切れる")
		default:
			fmt.Println(i)
		}
	} //listend
}
