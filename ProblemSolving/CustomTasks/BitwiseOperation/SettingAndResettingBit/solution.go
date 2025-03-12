package settingandresettingbit

func Solution(number int, index int, value bool) int {
	if value {
		number |= 1 << index
	} else {
		mask := ^(1 << index)
		number &= mask
	}
	return number
}
