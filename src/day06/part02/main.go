package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(name string) []string {
	var lines []string
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

type Object struct {
	key string
}

type Orbit struct {
	object     Object
	main       string
	satellites []string
}

func getOrbit(key string, orbits []Orbit) Orbit {
	var orbit Orbit
	for _, searchOrbit := range orbits {
		if searchOrbit.object.key == key {
			return searchOrbit
		}
	}
	orbit.object = Object{key}
	orbit.main = ""
	return orbit
}

func addOrbit(key string, orbits []Orbit) []Orbit {
	var orbit Orbit
	for _, searchOrbit := range orbits {
		if searchOrbit.object.key == key {
			return orbits
		}
	}
	orbit.object = Object{key}
	orbit.main = ""
	orbits = append(orbits, orbit)
	return orbits
}

func addSatellite(main Orbit, satellite Orbit, orbits []Orbit) []Orbit {
	var mainFound bool = false
	var satFound bool = false
	for i := 0; i < len(orbits); i++ {
		if !mainFound && orbits[i].object.key == main.object.key {
			orbits[i].satellites = append(orbits[i].satellites, satellite.object.key)
			mainFound = true
		} else if orbits[i].object.key == satellite.object.key {
			orbits[i].main = main.object.key
			satFound = true
		}
		if mainFound && satFound {
			break
		}
	}
	return orbits
}

func getOrbits(inputs []string) []Orbit {
	var orbits []Orbit
	for _, line := range inputs {
		var input []string = strings.Split(line, ")")
		var main Orbit = getOrbit(input[0], orbits)
		orbits = addOrbit(main.object.key, orbits)
		var satellite Orbit = getOrbit(input[1], orbits)
		orbits = addOrbit(satellite.object.key, orbits)
		orbits = addSatellite(main, satellite, orbits)
	}
	return orbits
}

func getOrbitCount(orbit Orbit, orbits []Orbit, count int) int {
	if orbit.main != "" {
		count++
		count = getOrbitCount(getOrbit(orbit.main, orbits), orbits, count)
	}

	return count
}

func getOrbitJumps(orbit Orbit, orbits []Orbit, jumps []string) []string {
	if orbit.main != "" {
		jumps = append(jumps, orbit.main)
		jumps = getOrbitJumps(getOrbit(orbit.main, orbits), orbits, jumps)
	}
	return jumps
}

func main() {

	var file = "input.txt"

	// file = "test.txt"
	var input []string = readFile(file)

	var orbits []Orbit = getOrbits(input)

	var san Orbit = getOrbit("SAN", orbits)
	var you Orbit = getOrbit("YOU", orbits)
	fmt.Println(san)
	fmt.Println(you)
	var sanJumps []string = getOrbitJumps(san, orbits, []string{})
	var youJumps []string = getOrbitJumps(you, orbits, []string{})
	fmt.Println(sanJumps)
	fmt.Println(youJumps)

	var shortestJump int = 0
	for youi := 0; youi < len(youJumps); youi++ {
		var foundCommon bool = false
		for sani := 0; sani < len(sanJumps); sani++ {
			if sanJumps[sani] == youJumps[youi] {
				shortestJump = shortestJump + sani
				foundCommon = true
				break
			}
		}
		if foundCommon {
			// first common path key
			shortestJump = shortestJump + youi
			break
		}
	}
	fmt.Println("shortest jump: ", shortestJump)

	// fmt.Println(orbits)
	// var total = 0
	// for _, orbit := range orbits {
	// 	var count int = getOrbitCount(orbit, orbits, 0)
	// 	total = total + count
	// 	fmt.Println("orbitCount:", count)
	// }
	// fmt.Println("total orbitCount:", total)

	// fmt.Println(orbits)
}
