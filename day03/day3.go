package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

type position struct {
	X, Y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func manhattanDistance(point position) int {
	return abs(point.X) + abs(point.Y)
}

func main() {
	file, _ := os.Open("input.txt")
	
	scanner := bufio.NewScanner(file)
	
	// Read line
	var wires = make(map[position]int)
	var minimum = math.MaxInt64
	secondLine := false

	for scanner.Scan() {
		steps := 0
		line := scanner.Text()
		moves := strings.Split(string(line), ",")
		currentX := 0
		currentY := 0

		for _, move := range moves {
			direction := move[:1]
			length, _ := strconv.Atoi(move[1:])

			xInc := 0
			yInc := 0

			switch direction {
			case "L":
				xInc = -1				
			case "R":
				xInc = 1
			case "U":
				yInc = 1
			case "D":
				yInc = -1
			default:
				panic("Unknown direction")
			}

			for i := 0; i < length; i++ {
				currentX += xInc
				currentY += yInc
				steps++

				pos := position{currentX, currentY}
				// dist := manhattanDistance(pos) // part 1
				prevSteps, exists := wires[pos]

				if secondLine && exists {
					// minimum = min(dist, minimum) // part 1
					minimum = min(steps + prevSteps, minimum) // part 2
				} else if !secondLine {
					wires[pos] = steps
				}
			}
		}
		secondLine = true

	}

	fmt.Println(minimum)
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}