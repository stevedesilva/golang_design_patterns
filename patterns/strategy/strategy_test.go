package strategy_test

import (
	"fmt"
	. "golangDesignPatterns/patterns/strategy"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestHtmlListStrategy(t *testing.T) {
	tp := NewTextProcessor(&MarkDownListStrategy{})
	tp.AppendList([]string{""})
	want := " * "
	fmt.Println(">", tp)
	fmt.Println(">", want)
	assert.Equal(t, tp, want)
}
