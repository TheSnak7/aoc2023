package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4() int {
	return sumWinningNumbers("../day4/input.txt")
}

func sumWinningNumbers(path string) int {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		card := scanner.Text()
		winnningNums, chosenNums := parseCard(card)

		points := 0

		for _, wN := range winnningNums {
			for _, cN := range chosenNums {
				if cN == wN {
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
				}
			}
		}

		sum += points
	}

	return sum
}

func parseCard(card string) ([]int, []int) {
	numbers := strings.TrimSpace((strings.Split(card, ":")[1]))
	splitNums := strings.Split(numbers, "|")
	winnignNumbersStringArray := strings.Fields(splitNums[0])
	chosenNumbersStringArray := strings.Fields(strings.TrimSpace(splitNums[1]))

	winnningNums := make([]int, len(winnignNumbersStringArray))
	chosenNums := make([]int, len(chosenNumbersStringArray))

	for i, num := range winnignNumbersStringArray {
		n, _ := strconv.Atoi(num)
		winnningNums[i] = n
	}
	for i, num := range chosenNumbersStringArray {
		n, _ := strconv.Atoi(num)
		chosenNums[i] = n
	}

	return winnningNums, chosenNums
}
