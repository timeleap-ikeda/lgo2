package main

import "fmt"

func main() {
	var i  = 20
	var f = float64(i)  // 型変換が必要。 「float32(i)」でもOK
	fmt.Println(i)
	fmt.Println(f)
}
	
