package lsp

/*If you have some API that takes a base class and works correctly with that base class,
then it should also work correctly with the derived class.

Liskov substitution principle  states that if you continue to use generalisations like interfaces,
for example, then you should not have inherited or you should not have implementors of
those generalisations break some of the assumptions which are set up at the higher level.
*/
type Sized interface {
	GetHeight() int
	SetHeight(height int)
	GetWidth() int
	SetWidth(width int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

// Use it should continue to work for both square and rectangle
func UseIt(sized Sized) (area int) {
	// by setting the height we are assuming only height is set
	// LSP violation : however square sets both height and width - resulting in incorrect area
	sized.SetHeight(10)
	return sized.GetWidth() * sized.GetHeight()
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	return &Square{
		Rectangle{width: size, height: size},
	}
}

// Violation: setting set both height and width
func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) GetHeight() int {
	return s.height
}

func (s *Square) GetWidth() int {
	return s.width
}

/*
Solutions
 */
type Square2 struct {
	size int
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}
