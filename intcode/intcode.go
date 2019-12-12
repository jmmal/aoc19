package intcode

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	Add         = 1
	Mul         = 2
	Input       = 3
	Output      = 4
	JumpIfTrue  = 5
	JumpIfFalse = 6
	LessThan    = 7
	Equals      = 8
	Rel         = 9
	Halt        = 99
)

const (
	Position  = 0
	Immediate = 1
	Relative  = 2
)

// Computer represents structure to run intcode
type Computer struct {
	Memory []int64
	IP     int64
	RB     int64
}

// Execute the opcode
func (computer Computer) Execute(input <-chan int64, output chan<- int64) {
	arr := make([]int64, 10000)
	copy(arr, computer.Memory)

	computer.Memory = arr
	computer.IP = 0
	computer.RB = 0

	for {
		op, modes := computer.getOpAndModes()

		switch op {
		case Add:
			a := computer.getArg(1, modes[0])
			b := computer.getArg(2, modes[1])
			c := computer.getArg(3, modes[2])

			computer.Memory[c] = computer.Memory[a] + computer.Memory[b]
			computer.IP += 4

		case Mul:
			a := computer.getArg(1, modes[0])
			b := computer.getArg(2, modes[1])
			c := computer.getArg(3, modes[2])
			computer.Memory[c] = computer.Memory[a] * computer.Memory[b]
			computer.IP += 4

		case Input:
			a := computer.getArg(1, modes[0])
			computer.Memory[a] = <-input
			computer.IP += 2

		case Output:
			a := computer.getArg(1, modes[0])
			output <- computer.Memory[a]
			computer.IP += 2

		case JumpIfTrue:
			a := computer.getArg(1, modes[0])
			b := computer.getArg(2, modes[1])

			if computer.Memory[a] != 0 {
				computer.IP = computer.Memory[b]
			} else {
				computer.IP += 3
			}
		case JumpIfFalse:
			a := computer.getArg(1, modes[0])
			b := computer.getArg(2, modes[1])

			if computer.Memory[a] == 0 {
				computer.IP = computer.Memory[b]
			} else {
				computer.IP += 3
			}
		case LessThan:
			a := computer.getArg(1, modes[0])
			b := computer.getArg(2, modes[1])
			c := computer.getArg(3, modes[2])

			if computer.Memory[a] < computer.Memory[b] {
				computer.Memory[c] = 1
			} else {
				computer.Memory[c] = 0
			}
			computer.IP += 4

		case Equals:
			a := computer.getArg(1, modes[0])
			b := computer.getArg(2, modes[1])
			c := computer.getArg(3, modes[2])

			if computer.Memory[a] == computer.Memory[b] {
				computer.Memory[c] = 1
			} else {
				computer.Memory[c] = 0
			}

			computer.IP += 4

		case Rel:
			a := computer.getArg(1, modes[0])
			computer.RB += computer.Memory[a]
			computer.IP += 2

		case Halt:
			close(output)
			return
		default:
			panic(fmt.Sprintf("Unknown instruction %d\n", op))
		}
	}
}

func (computer Computer) getOpAndModes() (int64, []int64) {
	instruction := computer.Memory[computer.IP]

	op := instruction % 100
	mode1 := (instruction / 100) % 10
	mode2 := (instruction / 1000) % 10
	mode3 := (instruction / 10000) % 10

	return op, []int64{mode1, mode2, mode3}
}

func (computer Computer) getArg(index int64, mode int64) int64 {
	switch mode {
	case Position:
		return computer.Memory[index+computer.IP]
	case Immediate:
		return index + computer.IP
	case Relative:
		return computer.Memory[index+computer.IP] + computer.RB

	}

	panic("Found unexpected mode")
}

// ReadProgram reads an int64Code program from file 'input.txt'
func ReadProgram() []int64 {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	strarr := strings.Split(string(file), ",")
	var ints = []int64{}
	for _, stringVal := range strarr {
		val, err := strconv.ParseInt(stringVal, 10, 0)

		if err != nil {
			panic(err)
		}

		ints = append(ints, val)
	}
	return ints
}
