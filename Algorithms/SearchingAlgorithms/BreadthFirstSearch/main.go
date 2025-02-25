package main

import (
	"algorithms-and-data-structures/Algorithms/SearchingAlgorithms/BreadthFirstSearch/tests"
	"fmt"
)

func main() {
	testsForSearchingNode(tests.TestCases)
}

func testsForSearchingNode(tests []tests.TestCase) {
	for index, test := range tests {
		funcResult := test.TreeLink.FindNode(test.Searching)
		fmt.Printf(
			"Test num: %d, IsNodeExist: %t, SearchingNodeValue: %d, FuncResult: %+v\n",
			index+1,
			test.Result,
			test.Searching,
			funcResult,
		)
	}
}
