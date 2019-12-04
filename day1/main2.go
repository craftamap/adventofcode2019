package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 1 Part 2")
	fmt.Println("by Fabian Siegel")

	file, _ := os.Open("input")
	defer file.Close()

	reader := bufio.NewReader(file)

	var numbers []int

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		number, err := strconv.Atoi(string(line))
		numbers = append(numbers, number)
	}

	var calcNumbers []int

	for _, element := range numbers {
		calcNumbers = append(calcNumbers, calcRecursiveFuel(element))
	}

	sum := 0
	for _, element := range calcNumbers {
		sum += element
		fmt.Print(element)
		fmt.Print("\n")
	}
	fmt.Print(sum)
}

func calcSimpleFuel(input int) int {
	a := float64(input / 3)
	a = math.Floor(a)
	a = a - 2
	return int(a)
}

func calcRecursiveFuel(currentMass int) int {
	returnVal := 0
	addedFuel := calcSimpleFuel(currentMass)
	for addedFuel > 0 {
		returnVal += addedFuel
		addedFuel = calcSimpleFuel(addedFuel)
	}
	return returnVal
}
