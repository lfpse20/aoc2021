package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	lines, diagram := getParameters()

	for _, line := range lines {
		// skip diagonal lines
		if line.x1 != line.x2 && line.y1 != line.y2 {
			continue
		}

		// horizontal line check
		if line.y1 == line.y2 {
			if line.x1 < line.x2 {
				for line.x1 <= line.x2 {
					diagram[line.y1][line.x1]++
					line.x1++
				}
			} else {
				for line.x2 <= line.x1 {
					diagram[line.y1][line.x2]++
					line.x2++
				}
			}
		} else {
			// vertical line
			if line.y1 < line.y2 {
				for line.y1 <= line.y2 {
					diagram[line.y1][line.x1]++
					line.y1++
				}
			} else {
				for line.y2 <= line.y1 {
					diagram[line.y2][line.x1]++
					line.y2++
				}
			}
		}
	}

	counter := 0
	for i := range diagram {
		for j := range diagram[i] {
			if diagram[i][j] > 1 {
				counter++
			}
		}
	}

	fmt.Println(counter)
}

func getParameters() ([]Line, [][]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("oh man... file not found")
	}

	scanner := bufio.NewScanner(file)

	lines := make([]Line, 0)
	maxX := 0
	maxY := 0

	// get hydro thermla vent lines
	for scanner.Scan() {
		fileLine := scanner.Text()
		hydroThermalLine := strings.Split(fileLine, " -> ")
		point1 := strings.Split(hydroThermalLine[0], ",")
		point2 := strings.Split(hydroThermalLine[1], ",")

		line := Line{
			x1: toInt(point1[0]),
			y1: toInt(point1[1]),
			x2: toInt(point2[0]),
			y2: toInt(point2[1]),
		}
		lines = append(lines, line)

		if maxX < line.x1 {
			maxX = line.x1
		}
		if maxX < line.x2 {
			maxX = line.x2
		}
		if maxY < line.y1 {
			maxY = line.y1
		}
		if maxY < line.y2 {
			maxY = line.y2
		}
	}

	// initialize hydro thermal vents diagram
	hydroThermalVents := make([][]int, 0)
	for i := 0; i < maxY+1; i++ {
		hydroThermalVents = append(hydroThermalVents, make([]int, maxX+1))
	}

	return lines, hydroThermalVents
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("blah blah")
	}

	return i
}
