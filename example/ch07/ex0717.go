package main

import (
	"fmt"
	"os"
)

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i2, ok := i.(int) // 「3.4.2 カンマokイディオム」参照 //liststart
	if !ok {
		err := fmt.Errorf("iの型（値:%v）が想定外です", i)
		fmt.Println(err.Error()) // エラーメッセージを表示（詳細は９章）
		os.Exit(1) // プログラムを終了
	}		
	fmt.Println(i2) //listend
}
