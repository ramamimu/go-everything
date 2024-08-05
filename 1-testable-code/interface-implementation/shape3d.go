// in this package will use mock to test unit testing
package shape

import "errors"

type Shape3D interface {
	SurfaceArea() float32
	// Volume() float32
}

type Cube struct {
	shape Shape2D
}

func NewCube(s Shape2D) (*Cube, error) {
	if s.Type() != RectangleShape {
		return nil, errors.New("invalid type for creating Cube instance")
	}

	return &Cube{
		shape: s,
	}, nil
}

func (c Cube) SurfaceArea() float32 {
	return 6 * c.shape.Area()
}
