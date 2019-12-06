package main

import (
	"fmt"
)

func computer(input int, procode []int, pointer int) int {
	var command int = procode[pointer]

	var opcode int = command
	var paramMode []int = []int{0, 0}
	if command > 99 {
		opcode = command % 100
		// if opcode == 99 {
		// 	// fmt.Println(procode)
		// 	return opcode
		// }
		// fmt.Println("opcode:", opcode)
		var p2 int = command / 1000
		var p1 int = command % 1000 / 100
		paramMode = []int{p1, p2}
		// fmt.Println("paramMode:", paramMode)
	}

	if opcode == 1 {
		var value1 int
		var value2 int

		if paramMode[0] == 0 {
			value1 = procode[procode[pointer+1]]
		} else {
			value1 = procode[pointer+1]
		}
		if paramMode[1] == 0 {
			value2 = procode[procode[pointer+2]]
		} else {
			value2 = procode[pointer+2]
		}
		procode[procode[pointer+3]] = value1 + value2
		return computer(input, procode, pointer+4)
	} else if opcode == 2 {
		var value1 int
		var value2 int

		if paramMode[0] == 0 {
			value1 = procode[procode[pointer+1]]
		} else {
			value1 = procode[pointer+1]
		}
		if paramMode[1] == 0 {
			value2 = procode[procode[pointer+2]]
		} else {
			value2 = procode[pointer+2]
		}
		procode[procode[pointer+3]] = value1 * value2
		return computer(input, procode, pointer+4)
	} else if opcode == 3 {
		var p1 int
		if paramMode[0] == 0 {
			p1 = procode[pointer+1]
		} else {
			p1 = pointer + 1
		}
		procode[p1] = input
		return computer(input, procode, pointer+2)
	} else if opcode == 4 {
		var p1 int
		if paramMode[0] == 0 {
			p1 = procode[pointer+1]
		} else {
			p1 = pointer + 1
		}
		fmt.Println("input: ", procode[p1])
		return computer(procode[p1], procode, pointer+2)
	} else if opcode == 5 {
		var value1 int
		var value2 int

		if paramMode[0] == 0 {
			value1 = procode[procode[pointer+1]]
		} else {
			value1 = procode[pointer+1]
		}
		if paramMode[1] == 0 {
			value2 = procode[procode[pointer+2]]
		} else {
			value2 = procode[pointer+2]
		}
		if value1 != 0 {
			return computer(input, procode, value2)
		}
		return computer(input, procode, pointer+3)
	} else if opcode == 6 {
		var value1 int
		var value2 int

		if paramMode[0] == 0 {
			value1 = procode[procode[pointer+1]]
		} else {
			value1 = procode[pointer+1]
		}
		if paramMode[1] == 0 {
			value2 = procode[procode[pointer+2]]
		} else {
			value2 = procode[pointer+2]
		}
		if value1 == 0 {
			return computer(input, procode, value2)
		}
		return computer(input, procode, pointer+3)
	} else if opcode == 7 {
		var value1 int
		var value2 int

		if paramMode[0] == 0 {
			value1 = procode[procode[pointer+1]]
		} else {
			value1 = procode[pointer+1]
		}
		if paramMode[1] == 0 {
			value2 = procode[procode[pointer+2]]
		} else {
			value2 = procode[pointer+2]
		}
		if value1 < value2 {
			procode[procode[pointer+3]] = 1
		} else {
			procode[procode[pointer+3]] = 0
		}
		return computer(input, procode, pointer+4)
	} else if opcode == 8 {
		var value1 int
		var value2 int

		if paramMode[0] == 0 {
			value1 = procode[procode[pointer+1]]
		} else {
			value1 = procode[pointer+1]
		}
		if paramMode[1] == 0 {
			value2 = procode[procode[pointer+2]]
		} else {
			value2 = procode[pointer+2]
		}
		if value1 == value2 {
			procode[procode[pointer+3]] = 1
		} else {
			procode[procode[pointer+3]] = 0
		}
		return computer(input, procode, pointer+4)
	} else if opcode == 99 {
		fmt.Println("output: ", input)
		return input
	}
	fmt.Println("Dropped out of loop", opcode, input)
	return input
}

func main() {

	var intcode []int = []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1101, 90, 64, 225, 1101, 15, 56, 225, 1, 14, 153, 224, 101, -147, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 3, 224, 1, 224, 223, 223, 2, 162, 188, 224, 101, -2014, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 223, 224, 223, 1001, 18, 81, 224, 1001, 224, -137, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 3, 224, 1, 223, 224, 223, 1102, 16, 16, 224, 101, -256, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 6, 224, 1, 223, 224, 223, 101, 48, 217, 224, 1001, 224, -125, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 3, 224, 1, 224, 223, 223, 1002, 158, 22, 224, 1001, 224, -1540, 224, 4, 224, 1002, 223, 8, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1101, 83, 31, 225, 1101, 56, 70, 225, 1101, 13, 38, 225, 102, 36, 192, 224, 1001, 224, -3312, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 4, 224, 1, 224, 223, 223, 1102, 75, 53, 225, 1101, 14, 92, 225, 1101, 7, 66, 224, 101, -73, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1101, 77, 60, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 7, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 329, 1001, 223, 1, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 344, 101, 1, 223, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 359, 101, 1, 223, 223, 7, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 374, 101, 1, 223, 223, 8, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 389, 1001, 223, 1, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 404, 101, 1, 223, 223, 1107, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 419, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 434, 1001, 223, 1, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 449, 1001, 223, 1, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 464, 101, 1, 223, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 479, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 494, 1001, 223, 1, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 509, 101, 1, 223, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 524, 101, 1, 223, 223, 107, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 539, 1001, 223, 1, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 554, 101, 1, 223, 223, 1008, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 569, 1001, 223, 1, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 599, 1001, 223, 1, 223, 108, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 614, 1001, 223, 1, 223, 108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 629, 101, 1, 223, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 644, 101, 1, 223, 223, 107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 659, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}

	// intcode = []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	// intcode = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
	// 	1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
	// 	999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}

	fmt.Println(computer(5, intcode, 0))

}
