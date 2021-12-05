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
		panic("file should be there")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	binaries := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		binaries = append(binaries, strings.Split(line, ""))
	}

	oxygenGeneratorRatings := copy2dSlice(binaries)
	c02ScrubberRatings := copy2dSlice(binaries)

	oxygenGeneratorRating := getValue(oxygenGeneratorRatings, getOxygenFn)
	c02ScrubberRating := getValue(c02ScrubberRatings, getCO2Fn)

	fmt.Println(binaryToDecimal(oxygenGeneratorRating) * binaryToDecimal(c02ScrubberRating))
}

func copy2dSlice(src [][]string) [][]string {
	dupe := make([][]string, len(src))
	for i := range src {
		dupe[i] = make([]string, len(src[i]))
		copy(dupe[i], src[i])
	}
	return dupe
}

func getValue(binaries [][]string, valueSpecificFn func(int, []string, [][]string) [][]string) string {
	bitTracker := make([]string, 0)
	bitFrequencyCounter := 0
	for i := 0; i < len(binaries[0]); i++ {
		if len(binaries) == 1 {
			break
		}

		for j := 0; j < len(binaries); j++ {
			if binaries[j][i] == "1" {
				bitFrequencyCounter++
				bitTracker = append(bitTracker, "1")
			} else {
				bitFrequencyCounter--
				bitTracker = append(bitTracker, "0")
			}
		}

		binaries = valueSpecificFn(bitFrequencyCounter, bitTracker, binaries)

		bitFrequencyCounter = 0
		bitTracker = bitTracker[:0]
	}

	return strings.Join(binaries[0], "")
}

func getOxygenFn(bitFrequencyCounter int, bitTracker []string, binaries [][]string) [][]string {
	if bitFrequencyCounter > 0 {
		for j := len(bitTracker) - 1; j >= 0; j-- {
			if bitTracker[j] != "1" {
				binaries = remove(binaries, j)
			}
		}
	} else if bitFrequencyCounter < 0 {
		for j := len(bitTracker) - 1; j >= 0; j-- {
			if bitTracker[j] != "0" {
				binaries = remove(binaries, j)
			}
		}
	} else {
		for j := len(bitTracker) - 1; j >= 0; j-- {
			if bitTracker[j] != "1" {
				binaries = remove(binaries, j)
			}
		}
	}

	return binaries
}

func getCO2Fn(bitFrequencyCounter int, bitTracker []string, binaries [][]string) [][]string {
	if bitFrequencyCounter > 0 {
		for j := len(bitTracker) - 1; j >= 0; j-- {
			if bitTracker[j] == "1" {
				binaries = remove(binaries, j)
			}
		}
	} else if bitFrequencyCounter < 0 {
		for j := len(bitTracker) - 1; j >= 0; j-- {
			if bitTracker[j] == "0" {
				binaries = remove(binaries, j)
			}
		}
	} else {
		for j := len(bitTracker) - 1; j >= 0; j-- {
			if bitTracker[j] == "1" {
				binaries = remove(binaries, j)
			}
		}
	}

	return binaries
}

func remove(s [][]string, i int) [][]string {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = nil
	s = s[:len(s)-1]
	return s
}

func binaryToDecimal(binary string) int64 {
	dec, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic("wtf")
	}
	return dec
}
