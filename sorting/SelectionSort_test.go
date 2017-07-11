package sorting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func SelectionSort(input []int) (returnValue []int) {

	returnValue = input

	for i := 0; i < len(input); i++ {
		var minJ = i
		for j := i + 1; j < len(input); j++ {
			if returnValue[j] < returnValue[minJ] {
				minJ = j
			}
		}

		var minJValue = returnValue[minJ]
		returnValue[minJ] = returnValue[i]
		returnValue[i] = minJValue
	}

	returnValue = input
	return
}

func TestSelectionSort(t *testing.T) {
	var input = []int{5, 4, 3, 2, 1}
	var expected = []int{1, 2, 3, 4, 5}
	var output = SelectionSort(input)

	assert.Equal(t, expected, output)
}
