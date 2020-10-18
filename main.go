package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := extractInput()
	input[1] = 12
	input[2] = 2
	result, err := computeIntcode(input)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while computing intcode: %s", err))
		os.Exit(1)
	}

	fmt.Println(result[0])
}

func extractInput() []int {
	fileBytes, err := ioutil.ReadFile("day2_input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), ",")

	var input = []int{}

	for _, i := range sliceData {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		input = append(input, j)
	}

	return input
}
