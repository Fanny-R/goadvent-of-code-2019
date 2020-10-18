package main

import (
	"errors"
	"fmt"
)

func computeIntcode(input []int) ([]int, error) {
	for i := 0; i < len(input); {
		switch input[i] {
		case 1:
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
			i += 4
		case 2:
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
			i += 4
		case 3:
			var first int
			fmt.Print("Input: ")
			fmt.Scanln(&first)
			input[input[i+1]] = first
			i += 2
		case 4:
			fmt.Println(fmt.Sprintf("Output: %d", input[input[i+1]]))
			i += 2
		case 99:
			return input, nil
		default:
			return nil, errors.New(fmt.Sprintf("Unknown opcode %d on position %d", input[i], i))
		}
	}

	return nil, errors.New("Oops, it seems that something went wrong")

}
