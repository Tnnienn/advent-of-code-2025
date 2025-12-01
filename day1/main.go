package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func nextIndex(currentIndex int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1 // Return -1 if the slice is empty
	}
	return (currentIndex + 1) % length
}

func prevIndex(currentIndex int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1 // Return -1 if the slice is empty
	}
	return (currentIndex - 1 + length) % length
}

func main() {
	fmt.Println("Executing Part 1")

	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var nums = []int{}

	for i := 0; i <= 99; i++ {
		nums = append(nums, i)
	}

	fmt.Println("Initial nums:", nums)

	index := 50

	counter1 := 0
	counter2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue // jump empty lines
		}
		//fmt.Println(line)
		command := line[:1]
		value := line[1:]
		val, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting value:", err)
			continue
		}
		switch command {
		case "L":
			for range val {
				index = prevIndex(index, nums)
				if index == 0 {
					counter2++
				}
			}
			if index == 0 {
				counter1++
			}
		case "R":
			for range val {
				index = nextIndex(index, nums)
				if index == 0 {
					counter2++
				}
			}
			if index == 0 {
				counter1++
			}
		default:
			fmt.Println("Unknown command:", command)
		}
	}

	fmt.Println("Final index:", index)
	fmt.Println("Counter part 1:", counter1)
	fmt.Println("Counter part 2:", counter2)
}
