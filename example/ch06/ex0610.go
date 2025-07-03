package main

import (
	"fmt"
	"runtime"
	"time"
)

type A struct {
	b *B
}

type B struct {
	c *C
}

type C struct {
	field string
}

func makeAPointer() *A {
	a := &A{&B{&C{"hello"}}}	// a, a.b, および a.b.c によって参照されているデータをヒープ上に確保

	// 何かがガベージコレクトされると実行される関数を指定する
	runtime.SetFinalizer(a.b.c, func(c *C) { fmt.Println("a.b.c（値:", c.field, "）を回収") })
	return a
}

func main() {
	aPointer := makeAPointer()

	runtime.GC() // ガベージコレクションを強制。aPointerがヒープ上のデータを参照しているので、ガベージはまだない

	time.Sleep(20) // finalizerが実行するための時間を与える（ガベージがまだないので回収は行われない）
	fmt.Println("aPointerをnilにする前:", aPointer)
	
	aPointer = nil //  aPointerが指すデータをガベージにする
	fmt.Println("aPointerをnilにした後:", aPointer)

	runtime.GC() // ガベージコレクションを強制
	time.Sleep(20) // finalizerが実行するための時間を与える（ガベージコレクションが行われる）
}
