// ch02/ex0201.go
package main
import (
	"fmt"
	"math/cmplx"
)

// 「虚数リテラル」を使った書き方を ex0201x.go に置きました

func main() { //liststart
	x := complex(2.5, 3.1) // :=は変数の宣言と代入を兼ねる。「2.2 変数の宣言」参照
	y := complex(10.2, 2)  // 変数xとyの型はcomplex128（複素数型）になる
	fmt.Println(x + y) // (12.7+5.1i)
	fmt.Println(x - y) // (-7.699999999999999+1.1i)
	fmt.Println(x * y) // (19.3+36.62i)
	fmt.Println(x / y) // (0.2934098482043688+0.24639022584228065i)
	fmt.Println(real(x)) // 2.5
	fmt.Println(imag(x)) // 3.1
	fmt.Println(cmplx.Abs(x)) // 3.982461550347976
}//listend
