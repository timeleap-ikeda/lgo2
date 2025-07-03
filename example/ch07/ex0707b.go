package main

import "fmt"

const (  //liststart1
	Field1 = 0
	Field2 = 1 + iota  // constの2行目なのでiotaは1になる
	Field3 = 20        // 明示的に値が指定されるとその値になる
	Field4             // 値の指定がないと直前の値が使われる
	Field5 = iota      // constの5行目なのでiotaは4になる 
	Field6 = iota      // constの6行目なのでiotaは5になる 
	Field7             // 値の指定がないと直前の値（つまりiota）が使われる -> 6
)

func main() {
	fmt.Println(Field1, Field2, Field3, Field4, Field5, Field6, Field7)
}  //listend1
// 結果→    0 2 20 20 4 5 6
