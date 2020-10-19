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
	inst := strconv.Itoa(input)

	if len(inst) <= 2 {
		return Instruction{opcode: stringToInt(inst)}
	}

	opcode := stringToInt(inst[len(inst)-2:])
	modes := inst[0 : len(inst)-2]
	modeParam1 := stringToInt(modes[len(modes)-1 : len(modes)])

	instruction := Instruction{opcode: opcode, modeParam1: modeParam1}

	if len(modes) >= 2 {
		instruction.modeParam2 = stringToInt(modes[len(modes)-2 : len(modes)-1])
	}
	if len(modes) >= 3 {
		instruction.modeParam3 = stringToInt(modes[len(modes)-3 : len(modes)-2])
	}

	return instruction
}

func stringToInt(input string) int {
	result, err := strconv.Atoi(input)

	if err != nil {
		// TODO do something more clever
		panic(err)
	}

	return result
}
