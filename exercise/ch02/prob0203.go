package main

import "fmt"

func main() {
	var b byte = 255  // 0xFF のほうがわかりやすいかも
	var smallI int32 = 2147483647 // 0x7FFFFFFF （同上）
	var bigI uint64 = 18446744073709551615 // 0xFFFFFFFFFFFFFFFF （同上）
	b = b + 1  // b++  あるいは b += 1  でも同じ
	smallI = smallI + 1  // 同上。
	bigI = bigI + 1 // 同上。
	fmt.Println(b)   // 0  // いずれもオーバーフローする。エラーにはならない
	fmt.Println(smallI) // -2147483648  
	fmt.Println(bigI) // 0
}
