package copy

import "fmt"

type Address struct {
	Street, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		Street:  a.Street,
		City:    a.City,
		Country: a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	clone := *p
	fmt.Printf("clone %v original %v\n", clone, p)
	clone.Address = p.Address.DeepCopy()
	copy(clone.Friends, p.Friends)
	return &clone
}
