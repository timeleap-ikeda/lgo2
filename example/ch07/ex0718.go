package main

import (
	"fmt"
	"io"
	"os"
)

type MyInt int

func main() {
	var i any
	fmt.Println("var i any")
	doThings(i)
	
	var mine MyInt = 20
	i = mine
	fmt.Println("\nvar mine MyInt = 20; i = mine")
	doThings(i)

	s := "これは文字列"
	fmt.Println("\ns := \"これは文字列\"")	
	doThings(s)

	s2 := []rune(s)
	fmt.Println("\ns2 := []rune(s)")	
	doThings(s2)

	fmt.Println("\ns2[0]")		
	doThings(s2[0])

	b := int(mine) < 20
	fmt.Println("\nb := int(mine) < 20")
	doThings(b)

	b = int(mine) == 20
	fmt.Println("\nb = int(mine) == 20")	
	doThings(b)
	
	type Person struct {
		FirstName string
		LastName string
		Age int
	}
	
	st := Person {
		FirstName: "John",
		LastName: "Baker",
		Age: 20,
	}
	fmt.Println("\nst := Person {...}")
	doThings(st)


	f, err := os.Open("ex0701.go")

	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	fmt.Println("\nf, err := os.Open(\"ex0701.go\")")
	doThings(f)

	finalInt := 100
	fmt.Println("\nfinalInt := 100")
	doThings(finalInt)
}

func doThings(i any) { //liststart1
	switch j := i.(type) {
	case nil: // iはnil。jの型はany
		fmt.Printf(" case nil; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case int: // jの型はint
		fmt.Printf(" case int; i:%d（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case MyInt: // jの型はMyInt
		fmt.Printf(" case MyInt; i:%d（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case io.Reader: // jの型はio.Reader
		fmt.Printf(" case io.Reader; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case string: // jは文字列
		fmt.Printf(" case string; i:%s（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case bool, rune: // iはboonかruneなので、jの型はany
		fmt.Printf(" case bool, rune; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	default: // iの型は不明。jの型はany
		fmt.Printf(" default; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	} //listend1
}
