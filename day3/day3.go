package day3

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func Day3() int {
	return findAdjacentPartNums("../day3/input.txt")
}

type PartNumber struct {
	start  int
	end    int
	number int
	marked bool
}

type CharCoords struct {
	x int
	y int
}

func findAdjacentPartNums(path string) int {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lineCount := 0
	var numberRanges [][]PartNumber
	var charsList []CharCoords

	//Parse the data
	for scanner.Scan() {
		line := scanner.Text()

		ranges, charsFound := parseLine(line, lineCount)

		numberRanges = append(numberRanges, ranges)
		charsList = append(charsList, charsFound...)

		lineCount += 1
	}

	//Mark every adjacent part number
	for _, char := range charsList {
		for y := char.y - 1; y <= char.y+1; y++ {
			//In bounds check
			if y >= 0 && y < lineCount {
				partsNumbers := &numberRanges[y]
				partsNumbersLength := len(numberRanges[y])
				for x := char.x - 1; x <= char.x+1; x++ {
					for i := range partsNumbersLength {
						partNumber := &((*partsNumbers)[i])
						if partNumber.number == 34 {
							fmt.Println("Checking 34")
						}

						if x >= partNumber.start && x <= partNumber.end {
							partNumber.marked = true
						}
					}
				}
			}
		}
	}

	//Sum all the marked numbers
	sum := 0

	for _, numberLine := range numberRanges {
		for _, partNumber := range numberLine {
			if partNumber.marked {
				sum += partNumber.number
			}
		}
	}

	return sum
}

func parseLine(line string, currentY int) ([](PartNumber), []CharCoords) {
	number := 0
	numberStart := -1

	var numberRanges []PartNumber
	var charsList []CharCoords

	for i, char := range line {
		if unicode.IsDigit(rune(char)) {
			number *= 10
			number += int(char - '0')

			if numberStart == -1 {
				numberStart = i
			}
		} else {
			if char != '.' {
				charCoords := &CharCoords{i, currentY}
				charsList = append(charsList, *charCoords)
			}
			if numberStart != -1 {
				numberRange := &PartNumber{numberStart, i - 1, number, false}
				numberRanges = append(numberRanges, *numberRange)

				numberStart = -1
				number = 0
			}

		}
	}

	if numberStart != -1 {
		numberRange := &PartNumber{numberStart, len(line) - 1, number, false}
		numberRanges = append(numberRanges, *numberRange)

		numberStart = -1
		number = 0
	}

	return numberRanges, charsList
}
