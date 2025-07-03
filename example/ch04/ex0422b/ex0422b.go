//
// 詳しくは 10章で説明しますが、標準以外のライブラリを使うにはモジュールにする必要があります
// 当面は このディレクトリで go run ex0422b.go を実行してください
package main

import (
	"fmt"
	"unicode/utf8" // utf8.RuneCountInString(word)で、wordの文字数を数えられる
	"github.com/rivo/uniseg"  // これを使うためにモジュールにする必要がある
)

func main() {
	words := []string{
		/*0*/	 "á",  // 書記素数 1  コードポイント1個
		/*1*/  "\u0061\u0301",  // 書記素数 2    コードポイント2個
		/*2*/  "\u00E1",  // "á"の16進表記   書記素数 1  コードポイント1個
		/*3*/  " ́a",   // 書記素数 2    コードポイント2個
		/*4*/  "\u0301\u0061",  //  ́a
		/*5*/  "\u0061",  // 参考  "a"
		/*6*/  "\u0301",  // 参考  " ́"
	}

	
	for i, word := range words {
		l0 := len(word) // バイト数
		l1 := utf8.RuneCountInString(word) // runeの数
		l2 := len([]rune(word)) // runeの数
		l3 := uniseg.GraphemeClusterCount(word) // 書記素数
		fmt.Println(i, ":", "word:", word, " l0(byte): ", l0, "; l1: ", l1, " l2(rune): ", l2,  "  l3(書記素): ", l3)
	}
}
