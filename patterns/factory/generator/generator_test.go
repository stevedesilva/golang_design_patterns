package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmployeeFactory(t *testing.T) {
	developerFactory := NewEmployeeFactory("developer", 100000)
	managerFactory := NewEmployeeFactory("manager", 150000)

	mike := developerFactory("Mike")
	steve := managerFactory("Steve")

	assert.Equal(t, mike.Position, "developer")
	assert.Equal(t, mike.AnnualIncome, 100000)
	assert.Equal(t, steve.Position, "manager")
	assert.Equal(t, steve.AnnualIncome, 150000)
}
