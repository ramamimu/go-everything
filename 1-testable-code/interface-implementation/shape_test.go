package shape_test

import (
	"testing"

	shape "github.com/ramamimu/go-everything/1-testable-code/interface-implementation"
	"github.com/stretchr/testify/assert"
)

func TestTriangleDataType(t *testing.T) {
	triangleShape := shape.NewTriangle(9, 4)

	triangle, ok := triangleShape.(shape.Triangle)
	if !ok {
		t.Fatalf("Expected type shape.Triangle but got %T", triangle)
	}
}

func TestTriangleCalculation(t *testing.T) {
	triangle := shape.NewTriangle(4, 3)
	area := triangle.Area()
	if area != 6 {
		t.Errorf("Expected 6 but got %f", area)
	}
}

func TestSquareDataType(t *testing.T) {
	squareShape := shape.NewSquare(9)

	square, ok := squareShape.(shape.Square)
	assert.Truef(t, ok, "Expected shape.Square but got %T", square)
}

func TestSquareCalculation(t *testing.T) {
	square := shape.NewSquare(9)
	area := square.Area()

	assert.Equalf(t, float32(81), area, "Expected 81 but got %f", area)
}

func TestRectangleDataType(t *testing.T) {
	rectangleShape := shape.NewRectangle(2, 4)
	rectangle, ok := rectangleShape.(shape.Rectangle)
	assert.Truef(t, ok, "Expected shape.Rectangle but got %T", rectangle)
}

func TestRectangleCalculation(t *testing.T) {
	rectangle := shape.NewRectangle(3, 5)
	area := rectangle.Area()
	assert.Equal(t, float32(15), area, "Expected 15 but got %f", area)
}
