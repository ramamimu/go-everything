package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockShape2D struct{ mock.Mock }

func newMockShape2D() *mockShape2D { return &mockShape2D{} }

func (m *mockShape2D) Area() float32 {
	args := m.Called()
	return float32(args.Int(0))
}

func (m *mockShape2D) Type() int {
	args := m.Called()
	return args.Int(0)
}

func TestCube_SurfaceArea(t *testing.T) {
	m := newMockShape2D()
	m.On("Area").Return(30)

	cube := NewCube(m)
	surfaceArea := cube.SurfaceArea()
	assert.Equalf(t, float32(180), surfaceArea, "Expected 180 but got %f", surfaceArea)
	assert.NotNil(t, cube, "Expected no error while create Cube instance")
}

func TestTringularPrism_SurfaceArea(t *testing.T) {
	trianglesMock := [2]mockShape2D{*newMockShape2D(), *newMockShape2D()}
	rectanglesMock := [3]mockShape2D{*newMockShape2D(), *newMockShape2D(), *newMockShape2D()}

	areaFuncName := "Area"
	trianglesMock[0].On(areaFuncName).Return(1)
	trianglesMock[1].On(areaFuncName).Return(2)
	rectanglesMock[0].On(areaFuncName).Return(1)
	rectanglesMock[1].On(areaFuncName).Return(2)
	rectanglesMock[2].On(areaFuncName).Return(3)

	tp := NewTringularPrism(&trianglesMock[0], &trianglesMock[1], &rectanglesMock[0], &rectanglesMock[1], &rectanglesMock[2])
	tpSurfaceArea := tp.SurfaceArea()
	assert.Equal(t, float32(9), tpSurfaceArea, "Expected 5 but got %f", tpSurfaceArea)

}
