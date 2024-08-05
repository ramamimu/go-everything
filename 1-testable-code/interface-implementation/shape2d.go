package shape

type Shape2D interface {
	Area() float32
	Type() int
}

type Triangle struct {
	base   float32
	height float32
}

func (t Triangle) Area() float32 {
	return t.base * t.height / 2
}

func (t Triangle) Type() int {
	return TriangleShape
}

func NewTriangle(base float32, height float32) Shape2D {
	return Triangle{
		base:   base,
		height: height,
	}
}

type Square struct {
	sides float32
}

func (s Square) Area() float32 {
	return s.sides * s.sides
}

func (s Square) Type() int {
	return SquareShape
}

func NewSquare(sides float32) Shape2D {
	return Square{
		sides: sides,
	}
}

type Rectangle struct {
	sides1 float32
	sides2 float32
}

func (r Rectangle) Area() float32 {
	return r.sides1 * r.sides2
}

func (r Rectangle) Type() int {
	return RectangleShape
}

func NewRectangle(side1 float32, side2 float32) Shape2D {
	return Rectangle{
		sides1: side1,
		sides2: side2,
	}
}
