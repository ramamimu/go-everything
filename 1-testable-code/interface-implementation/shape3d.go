// in this package will use mock to test unit testing
package shape

type Shape3D interface {
	SurfaceArea() float32
	// Volume() float32
}

type Cube struct {
	rectangle RectangleShape2D
}

func NewCube(r RectangleShape2D) Shape3D {
	return &Cube{
		rectangle: r,
	}
}

func (c Cube) SurfaceArea() float32 {
	return 6 * c.rectangle.Area()
}

type TringularPrism struct {
	triangles  [2]TriangleShape2D
	rectangles [3]RectangleShape2D
}

func NewTringularPrism(t1, t2 TriangleShape2D, r1, r2, r3 RectangleShape2D) Shape3D {
	return &TringularPrism{
		triangles:  [2]TriangleShape2D{t1, t2},
		rectangles: [3]RectangleShape2D{r1, r2, r3},
	}
}

func (tp TringularPrism) SurfaceArea() float32 {
	totalArea := float32(0)
	for _, triangle := range tp.triangles {
		totalArea += triangle.Area()
	}

	for _, rectangle := range tp.rectangles {
		totalArea += rectangle.Area()
	}

	return totalArea
}
