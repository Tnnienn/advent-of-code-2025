package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2() {
	fmt.Println("Day 2, Part 2")

	data, err := os.ReadFile("day2/input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	text := strings.TrimSpace(string(data))
	ranges := strings.Split(text, ",")

	sum := 0

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}

		ids := strings.Split(r, "-")
		if len(ids) != 2 {
			fmt.Println("Invalid range:", r)
			continue
		}

		firstId, err := strconv.Atoi(strings.TrimSpace(ids[0]))
		if err != nil {
			fmt.Println("Error converting first ID:", err)
			continue
		}
		lastId, err := strconv.Atoi(strings.TrimSpace(ids[1]))
		if err != nil {
			fmt.Println("Error converting last ID:", err)
			continue
		}

		for id := firstId; id <= lastId; id++ {
			sId := strconv.Itoa(id)
			length := len(sId)

			found := false

			for val := 1; val <= length/2; val++ {
				if length%val != 0 {
					continue
				}

				base := sId[0:val]
				ok := true

				for j := val; j < length; j += val {
					if sId[j:j+val] != base {
						ok = false
						break
					}
				}

				if ok {
					found = true
					break
				}
			}

			if found {
				sum += id
			}
		}
	}

	fmt.Println("Final sum part 2:", sum)
}
