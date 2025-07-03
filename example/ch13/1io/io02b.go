// ex1102.goをUTF8の文字についてカウントするようにしたものです
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"compress/gzip"
)

// 文字（rune）ごとの出現回数を数える
func countCharacters(r io.Reader) (map[rune]int, error) {
	out := map[rune]int{}  // 文字 -> 出現回数 のマップ
	br := bufio.NewReader(r)  // バッファ付きリーダーを使用

	for {
		ch, _, err := br.ReadRune()   // 1文字（rune）ずつ読み取る
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}

		// 文字が空白でない場合のみカウント
		if !unicode.IsSpace(ch) {
			out[ch]++
		}
	}
}

// GZip圧縮されたファイルを開いて `gzip.Reader` を作成
func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		r.Close()
		return nil, nil, err
	}
	return gr, func() { // クリーンアップ関数
		gr.Close()
		r.Close()
	}, nil
}

// GZipファイルを処理し、文字数をカウントする
func countCharactersInGzipFile() error {
	r, closer, err := buildGZipReader("my_data_rune.txt.gz")
	if err != nil {
		return err
	}
	defer closer()

	counts, err := countCharacters(r)
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
	countCharactersInGzipFile()
}
