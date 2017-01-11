package main

import (
	"fmt"
	// "github.com/tushar00jain/gocode/string"
)

type Person struct {
	Name string
}

type Saiyan struct {
	*Person
	Power int
}

func main() {
	// fmt.Printf("Hello, world.\n")
	// fmt.Println(string.Reverse("Hello, world."))

	goku := &Saiyan{&Person{"Goku"}, 9000}
	goku.Introduce()
	goku.Person.Introduce()

}

func Super(s *Saiyan) {
	s = &Saiyan{&Person{"Gohan"}, 1000}
}

func (p *Person) Introduce() {
	fmt.Println("Person", p.Name)
}

func (s *Saiyan) Introduce() {
	fmt.Println("Saiyan", s.Name)
}
