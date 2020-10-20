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

var intcode []int
var i int
var relativeBase int

func computeIntcode(input []int) ([]int, error) {
	intcode = input
	i = 0
	relativeBase = 0

	for i < len(intcode) {
		instruction := extractInstructionData(intcode[i])
		switch instruction.opcode {
		case 1:
			// add
			value1, value2 := getValuesFromParamsMode(instruction.modeParam1, instruction.modeParam2)
			intcode[intcode[i+3]] = value1 + value2
			i += 4
		case 2:
			// multiply
			value1, value2 := getValuesFromParamsMode(instruction.modeParam1, instruction.modeParam2)
			intcode[intcode[i+3]] = value1 * value2
			i += 4
		case 3:
			// intput
			var first int
			fmt.Print("intcode: ")
			fmt.Scanln(&first)
			intcode[intcode[i+1]] = first
			i += 2
		case 4:
			// output
			value1 := intcode[i+1]
			if instruction.modeParam1 == 0 {
				value1 = intcode[value1]
			}

			fmt.Println(fmt.Sprintf("Output: %d", value1))
			i += 2
		case 5:
			// jump if true
			value1, value2 := getValuesFromParamsMode(instruction.modeParam1, instruction.modeParam2)
			if value1 != 0 {
				i = value2
			} else {
				i += 3
			}
		case 6:
			// jump if false
			value1, value2 := getValuesFromParamsMode(instruction.modeParam1, instruction.modeParam2)
			if value1 == 0 {
				i = value2
			} else {
				i += 3
			}
		case 7:
			// less than
			value1, value2 := getValuesFromParamsMode(instruction.modeParam1, instruction.modeParam2)
			if value1 < value2 {
				intcode[intcode[i+3]] = 1
			} else {
				intcode[intcode[i+3]] = 0
			}
			i += 4
		case 8:
			// equals
			value1, value2 := getValuesFromParamsMode(instruction.modeParam1, instruction.modeParam2)
			if value1 == value2 {
				intcode[intcode[i+3]] = 1
			} else {
				intcode[intcode[i+3]] = 0
			}
			i += 4
		case 9:
			// adjusts the relative base
			relativeBase += instruction.modeParam1
			i += 2
		case 99:
			return intcode, nil
		default:
			return nil, errors.New(fmt.Sprintf("Unknown opcode %d on position %d", intcode[i], i))
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

func getValuesFromParamsMode(modeParam1, modeParam2 int) (int, int) {
	value1 := intcode[i+1]
	if modeParam1 == 0 {
		value1 = intcode[value1]
	} else if modeParam1 == 2 {
		value1 = intcode[relativeBase+value1]
	}

	value2 := intcode[i+2]
	if modeParam2 == 0 {
		value2 = intcode[value2]
	} else if modeParam2 == 2 {
		value1 = intcode[relativeBase+value2]
	}

	return value1, value2
}
