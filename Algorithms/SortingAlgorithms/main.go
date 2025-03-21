package main

import (
	insertionsort "algorithms-and-data-structures/Algorithms/SortingAlgorithms/insertionSort"
	"fmt"
	"reflect"
)

func main() {
	TestSortArray()
}

func TestSortArray() {
	var testCases = []struct {
		input    []int
		expected []int
	}{
		{[]int{6, 4, 3, 1, 6, 2}, []int{1, 2, 3, 4, 6, 6}},
		{[]int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}

	for _, tc := range testCases {
		insertionsort.Sorted(tc.input)
		if !reflect.DeepEqual(tc.input, tc.expected) {
			fmt.Printf("Result = %v; expected %v\n", tc.input, tc.expected)
		}
	}
}
