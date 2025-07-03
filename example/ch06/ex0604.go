package main

import "fmt"


func main() {

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}
	
	p := person{ //liststart2
		FirstName:  "Pat",
		MiddleName: makePointer("Perry"), // これならうまくいく
		LastName:   "Peterson",
	}
	fmt.Println(p) // {Pat 0xc000010250 Peterson}
	fmt.Println(*p.MiddleName) // Perry  //listend2
}

func makePointer[T any](t T) *T { // ヘルパー関数  //liststart1
	return &t
} //listend1

