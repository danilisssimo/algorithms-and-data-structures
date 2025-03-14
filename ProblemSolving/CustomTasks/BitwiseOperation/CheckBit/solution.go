package checkbit

func Solution(num int, position int) bool {
	return (num>>(position-1))&1 == 1
}
