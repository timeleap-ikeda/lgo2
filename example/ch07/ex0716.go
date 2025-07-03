package main

import (
	"fmt"
)

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine  // mineの型はMyInt。iはanyなので代入できる  //liststart
	i2 := i.(int) // iに代入された型はMyIntなので、パニック！
	fmt.Println(i2)  //listend
}
