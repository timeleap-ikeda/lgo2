package main

import (
	"fmt"
)

type MyInt int

func main() {
	var i any //liststart1
	var mine MyInt = 20
	i = mine
	i2 := i.(string) // iをstring型だと仮定して値をもらおうとするが... パニック！  
	fmt.Println(i2) //listend1
}
