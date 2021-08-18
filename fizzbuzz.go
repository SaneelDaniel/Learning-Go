package main

import (
	"fmt"
)

const aConstVar int = 1111

func main() {

	myInt := 1

	for myInt <= 20 {
		if myInt%3 == 0 && myInt%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if myInt%3 == 0 {
			fmt.Println("Fizz")
		} else if myInt%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%v\n", myInt)
		}
		myInt++
	}
}
