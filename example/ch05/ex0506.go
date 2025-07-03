package main

import (
	"fmt"
	"errors"
	"os"
)

//liststart
func divAndRemainder(num int, denom int) (result int, remainder int,
	err error) {
	if denom == 0 {
		return num, denom, errors.New("0で割ることはできません")
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}
//listend

func callDivAndRemainder(num int, denom int) {
	x, y, z := divAndRemainder(num, denom)
	if z!=nil {
		fmt.Print(x, "÷", y, "：")
		fmt.Println(z)
		os.Exit(1)
	}
	fmt.Print(num, "÷", denom, " = ", x, "余り" , y, "\n")
}

func main() {
	callDivAndRemainder(5,2)
	callDivAndRemainder(10,0)
}
