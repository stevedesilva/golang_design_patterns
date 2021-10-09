package ocp

type Product struct {
	name   string
	colour Colour
	size   Size
}

type Colour int
const (
	red Colour = iota
	green
	blue
)

type Size int
const (
	small Size = iota
	medium
	large
)
