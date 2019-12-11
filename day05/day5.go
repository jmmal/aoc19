package main

import (
	"fmt"
	"advent-of-code-2019/intcode"
)

func main() {
	var program = intcode.ReadProgram()
	
	input := make(chan int64, 1)
	output := make(chan int64, 20)
	halt := make(chan bool)

	computer := intcode.Computer{Memory: program, IP: 0, RB: 0}

	go computer.Execute(input, output, halt)

	input <- 1

	for val := range output {
		fmt.Println(val)
	}
}
