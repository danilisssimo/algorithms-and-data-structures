package timsort

import (
	insertionsort "algorithms-and-data-structures/Algorithms/SortingAlgorithms/insertionSort"
)

func Sorted(slice []int) {
	sliceLen := len(slice)
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

	for sliceSize := minRun; sliceSize < sliceLen; sliceSize += minRun {
		nextEndSubSliceIndex := sliceSize + minRun
		if nextEndSubSliceIndex > sliceLen {
			nextEndSubSliceIndex = sliceLen
		}

		pLeftSlice := slice[0:sliceSize]
		pRightSlice := slice[sliceSize:nextEndSubSliceIndex]

		leftSliceIndex, rightSliceIndex := 0, 0
		var leftSlice, rightSlice []int = make([]int, len(pLeftSlice)), make([]int, len(pRightSlice))

		copy(leftSlice, pLeftSlice)
		copy(rightSlice, pRightSlice)

		for leftSliceIndex+rightSliceIndex < nextEndSubSliceIndex {
			if rightSliceIndex >= len(rightSlice) ||
				(leftSliceIndex < len(leftSlice) && leftSlice[leftSliceIndex] < rightSlice[rightSliceIndex]) {
				slice[leftSliceIndex+rightSliceIndex] = leftSlice[leftSliceIndex]
				leftSliceIndex++
			} else {
				slice[leftSliceIndex+rightSliceIndex] = rightSlice[rightSliceIndex]
				rightSliceIndex++
			}
		}
	}
}

func getMinrun(n int) int {
	var lastBit int = 0
	for n >= 64 {
		lastBit |= n & 1
		n >>= 1
	}
	return n + lastBit
}
