package deepcopy

type Address struct {
	Street, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}
