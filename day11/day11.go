package main

import (
	"advent-of-code-2019/intcode"
	"fmt"
)

const (
	UP    = 1
	DOWN  = 2
	LEFT  = 3
	RIGHT = 4
)

const (
	BLACK = 0
	WHITE = 1
)

const (
	LTURN = 1
	RTURN = 0
)

func main() {
	input := intcode.ReadProgram()

	paint(input)
}

type position struct {
	x int
	y int
}


func paint(program []int64) {
	sensor := make(chan int64, 1)
	nextMove := make(chan int64, 2)

	visitedPanels := make(map[position]bool)

	grid := getGrid()

	computer := intcode.Computer{Memory: program}

	go computer.Execute(sensor, nextMove)

	x := 50
	y := 50
	currentDirection := UP
	
	grid[x][y] = WHITE
	sensor <- grid[x][y] // Seed with (black) for part 1


	for color := range nextMove {
		turn := <-nextMove

		visitedPanels[position{x, y}] = true

		grid[x][y] = color

		switch currentDirection {
		case UP:
			if turn == LTURN {
				x--
				currentDirection = LEFT
			} else {
				x++
				currentDirection = RIGHT
			}
			break

		case LEFT:
			if turn == LTURN {
				y++
				currentDirection = DOWN
			} else {
				y--
				currentDirection = UP
			}
			break

		case RIGHT:
			if turn == LTURN {
				y--
				currentDirection = UP
			} else {
				y++
				currentDirection = DOWN
			}
			break

		case DOWN:
			if turn == LTURN {
				x++
				currentDirection = RIGHT
			} else {
				x--
				currentDirection = LEFT
			}
		default:
			panic("Unknown direction")
		}

		col := grid[x][y]
		sensor <- col
	}

	fmt.Println(len(visitedPanels))
	print(grid, x, y)
}

func print(grid [][]int64, x int, y int) {
	for _, row := range grid {
		for _, panel := range row {
			switch panel {
			case BLACK:
				fmt.Printf(" ")
			case WHITE:
				fmt.Printf("#")
			default:
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func getGrid() [][]int64 {
	grid := make([][]int64, 100)

	for i := 0; i < 100; i++ {
		grid[i] = make([]int64, 100)
	}
	
	return grid
}
