package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func quicksort(arr []int) []int {
	// base case
	// if arr has 0 or 1 elements, it's already sorted
	if len(arr) < 2 {
		return arr
	} else {

		// recursive case

		// picks a random pivot
		rand.Seed(time.Now().Unix())
		pivotIndex := rand.Intn(len(arr))
		pivot := arr[pivotIndex]

		// removes pivot from slice
		arr = append(arr[:pivotIndex], arr[pivotIndex+1:]...)

		var left []int
		var right []int

		for _, item := range arr {
			// all elements lesser than the pivot are on the left
			if item <= pivot {
				left = append(left, item)
			} else {
				// all elements greater then the pivot are on the right
				right = append(right, item)
			}
		}

		// we recursively sort all the lefts, append the middle (pivot) and all the rights
		left = quicksort(left)
		left = append(left, pivot)
		right = quicksort(right)

		return append(left, right...)
	}
}

func Test_Quicksort(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output []int
	}{
		{
			desc:   "test zero elements",
			input:  []int{},
			output: []int{},
		},
		{
			desc:   "test one element",
			input:  []int{1},
			output: []int{1},
		},
		{
			desc:   "test two elements",
			input:  []int{10, 9},
			output: []int{9, 10},
		},
		{
			desc:   "test three+ elements",
			input:  []int{9, 5, 6, 8, 3},
			output: []int{3, 5, 6, 8, 9},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := quicksort(tC.input)

			if !reflect.DeepEqual(r, tC.output) {
				t.Errorf("expected %v, got %v", tC.output, r)
			}
		})
	}
}
