
package main

import "fmt"

type EmptyStruct struct{}

func main() {
	intSet := map[int]EmptyStruct{} //「int→空の構造体」のマップ。要素なし //liststart
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = EmptyStruct{}
	}
	if _, ok := intSet[5]; ok {
		fmt.Println("5は含まれている")
	}
	if _, ok := intSet[500]; ok {
		fmt.Println("500は含まれている")
	} //listend
}
