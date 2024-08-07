package suite1_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DummyTestSuite struct {
	suite.Suite
	initNumber int
}

func TestDummy(t *testing.T) {
	suite.Run(t, new(DummyTestSuite))
}

// executes before the test suite begins
func (dummy *DummyTestSuite) SetupSuite() {
	fmt.Println("setup suite called")
	dummy.initNumber = 2
}

// executes after all tests executed
func (dummy *DummyTestSuite) TearDownSuite() {
	fmt.Println("tear down suite called")
}

// executes before each test case
func (suite *DummyTestSuite) SetupTest() {
	// reset StartingNumber to one
	fmt.Println("SetupTest called")
	suite.initNumber += 1
}

// executes after each test case
func (suite *DummyTestSuite) TearDownTest() {
	fmt.Println("TearDownTest called")
}

func (d *DummyTestSuite) TestDummyOne() {
	val := d.initNumber + 1
	assert.Equal(d.T(), 4, val, "expected 3 but got %d", val)
}
