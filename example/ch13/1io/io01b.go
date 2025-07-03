// io01.goをUTF8の文字についてカウントするようにしたものです
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

// 文字（rune）ごとの出現回数を数える
func countCharacters(r io.Reader) (map[rune]int, error) {
	out := map[rune]int{} // 文字 -> 出現回数 のマップ
	br := bufio.NewReader(r) // バッファ付きリーダーを使用

	for {
		ch, _, err := br.ReadRune() // 1文字（rune）ずつ読み取る
		if err == io.EOF {
			return out, nil // EOFなら結果を返す
		}
		if err != nil {
			return nil, err // エラー発生時はnilを返す
		}

		// 文字が空白でない場合のみカウント
		if !unicode.IsSpace(ch) {
			out[ch]++
		}
	}
}

// 文字列を処理して文字カウントを行う
func countCharactersInString() error {
	s := "東京特許許可局から許可がおりた特許について特許庁に相談に行った。 😊"
	sr := strings.NewReader(s) // io.Reader を作成
	counts, err := countCharacters(sr)
	if err != nil {
		return err
	}

	// 結果を出力
	for char, count := range counts {
		fmt.Printf("%c: %d\n", char, count)
	}
	return nil
}

func main() {
	countCharactersInString()
}
