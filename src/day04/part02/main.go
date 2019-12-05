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
		last, _ := strconv.Atoi(string(pwdArray[len(pwdArray)-1]))
		var okToAdd = false
		var equalCount int = 1
		var hasEqualPair = false

		for indx := len(pwdArray) - 2; indx >= 0; indx-- {
			current, _ := strconv.Atoi(string(pwdArray[indx]))
			if current == last {
				equalCount++
				okToAdd = true
				if indx == 0 {
					if equalCount == 2 && current == last {
						hasEqualPair = true
					}
				} else {
					next, _ := strconv.Atoi(string(pwdArray[indx-1]))
					if equalCount == 2 && current != next {
						hasEqualPair = true
					}
				}
			} else if current > last {
				okToAdd = false
				break
			} else {
				equalCount = 1
			}
			last = current
		}

		if okToAdd && hasEqualPair {
			list = append(list, potentialPassword)
		}
		hasEqualPair = false
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
