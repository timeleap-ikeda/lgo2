package main

import (
	"fmt"
)

type Adder struct { // 構造体Adderの定義
	start int  // int型のフィールドstartをもつ
}

func (a Adder) AddTo(val int) int { // 型Adderをレシーバーとするメソッドを定義
	return a.start + val // フィールドstartの値に、引数の値valを足して戻す
}

func main() {
	myAdder := Adder{start: 10} // startの値を10にしてAdderのインスタンスを生成
	fmt.Println("myAdder.AddTo(5):", myAdder.AddTo(5)) // 15

	f1 := myAdder.AddTo // Adder型の変数myAdderのメソッドAddToをf1に代入
	fmt.Println("f1(10):", f1(10)) // 20

	f2 := Adder.AddTo //型Adderをレシーバーとして定義されているメソッドAddToをf2に代入
	fmt.Println("f2(myAdder, 15):", f2(myAdder, 15)) // 25

} 
