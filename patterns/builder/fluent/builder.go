package fluent

import (
	"fmt"
	"strings"
)

const indentSize = 2

// HtmlElement - elements are not public
type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

// HtmlBuilder - builder functions are public
type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(name string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: name,
		root: HtmlElement{
			name:     name,
			text:     "",
			elements: []HtmlElement{},
		},
	}
}

func (h *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{name: childName, text: childText, elements: []HtmlElement{}}
	h.root.elements = append(h.root.elements, e)
}

func (h *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{name: childName, text: childText, elements: []HtmlElement{}}
	h.root.elements = append(h.root.elements, e)
	return h
}

func (h *HtmlBuilder) String() string {
	return h.root.String()
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}
