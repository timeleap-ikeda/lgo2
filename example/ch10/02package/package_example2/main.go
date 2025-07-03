package main

import (
	"fmt" // 標準ライブラリ

	formatter "github.com/mushahiroyuki/package_example2/do-format"
	"github.com/mushahiroyuki/package_example2/math"
)

func main() {
	num := math.Double(2)
	output := formatter.Number(num)
	fmt.Println(output)
}
