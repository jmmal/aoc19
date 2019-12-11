package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"github.com/gitchander/permutation"
	"advent-of-code-2019/intcode"
)

func execute(mem []int, phases []int ) int {
	a  := make(chan int, 1)
	ab := make(chan int)
	bc := make(chan int)
	cd := make(chan int)
	de := make(chan int)

	halt := make(chan bool)

	go intcode.Execute(mem, a,  ab, halt)
	go intcode.Execute(mem, ab, bc, halt)
	go intcode.Execute(mem, bc, cd, halt)
	go intcode.Execute(mem, cd, de, halt)
	go intcode.Execute(mem, de, a,  halt)

	a <- phases[0]
	ab <- phases[1]
	bc <- phases[2]
	cd <- phases[3]
	de <- phases[4]

	a<-0

	for i := 0; i < 5; i++ {
		<-halt
	}

	return <-a
}

func main() {
	program := readInput()

	phases := []int{5,6,7,8,9}

	perm := permutation.New(permutation.IntSlice(phases))
	best := -10000

	for perm.Next() {
		result := execute(program, phases)
		best = max(best, result)
	}

	fmt.Println(best)
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func readInput() []int {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	strarr := strings.Split(string(file), ",")
	var ints = []int{}
	for _, stringVal := range strarr {
		val, err := strconv.Atoi(stringVal)

		if err != nil {
			panic(err)
		}

		ints = append(ints, val)
	}
	return ints
}
