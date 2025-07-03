package main

import (
	"fmt"
	"errors"
)

//liststart
func divAndRemainder(num, denom int) (result int, remainder int, err error) {	
	result, remainder = 20, 30  // 適当な値を代入
	if denom == 0 {
		return num, denom, errors.New("0で割ることはできません")
	}
	return num / denom, num % denom, nil
}

func main() {
	rs, rm, _ := divAndRemainder(5,2)
	fmt.Println(rs, rm)
}
