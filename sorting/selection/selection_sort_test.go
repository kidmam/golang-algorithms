package selection

import (
	"reflect"
	"testing"
)

func getSmallestIndex(s []int) int {
	smallestIndex := 0
	smallestItem := s[0]

	for index, item := range s {
		if item < smallestItem {
			smallestIndex = index
			smallestItem = item
		}
	}

	return smallestIndex
}

func selectionSort(toSort []int) (sorted []int) {
	for index := len(toSort) - 1; index >= 0; index-- {

		// gets the index from the smallest value
		smallestIndex := getSmallestIndex(toSort)

		// append the value from the smallest item
		sorted = append(sorted, toSort[smallestIndex])

		// remove the item just appended
		toSort = append(toSort[:smallestIndex], toSort[smallestIndex+1:]...)
	}

	return sorted
}

func Test_SelectionSort(t *testing.T) {
	tt := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 2, 9, 12, 4, 6, 0, 6},
			expected: []int{0, 2, 4, 5, 6, 6, 9, 12},
		},
		{
			input:    []int{9, 1, 2, 3, 4, 5, 1, 2, 3, 1},
			expected: []int{1, 1, 1, 2, 2, 3, 3, 4, 5, 9},
		},
	}

	for _, test := range tt {
		t.Run("", func(t *testing.T) {
			sorted := selectionSort(test.input)

			if !reflect.DeepEqual(test.expected, sorted) {
				t.Errorf("expected %v, got %v", test.expected, sorted)
			}

		})
	}
}

func Test_FindSmallestIndex(t *testing.T) {
	tt := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 10, 9, 7, 5, 0, 4},
			expected: 6,
		},
		{
			input:    []int{10, 9, 8, 7, 6},
			expected: 4,
		},
		{
			input:    []int{0, 9, 8, 7, 6},
			expected: 0,
		},
	}

	for _, test := range tt {
		t.Run("", func(t *testing.T) {
			smallest := getSmallestIndex(test.input)

			if smallest != test.expected {
				t.Errorf("expected %d, got %d", test.expected, smallest)
			}
		})
	}
}
