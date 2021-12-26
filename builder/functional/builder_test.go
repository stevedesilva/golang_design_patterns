package functional_test

import (
	"golangDesignPatterns/builder/functional"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonBuilder_Build(t *testing.T) {
	b := functional.PersonBuilder{}
	got := b.Called("Steve").WorksAs("CEO").Build()
	assert.Equal(t, "Steve CEO", got.String())
}
