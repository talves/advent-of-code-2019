package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkPasswords(lower int, upper int) []int {
	var list []int = []int{}

	for potentialPassword := lower; potentialPassword <= upper; potentialPassword++ {
		pwdArray := strconv.Itoa(potentialPassword)
		var last = pwdArray[len(pwdArray)-1]
		var okToAdd = false

		for indx := len(pwdArray) - 2; indx >= 0; indx-- {
			current := pwdArray[indx]
			if current == last {
				okToAdd = true
			} else if current > last {
				okToAdd = false
				break
			}
			last = current
		}
		// fmt.Println(okToAdd, pwdArray)

		if okToAdd {
			list = append(list, potentialPassword)
		}
	}

	return list
}

func main() {
	var input string = "165432-707912"
	var stringSplit []string = strings.Split(input, "-")
	lowerRange, err := strconv.Atoi(stringSplit[0])
	if err != nil {
		panic(err)
	}
	upperRange, err := strconv.Atoi(stringSplit[1])
	if err != nil {
		panic(err)
	}

	var passwords []int = checkPasswords(lowerRange, upperRange)

	fmt.Println(passwords, len(passwords))
}
