package main

import "fmt"

func main() {
	const value = 10
	i := value  // valueのデフォルトの型はintになるのでこれでOK。「var i int = value」でもOK
	var f float64 = value  // こちらはfloat64（あるいはfloat32）を指定する必要あり
	fmt.Println(i)
	fmt.Println(f)
}
	
