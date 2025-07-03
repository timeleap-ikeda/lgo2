package main

import "fmt"

func div(num int, denom int) int { // 分子（numerator）, 分母（denominator)
	if denom == 0 {
		return 0
	}
	return num / denom
}

func main() {
	result := div(5, 2)
	fmt.Println(result)
}
