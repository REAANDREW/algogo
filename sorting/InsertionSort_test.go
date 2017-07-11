package sorting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func InsertionSort(input []int) []int {

	for i := 1; i < len(input); i++ {
		var j = i
		for j > 0 {
			if input[j] > input[j-1] {
				break
			}
			var temp = input[j]
			input[j] = input[j-1]
			input[j-1] = temp
			j--
		}
	}

	return input
}

func InsertionSortReverse(input []int) []int {

	for i := 1; i < len(input); i++ {
		var j = i
		for j > 0 {
			if input[j] < input[j-1] {
				break
			}
			var temp = input[j]
			input[j] = input[j-1]
			input[j-1] = temp
			j--
		}
	}

	return input
}

func TestInsertionSort(t *testing.T) {

	var input = []int{5, 4, 3, 2, 1}
	var expected = []int{1, 2, 3, 4, 5}
	var output = InsertionSort(input)

	assert.Equal(t, expected, output)
}

func TestInsertionSortReverse(t *testing.T) {
	var input = []int{1, 2, 3, 4, 5}
	var expected = []int{5, 4, 3, 2, 1}
	var output = InsertionSortReverse(input)

	assert.Equal(t, expected, output)
}
