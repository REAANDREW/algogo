package sorting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Merge(l []int, r []int) (returnValue []int) {
	returnValue = make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			returnValue = append(returnValue, r...)
			break
		}
		if len(r) == 0 {
			returnValue = append(returnValue, l...)
			break
		}
		if l[0] <= r[0] {
			returnValue = append(returnValue, l[0])
			l = l[1:]
		} else {
			returnValue = append(returnValue, r[0])
			r = r[1:]
		}
	}
	return
}

func TestMergeReturnsLeftWhenRightIsEmpty(t *testing.T) {
	var l = []int{1}
	var r = []int{}
	var output = Merge(l, r)
	var expected = []int{1}
	assert.Equal(t, expected, output)
}

func TestMergeReturnsRightWhenLeftIsEmpty(t *testing.T) {
	var l = []int{}
	var r = []int{1}
	var output = Merge(l, r)
	var expected = []int{1}
	assert.Equal(t, expected, output)
}

func TestMergeReturnsLeftThenRight(t *testing.T) {
	var l = []int{1}
	var r = []int{2}
	var output = Merge(l, r)
	var expected = []int{1, 2}
	assert.Equal(t, expected, output)
}

func TestMergeReturnsRightThenLeft(t *testing.T) {
	var l = []int{2}
	var r = []int{1}
	var output = Merge(l, r)
	var expected = []int{1, 2}
	assert.Equal(t, expected, output)
}

func TestMergeReturnsLeftThenRightThenLeft(t *testing.T) {
	var l = []int{1, 3}
	var r = []int{2}
	var output = Merge(l, r)
	var expected = []int{1, 2, 3}
	assert.Equal(t, expected, output)
}

func TestMergeReturnsRightThenLeftThenRight(t *testing.T) {
	var l = []int{2}
	var r = []int{1, 3}
	var output = Merge(l, r)
	var expected = []int{1, 2, 3}
	assert.Equal(t, expected, output)
}

func MergeSort(input []int) (returnValue []int) {
	if len(input) <= 1 {
		return input
	}
	var n = len(input) / 2
	var l = MergeSort(input[:n])
	var r = MergeSort(input[n:])

	return Merge(l, r)
}

func TestMergeSort(t *testing.T) {
	var input = []int{5, 4, 3, 2, 1, 9, 8, 7, 6}
	var expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var output = MergeSort(input)

	assert.Equal(t, expected, output)
}
