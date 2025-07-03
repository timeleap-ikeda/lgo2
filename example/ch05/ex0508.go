package main

import (
	"fmt"
	"errors"
	"os"
)

//liststart
func divAndRemainder(num int, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("0で割ることはできません")
		return 
	}
	result, remainder = num / denom, num % denom
	return
}
//listend

func callDivAndRemainder(num int, denom int) {
	x, y, z := divAndRemainder(num, denom)
	if z!=nil {
		fmt.Print(num, "÷", denom, "：")
		fmt.Println(z)
		os.Exit(1)
	}
	fmt.Print(num, "÷", denom, " = ", x, "余り" , y, "\n")
}

func main() {
	callDivAndRemainder(5,2)
	callDivAndRemainder(10,0)
}

