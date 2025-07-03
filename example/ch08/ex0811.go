package main

import (
	"fmt"
)


type Thinger interface {  //liststart1
	Thing()
}

type ThingerInt int

func (t ThingerInt) Thing() {
	fmt.Println("ThingInt:", t)
}

type ThingerSlice []int

func (t ThingerSlice) Thing() {
	fmt.Println("ThingSlice:", t)
}  //listend1


func Comparer[T comparable](t1, t2 T) {  //liststart2
	if t1 == t2 {
		fmt.Println("等しい!")
	}
}  //listend2

func main() {
	var a int = 10  //liststart3
	var b int = 10
	Comparer(a, b) // 等しい!
	
	var a2 ThingerInt = 20
	var b2 ThingerInt = 20
	Comparer(a2, b2) // 等しい!  //listend3
}
