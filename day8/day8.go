package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var result int
	
	input := input()
	picture := getBlankPicture()
	
	fewest := 1000000000
	
	i := 0

	for i < len(input) {
		counts := [3]int{0, 0, 0}

		for row := 0; row < 6; row++ {
			for col := 0; col < 25; col++ {
				pixel := string(input[i])
				update := picture[row][col] == "" || picture[row][col] == "2"

				switch pixel {
				case "0":
					counts[0]++
				case "1":
					counts[1]++
				case "2":
					counts[2]++
				default:
				}

				if update {
					picture[row][col] = pixel
				}

				i++
			}
		}

		if counts[0] < fewest {
			fewest = counts[0]
			result = counts[1] * counts[2]
		}
	}

	fmt.Println(result)

	for _, row := range picture {
		fmt.Println(row)
	}
}

func input() string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var input string

	for scanner.Scan() {
		input = scanner.Text()
	}

	return input
}

func getBlankPicture() [][]string {
	picture := make([][]string, 6)
	for i := 0; i < 6; i++ {
		picture[i] = make([]string, 25)
	}

	return picture
}