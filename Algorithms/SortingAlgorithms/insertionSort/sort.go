package insertionsort

import "fmt"

func Sorted(slice []int) {
	for index := 1; index < len(slice); index++ {
		num := slice[index]
		subindex := index - 1
		for subindex >= 0 && num < slice[subindex] {
			tmp := slice[subindex]
			slice[subindex] = num
			slice[subindex+1] = tmp
			subindex--
		}
		fmt.Println(slice)
	}
	fmt.Println()
}
