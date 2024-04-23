package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	num1 := 3
	num2 := 4
	expect := 7
	result := Sum(num1, num2)
	assert.Equal(t, result, expect, "should be equal")
}

func TestSub(t *testing.T) {
	num1 := 3
	num2 := 4
	expect := -1
	result := Sub(num1, num2)
	assert.Equal(t, result, expect, "should be equal")
}

func TestSliceOrder(t *testing.T) {
	t.Run("Testing if array is sort", func(t *testing.T) {
		num1 := []int{1, 5, 2, 3, 7}
		expect := []int{1, 2, 3, 5, 7}
		result := OrdemCres(num1)
		assert.Equal(t, result, expect, "should be equal")
	})

	t.Run("Testing if array is sort and original stay equals", func(t *testing.T) {
		// variavel
		num1 := []int{1, 5, 2, 3, 7}
		original := []int{1, 5, 2, 3, 7}
		expect := []int{1, 2, 3, 5, 7}

		// exec
		result := OrdemCres(num1)

		// comp
		assert.Equal(t, result, expect, "should be equal")
		assert.Equal(t, num1, original, "should be equal")
	})
}

func TestDivide(t *testing.T) {
	t.Run("Divide a number normal", func(t *testing.T) {
		num1 := 4
		num2 := 2
		expect := 2
		result, err := Divide(num1, num2)
		assert.Equal(t, result, expect, "should be equal")
		assert.Nil(t, err, "should be nil")
	})

	t.Run("Divide a number by zero and erro", func(t *testing.T) {
		num1 := 4
		num2 := 0
		expect := 0
		result, err := Divide(num1, num2)
		assert.Equal(t, result, expect, "should be equal")
		assert.NotNil(t, err, "should not be nil")

	})
}
