package functional

import "fmt"

/*
  Extending a builder that we already have
  Lazily apply actions
*/
type Person struct {
	name, position string
}

func (p *Person) String() string {
	return fmt.Sprintf("%s %s", p.name, p.position)
}

type PersonMod func(p *Person)

type PersonBuilder struct {
	actions []PersonMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) WorksAs(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	person := Person{}
	for _, a := range b.actions {
		a(&person)
	}
	return &person
}
