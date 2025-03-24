package timsort

import (
	insertionsort "algorithms-and-data-structures/Algorithms/SortingAlgorithms/insertionSort"
	"fmt"
)

func Sorted(slice []int) []int {
	sliceLen := len(slice)
	sortedSlice := make([]int, sliceLen)
	minRun := getMinrun(sliceLen)
	var index int = 0
	for index <= sliceLen {
		nextIndexSubSlice := index + minRun
		if nextIndexSubSlice > sliceLen {
			nextIndexSubSlice = sliceLen
		}
		insertionsort.Sorted(slice[index:nextIndexSubSlice])
		index += minRun
	}

	// Need to save between result
	for sliceSize := minRun; sliceSize < sliceLen; sliceSize += minRun {
		nextEndSubSliceIndex := sliceSize + minRun
		if nextEndSubSliceIndex > sliceLen {
			nextEndSubSliceIndex = sliceLen
		}

		leftSliceIndex, rightSliceIndex := 0, sliceSize
		fmt.Printf("\n%d %d\n", sliceSize, nextEndSubSliceIndex)
		fmt.Printf("\n%+v\n%+v\n", slice[0:sliceSize], slice[sliceSize:nextEndSubSliceIndex])

		for leftSliceIndex+rightSliceIndex-sliceSize < nextEndSubSliceIndex {
			if (leftSliceIndex < sliceSize) && (slice[leftSliceIndex] < slice[rightSliceIndex]) {
				sortedSlice[leftSliceIndex+rightSliceIndex-sliceSize] = slice[leftSliceIndex]
				leftSliceIndex++
			} else {
				sortedSlice[leftSliceIndex+rightSliceIndex-sliceSize] = slice[rightSliceIndex]
				rightSliceIndex++
			}
			fmt.Println(sortedSlice)
			// if slice[leftSliceIndex] <= slice[rightSliceIndex] {
			// 	sortedSlice[leftSliceIndex+rightSliceIndex-sliceSize+1] = slice[leftSliceIndex]
			// 	leftSliceIndex++
			// } else {
			// 	sortedSlice[leftSliceIndex+rightSliceIndex-sliceSize+1] = slice[rightSliceIndex]
			// 	rightSliceIndex++
			// }
		}
	}
	return sortedSlice
}

func getMinrun(n int) int {
	var lastBit int = 0
	for n >= 64 {
		lastBit |= n & 1
		n >>= 1
	}
	return n + lastBit
}

// [
// 	18 58 93 130 136 173 209 211 224 230 255
// 	308 440 455 510 511 514 554 576 576 583
// 	596 603 616 649 708 744 746 768 816 854
// 	872 895 904 944

// 	15 69 121 161 162 213
// 	266 369 373 392 421 443 474 541 543 564
// 	592 597 620 625 668 680 720 729 755 758
// 	765 776 809 826 844 846 864 886 998

// 	2 11
// 	24 30 36 55 88 99 119 171 174 175
// 	231 240 254 396 411 412 415 448 474 500
// 	532 581 598 663 670 703 798 799 799 817
// 	872 979 996

// 	41 72 92 100 142 176 200 241
// 	282 341 392 394 394 403 457 458 483 503
// 	512 521 523 533 546 619 636 650 655 655
// 	663 685 724 743 872 908 920
// ]
