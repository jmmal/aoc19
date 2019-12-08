package intcode

import (
	"fmt"
)

// Execute the opcode
func Execute(mem []int, input <-chan int, output chan<- int, halt chan<- bool) {
	arr := make([]int, len(mem))
	copy(arr, mem)

	ip := 0
	
	for {
		op, modes := getOpAndModes(arr[ip])
		
		switch op {
		case 1:
			a := getArg(ip+1, arr, modes[0])
			b := getArg(ip+2, arr, modes[1])
			c := arr[ip+3]
			
			arr[c] = a + b
			ip += 4
		case 2:
			a := getArg(ip+1, arr, modes[0])
			b := getArg(ip+2, arr, modes[1])
			c := arr[ip+3]
			arr[c] = a * b
			ip += 4
		case 3:
			arr[arr[ip+1]] = <-input
			ip += 2
		case 4:
			a := getArg(ip+1, arr, modes[0])
			output<-a
			ip += 2
		case 5:
			a := getArg(ip+1, arr, modes[0])
			b := getArg(ip+2, arr, modes[1])
			
			if a != 0 {
				ip = b
				} else {
				ip += 3
			}
		case 6:
			a := getArg(ip+1, arr, modes[0])
			b := getArg(ip+2, arr, modes[1])

			if a == 0 {
				ip = b
			} else {
				ip += 3
			}
		case 7:
			a := getArg(ip+1, arr, modes[0])
			b := getArg(ip+2, arr, modes[1])
			c := arr[ip+3]

			if a < b {
				arr[c] = 1
			} else {
				arr[c] = 0
			}
			ip += 4
		case 8:
			a := getArg(ip+1, arr, modes[0])
			b := getArg(ip+2, arr, modes[1])
			c := arr[ip+3]

			if a == b {
				arr[c] = 1
			} else {
				arr[c] = 0
			}

			ip += 4
		case 99:
			halt <- true 
			return
		default:
			panic(fmt.Sprintf("Unknown instruction %d\n", op))
		}
	}
}

func getOpAndModes(instruction int) (int, []int) {
	op := instruction % 100
	mode1 := (instruction / 100) % 10
	mode2 := (instruction / 1000) % 10
	mode3 := (instruction / 10000) % 10

	return op, []int{mode1, mode2, mode3}
}

func getArg(index int, memory []int, mode int) int {
	if mode == 0 {
		return memory[memory[index]]
	} else if mode == 1 {
		return memory[index]
	}

	panic("Found unexpected mode")
}
