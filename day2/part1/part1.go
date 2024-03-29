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
	fmt.Println("Day 2 Part 1")
	fmt.Println("by Fabian Siegel")

	resetMemory()

	memory[1] = 12
	memory[2] = 2

	for memory[isp] != 99 {
		if memory[isp] == 1 {
			add()
		} else if memory[isp] == 2 {
			multiply()
		}
	}

	fmt.Println(memory)
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
}
