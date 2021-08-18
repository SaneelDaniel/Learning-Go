package main

import (
	"fmt"
)

func main() {
	count := 0
	arr := make([]Number, 0)
	count, arr = CheckEvenEnded(count, arr)

	defer fmt.Println("The Total Values :", count)
	defer DoPrinting(arr, 4)
}

// DoPrinting function prints the values in the array
// ( Within the 0 to Limit-1 range )
func DoPrinting(arr []Number, limit int) {
	for i, number := range arr {
		if i > limit-1 {
			break
		}
		fmt.Printf("The Number Is: %d, and The Total Is: %d\n", number.Total, number.Number)
	}
}

//CheckEvenEnded function checks for even ended squares
// for numbers from 100-500 and appends even ended squares
// to a list
func CheckEvenEnded(count int, arr []Number) (int, []Number) {
	for i := 100; i < 500; i++ {
		tot := i * i

		str := fmt.Sprintf("%d", tot)
		intVal := i
		if str[0] == str[len(str)-1] {
			count++
			num := Number{tot, intVal}
			arr = append(arr, num)
		}
	}

	return count, arr
}

// Number "This is a number struct"
type Number struct {
	Total  int
	Number int
}
