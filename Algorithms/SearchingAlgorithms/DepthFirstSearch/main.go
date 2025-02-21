package main

import (
	"fmt"
	test "main/tests"
)

func main() {
	for index, test := range test.TestCases {
		result := test.TreeLink.FindNode(test.Searching)
		fmt.Printf("Test num: %d, My result: %+v, TestCase result: %t\n", index+1, result, test.Result)
	}
	// result := test.TestCases[5].TreeLink.FindNode(test.TestCases[5].Searching)
	// fmt.Printf("Test num: 6, My result: %t, TestCase result: %t\n", result, test.TestCases[5].Result)
}
