package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"} //liststart
	y := x[:2]
	fmt.Println(y) // [a b]  // xの先頭から2つ。つまりx[0]とx[1]と同じ
	z := x[1:]
	fmt.Println(z) // [b c d]  // xの2つめから残り全部。x[1]からx[3]までと同じ
	x[1] = "y"
	fmt.Println("x:", x)  // x: [a y c d]
	fmt.Println("y:", y)  // y: [a y]
	fmt.Println("z:", z)  // z: [y c d]
	y[0] = "x" // x[1]も変わる	
	fmt.Println("----")
	fmt.Println("x:", x)  // x: [x y c d]
	fmt.Println("y:", y)  // y: [x y]  // xの先頭から2つ
	fmt.Println("z:", z)  // z: [y c d]  // xの2つめから残り全部
	z[1] = "z"
	fmt.Println("----")
	fmt.Println("x:", x) // x: [x y z d]
	fmt.Println("y:", y) // y: [x y]
	fmt.Println("z:", z) // z: [y z d] //listend
}

/*
     ↓のように、同じ場所を共有している
    x[0]  x[1]  x[2]   x[3]
    y[0]  y[1]
          z[0]  z[1]   z[2]
-----------------------------
↓変化するもの
      ↓  その値     
x:   "a"   "b"   "c"   "d"    ← x := []string{"a", "b", "c", "d"}
y:   "a"   "b"                ← y := x[:2]
z:         "b"   "c"   "d"    ← z := x[1:]
x:   "a"   "y"   "c"   "d"    ← x[1] = "y"
y:   "x"   "y"                ← y[0] = "x"
z:         "y"   "z"   "d"    ← z[1] = "z"    

最終的な値
   {-------xの範囲---------}  → ["x" "y" "z" "d"]
   {--yの範囲--}              → ["x" "y"]
          {----zの範囲-----}  → ["y" "z" "d"]
*/
