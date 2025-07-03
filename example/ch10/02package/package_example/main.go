package main

import (
	"fmt" // 標準ライブラリ

	format "github.com/mushahiroyuki/package_example/do-format"
	// ↑先頭のformatはなくても「format.xxx」で参照できる。
	// do-format/formatter.goで「package format」と宣言しているため
	"github.com/mushahiroyuki/package_example/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
