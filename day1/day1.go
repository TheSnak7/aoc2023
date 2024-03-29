package day1

import (
	"log"
	"os"
	"unicode"
)

func Day1() int {
	content, err := os.ReadFile("../day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	currentDigit := -1
	lineSum := 0
	totalSum := 0

	for i, char := range content {
		if unicode.IsDigit(rune(char)) {
			if currentDigit == -1 {
				currentDigit = int(char) - '0'
				lineSum += currentDigit * 10
			} else {
				currentDigit = int(char) - '0'
			}
		}
		if char == '\n' || i == len(content)-1 {
			lineSum += currentDigit

			totalSum += lineSum

			lineSum = 0
			currentDigit = -1
			i += 2
		}
	}
	return totalSum
}
