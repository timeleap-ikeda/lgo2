package pet

import "github.com/learning-go-book-2e/ch10/sample_code/circular_dependency_example/person"  //liststart1
	//listend1

type Pet struct {
	Name      string
	Type      string
	OwnerName string
}

var owners = map[string]person.Person{  //liststart2
	"Bob":   {"Bob", 30, "Fluffy"},
	"Julia": {"Julia", 40, "Rex"},
}  //listend2

func (p Pet) Owner() person.Person {
	return owners[p.OwnerName]
}
