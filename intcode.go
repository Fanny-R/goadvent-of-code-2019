package main

import (
	"errors"
	"fmt"
	"strconv"
)

func computeIntcode(input []int) ([]int, error) {
	for i := 0; i < len(input); {
		opcode, _, _, _ := extractInstructionData(input[i])
		switch opcode {
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

func extractInstructionData(instruction int) (int, int, int, int) {
	instructions := strconv.Itoa(instruction)

	if len(instructions) <= 2 {
		opcode, _ := strconv.Atoi(instructions)
		return opcode, 0, 0, 0
	}

	mode1, mode2, mode3 := extractModes(instructions[0 : len(instructions)-2])
	opcode, _ := strconv.Atoi(instructions[len(instructions)-2:])

	return opcode, mode1, mode2, mode3
}

func extractModes(modes string) (int, int, int) {
	modeParam1, _ := strconv.Atoi(modes[len(modes)-1 : len(modes)])
	modeParam2 := 0
	modeParam3 := 0
	if len(modes) >= 2 {
		modeParam2, _ = strconv.Atoi(modes[len(modes)-2 : len(modes)-1])
	}
	if len(modes) >= 3 {
		modeParam3, _ = strconv.Atoi(modes[len(modes)-3 : len(modes)-2])
	}

	return modeParam1, modeParam2, modeParam3
}
