package main

import (
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4}

	doubleAt(values, 2)
	fmt.Println("Values: ", values)

	val := 10

	// local function to double val without pointer
	doubleVal(val)
	fmt.Println("Val: ", val)

	// local function to double the value with pointer
	doubleValPtr(&val)
	fmt.Println("Pointer Val: ", val)
}

func doubleAt(arr []int, index int) {
	arr[index] *= 2
}

func doubleVal(value int) {
	value *= 2
}

func doubleValPtr(value *int) {
	*value *= 2
}
