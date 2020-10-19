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
			t.Errorf("got %q, want %q", output, tt.out)
		}
		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("got %q, want %q", err, tt.err)
		}
	}
}

var instructionextractiontests = []struct {
	instruction int
	opcode      int
	mode1       int
	mode2       int
	mode3       int
}{
	{99, 99, 0, 0, 0},
	{2, 2, 0, 0, 0},
	{1002, 2, 0, 1, 0},
	{11102, 2, 1, 1, 1},
	{10002, 2, 0, 0, 1},
	{102, 2, 1, 0, 0},
}

func TestExtractInstructionData(t *testing.T) {
	for _, tt := range instructionextractiontests {
		// TODO return an object instead of that mess
		opcode, mode1, mode2, mode3 := extractInstructionData(tt.instruction)

		if tt.opcode != opcode {
			t.Errorf("Wrong opcode : got %d, want %d", opcode, tt.opcode)
		}
		if tt.mode1 != mode1 {
			t.Errorf("Wrong mode1 : got %d, want %d", mode1, tt.mode1)
		}
		if tt.mode2 != mode2 {
			t.Errorf("Wrong mode2 : got %d, want %d", mode2, tt.mode2)
		}
		if tt.mode3 != mode3 {
			t.Errorf("Wrong mode3 : got %d, want %d", mode3, tt.mode3)
		}
	}
}
