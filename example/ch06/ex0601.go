package main

import "fmt"

func main() {
	x := 10  //liststart1
	pointerToX := &x
	fmt.Println(pointerToX) // アドレスが表示される
	fmt.Println(*pointerToX) // 10  // デリファレンスする
	z := 5 + *pointerToX
	fmt.Println(z) // 15   //listend1

	var y *int  //liststart2
	fmt.Println(y == nil) // true
	// fmt.Println(*y)  // ←パニックになる  //listend2

	{
		fmt.Println("--- new ---")
		var x = new(int)  //liststart3  // xの参照先にはintのゼロ値（0)が記憶される
		fmt.Println(x == nil) // false
		fmt.Println(x) // 0x1400000e130（例：xのアドレスが表示される）
		fmt.Println(*x == 0) // true
		fmt.Println(*x)   // 0  //listend3
	}
}
