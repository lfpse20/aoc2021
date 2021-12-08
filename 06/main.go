package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("ooooh noooo")
	}

	scanner := bufio.NewScanner(file)

	internalTimers := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, ",")

		for _, input := range inputs {
			internalTimers = append(internalTimers, toInt(input))
		}
	}

	groupedTimers := make(map[int]int, 8)
	for _, v := range internalTimers {
		groupedTimers[v]++
	}

	days := 256
	for i := 0; i < days; i++ {
		nextDayCounts := make(map[int]int, 8)
		for j := 8; j >= 0; j-- {
			count := groupedTimers[j]
			if j == 0 {
				nextDayCounts[6] += count
				nextDayCounts[8] = count
			} else {
				nextDayCounts[j-1] = count
			}
		}

		for j := 0; j <= 8; j++ {
			groupedTimers[j] = nextDayCounts[j]
		}
	}

	total := 0
	for _, count := range groupedTimers {
		total += count
	}

	fmt.Println("total:", total)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("whyyyyy??")
	}

	return i
}
