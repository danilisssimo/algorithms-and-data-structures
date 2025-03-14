package main

import (
	checkbit "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/CheckBit"
	checkparityofnum "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/CheckParityOfNum"
	divisionby2 "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/DivisionBy2"
	finduniquenum "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/FindUniqueNum"
	invertingthebits "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/InvertingTheBits"
	multiplicationby2 "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/MultiplicationBy2"
	settingandresettingbit "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/SettingAndResettingBit"
	singlebitcount "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/SingleBitCount"
	swap2varible "algorithms-and-data-structures/ProblemSolving/CustomTasks/BitwiseOperation/Swap2Varible"
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
	testForSingleBitCount()
	testForSettingAndResettingBit()
	testForSwap2Varible()
	testForFindUniqueNum()
	testForCheckBit()
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

func testForSingleBitCount() {
	fmt.Println("--- TEST FOR SINGLE BIT COUNT ---")
	var tests []int = []int{33, 12, 9, 34, 0, 4, 321, 8123849123}
	for _, data := range tests {
		fmt.Printf("Num: %b, Result:%d;\n", data, singlebitcount.Solution(data))
	}
}

func testForSettingAndResettingBit() {
	fmt.Println("--- TEST FOR SETTING AND RESETTING BIT ---")
	var testsForSet []int = []int{33, 12, 9, 34, 4, 321, 8123849123}
	fmt.Println("-- test for set bit --")
	for _, data := range testsForSet {
		fmt.Printf("Num: %b, Result:%b;\n", data, settingandresettingbit.Solution(data, 3, true))
	}
	var testsForReset []int = []int{15, 45, 37, 20, 4, 2228, 1810438}
	fmt.Println("-- test for reset bit --")
	for _, data := range testsForReset {
		fmt.Printf("Num: %b, Result:%b;\n", data, settingandresettingbit.Solution(data, 2, false))
	}
}

func testForSwap2Varible() {
	fmt.Println("--- TEST FOR SWAP 2 VARIBLES ---")
	var tests [][]int = [][]int{{33, 12}, {15, 45}, {11, 91}, {67, 12}, {90, 80}, {14, 43}}
	for _, data := range tests {
		a, b := swap2varible.Solution(data[0], data[1])
		fmt.Printf("Num: %+v, Result:[%d %d];\n", data, a, b)
	}
}

func testForFindUniqueNum() {
	fmt.Println("--- TEST FOR SWAP 2 VARIBLES ---")
	var tests [][]int = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9}, // 10 встречается один раз
		{10, 20, 30, 40, 50, 10, 20, 30, 40},                       // 50 встречается один раз
		{7, 3, 5, 3, 7},                                            // 5 встречается один раз
		{1, 1, 2, 2, 3, 3, 4, 4, 5},                                // 5 встречается один раз
		{42},                                                       // 42 встречается один раз (массив из одного элемента)
		{100, 200, 300, 100, 200},                                  // 300 встречается один раз
	}
	for _, data := range tests {
		res := finduniquenum.Solution(data)
		fmt.Printf("Num: %+v, Result: %d;\n", data, res)
	}
}

func testForCheckBit() {
	fmt.Println("--- TEST FOR CHECK BIT ---")
	var tests []int = []int{33, 12, 9, 34, 0, 4, 321, 8123849123}
	for _, data := range tests {
		fmt.Printf("Num: %b %d, Result: %t;\n", data, 1, checkbit.Solution(data, 1))
	}
}
