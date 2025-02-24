package main

import (
	tests "algorithms-and-data-structures/Algorithms/SearchingAlgorithms/DepthFirstSearch/tests"
	"fmt"
)

func main() {
	for index, test := range tests.TestCases {
		result := test.TreeLink.FindNode(test.Searching)
		fmt.Printf("Test num: %d, My result: %+v, TestCase result: %t\n", index+1, result, test.Result)
	}
}
