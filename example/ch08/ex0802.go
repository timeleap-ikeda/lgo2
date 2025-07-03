package main

import (
	"fmt"
)

type Stack[T any] struct { // 型Stackの宣言。ジェネリクス版 //liststart1
	vals []T  // 型Tのスライスで「（たとえばお皿が積まれる）棚」を表す
}

func (s *Stack[T]) Push(val T) { // メソッドPushの定義
	s.vals = append(s.vals, val)  // 型Tの皿valを積む（Push）
}

func (s *Stack[T]) Pop() (T, bool) { // メソッドPopの定義
	if len(s.vals) == 0 { // 棚が空（から）のとき
		var zero T  // 変数zeroの値は、型Tのゼロ値になる
		return zero, false  // ゼロ値と取り出せないことを示すfalseを返す
	}
	top := s.vals[len(s.vals)-1] // 棚から1個（topを）取り出す
	s.vals = s.vals[:len(s.vals)-1]  // スライスの長さを1減らす
	return top, true
} //listend1

func main() { //liststart2
	var intStack Stack[int] // int型のスタック
	intStack.Push(10) // 10をスタック（棚）に積む（「プッシュ」）
	intStack.Push(20)
	intStack.Push(30)
	v, ok := intStack.Pop() // 一番上の要素を取り出す（「ポップ」）
	fmt.Println(v, ok) // 30 true
	v, ok = intStack.Pop() // もうひとつ上からポップ
	fmt.Println(v, ok) // 20 true
	intStack.Push(15)
	v, ok = intStack.Pop() // 今プッシュしたばかりの要素をポップ
	fmt.Println(v, ok) // 15 true
	v, ok = intStack.Pop() // もうひとつポップ
	fmt.Println(v, ok) // 10 true
	v, ok = intStack.Pop()  // 値がないときはokにfalseが返る
	fmt.Println(v, ok) // 0 false  // vはゼロ値（0）
} //listend2
