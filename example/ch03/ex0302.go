package main

import "fmt"

func main() {
	{
		var data []int  // スライスのゼロ値nilで初期化される（nilスライス）  //liststart1
		fmt.Println(data == nil)  // true
		fmt.Println(data)  // []
		fmt.Println(len(data))  // 0  //listend1

		x := []int{} // 空のスライスリテラル  //liststart2
		fmt.Println(x == nil)  // false
		fmt.Println(x)  // []   // (nilのときと出力結果は変わらない）
		fmt.Println(len(x))  // 0  //listend2
	}
}
