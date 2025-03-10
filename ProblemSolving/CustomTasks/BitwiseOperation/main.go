package main

import (
	checkparityofnum "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/CheckParityOfNum"
	divisionby2 "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/DivisionBy2"
	invertingthebits "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/InvertingTheBits"
	multiplicationby2 "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/MultiplicationBy2"
	thepoweroftwo "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/ThePowerOfTwo"
	"algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/utils"
	"fmt"
)

func main() {
	testForCheckParityOfNum()
	testForMultiplicationBy2()
	testForDivisionBy2()
	testForInvertingTheBits()
	testForThePowerOfTwo()
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

func testForInvertingTheBits() {
	fmt.Println("--- TEST FOR INVERTING THE BITS ---")
	var tests []int = []int{1, 3, 5, 2, 6, 8}
	for _, data := range tests {
		fmt.Printf("Num: ")
		utils.PrintAsBinary(data)
		fmt.Printf("Res: ")
		utils.PrintAsBinary(invertingthebits.Solution(data))
	}
}

func testForThePowerOfTwo() {
	fmt.Println("--- TEST FOR THE POWER OF TWO ---")
	var tests []int = []int{1, 3, 5, 2, 6, 8}
	for _, data := range tests {
		fmt.Printf("Num: %d, Result: %t;\n", data, thepoweroftwo.Solution(data))
	}
}
