package main

import (
	"fmt"
	"strconv"
)

func isValid(password string) bool {
	foundAdjacents := false
	adjacentCount := 0

	for i := 0; i < len(password) - 1; i++ {
		left, right := password[i], password[i + 1]

		if left > right {
			return false
		}

		if left == right {
			adjacentCount++
		} else if adjacentCount == 1 {
			foundAdjacents = true
		} else {
			adjacentCount = 0
		}
	}

	if adjacentCount == 1 {
		foundAdjacents = true
	}

	return foundAdjacents
}

func main() {
	count := 0

	for val := 347312; val <= 805915; val++ {
		password := strconv.Itoa(val)
		if isValid(password) {
			count++
		}
	}

	fmt.Println(count)
}