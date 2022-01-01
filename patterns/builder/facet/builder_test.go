package facet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPersonBuilder(t *testing.T) {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("10 Downing Street").
		In("London").
		WithPostcode("EC1 12U").
		Work().
		At("City Of London").
		AsA("MP").
		Earning(123000)
	person := pb.Build()
	want := &Person{
		StreetAddress: "10 Downing Street",
		Postcode:      "EC1 12U",
		City:          "London",
		CompanyName:   "City Of London",
		Position:      "MP",
		AnnualIncome:  123000,
	}
	assert.Equal(t, want, person)
}
