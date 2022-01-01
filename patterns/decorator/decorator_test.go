package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Decorate_should_have_size(t *testing.T) {
	size := TreeSizeDecorator{Decorator: &Tree{"Fur tree"}, Size: "small"}
	got := size.Decorate()
	assert.Equal(t, "Fur tree small", got)
}

func TestTree_Decorate_should_have_colour_and_size(t *testing.T) {
	colour := TreeColourDecorator{Decorator: &TreeSizeDecorator{Decorator: &Tree{"Fur tree"}, Size: "small"}, Decoration: "red"}
	got := colour.Decorate()
	assert.Equal(t, "Fur tree small red", got)
}
