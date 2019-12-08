package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)


func part1(mass int) int {
 return (mass / 3) - 2
}

func part2(mass int) int {
	fuel := part1(mass)
	total := 0

	for fuel > 0 {
		total += fuel
		fuel = part1(fuel)
	}
	
	return total
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")

	result := 0

	for _, value := range strings.Split(string(file), "\n") {
		if value == "" {
			continue
		}

		mass, _ := strconv.Atoi(value)
		
		result += part1(mass)
		// total += Part2(mass)
	}

	fmt.Println(result)
}

