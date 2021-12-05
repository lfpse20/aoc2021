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
		panic("failed to open")
	}
	defer file.Close()

	horizontalPosition := 0
	depthPosition := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")

		switch command[0] {
		case "forward":
			horizontalPosition += toInt(command[1])
			depthPosition += aim * toInt(command[1])
		case "up":
			aim -= toInt(command[1])
		case "down":
			aim += toInt(command[1])
		default:
			panic(fmt.Sprintf("wtf is this command %s", command[0]))
		}
	}

	fmt.Println(horizontalPosition * depthPosition)
}

func toInt(number string) int {
	num, err := strconv.Atoi(number)
	if err != nil {
		panic("not a number for some reason")
	}

	return num
}
