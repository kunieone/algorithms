package designmode

import "fmt"

type IPerson interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

// Here, NewPerson returns an interface, and not the person struct itself
func NewIPerson(name string, age int) IPerson {
	return person{
		name: name,
		age:  age,
	}
}
