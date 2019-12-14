package main

import (
	"advent-of-code-2019/intcode"
	"fmt"
)

const (
	empty = 0
	wall  = 1
	block = 2
	hp    = 3
	ball  = 4
)

func main() {
	program := intcode.ReadProgram()
	program[0] = 2

	run(program)
}

func run(program []int64) {
	input := make(chan int64)
	output := make(chan int64)
	waiting := make(chan bool)

	grid := getGrid()

	computer := intcode.Computer{Memory: program}

	go computer.Execute(input, output, waiting)

	bc := 0
	ballX := int64(0)
	paddleX := int64(0)

	for {
		select {
		case w := <-waiting:
			if w && ballX > paddleX {
				input <- 1
			} else if ballX < paddleX {
				input <- -1
			} else {
				input <- 0
			}
			break

		case x := <-output:
			y := <-output
			id := <-output

			if x == -1 && y == 0 {
				fmt.Printf("Score: %d\n", id)
				break
			}

			// part 1
			if id == 2 {
				bc++
			}

			switch id {
			case ball:
				ballX = x
				break

			case hp:
				paddleX = x
			}

			grid[x][y] = id

		default:
			break
		}
	}

	fmt.Println(bc)
}

func getGrid() [][]int64 {
	grid := make([][]int64, 50)

	for i := 0; i < 50; i++ {
		grid[i] = make([]int64, 50)
	}

	return grid
}
