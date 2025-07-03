package main

import (
	"fmt"
	"errors"
)

func doubleEven(i int) (int, error) {  // 偶数なら2倍して戻す //liststart1
	if i % 2 != 0 {  // 偶数ではないのでエラーを戻す
		return 0, errors.New("処理対象は偶数のみです")
	}
	return i * 2, nil  // 2倍して戻す
}

func main() {
	result, err := doubleEven(1)
	if err != nil {
		fmt.Println(err) // 処理対象は偶数のみです
	}
	fmt.Println(result)
}  //listend1
