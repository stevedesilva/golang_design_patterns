package main

import (
	"fmt"
	copy2 "golangDesignPatterns/patterns/prototype/copy"
	s "golangDesignPatterns/patterns/prototype/serialisation"
)

func main() {
	//copyExample()
	serialisationExample()
}

func copyExample() {
	fmt.Printf("hello.\n")
	personA := copy2.Person{
		Name: "Person A",
		Address: &copy2.Address{
			Street:  "54 Olinda Road",
			City:    "London",
			Country: "England",
		},
	}

	//personB := personA
	//
	//personB.Name = "Person B"
	//personB.Address = &copy2.Address{
	//	Street:  "Somewhere over the rainbow",
	//	City:    "Way up high",
	//	Country: "USA",
	//}

	personB := personA.DeepCopy()

	personA.Address.Street = "13 Obrien Road"
	personA.Address.City = "Dublin"
	personA.Address.Country = "Ireland"

	fmt.Printf("Person A= %v\n", personA)
	fmt.Printf("Person B= %v\n", personB)

	fmt.Printf("Person A address = %v\n", personA.Address)
	fmt.Printf("Person B address %v\n", personB.Address)
}

func serialisationExample() {
	fmt.Printf("hello.\n")
	personA := s.Person{
		Name: "Person A",
		Address: &s.Address{
			Street:  "54 Olinda Road",
			City:    "London",
			Country: "England",
		},
	}

	personB := personA.DeepCopy()

	personA.Address.Street = "13 Obrien Road"
	personA.Address.City = "Dublin"
	personA.Address.Country = "Ireland"

	fmt.Printf("Person A= %v\n", personA)
	fmt.Printf("Person B= %v\n", personB)

	fmt.Printf("Person A address = %v\n", personA.Address)
	fmt.Printf("Person B address %v\n", personB.Address)
}
