package main

import (
	"fmt"
	"strings"
)

func main() {
	countMe := `
		There are too many words which I wish to put, 
		but I will buy myself a BMW before my birthday, next year!
	`

	words := map[string]int{}

	wordsArr := strings.Fields(countMe)

	for _, word := range wordsArr {
		words[strings.ToLower(word)]++
	}

	fmt.Printf("The counts array: \n", words)
}
