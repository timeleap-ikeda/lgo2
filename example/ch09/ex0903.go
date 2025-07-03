package main

import (
	"fmt"
)

func doubleEven(i int) (int, error) {  //liststart1
	if i % 2 != 0 {
		return 0, fmt.Errorf("%dは偶数ではありません", i)
	}
	return i * 2, nil
}  //listend1

func main() {
	result, err := doubleEven(1)
	if err != nil {
		fmt.Println(err) // 1は偶数ではありません
	}
	fmt.Println(result) // 0
}
