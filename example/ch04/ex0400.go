// 4章のリスト番号の付いていないコードを試すためのものです。
// エラーになる例は、コメントになっています。

package main

import "fmt"

func main() {
	fmt.Println("===== 4.3.1 for文 =====")

	fmt.Println("----- 初期化パートの省略 -----")
	i := 0
	for ; i < 10; i++ {
    fmt.Println(i)
	}

	fmt.Println("----- 再設定パートの省略 -----")	
	for i := 0; i < 10; {
    fmt.Println(i)
    if i % 2 == 0 {
			i++ // iを1増やす
    } else {
			i+=2 // iを2増やす
    }
	}

	fmt.Println("----- 4.3.2 条件のみのfor 文 -----")
	{
		i := 1
		for i < 100 {
			fmt.Println(i)
			i = i * 2
		}
	}
	
	fmt.Println("----- 4.3.4 breakとcontinue -----")
	{
		x := 0
		for	{
			fmt.Println("x: ", x)
			x++
	    if 0<=x {
				break
			}
		}
	}
}
