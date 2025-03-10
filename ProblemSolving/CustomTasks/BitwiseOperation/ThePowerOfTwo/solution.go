package thepoweroftwo

func Solution(number int) bool {
	return number&(number-1) == 0
}
