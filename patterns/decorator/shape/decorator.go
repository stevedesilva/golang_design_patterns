package shape

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Radius of the circle %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %.2f", s.Side)
}

type ColorShapeDecorator struct {
	Shape  Shape
	Colour string
}

func (c *ColorShapeDecorator) Render() string {
	return fmt.Sprintf("%s has the colour %s", c.Shape.Render(), c.Colour)
}