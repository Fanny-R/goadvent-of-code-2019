package main

import (
	"errors"
	"fmt"
)

func computeIntcode(input []int) ([]int, error) {
	for i := 0; i < len(input); i += 4 {
		switch input[i] {
		case 1:
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
		case 2:
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		case 99:
			return input, nil
		default:
			return nil, errors.New(fmt.Sprintf("Unknown opcode %d on position %d", input[i], i))
		}
	}

	return nil, errors.New("Oops, it seems that something went wrong")

}
