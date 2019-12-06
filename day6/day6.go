package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type satellite struct {
	name string
	orbits string
	orbittedBy []string
}

func getSatelitte(smap map[string]satellite, s1 string) satellite {
	if satellite, ok := smap[s1]; ok {
		return satellite
	}
	
	s := satellite{s1, "", []string{}}
	smap[s1] = s
	return s
}

func part1(smap map[string]satellite, curr satellite, visitedCount int) (total int) {	
	for _, s := range curr.orbittedBy {
		total += part1(smap, smap[s], visitedCount + 1)
	}

	return total + visitedCount
}

func part2(smap map[string]satellite, start string, end string) int {
	q := make([]string, 0)
	dist := make(map[string]int)

	q = append(q, start)
	dist[start] = 0

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		jumps, _ := dist[front]

		if front == end {
			return jumps
		}

		element := smap[front].orbits
		if _, ok := dist[element]; !ok {
			q = append(q, element)
			dist[element] = jumps + 1
		}

		for _, element := range smap[front].orbittedBy {
			if _, ok := dist[element]; !ok {
				q = append(q, element)
				dist[element] = jumps + 1
			}
		}
	}

	return -1
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	satellites := make(map[string]satellite)

	for scanner.Scan() {
		line := scanner.Text()
		objs := strings.Split(string(line), ")")

		s1 := getSatelitte(satellites, objs[0])
		s2 := getSatelitte(satellites, objs[1])

		s1.orbittedBy = append(s1.orbittedBy, s2.name)
		s2.orbits = s1.name
		satellites[s1.name] = s1
		satellites[s2.name] = s2
	}

	com, _ := satellites["COM"]
	san, _ := satellites["SAN"]
	you, _ := satellites["YOU"]

	result1 := part1(satellites, com, 0)
	result2 := part2(satellites, you.orbits, san.orbits)

	fmt.Println(result1)
	fmt.Println(result2)
}