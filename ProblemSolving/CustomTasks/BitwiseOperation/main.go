package main

import (
	checkparityofnum "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/CheckParityOfNum"
	divisionby2 "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/DivisionBy2"
	multiplicationby2 "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/MultiplicationBy2"
	"fmt"
)

func main() {
	testForCheckParityOfNum()
	testForMultiplicationBy2()
	testForDivisionBy2()
}

func testForCheckParityOfNum() {
	fmt.Println("--- TEST FOR CHECK PARITY OF NUMBER ---")
	var tests []int = []int{1, 3, 5, 2, 6, 8}
	for _, data := range tests {
		fmt.Printf("Num: %d, Result: %t;\n", data, checkparityofnum.Solution(data))
	}
}

func testForMultiplicationBy2() {
	fmt.Println("--- TEST FOR MULTIOLICATION BY 2 ---")
	var tests []int = []int{1, 3, 5, 2, 6, 8}
	for _, data := range tests {
		fmt.Printf("Num: %d, Result: %d;\n", data, multiplicationby2.Solution(data))
	}
}

func testForDivisionBy2() {
	fmt.Println("--- TEST FOR DIVISION BY 2 ---")
	var tests []int = []int{1, 3, 5, 2, 6, 8}
	for _, data := range tests {
		fmt.Printf("Num: %d, Result: %d;\n", data, divisionby2.Solution(data))
	}
}
