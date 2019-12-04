package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var isp int = 0
var memory []int

func main() {
	fmt.Println("Day 2 Part 2")
	fmt.Println("by Fabian Siegel")
	memory = append(memory, 0)
	noun := 0
	verb := 0
	nonBreaker := true

	for noun = 0; noun < 100 || nonBreaker; noun++ {
		for verb = 0; verb < 100; verb++ {
			resetMemory()

			memory[1] = noun
			memory[2] = verb

			operate()

			if memory[0] == 19690720 {
				fmt.Println(memory)
				fmt.Println(100*noun + verb)
				nonBreaker = false
				break
			}
		}

	}

}

func operate() {

	for memory[isp] != 99 {
		if memory[isp] == 1 {
			add()
		} else if memory[isp] == 2 {
			multiply()
		}
	}

}

func add() {
	memory[memory[isp+3]] = memory[memory[isp+1]] + memory[memory[isp+2]]
	isp += 4
}

func multiply() {
	memory[memory[isp+3]] = memory[memory[isp+1]] * memory[memory[isp+2]]
	isp += 4

}

func resetMemory() {
	memory = nil
	file, _ := os.Open("input")
	defer file.Close()

	reader := bufio.NewReader(file)

	line, _, _ := reader.ReadLine()
	lineS := string(line)

	split := strings.Split(lineS, ",")

	for _, v := range split {
		number, _ := strconv.Atoi(v)
		memory = append(memory, number)

	}
	isp = 0
}
