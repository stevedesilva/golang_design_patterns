package copy

import (
	"bytes"
	"encoding/gob"
)

type Address struct {
	Street, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(p)
	println("buffer content" + string(b.Bytes()))

	d := gob.NewDecoder(&b)
	clone := Person{}
	_ = d.Decode(&clone)

	return &clone
}
