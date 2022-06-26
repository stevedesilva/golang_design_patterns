package strategy_test

import (
	. "golangDesignPatterns/patterns/strategy"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMarkdownListStrategy(t *testing.T) {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList([]string{"t1", "t2", "t3"})
	builder := strings.Builder{}
	builder.WriteString(" * t1\n * t2\n * t3\n")
	want := builder.String()
	got := tp.String()
	assert.Equal(t, want, got)
}

func TestHtmlListStrategy(t *testing.T) {
	tp := NewTextProcessor(&HtmlListStrategy{})
	tp.AppendList([]string{"t1", "t2", "t3"})
	builder := strings.Builder{}
	builder.WriteString("<ul>\n  <li>t1</li>\n  <li>t2</li>\n  <li>t3</li>\n</ul>\n")
	want := builder.String()
	got := tp.String()
	assert.Equal(t, want, got)
}

func TestHtmlAndMarkupListStrategies(t *testing.T) {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList([]string{"t1", "t2", "t3"})
	builder := strings.Builder{}
	builder.WriteString(" * t1\n * t2\n * t3\n")
	want := builder.String()
	got := tp.String()
	assert.Equal(t, want, got)
	tp.Reset()
	tp.SetOutputFormat(Html)
	tp.AppendList([]string{"t1", "t2", "t3"})
	builder = strings.Builder{}
	builder.WriteString("<ul>\n  <li>t1</li>\n  <li>t2</li>\n  <li>t3</li>\n</ul>\n")
	want = builder.String()
	got = tp.String()
	assert.Equal(t, want, got)
}
