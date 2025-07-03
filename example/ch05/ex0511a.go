package main

import "fmt"

func main() {  //liststart
	a := 20
	f := func() {
		fmt.Println(a) // 20
		a := 30  // 新たに変数aを宣言。main()のaはシャドーイングされる
		fmt.Println(a) // 30
	} // ここまで関数の定義
	f() 
	fmt.Println(a) // 20
}  //listend

