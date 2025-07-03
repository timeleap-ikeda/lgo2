package main

import (
	"fmt"
)


type Thinger interface {
	Thing()
}

type ThingerInt int

func (t ThingerInt) Thing() {
	fmt.Println("ThingInt:", t)
}

type ThingerSlice []int

func (t ThingerSlice) Thing() {
	fmt.Println("ThingSlice:", t)
}


func Comparer[T comparable](t1, t2 T) {
	if t1 == t2 {
		fmt.Println("等しい！")
	}
}

func main() {
	var a int = 10
	var b int = 10
	Comparer(a, b) // 等しい！
	
	var a3 ThingerSlice = []int{1, 2, 3}  //liststart1
	var b3 ThingerSlice = []int{1, 2, 3}
	Comparer(a3, b3) // コンパイル時のエラー: 
	// "ThingerSlice does not satisfy comparable"  //listend1
}
