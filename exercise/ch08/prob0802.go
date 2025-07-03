package main

import (
	"fmt"
	"strconv"
)

/*
2. Printableというジェネリックなインタフェースを定義せよ
  ・fmt.Stringerを実装し、基底型がintまたはfloat64である型に一致する
  ・このインタフェースを満たす型を定義する
  ・Printableを受け取り、fmt.Printlnを使ってその値を画面（標準出力）に出力する関数を定義する
*/

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

func (pi PrintInt) String() string {
	return strconv.Itoa(int(pi))
}

type PrintFloat float64

func (pf PrintFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

func PrintIt[T Printable](t T) {
	fmt.Println(t)
}

func main() {
	var i PrintInt = 20
	PrintIt(i)

	var f PrintFloat = 10.23
	PrintIt(f)
}
