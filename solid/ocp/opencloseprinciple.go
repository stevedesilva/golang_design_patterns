package ocp

/*

OCP: Open Closed Principle

So the times in this case, the interface type is open for extension,
meaning you can implement this interface, but it's close to a modification, which means that you are unlikely to ever modify the
specification interface and in a similar fashion, you are unlikely to ever modify better filter because there's no reason for us to do so.

It's very flexible. It takes a bunch of products and a specification and that's pretty much all there is to it.

So it is better to just use this this idea of just implementing interfaces and making new types rather

than just working with a single type and just extending it over and over and over again.
 */
type Specification interface {
	IsSatisfied(p Product) bool
}

type ColourSpecification struct {
	colour Colour
}

func (c ColourSpecification) IsSatisfied(p Product) bool {
	return p.colour == c.colour
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p Product) bool {
	return p.size == s.size
}

type BetterFilter struct {
}

func (b *BetterFilter) Filter(product []Product, specification Specification)(result []Product){
	//result := make([]Product,0)
	for _,p := range product {
		if specification.IsSatisfied(p) {
			result = append(result, p)
		}
	}
	return
}

// Composite design pattern
type AndSpecification struct {
	first,second Specification
}

func (a AndSpecification) IsSatisfied(p Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}