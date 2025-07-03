package main

import (
	"fmt"  // 標準ライブラリ
	"github.com/mushahiroyuki/package_example_simple2/math"
	"github.com/mushahiroyuki/package_example_simple2/format"
)

func main() {
	num := math.Double(3)
	output := format.AddNote("2倍すると", num)
	fmt.Println(output)

	num = math.Square(3)
	output = format.AddNote("2乗すると", num)
	fmt.Println(output)
}
