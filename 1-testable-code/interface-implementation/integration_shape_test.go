package shape_test

import (
	"testing"

	shape "github.com/ramamimu/go-everything/1-testable-code/interface-implementation"
	"github.com/stretchr/testify/assert"
)

// ==== negative cube ====
// create cube with non rectangle instance
func TestCubeIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	square := shape.NewSquare(1)
	cube := shape.NewCube(square)

	surfaceArea := cube.SurfaceArea()
	assert.Equalf(t, float32(6), surfaceArea, "expected 6 but got %f", surfaceArea)
}

func TestTringularPrismIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	triangle1 := shape.NewTriangle(1, 2)   // 1			1
	triangle2 := shape.NewTriangle(3, 4)   // 6			7
	rectangle1 := shape.NewRectangle(1, 2) // 2			9
	rectangle2 := shape.NewRectangle(2, 2) // 4			13
	rectangle3 := shape.NewRectangle(3, 4) // 12		25

	tringularPrism := shape.NewTringularPrism(triangle1, triangle2, rectangle1, rectangle2, rectangle3)
	surfaceArea := tringularPrism.SurfaceArea()
	assert.Equalf(t, float32(25), surfaceArea, "expected 25 but got %f", surfaceArea)
}
