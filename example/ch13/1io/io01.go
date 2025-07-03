// io01b.go にUTF8の例があります

package main

import (
	"io"
	"fmt"
	"strings"
)

//liststart1
// 入力中の文字（アルファベット）について、出現回数を数える
// 「文字列 -> 整数」のマップ（とエラー）を返す
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{} // この関数の戻り値のマップ
	for {
		n, err := r.Read(buf) // 読み込み。io.Readerの唯一のメソッド
		for _, b := range buf[:n] { // 各文字についてループ
			// io.EOFがtrueでも 0<n ならこの部分は実行される
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF { // ファイル（文字列）の終わりに到達していたら
			return out, nil // マップを戻す
		}
		if err != nil {
			return nil, err
		}
	}
}
//listend1

func countLettersInString() error {
	s := "The quick brown fox jumped over the lazy dog" //liststart2
	sr := strings.NewReader(s)  // 文字列sからio.Readerを作る
	counts, err := countLetters(sr)
	if err != nil {
		return err
	}
	fmt.Println(counts) //listend2
	return nil
}

func main() {
	countLettersInString()
}
