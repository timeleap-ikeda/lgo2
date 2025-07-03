package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(BoolInt{}),    //liststart2
		unsafe.Offsetof(BoolInt{}.b),
		unsafe.Offsetof(BoolInt{}.i))
	fmt.Println(unsafe.Sizeof(IntBool{}),
		unsafe.Offsetof(IntBool{}.i),
		unsafe.Offsetof(IntBool{}.b))  //listend2
	fmt.Println()
	fmt.Println(unsafe.Sizeof(BoolIntBool{}),
		unsafe.Offsetof(BoolIntBool{}.b),
		unsafe.Offsetof(BoolIntBool{}.i),
		unsafe.Offsetof(BoolIntBool{}.b2))
	fmt.Println(unsafe.Sizeof(BoolBoolInt{}),
		unsafe.Offsetof(BoolBoolInt{}.b),
		unsafe.Offsetof(BoolBoolInt{}.b2),
		unsafe.Offsetof(BoolBoolInt{}.i))
	fmt.Println(unsafe.Sizeof(IntBoolBool{}),
		unsafe.Offsetof(IntBoolBool{}.i),
		unsafe.Offsetof(IntBoolBool{}.b),
		unsafe.Offsetof(IntBoolBool{}.b2))
}

type BoolInt struct {  //liststart1
	b bool
	i int64
}

type IntBool struct {
	i int64
	b bool
}  //listend1

type BoolIntBool struct {  //liststart3
	b  bool
	i  int64
	b2 bool
}

type BoolBoolInt struct {
	b  bool
	b2 bool
	i  int64
}

type IntBoolBool struct {
	i  int64
	b  bool
	b2 bool
}  //listend3
