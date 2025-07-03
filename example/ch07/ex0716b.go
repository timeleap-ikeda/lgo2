package main

import (
	"fmt"
)

type MyInt int

func main() {
	var i any  //liststart
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt) // mineの値はMyInt型なので、intとしては扱えない
	fmt.Println(i2 + 1) //listend
}
