package constructorstruct

import "errors"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) (*Person, error) {
	if age > 120 {
		return nil, errors.New("invalid age")
	}
	return &Person{
		Name:     name,
		Age:      age,
		EyeCount: 2, // Default value uses function
	}, nil
}
