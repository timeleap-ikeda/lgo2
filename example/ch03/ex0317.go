package main

import "fmt"

func main() {
	var s string = "Hello there"  //liststart
//                01234567890
	var b byte = s[6]
	fmt.Println("b:", b) // b: 116
	fmt.Printf("b (%%cを指定): %c\n", b) // b (%cを指定): t
	var c rune = rune(s[6])
	fmt.Println("c:", c) // // c: 116
	fmt.Printf("c (%%cを指定): %c\n", c) // c (%cを指定): t
	
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2) // o t
	fmt.Println(s3) // Hello
	fmt.Println(s4) // there  //listend
}
