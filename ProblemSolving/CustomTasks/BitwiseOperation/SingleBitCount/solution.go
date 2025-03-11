package singlebitcount

func Solution(number int) int {
	var count int = 0
	for number != 0 {
		count += number & 1
		number >>= 1
	}
	return count
}
