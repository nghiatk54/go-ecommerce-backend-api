package basic

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 3
	// )

	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), output %d, actual = %d ", input, output, actual)
	// }

	assert.Equal(t, AddOne(2), 4, "AddOne(2) should be 3")

}

func TestAddOne2(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 3
	// )

	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), output %d, actual = %d ", input, output, actual)
	// }

	assert.Equal(t, AddOne2(2), 3, "AddOne2(2) should be 3")

}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not executing")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("Executing")
}
