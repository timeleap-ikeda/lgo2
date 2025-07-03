package main

import "fmt"

func main() {
	{  //liststart
		var s string = "Hello ☀"
		var s2 string = s[4:7]
		var s3 string = s[:5]
		var s4 string = s[6:]
		fmt.Println(s)  // Hello ☀
		fmt.Println(s2) // o ?（文字化け）
		fmt.Println(s3) // Hello
		fmt.Println(s4) // ☀
		fmt.Println("len(s):", len(s))
	}
	{
		var s string = "こんにちは"
		fmt.Println("s:", s) // s: こんにちは
		fmt.Println("s[0]:", s[0]) // s[0]: 227
		var b byte = s[6]
		fmt.Println("b:", b) // b: 227
		var c rune = rune(s[6])
		fmt.Println("c:", c) // c: 227
	}
	{
		var s string = "こんにちは、みなさん"
		var s2 string = s[4:7]
		var s3 string = s[:5]
		var s4 string = s[6:]
		fmt.Println("s:", s) // s: こんにちは、みなさん
		fmt.Println("s2:", s2) // s2: ???
		fmt.Println("s3:", s3) // s3: こ?
		fmt.Println("s4:", s4) // s4: にちは、みなさん
	}  //listend
}
