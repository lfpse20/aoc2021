package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoNumber struct {
	number   string
	isMarked bool
}

func main() {
	// get bingo data
	numbers, bingoBoards := parseBingoInputs()

out:
	// loop through drawn numbers
	for _, drawnNumber := range numbers {

		// mark number on bingo boards
		for _, bingoBoard := range bingoBoards {
			markNumberOnBoard(drawnNumber, bingoBoard)
		}

		// check if any board won
		for _, bingoBoard := range bingoBoards {
			if hasBoardWon(bingoBoard) {

				// check total of unmarked numbers for that board
				unmarkedSum := sumUnmarkedNumbers(bingoBoard)

				// multiply total with last number drawn
				fmt.Println(unmarkedSum * toInt(drawnNumber))
				break out
			}
		}
	}
}

func parseBingoInputs() ([]string, map[int][][]BingoNumber) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("file should be here")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := make([]string, 0)
	bingoBoards := make(map[int][][]BingoNumber, 0)
	boardNumber := 1
	for scanner.Scan() {
		line := scanner.Text()

		// is drawn numbers?
		if strings.Contains(line, ",") {
			numbers = strings.Split(line, ",")
			continue
		}

		// line break?
		if line == "" {
			boardNumber++
			continue
		}

		// must be bingo board line
		bingoLine := toBingoNumbers(line)
		if _, ok := bingoBoards[boardNumber]; ok {
			bingoBoards[boardNumber] = append(bingoBoards[boardNumber], [][]BingoNumber{bingoLine}...)
		} else {
			bingoBoards[boardNumber] = [][]BingoNumber{bingoLine}
		}
	}

	return numbers, bingoBoards
}

func toBingoNumbers(line string) []BingoNumber {
	line = cleanUpLine(line)
	numbers := strings.Split(line, " ")

	bingoNumbers := make([]BingoNumber, 0, len(numbers))
	for _, number := range numbers {
		bingoNumbers = append(bingoNumbers, BingoNumber{
			number:   number,
			isMarked: false,
		})
	}
	return bingoNumbers
}

func cleanUpLine(line string) string {
	line = strings.ReplaceAll(line, "  ", " ")
	line = strings.TrimSpace(line)
	return line
}

func markNumberOnBoard(drawnNumber string, bingoBoard [][]BingoNumber) {
	for i := range bingoBoard {
		for j := range bingoBoard[i] {
			if bingoBoard[i][j].number == drawnNumber {
				bingoBoard[i][j].isMarked = true
			}
		}
	}
}

func hasBoardWon(bingoBoard [][]BingoNumber) bool {

	markedNumberCount := 0

	// check horizontal lines
	for i := range bingoBoard {
		for j := range bingoBoard[i] {
			if bingoBoard[i][j].isMarked {
				markedNumberCount++
			}
		}
		if markedNumberCount == 5 {
			return true
		}
		markedNumberCount = 0
	}

	// check vertical lines
	for i := range bingoBoard {
		for j := range bingoBoard[i] {
			if bingoBoard[j][i].isMarked {
				markedNumberCount++
			}
		}
		if markedNumberCount == 5 {
			return true
		}
		markedNumberCount = 0
	}

	return false
}

func sumUnmarkedNumbers(bingoBoard [][]BingoNumber) int {
	sum := 0
	for i := range bingoBoard {
		for j := range bingoBoard[i] {
			if !bingoBoard[i][j].isMarked {
				sum += toInt(bingoBoard[i][j].number)
			}
		}
	}
	return sum
}

func toInt(number string) int {
	num, err := strconv.Atoi(number)
	if err != nil {
		panic("should be an int")
	}
	return num
}
