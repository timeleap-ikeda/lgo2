package main

import "fmt"

func main() {
	{
		var x int = 10  // 型と初期値を指定する
		fmt.Printf("%T:%v\n",x, x)
	}

	{
		var x = 10  // 型を省略する（デフォルトの型になる。この場合int）
		fmt.Printf("%T:%v\n",x, x)		
	}
	{
		var x int  // 初期値を省略する（intの「ゼロ値」が入る。この場合0）
		fmt.Printf("%T:%v\n",x, x)		
	}
	{
		var x, y int = 10, 20  // 多重代入
		fmt.Printf("%T:%v  %T:%v \n",x, x, y, y)
	}
	{
		var x, y int  // 初期値の省略（ゼロ値が入る）
		fmt.Printf("%T:%v  %T:%v \n",x, x, y, y)		
	}
	{
		var x, y = 10, "hello"  // 型の異なる変数の宣言
		fmt.Printf("%T:%v  %T:%v \n",x, x, y, y)
	}
	{
		x := 10  // 「var x = 10」と同じ。関数内でのみ使える。一番簡単
		fmt.Printf("%T:%v\n",x, x)
	}
	{
	// 宣言リスト
		var (
			x    int
			y        = 20
			d, e     = 40, "hello"  // eは文字列
	  )
		fmt.Printf("%T:%v  %T:%v \n",x, x, y, y)
		fmt.Printf("%T:%v  %T:%v \n",d, d, e, e)
	}
}
