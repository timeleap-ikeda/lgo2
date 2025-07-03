package main

import "fmt"

func f1(a string) int { //liststart
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

func main() {
	var myFuncVariable func(string) int
	myFuncVariable = f1
	result := myFuncVariable("こんにちは")
	fmt.Println(result) // 5

	myFuncVariable = f2
	result = myFuncVariable("こ")
	fmt.Println(result)
	result = myFuncVariable("ん")
	fmt.Println(result)
	result = myFuncVariable("こん")
	fmt.Println(result)
	result = myFuncVariable("こんにちは")
	fmt.Println(result)
}
