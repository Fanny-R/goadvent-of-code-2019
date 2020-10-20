package main

import (
	"errors"
	"reflect"
	"testing"
)

var intcodetests = []struct {
	in  []int
	out []int
	err error
}{
	{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}, nil},
	{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}, nil},
	{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}, nil},
	{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, nil},
	{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, nil},
	{[]int{0, 0, 0, 0, 99}, nil, errors.New("Unknown opcode 0 on position 0")},
}

func TestComputeIntcode(t *testing.T) {
	for _, tt := range intcodetests {
		output, err := computeIntcode(tt.in)

		if !reflect.DeepEqual(output, tt.out) {
			t.Errorf("got %d, want %d", output, tt.out)
		}
		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("got %q, want %q", err, tt.err)
		}
	}
}

var instructionextractiontests = []struct {
	in  int
	out Instruction
}{
	{99, Instruction{99, 0, 0, 0}},
	{2, Instruction{2, 0, 0, 0}},
	{1002, Instruction{2, 0, 1, 0}},
	{11102, Instruction{2, 1, 1, 1}},
	{10002, Instruction{2, 0, 0, 1}},
	{102, Instruction{2, 1, 0, 0}},
}

func TestExtractInstructionData(t *testing.T) {
	for _, tt := range instructionextractiontests {
		extractedInstruction := extractInstructionData(tt.in)
		// TODO actually compare the 2 structs
		if extractedInstruction.opcode != tt.out.opcode {
			t.Errorf("Wrong opcode : got %d, want %d", extractedInstruction.opcode, tt.out.opcode)
		}
		if extractedInstruction.modeParam1 != tt.out.modeParam1 {
			t.Errorf("Wrong mode1 : got %d, want %d", extractedInstruction.modeParam1, tt.out.modeParam1)
		}
		if extractedInstruction.modeParam2 != tt.out.modeParam2 {
			t.Errorf("Wrong mode2 : got %d, want %d", extractedInstruction.modeParam2, tt.out.modeParam2)
		}
		if extractedInstruction.modeParam3 != tt.out.modeParam3 {
			t.Errorf("Wrong mode3 : got %d, want %d", extractedInstruction.modeParam3, tt.out.modeParam3)
		}
	}
}
