package finduniquenum

func Solution(slice []int) int {
	var result int = 0
	for _, data := range slice {
		result ^= data
	}
	return result
}
