package main

import (
	"fmt"
	"consterr"
)


const (
	ErrFoo = consterr.Sentinel("foo error")
	ErrBar = consterr.Sentinel("bar error")
)

func main() {
	fmt.Println(ErrFoo);
	fmt.Println(ErrBar);
}
