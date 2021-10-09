package ocp


type Filter struct {
  //
}
// initial requirement was filter by colour
func (f *Filter) FilterByColour(products []Product, color Colour) []Product{
	result := make([]Product,0)
	for _,v := range products {
		if color == v.colour {
			result = append(result,v)
		}
	}
	return result
}

// violation of the open close principle
// after some time a new requirement was filter by size which required modification of type
func (f *Filter) FilterBySize(products []Product, size Size) []Product{
	result := make([]Product,0)
	for _,v := range products {
		if size == v.size {
			result = append(result,v)
		}
	}
	return result
}

// after some time a new requirement was filter by size and colour which required modification of type
func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Colour) []Product{
	result := make([]Product,0)
	for _,v := range products {
		if size == v.size && color == v.colour {
			result = append(result,v)
		}
	}
	return result
}
