package main

import (
	"fmt"
)

func computer(input []int, start int) []int {
	var end int = start + 4
	// our opcode block
	var opcode []int = input[start:end]
	// fmt.Println("'opcode'", opcode)
	// fmt.Println(start, end)

	var command = opcode[0]
	var position1 int = opcode[1]
	var position2 int = opcode[2]
	var position int = opcode[3]

	if command == 1 {
		input[position] = input[position1] + input[position2]
	} else if command == 2 {
		input[position] = input[position1] * input[position2]
	}

	if input[end] == 99 {
		return input
	}
	return computer(input, end)

}

func main() {

	var answer int = 19690720
	// var answer int = 3058646
	var oktoend = false
	var intcode []int
	var resetCode int

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			intcode = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 6, 19, 2, 19, 6, 23, 1, 23, 5, 27, 1, 9, 27, 31, 1, 31, 10, 35, 2, 35, 9, 39, 1, 5, 39, 43, 2, 43, 9, 47, 1, 5, 47, 51, 2, 51, 13, 55, 1, 55, 10, 59, 1, 59, 10, 63, 2, 9, 63, 67, 1, 67, 5, 71, 2, 13, 71, 75, 1, 75, 10, 79, 1, 79, 6, 83, 2, 13, 83, 87, 1, 87, 6, 91, 1, 6, 91, 95, 1, 10, 95, 99, 2, 99, 6, 103, 1, 103, 5, 107, 2, 6, 107, 111, 1, 10, 111, 115, 1, 115, 5, 119, 2, 6, 119, 123, 1, 123, 5, 127, 2, 127, 6, 131, 1, 131, 5, 135, 1, 2, 135, 139, 1, 139, 13, 0, 99, 2, 0, 14, 0}
			intcode[1] = noun
			intcode[2] = verb
			var output []int = computer(intcode, 0)

			oktoend = output[0] == answer
			if oktoend {
				resetCode = noun*100 + verb
				break
			}
		}
		if oktoend {
			break
		}
	}

	fmt.Println(resetCode)

}
