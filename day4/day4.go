package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4() int {
	return calculateTotalScratchcards("../day4/input.txt")
}

// Part 1
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

// Part 2
func calculateTotalScratchcards(path string) int {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	//Assume at least one card is always present
	cardIndex := 0
	matches := make([]int, 0)

	for scanner.Scan() {
		card := scanner.Text()
		cardIndex++

		winnningNums, chosenNums := parseCard(card)

		matched := 0

		//Assume all numbers are unique on each side
		for _, wN := range winnningNums {
			for _, cN := range chosenNums {
				if cN == wN {
					matched++
				}
			}
		}
		matches = append(matches, matched)
	}

	cardsWon := 0

	cards := make([]int, len(matches))

	for i := range cards {
		cards[i] = 1
	}

	for i, matchCount := range matches {
		//fmt.Printf("Card %v (%v) has %v matches\n", i+1, cards[i], matchCount)
		cardCount := cards[i]
		cardsWon += cardCount

		for j := range matchCount {
			//Bounds check
			if j+i < len(cards) {
				cards[j+i+1] += cardCount
				//fmt.Printf("  Copy won of Card %v\n", j+i+2)
			}
		}
	}

	return cardsWon
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
