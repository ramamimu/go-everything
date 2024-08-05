package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCube struct{ mock.Mock }

func newMockCube() *mockCube { return &mockCube{} }

func (m *mockCube) Area() float32 {
	args := m.Called()
	return float32(args.Int(0))
}

func (m *mockCube) Type() int {
	args := m.Called()
	return args.Int(0)
}

func TestCube_NonRectangleShape(t *testing.T) {
	m := newMockCube()
	m.On("Area").Return(0)
	m.On("Type").Return(SquareShape)

	_, err := NewCube(m)
	assert.Error(t, err, "Expected error while create Cube instance due to not coming from Rectangle instance")
}

func TestCube_SurfaceArea(t *testing.T) {
	m := newMockCube()
	m.On("Area").Return(30)
	m.On("Type").Return(RectangleShape)

	cube, err := NewCube(m)
	surfaceArea := cube.SurfaceArea()
	assert.Equalf(t, float32(180), surfaceArea, "Expected 180 but got %f", surfaceArea)
	assert.NoError(t, err, "Expected no error while create Cube instance")
}
