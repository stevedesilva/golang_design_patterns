package main

import (
	"fmt"
	. "golangDesignPatterns/patterns/prototype"
)

func main() {
	fmt.Printf("hello.\n")
	personA := Person{
		Name: "Person A",
		Address: &Address{
			Street:  "54 Olinda Road",
			City:    "London",
			Country: "England",
		},
	}

	personB := personA

	personB.Name = "Person B"
	personB.Address = &Address{
		Street:  "Somewhere over the rainbow",
		City:    "Way up high",
		Country: "USA",
	}

	personA.Address.Street = "13 Obrien Road"
	personA.Address.City = "Dublin"
	personA.Address.Country = "Ireland"

	fmt.Printf("Person A= %v\n", personA)
	fmt.Printf("Person B= %v\n", personB)

	fmt.Printf("Person A address = %v\n", personA.Address)
	fmt.Printf("Person B address %v\n", personB.Address)
}
