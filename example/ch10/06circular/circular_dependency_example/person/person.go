package person

import "github.com/learning-go-book-2e/ch10/sample_code/circular_dependency_example/pet"  //liststart1
	//listend1

type Person struct {
	Name    string
	Age     int
	PetName string
}

var pets = map[string]pet.Pet{  //liststart2
	"Fluffy": {"Fluffy", "Cat", "Bob"},
	"Rex":    {"Rex", "Dog", "Julia"},
}  //listend2

func (p Person) Pet() pet.Pet {
	return pets[p.PetName]
}
