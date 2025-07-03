package text

import (
	// 	"io/ioutil"  // 古い。非推奨
	"os"  // こちらが新しい
	"unicode/utf8"
)

func CountCharacters(fileName string) (int, error) {
	data, err := os.ReadFile(fileName)
	// data, err := ioutil.ReadFile(fileName) // 古い。非推奨
	if err != nil {
		return 0, err
	}
	return utf8.RuneCount(data), nil
}
