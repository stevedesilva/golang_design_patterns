package dip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResearch_Investigate(t *testing.T) {
	// Given
	steve := Person{"Steve"}
	ben := Person{"Ben"}
	josephine := Person{"Josephine"}

	// create relationship
	relationships := &Relationships{}
	relationships.AddParentAndChild(&steve, &ben)
	relationships.AddParentAndChild(&steve, &josephine)

	// research
	research := Research{
		browser: relationships,
	}
	expected := []string{"Steve has a child called Ben", "Steve has a child called Josephine"}
	actual := research.Investigate("Steve")
	assert.Equal(t, expected, actual)
}
