package main

import (
	"fmt"
)

// Map はマッピング関数fを使って[]T1を[]T2に変換（「マップ」）する //liststart1
// 2つの型パラメータT1、T2をとる（T1とT2は違ってもよい）
// 任意の型のスライスを扱える
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s)) // 変換結果を記憶するスライス
	for i, v := range s {
		r[i] = f(v)  // sの各要素を関数fで変換し結果をrに記憶
	}
	return r  // 変換結果が記憶されたスライスを返す
}

// Reduceは[]T1を関数を用いて1個の値に「縮約（ひとつの値で代表）」する
// （この場合は、合計をとる）
func Reduce[T1, T2 any] (  // 型パラメータは2つ
	s []T1,  // Reduceの引数1
	initializer T2, // Reduceの引数2
	f func(T2, T1) T2) T2 { // Reduceの引数3 関数。T2とT1を引数としてT2の値を返す関数。下の例ではT2にはそれまでの合計、T2には新しい値が渡る。結果は合計
       // Reduceの戻り値の型はT2
	r := initializer  // r 最終的な結果（Result）。初期化する。下の例では0になる
	for _, v := range s {
		r = f(r, v)  // 
	}
	return r
}

// Filterは指定された関数を使ってスライスの要素をフィルタリングする
// この場合は関数fがtrueを返した要素だけからなるスライスを返す
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
} //listend1


func main() {
	words := []string{"One", "Potato", "Two", "Potato", "Three"} //liststart2
	filtered := Filter(words, func(s string) bool {
		return s != "Potato"
	})
	fmt.Println(filtered) // [One Two Three]

	lengths := Map(filtered, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)  // [3 3 5]
	
	sum := Reduce(lengths, // 第1引数 この場合は上でできた[3 3 5]
		0, // 第2引数。初期値0
		func(acc int, val int) int {
			return acc + val } ) // それまでの合計にvalを足す
	fmt.Println(sum) // 11  // ←3+3+5 //listend2

	var s = []int{10, 20, 30}
	s = Map(s, func(i int) int { // 自乗する
		return i*i
	})
	fmt.Println(s) // [100 400 900]
}
