package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Instruction struct {
	opcode     int
	modeParam1 int
	modeParam2 int
	modeParam3 int
}

func computeIntcode(input []int) ([]int, error) {
	for i := 0; i < len(input); {
		instruction := extractInstructionData(input[i])
		switch instruction.opcode {
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

func extractInstructionData(input int) Instruction {
	instruction := strconv.Itoa(input)

	if len(instruction) <= 2 {
		opcode, _ := strconv.Atoi(instruction)
		return Instruction{opcode: opcode}
	}

	mode1, mode2, mode3 := extractModes(instruction[0 : len(instruction)-2])
	opcode, _ := strconv.Atoi(instruction[len(instruction)-2:])

	return Instruction{opcode, mode1, mode2, mode3}
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
