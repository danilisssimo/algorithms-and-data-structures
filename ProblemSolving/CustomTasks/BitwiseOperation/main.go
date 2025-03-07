package main

import (
	checkparityofnum "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/CheckParityOfNum"
	multiplicationby2 "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/MultiplicationBy2"
	"fmt"
)

func main() {
	testForCheckParityOfNum()
	testForMultiplicationBy2()
}

func testForCheckParityOfNum() {
	fmt.Println("--- TEST FOR CHECK PARITY OF NUMBER ---")
	var tests []int = []int{1, 3, 5, 2, 6, 8}
	for _, data := range tests {
		fmt.Printf("Num: %d, Result: %t;\n", data, checkparityofnum.Solving(data))
	}
}

func testForMultiplicationBy2() {
	fmt.Println("--- TEST FOR MULTIOLICATION BY 2 ---")
	var tests []int = []int{1, 3, 5, 2, 6, 8}
	for _, data := range tests {
		fmt.Printf("Num: %d, Result: %d;\n", data, multiplicationby2.Solving(data))
	}
}
