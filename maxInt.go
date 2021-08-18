package main

import (
	"fmt"
)

func main() {
	numsArr := []int{50, 500, 12, 0, 12876921387, 88, -99}

	max, maxIndex := findMax(numsArr)

	fmt.Printf("The MAx Val: %d, Max Index: %d\n", max, maxIndex)
}

func findMax(myArg []int) (int, int) {

	max := -9999999999999999
	maxIndex := -1

	for i, num := range myArg {
		if num > max {
			max = num
			maxIndex = i
		}
	}

	return max, maxIndex
}
