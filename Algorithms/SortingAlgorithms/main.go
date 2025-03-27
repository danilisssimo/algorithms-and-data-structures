package main

import (
	timsort "algorithms-and-data-structures/Algorithms/SortingAlgorithms/TimSort"
	insertionsort "algorithms-and-data-structures/Algorithms/SortingAlgorithms/insertionSort"
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	TestCasesForInsertionSort()
	TestCasesForTimsort()
}

func TestCasesForInsertionSort() {
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

func TestCasesForTimsort() {
	var testCase []int = []int{
		603, 230, 872, 816, 904, 308, 768,
		255, 616, 18, 209, 708, 895, 130,
		224, 440, 649, 173, 514, 944, 744,
		510, 576, 596, 554, 583, 211, 746,
		58, 854, 136, 576, 511, 455, 93,
		755, 886, 15, 69, 121, 541, 162,
		844, 564, 161, 474, 373, 625, 826,
		680, 597, 720, 213, 998, 864, 729,
		758, 809, 543, 392, 620, 443, 369,
		592, 765, 846, 266, 776, 421, 668,
		798, 817, 581, 996, 799, 2, 11, 231,
		415, 254, 24, 799, 99, 174, 703, 88,
		412, 171, 55, 532, 598, 872, 240,
		119, 474, 411, 448, 663, 396, 36, 30,
		500, 670, 979, 175, 523, 655, 512,
		403, 619, 483, 341, 685, 282, 920,
		533, 142, 457, 176, 41, 743, 724, 241,
		100, 636, 200, 663, 394, 458, 394,
		503, 72, 908, 655, 392, 546, 92,
		650, 521, 872,
	}
	timsort.Sorted(testCase)
	fmt.Println(testCase)
}

func genTest(length int) {
	var test []int = make([]int, length)
	fmt.Printf("{ ")
	for index := range test {
		test[index] = rand.Intn(1000)
		fmt.Printf("%d, ", test[index])
	}
	fmt.Printf("}")
}
