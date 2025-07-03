package main

import "fmt"

func UpdateSlice(ss []string, s string) {
	ss[len(ss)-1] = s
	fmt.Println("in UpdateSlice:", ss)
}

func GrowSlice(ss []string, s string) {
	ss = append(ss, s)
	fmt.Println("in GrowSlice:", ss)
}

func main() {
	s := []string{"a", "b", "c"}
	UpdateSlice(s, "d")
	fmt.Println("in main after UpdateSlice:", s)
	GrowSlice(s, "e")
	fmt.Println("in main, after GrowSlice:", s)
}
