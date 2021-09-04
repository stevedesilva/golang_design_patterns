package decorator

import "fmt"

type Decorator interface {
	Decorate() string
}

type Tree struct {
	Type string
}

func (t *Tree) Decorate() string {
	return t.Type
}

type TreeColourDecorator struct {
	Decorator Decorator
	Decoration string
}

func (t *TreeColourDecorator) Decorate() string {
	return fmt.Sprintf("%s %s",t.Decorator.Decorate(), t.Decoration)
}
type TreeSizeDecorator struct {
	Decorator Decorator
	Size string
}

func (t *TreeSizeDecorator) Decorate() string {
	return fmt.Sprintf("%s %s",t.Decorator.Decorate(), t.Size)
}