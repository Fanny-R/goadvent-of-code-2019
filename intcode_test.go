package main

import (
	"errors"
	"fmt"
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
	{[]int{0, 0, 0, 0, 99}, nil, errors.New("Unknown opcode 0")},
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
