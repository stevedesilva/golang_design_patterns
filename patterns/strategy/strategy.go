package strategy

import (
	"strings"
)

type OutputFormat int

const (
	MarkDown OutputFormat = iota
	Html
	Xml
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

//MarkDown
type MarkDownListStrategy struct {
}

func (m *MarkDownListStrategy) Start(builder *strings.Builder) {
}

func (m *MarkDownListStrategy) End(builder *strings.Builder) {
}

func (m *MarkDownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("*" + item + "\n")
}

// Html
type HtmlListStrategy struct{}

func (h HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (h HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(" <li>" + item + "</li>\n")
}

type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
	return &TextProcessor{builder: strings.Builder{}, listStrategy: listStrategy}
}

func (t *TextProcessor) SetOutputStrategy(format OutputFormat) {
	switch format {
	case MarkDown:
		t.listStrategy = &MarkDownListStrategy{}
	case Html:
		t.listStrategy = &HtmlListStrategy{}
	}
}

func (t *TextProcessor) AppendList(items []string) {
	t.listStrategy.Start(&t.builder)
	for _, item := range items {
		t.listStrategy.AddListItem(&t.builder, item)
	}
	t.listStrategy.End(&t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}
