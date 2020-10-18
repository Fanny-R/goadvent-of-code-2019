package main

import (
	"fmt"
	"os"
)

func main() {
	input := []int{8, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	fmt.Println(input)

	result, err := computeIntcode(input)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while computing intcode: %s", err))
		os.Exit(1)
	}

	fmt.Println(result)
}
