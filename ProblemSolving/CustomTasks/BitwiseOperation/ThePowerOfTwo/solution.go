package thepoweroftwo

func Solution(number int) bool {
	return number > 0 && number&(number-1) == 0
}
