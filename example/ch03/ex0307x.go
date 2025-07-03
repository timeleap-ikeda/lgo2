package main

import "fmt"

func main() {
	x := make([]string, 0, 5) 
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]  // ①（下の図参照）
	fmt.Println("x:", x) // x: [a b c d]
	fmt.Println("y:", y) // y: [a b]
	fmt.Println("z:", z) // z: [c d] 
	fmt.Println(len(x), len(y), len(z)) // 4 2 2
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3
	
	y = append(y, "i", "j", "k") // ②
	fmt.Println("x:", x) // x: [a b i j]  // len(x)は4。
	fmt.Println("y:", y) // y: [a b i j k]
	fmt.Println("z:", z) // z: [i j]
	fmt.Println(len(x), len(y), len(z)) // 4 5 2
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3
	
	x = append(x, "x") // ③
	fmt.Println("x:", x) // x: [a b i j x]
	fmt.Println("y:", y) // y: [a b i j x]
	fmt.Println("z:", z) // z: [i j]
	fmt.Println(len(x), len(y), len(z)) // 5 5 2
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3
	
	z = append(z, "y") // ④
	fmt.Println("x:", x) // x: [a b i j y]
	fmt.Println("y:", y) // y: [a b i j y]
	fmt.Println("z:", z) // z: [i j y]
	fmt.Println(len(x), len(y), len(z)) // 5 5 3
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3 
}

/*  *5* など前後に「*」が付いているものは、そこで値が変わったことを示します

                                       lenの値    capの値
①の後                                  x   y   z    x   y   z
   a   b   c   d                       4   2   2    5   5   3
   <---- x ---->
   <-y->   <-z->
②の後  y = append(y, "i", "j", "k") // 3番目以降が i j k になる
   a   b   i   j   k                   4  *5*  2    5   5   3
   <---- x ----> // a b i j
   <------ y ------>  // a b i j k    len(y)は5
           <-z->      // i j          len(z)は変わらず2
③の後  x = append(x, "x")  // len(x)は4なので、5番目kがxに変わる
   a   b   i   j   x                  *5*  5   2    5   5   3
   <------ x ------>  // a b i j x    len(x)は4→5に変わる
   <------ y ------>  // a b i j x   
           <-z->      // i j          len(z)は変わらず2
④の後  z = append(z, "y") // len(z)は2なので、後ろにあるxがyに変わる
   a   b   i   j   y                   5   5  *3*    5   5   3
   <------ x ------>  // a b i j y     
   <------ y ------>  // a b i j y   
           <-- z -->  // i j y         len(z)は2→3に変わる

*/
