package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type asteroid struct {
	x float64
	y float64
}

func getBestPosition(asteroids []asteroid) (int, asteroid) {
	best := 0
	var bp asteroid

	for _, startAsteroid := range asteroids {
		count := 0

		for _, endAsteroid := range asteroids {
			if startAsteroid == endAsteroid {
				continue
			}

			angleToEnd := calculateAngle(startAsteroid, endAsteroid)
			distToEnd := dist(startAsteroid, endAsteroid)

			reachable := true

			for _, middleAsteroid := range asteroids {
				if middleAsteroid == startAsteroid || middleAsteroid == endAsteroid {
					continue
				}

				distToMiddle := dist(startAsteroid, middleAsteroid)
				middleAngle := calculateAngle(startAsteroid, middleAsteroid)

				if angleToEnd == middleAngle && distToMiddle < distToEnd {
					reachable = false
					break
				}
			}

			if reachable {
				count++
			}
		}

		// fmt.Printf("Count for %v is %d\n", startAsteroid, count)
		if count > best {
			best = count
			bp = startAsteroid
		}

	}

	return best, bp
}

func vaporise(asteroids []asteroid, baseStation asteroid) {
	asteroidAngles := make(map[float64][]asteroid)
	angles := []float64{}

	for _, asteroid := range asteroids {
		if asteroid == baseStation {
			continue
		}

		angle := calculateAngle(asteroid, baseStation)

		asteroidAngles[angle] = append(asteroidAngles[angle], asteroid)
		angles = append(angles, angle)
	}

	sort.Float64s(angles)

	angleIndex := 0
	prevAngle := float64(-1)
	vc := 0
	for vc < 201 {
		angle := angles[angleIndex]
		targets := asteroidAngles[angle]

		closestDistance := math.MaxFloat64
		var asteroidToVaporise asteroid
		var vaporiseIndex int

		for asteroidIndex, asteroid := range targets {
			distance := dist(asteroid, baseStation)

			if distance < closestDistance {
				closestDistance = distance
				asteroidToVaporise = asteroid
				vaporiseIndex = asteroidIndex
			}
		}

		fmt.Printf("%d: vaporising asteroid %v\n", vc+1, asteroidToVaporise)

		newOptions := append(targets[:vaporiseIndex], targets[vaporiseIndex+1:]...)
		asteroidAngles[angle] = newOptions

		angles = append(angles[:angleIndex], angles[angleIndex+1:]...)

		prevAngle = angle
		vc++

		for angles[angleIndex] == prevAngle {
			angleIndex = (angleIndex + 1) % len(angles)
		}
	}
}

func main() {
	asteroids := readInput()

	bp, result := getBestPosition(asteroids)

	fmt.Println(bp)
	fmt.Println(result)

	vaporise(asteroids, result)
}


func dist(start, end asteroid) float64 {
	return math.Sqrt(math.Pow(end.x-start.x, 2) + math.Pow(end.y-start.y, 2))
}

// Calculates the angle (in degrees between two points
// Angles start from the 12 O'Clock position and rotate
// clockwise
func calculateAngle(pos1, pos2 asteroid) float64 {
	deltaX := pos1.x - pos2.x
	deltaY := pos1.y - pos2.y

	angle := math.Atan2(deltaY, deltaX)*180/math.Pi + 90

	if angle < 0 {
		angle += 360
	}

	return angle
}

func readInput() []asteroid {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var asteroids []asteroid

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		for x, pos := range line {
			if string(pos) == "#" {
				asteroids = append(asteroids, asteroid{float64(x), float64(y)})
			}
		}

		y++
	}

	return asteroids
}
