package main

import "fmt"

func main() {
	a := 20
	f := func() {  // func()からが関数の定義。定義を変数fに代入
		fmt.Println(a)  // 外側で定義された変数にアクセス
		a = 30
	}  // ここまで関数の定義。
	f()  // 20     // 変数fに代入された関数を実行 
	fmt.Println(a) // 30
}
