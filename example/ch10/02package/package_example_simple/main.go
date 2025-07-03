package main

import (
	"fmt"  // 標準ライブラリ
	"github.com/mushahiroyuki/package_example_simple/math"
	"github.com/mushahiroyuki/package_example_simple/format"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
