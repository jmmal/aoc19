package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(arr []int) int {
	var counter = 0

	for {
		if counter+3 >= len(arr) {
			break
		}

		operation := arr[counter]

		if operation == 99 {
			break
		}

		a := arr[counter+1]
		b := arr[counter+2]
		dest := arr[counter+3]

		switch operation {
		case 1:
			arr[dest] = arr[a] + arr[b]
		case 2:
			arr[dest] = arr[a] * arr[b]
		default:
			return arr[0]
		}

		counter += 4
	}

	return arr[0]
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")

	var input = readInts(file)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			input[1] = i
			input[2] = j

			clone := make([]int, len(input))
			copy(clone, input)

			var result = part1(clone)
			if result == 19690720 {
				fmt.Println(100*i + j)
				return
			}
		}
	}
}

func readInts(file []byte) []int {
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
