package fluent

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsBuilderShouldPElement(t *testing.T) {

	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString("Hello")
	sb.WriteString("</p>")
	assert.Equal(t, "<p>Hello</p>", sb.String())
	assert.Equal(t, sb.Len(), 12)

}

func TestStringsBuilderShouldPrintUnOrderList(t *testing.T) {
	words := []string{"hello", "world"}
	sb := strings.Builder{}
	sb.WriteString("<ul>")
	for _, w := range words {
		sb.WriteString("<li>")
		sb.WriteString(w)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	got := sb.String()
	fmt.Println(got)
	assert.Equal(t, "<ul><li>hello</li><li>world</li></ul>", got)
}

func TestBuilderShouldPrintUnOrderList(t *testing.T) {
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	got := b.String()
	fmt.Println(got)
	want := `<ul>
  <li>
    hello
  </li>
  <li>
    world
  </li>
</ul>
`
	assert.Equal(t, want, got)
}

func TestFluentBuilderShouldPrintUnOrderList(t *testing.T) {
	b := NewHtmlBuilder("ul").
		AddChildFluent("li", "hello").
		AddChildFluent("li", "world")
	got := b.String()
	fmt.Println(got)
	want := `<ul>
  <li>
    hello
  </li>
  <li>
    world
  </li>
</ul>
`
	assert.Equal(t, want, got)
}
