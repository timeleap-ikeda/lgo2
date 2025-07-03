package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	emp1 := Employee{"清", "少納言", 12345}
	emp2 := Employee{
		firstName: "紫",
		lastName:  "式部",
		id:        65432,
	}
	var emp3 Employee
	emp3.firstName = "小野"
	emp3.lastName = "小町"
	emp3.id = 98765

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
