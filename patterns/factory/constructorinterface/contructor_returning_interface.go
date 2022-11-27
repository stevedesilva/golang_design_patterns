package constructorinterface

import "fmt"

type Person interface {
	SayHello()
}
type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi my name is %s, I am %d years old.\n", p.name, p.age)
}

type oldPerson struct {
	name string
	age  int
}

func (p *oldPerson) SayHello() {
	fmt.Printf("I'm too old to talk.\n")
}

func NewPerson(name string, age int) Person {
	if age > 99 {
		return &oldPerson{name: name, age: age}
	} else {

		return &person{name: name, age: age}
	}
}
