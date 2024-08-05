package shape

type Shape interface {
	Area() float32
}

type Triangle struct {
	base   float32
	height float32
}

func (t Triangle) Area() float32 {
	return t.base * t.height / 2
}

func NewTriangle(base float32, height float32) Shape {
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

func NewSquare(sides float32) Shape {
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

func NewRectangle(side1 float32, side2 float32) Shape {
	return Rectangle{
		sides1: side1,
		sides2: side2,
	}
}
