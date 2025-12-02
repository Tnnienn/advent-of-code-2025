package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	fmt.Println("Day 2, Part 1")
	data, err := os.ReadFile("day2/input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	ranges := strings.Split(string(data), ",")

	sum := 0

	for _, r := range ranges {
		//fmt.Println("Processing range:", r)
		ids := strings.Split(r, "-")
		firstId, err := strconv.Atoi(ids[0])
		if err != nil {
			fmt.Println("Error converting first ID:", err)
			continue
		}
		lastId, err := strconv.Atoi(ids[1])
		if err != nil {
			fmt.Println("Error converting last ID:", err)
			continue
		}
		for id := firstId; id <= lastId; id++ {
			//fmt.Println("Processing ID:", id)
			sId := strconv.Itoa(id)
			if len(sId)%2 != 0 {
				continue
			}
			found := true
			for i := 0; i < len(sId)-len(sId)/2; i++ {
				digit1, _ := strconv.Atoi(string(sId[i]))
				digit2, _ := strconv.Atoi(string(sId[(len(sId)/2)+i]))
				if digit1 != digit2 {
					found = false
					break
				}
			}
			if found {
				//fmt.Println("Invalid id: ", id)
				sum += id
			}

		}

	}
	fmt.Println("Final sum part 1:", sum)

}
