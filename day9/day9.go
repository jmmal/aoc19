package main

import (
	"fmt"
	"advent-of-code-2019/intcode"
)

func boost(mem []int64) {
	input := make(chan int64, 1)
	output := make(chan int64, 20)
	halt := make(chan bool)

	computer := intcode.Computer{Memory: mem, IP: 0, RB :0}

	go computer.Execute(input, output, halt)

	input <- 1
	// input <- 2
	for val := range output {
		fmt.Println(val)
	}
}

func main() {
	program := intcode.ReadProgram()
	boost(program)
}
