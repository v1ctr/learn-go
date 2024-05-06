package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	// You can safely return a pointer to local variable as a local variable will survive the scope of the function.
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// It’s idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

	// If a struct type is only used for a single value, we don’t have to give it a name.
	// The value can have an anonymous struct type.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
