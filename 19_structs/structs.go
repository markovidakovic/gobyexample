package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"bob", 20})

	fmt.Println(person{name: "alice", age: 30})

	fmt.Println(person{name: "fred"})

	fmt.Println(&person{name: "ann", age: 40})

	fmt.Println(newPerson("john"))

	s := person{name: "sean", age: 50}
	fmt.Println(s.name)

	sp := &s            // struct pointer
	fmt.Println(sp.age) // the pointers are automatically dereferenced

	sp.age = 51
	fmt.Println(sp.age)
}
