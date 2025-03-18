package timsort

func Sorted(arr []int) []int {
	var sortedSlice []int
	sliceLen := len(arr)
	minRun := getMinrun(sliceLen)
	var index int = 0
	for index <= sliceLen {
		var subIndex int = index
		// Here insertion sort
		for subIndex <= minRun+subIndex {
			if arr[subIndex] > arr[subIndex+1] {

			}
		}
		// Here merging
	}
}

func insertionSort(slice []int) {}
func merging(slice []int) []int {}

// Find min subarray length for Timsort
func getMinrun(n int) int {
	var lastBit int = 0
	for n <= 64 {
		lastBit |= n & 1
		n >>= 1
	}
	return n + lastBit
}
