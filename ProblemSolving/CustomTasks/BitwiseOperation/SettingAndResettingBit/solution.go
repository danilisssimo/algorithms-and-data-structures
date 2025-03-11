package settingandresettingbit

func Solution(number int, index int, value bool) int {
	currentValue := 0
	if value {
		currentValue = 1
	}
	count := 0
	temp := 0
	for number != 0 {
		if count == index {
			number |= number & currentValue
		}
		temp |= number & 1
		temp <<= 1
		number >>= 1
		count += 1
	}
	temp >>= 1
	return temp
}
