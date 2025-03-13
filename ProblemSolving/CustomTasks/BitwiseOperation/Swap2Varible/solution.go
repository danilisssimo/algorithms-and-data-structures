package swap2varible

import (
	"fmt"
)

func Solution(a int, b int) (int, int) {
	fmt.Printf("Varible A: %8b\n", a)
	fmt.Printf("Varible B: %8b\n\n", b)
	a ^= b
	fmt.Printf("Varible A: %8b\n", a)
	fmt.Printf("Varible B: %8b\n\n", b)
	b ^= a
	fmt.Printf("Varible A: %8b\n", a)
	fmt.Printf("Varible B: %8b\n\n", b)
	a ^= b
	fmt.Printf("Varible A: %8b\n", a)
	fmt.Printf("Varible B: %8b\n\n", b)
	return a, b
}
