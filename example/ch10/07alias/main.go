package main

import "fmt"

type Foo struct { //liststart1
	x int
	S string
}

func (f Foo) Hello() string {
	return "hello"
}

func (f Foo) goodbye() string {
	return "goodbye"
} //listend1


type Bar = Foo //liststart2
//listend2

func MakeBar() Bar { //liststart3
	bar := Bar {
		x: 20,
		S: "Hello",
	}
	var f Foo = bar
	fmt.Println(f.Hello())
	return bar
} //listend3

func main() {
	a :=MakeBar()
	fmt.Println(a)
}
