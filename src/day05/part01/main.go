package main

import (
	"fmt"
)

func computer(input int, opcode []int, pointer int) int {
	var command = opcode[pointer]
	if command == 99 {
		return input
	}

	if command == 1 {
		opcode[opcode[pointer+3]] = opcode[opcode[pointer+1]] + opcode[opcode[pointer+2]]
	} else if command == 2 {
		opcode[opcode[pointer+3]] = opcode[opcode[pointer+1]] * opcode[opcode[pointer+2]]
	} else if command == 3 {
		opcode[opcode[pointer+1]] = input
		return computer(input, opcode, pointer+2)
	} else if command == 4 {
		return computer(opcode[opcode[pointer+1]], opcode, pointer+2)
	}
	return computer(input, opcode, pointer+4)
}

func main() {

	var intcode []int = []int{3, 0, 4, 0, 99}

	fmt.Println(computer(99, intcode, 0))

}
