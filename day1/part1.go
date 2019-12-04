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
	fmt.Println("Day 1 Part 1")
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
		nNumber := float64(element / 3)
		nNumber = math.Floor(nNumber)
		nNumber = nNumber - 2
		calcNumbers = append(calcNumbers, int(nNumber))
	}

	sum := 0
	for _, element := range calcNumbers {
		sum += element
	}
	fmt.Print(sum)
}
