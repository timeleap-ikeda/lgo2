package main
import (
	"fmt"
)

func main() {
	var x int = 10 //liststart
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, sum2) // 40.2 40  //listend
}
