package main

import (
	"fmt"
	"math"
)

func main() {

	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Printf("Error Occured: %s\n", err)
	} else {
		fmt.Println("SQRT Success: ", math.Round(s1))
	}

	s2, err := sqrt(-2.0)
	if err != nil {
		fmt.Printf("Error Occured: %s\n", err)
	} else {
		fmt.Println("SQRT Success: ", math.Round(s2))
	}
}

func sqrt(val float64) (float64, error) {
	if val < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value (%f) not permitted", val)
	}

	return math.Sqrt(val), nil
}
