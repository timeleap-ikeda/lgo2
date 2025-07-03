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
		fmt.Println("等しい!")
	}
}

func main() {
	var a int = 10
	var b int = 10
	Comparer(a, b) // 等しい!

	var a2 ThingerInt = 20
	var b2 ThingerInt = 20
	Comparer(a2, b2) // 等しい!
	
	var a4 Thinger = a2  //liststart1
	var b4 Thinger = b2
	Comparer(a4, b4) // 等しい!  //listend1
}


