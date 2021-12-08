package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic("ooooh noooo")
	}

	scanner := bufio.NewScanner(file)

	positions := make([]int, 0)
	length := 0
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, ",")

		for _, input := range inputs {
			val := toInt(input)

			if length < val {
				length = val
			}
			positions = append(positions, val)
		}
	}

	groupedPositions := make(map[int]int, length)
	for _, position := range positions {
		groupedPositions[position]++
	}

	minFuelCost := calculateFuelCostAt(groupedPositions, 0)
	for i := 1; i < length; i++ {
		fuelCost := calculateFuelCostAt(groupedPositions, i)
		if minFuelCost > fuelCost {
			minFuelCost = fuelCost
		}
	}

	fmt.Println(minFuelCost)
}

func calculateFuelCostAt(positions map[int]int, goalPosition int) int {
	fuelCost := 0
	for position, countAtPosition := range positions {
		if countAtPosition == 0 {
			continue
		}
		diff := math.Abs(float64(position - goalPosition))
		gasCostPerDiff := 0
		for i := 1; i < int(diff); i++ {
			gasCostPerDiff += i
		}
		diff += float64(gasCostPerDiff)
		fuelCost += int(diff * float64(countAtPosition))
	}
	return fuelCost
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("whyyyyy??")
	}

	return i
}
